CREATE TABLE IF NOT EXISTS "notes" (
    "id" UUID PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "file_path" TEXT NOT NULL,
    "user_id" VARCHAR(255) NOT NULL,
    "university_id" VARCHAR(255) NOT NULL,
    "course_code" VARCHAR(50),
    "status" note_status NOT NULL DEFAULT 'pending',
    "download_count" BIGINT NOT NULL DEFAULT 0,
    "average_rating" NUMERIC(3, 2) NOT NULL DEFAULT 0.00,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "deleted_at" TIMESTAMPTZ
);
