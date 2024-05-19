package main

import (
	"api/connect_to_db"
	"api/exchange_rate"
	"api/get_emails"
	"api/send_email"
	"api/subscribe_email"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
)

func getRate(c *gin.Context) {
	// Get the current exchange rate
	var url = "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=USD&json"
	rate, err := exchange_rate.GetRateFromUrl(url)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"rate": rate,
		})
	}
}

// TODO: add validation
func subscribeEmail(c *gin.Context) {
	// Connect to the database
	db, err := connect_to_db.ConnectDB()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	// Insert the email into the database
	email := c.PostForm("email")
	err = subscribe_email.InsertEmail(db, email)
	if err != nil {
		if err.Error() == "already subscribed" {
			c.JSON(409, gin.H{
				"error": "already subscribed",
			})
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "subscribed",
		})
	}
}

// TODO: implenent async sending
func sendEmails() {
	// Connect to the database
	db, err := connect_to_db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Select emails from the database
	emails, err := get_emails.SelectEmailsFromDB(db)
	if err != nil {
		log.Fatalf("Error selecting emails from the database: %v", err)
	} else {
		// Get the current exchange rate
		subject := "CURRENT EXCHANGE RATE"
		var url = "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=USD&json"
		rate, err := exchange_rate.GetRateFromUrl(url)
		if err != nil {
			log.Fatalf("Error getting exchange rate: %v", err)
		}
		var body = fmt.Sprintf("Current exchange rate is %f", rate)

		// Send an email to each recipient
		for _, email := range emails {
			err := send_email.SendEmail(email, subject, body)
			if err != nil {
				log.Printf("Error sending email to %s: %v", email, err)
			} else {
				fmt.Printf("Email sent to %s successfully\n", email)
			}
		}
	}
}

func main() {
	router := gin.Default()

	router.GET("/rate", getRate)
	router.POST("/subscribe", subscribeEmail)

	// Schedule the sendEmails function to run every day at 11:00 UTC (14:00 Kyiv time)
	c := cron.New()
	_, err := c.AddFunc("00 11 * * *", sendEmails)
	if err != nil {
		fmt.Println("Error scheduling sendEmails:", err)
		return
	}
	c.Start()
	defer c.Stop()

	router.Run(":8080")
}
