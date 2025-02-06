DROP TABLE IF EXISTS public.materials;
CREATE TABLE public.materials
(
    id               SERIAL PRIMARY KEY,
    title            VARCHAR(255) NOT NULL,
    description      TEXT,
    content_url      VARCHAR(512) NOT NULL,
    difficulty_level VARCHAR(20) CHECK (difficulty_level IN ('beginner', 'intermediate', 'advanced')),
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT
ON COLUMN materials.content_url IS 'Ссылка на материал (видео, статья, тест и т.д.)';