DROP TABLE IF EXISTS public.users;
CREATE TABLE public.users
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(50) UNIQUE  NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);