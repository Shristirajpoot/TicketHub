CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);
--hi 
CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "organizer_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" text,
  "venue" varchar NOT NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "tickets" (
  "id" bigserial PRIMARY KEY,
  "event_id" bigint NOT NULL,
  "type" varchar NOT NULL,
  "price" numeric NOT NULL,
  "quantity_available" integer NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "ticket_id" bigint NOT NULL,
  "quantity" integer NOT NULL,
  "total_amount" numeric NOT NULL,
  "status" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "audit_logs" (
  "id" bigserial PRIMARY KEY,
  "action" varchar NOT NULL,
  "entity_type" varchar NOT NULL,
  "entity_id" bigint NOT NULL,
  "user_id" bigint,
  "timestamp" timestamptz DEFAULT (now())
);

CREATE UNIQUE INDEX ON "tickets" ("event_id", "type");

COMMENT ON COLUMN "users"."role" IS 'organizer/user';

COMMENT ON COLUMN "tickets"."type" IS 'e.g., General, VIP';

COMMENT ON COLUMN "transactions"."status" IS 'pending/success/failed';

COMMENT ON COLUMN "audit_logs"."action" IS 'create/update/delete';

COMMENT ON COLUMN "audit_logs"."entity_type" IS 'event/ticket/transaction';

ALTER TABLE "events" ADD FOREIGN KEY ("organizer_id") REFERENCES "users" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
