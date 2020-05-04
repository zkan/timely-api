CREATE TABLE IF NOT EXISTS "public"."tasks" (
    "id" SERIAL NOT NULL,
    "name" varchar(100),
    "category" varchar(100),
    "start" timestamp,
    "end" timestamp,
    PRIMARY KEY ("id")
);