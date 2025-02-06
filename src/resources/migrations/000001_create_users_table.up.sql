DROP TABLE IF EXISTS public.users;
CREATE TABLE public.users
(
    user_id           SERIAL PRIMARY KEY,
    username          VARCHAR(50) UNIQUE  NOT NULL,
    email             VARCHAR(255) UNIQUE NOT NULL,
    password_hash     VARCHAR(255)        NOT NULL,
    registration_date TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    last_login        TIMESTAMP,
    role              VARCHAR(20) DEFAULT 'user' CHECK (role IN ('user', 'admin', 'author'))
);

COMMENT ON COLUMN users.role IS 'Роль: user, admin, author';