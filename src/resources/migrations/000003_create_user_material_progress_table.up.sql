DROP TABLE IF EXISTS public.user_material_progress;
CREATE TABLE public.user_material_progress
(
    id         SERIAL PRIMARY KEY,
    user_id             INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    material_id         INT NOT NULL REFERENCES materials (id) ON DELETE CASCADE,
    is_completed        BOOLEAN   DEFAULT FALSE,
    completion_date     TIMESTAMP,
    progress_percentage INT CHECK (progress_percentage BETWEEN 0 AND 100),
    last_accessed       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, material_id) -- Чтобы избежать дублирования прогресса
);

CREATE INDEX idx_user_progress ON user_material_progress (user_id);
CREATE INDEX idx_material_progress ON user_material_progress (material_id);