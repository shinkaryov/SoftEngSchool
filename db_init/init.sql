-- Description: This script is used to initialize the database.
-- It creates the necessary tables and inserts some initial data.
CREATE TABLE IF NOT EXISTS public.subscriptions
(
    email VARCHAR(255) PRIMARY KEY,
    is_subscribed INTEGER NOT NULL
);

INSERT INTO public.subscriptions (email, is_subscribed) VALUES
    ('shynkarov.oleksandr@lll.kpi.ua', 1);