BEGIN;

CREATE TABLE "user" (
	"id" CHAR(36) NOT NULL PRIMARY KEY,
	"fname" VARCHAR(55) NOT NULL,
    "lname" VARCHAR(55) NOT NULL,
	"username" VARCHAR NOT NULL, 
	"password" VARCHAR(255) NOT NULL,
	"user_type" VARCHAR NOT NULL,
	"address" VARCHAR(255) NOT NULL,
	"phone" VARCHAR(55) NOT NULL,
	"created_at" TIMESTAMP DEFAULT now(),
	"updated_at" TIMESTAMP,
	"deleted_at" TIMESTAMP
);

COMMIT;