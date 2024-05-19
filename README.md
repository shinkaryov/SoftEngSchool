# Software Engineering School Test Task

To start this application follow this steps.

  1. Clone this repository to your local folder
  2. Install Docker and Docker compose if not installed yet
  3. Configure SMTP setting for email account to send letters as in this [guide](https://www.gmass.co/blog/gmail-smtp/)
  4. Create .env file in the main directory with keys POSTGRES_USER, POSTGRES_DB, POSTGRES_PASSWORD, EMAIL_ADDRESS and EMAIL_PASSWORD
  5. Run ```docker compose up```

Now 3 services are running:
  1. PostgreSQL Database with table subscriptions with 1 row in it.
  2. SMTP server to send emails
  3. API ([docs](https://github.com/Falko05/se_school/blob/main/gses2swagger.yaml))

Example of API methods:
  1. <img width="973" alt="image" src="https://github.com/shinkaryov/SoftEngSchool/assets/81474894/c846dab5-2c8c-4c76-bcde-787496a06ca7">

  2. <img width="973" alt="image" src="https://github.com/shinkaryov/SoftEngSchool/assets/81474894/aca1b32d-fb9b-4c65-aa56-aa795e402e40"><img width="973" alt="image" src="https://github.com/shinkaryov/SoftEngSchool/assets/81474894/3566d5d9-8cce-4fdd-b212-6e11f67311ed">

  3. At 14:00 Kyiv time app sends current exchange rates to all subscribed users iterating emails extracted from table subscriptions

