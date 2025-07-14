CREATE TABLE "notes" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "title" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "file_path" VARCHAR(255) NOT NULL,
  "user_id" VARCHAR(255) NOT NULL, -- Diğer servisten geleceği için şimdilik VARCHAR
  "course_code" VARCHAR(50),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);