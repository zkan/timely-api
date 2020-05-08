CREATE TABLE IF NOT EXISTS "public"."tasks" (
    "id" SERIAL NOT NULL,
    "name" varchar(100),
    "category" varchar(100),
    "author" varchar(100),
    "started_at" timestamp,
    "ended_at" timestamp,
    PRIMARY KEY ("id")
);
