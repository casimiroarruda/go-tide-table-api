CREATE SCHEMA IF NOT EXISTS tide_tracker;

CREATE EXTENSION IF NOT EXISTS "postgis" SCHEMA "tide_tracker";
CREATE EXTENSION IF NOT EXISTS "pgcrypto" SCHEMA "tide_tracker";

CREATE TABLE IF NOT EXISTS "tide_tracker"."location"(
   "id" uuid DEFAULT gen_random_uuid() NOT NULL,
   "marine_id" character varying(2),
   "name" character varying(100),
   "point" "tide_tracker".geography NOT NULL,
   "mean_sea_level" real,
   "timezone" character varying(30),
   CONSTRAINT "location_pkey" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "location_marine_id" ON "tide_tracker"."location" ("marine_id");
CREATE INDEX "location_point" ON "tide_tracker"."location" ("point");

CREATE TABLE IF NOT EXISTS "tide_tracker"."tide" (
    "location_id" UUID NOT NULL,
    "time" TIMESTAMP WITH TIME ZONE NOT NULL,
    "height" NUMERIC(5,2) NULL,
    "type" VARCHAR(4) NULL,
    PRIMARY KEY (location_id, time),
    FOREIGN KEY (location_id) REFERENCES tide_tracker.location(id)
);