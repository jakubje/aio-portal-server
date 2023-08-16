-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-08-09T22:32:38.680Z

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "is_email_verified" bool NOT NULL DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT ('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "userID" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE TABLE "watchlists" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "account_id" bigint NOT NULL
);

CREATE TABLE "watchlist_coins" (
  "id" BIGSERIAL PRIMARY KEY,
  "watchlist_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "symbol" varchar NOT NULL,
  "rank" smallint NOT NULL
);

CREATE TABLE "portfolios" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "holdings" integer NOT NULL DEFAULT 0,
  "change_24h" integer NOT NULL DEFAULT 0,
  "profit_loss" integer NOT NULL DEFAULT 0
);

CREATE TABLE "transactions" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "account_id" bigint NOT NULL,
  "portfolio_id" bigint NOT NULL,
  "type" int NOT NULL,
  "symbol" varchar NOT NULL,
  "price_per_coin" float NOT NULL,
  "quantity" float NOT NULL,
  "time_transacted" timestamp NOT NULL,
  "time_created" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "football" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "team" varchar NOT NULL,
  "league" varchar NOT NULL,
  "country" varchar NOT NULL
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "watchlists" ("account_id");

CREATE INDEX ON "watchlist_coins" ("symbol");

CREATE INDEX ON "portfolios" ("account_id");

CREATE INDEX ON "transactions" ("type");

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("userID") REFERENCES "users" ("id");

ALTER TABLE "watchlists" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "watchlist_coins" ADD FOREIGN KEY ("watchlist_id") REFERENCES "watchlists" ("id");

ALTER TABLE "portfolios" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("portfolio_id") REFERENCES "portfolios" ("id");

ALTER TABLE "football" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");
