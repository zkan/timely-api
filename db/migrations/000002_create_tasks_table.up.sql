CREATE TABLE IF NOT EXISTS "public"."tasks" (
    "id" SERIAL NOT NULL,
    "category_id" int4,
    "name" varchar(100),
    "started_at" timestamp,
    "ended_at" timestamp,
    PRIMARY KEY ("id")
);
