CREATE TABLE IF NOT EXISTS "user" (
	"id" CHAR(36) PRIMARY KEY,
	"username" VARCHAR(255) UNIQUE NOT NULL,
	"password" VARCHAR(255) NOT NULL,
	"user_type" VARCHAR(255),
	"created_at" TIMESTAMP DEFAULT now(),
	"updated_at" TIMESTAMP,
	"deleted_at" TIMESTAMP
);