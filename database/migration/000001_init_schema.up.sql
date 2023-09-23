CREATE TABLE "campaigns" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "Content" varchar NOT NULL,
  "Contacts" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "contacts" (
  "id" bigserial PRIMARY KEY,
  "email" bigint NOT NULL,
);

ALTER TABLE "contacts" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns" ("id");