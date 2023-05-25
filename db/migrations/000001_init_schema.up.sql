CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "watchlists" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "account_id" bigint NOT NULL
);

CREATE TABLE "watchlist_coins" (
  "id" bigserial PRIMARY KEY,
  "watchlist_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "symbol" varchar NOT NULL,
  "rank" smallint NOT NULL
);

CREATE TABLE "portfolios" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "holdings" integer NOT NULL DEFAULT 0,
  "change_24h" integer NOT NULL DEFAULT 0,
  "profit_loss" integer NOT NULL DEFAULT 0
);

CREATE TABLE "coin" (
  "id" bigserial PRIMARY KEY,
  "coin_name" varchar NOT NULL,
  "coin_symbol" varchar NOT NULL,
  "quantity" integer NOT NULL DEFAULT 0,
  "time_created" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" uuid DEFAULT uuid_generate_v4 (),
  "account_id" bigint NOT NULL,
  "portfolio_id" bigint NOT NULL,
  "coin_id" bigint NOT NULL,
  "type" int NOT NULL,
  "symbol" varchar NOT NULL,
  "price_per_coin" float NOT NULL,
  "quantity" float NOT NULL,
  "time_transacted" timestamp NOT NULL,
  "time_created" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "football" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "team" varchar NOT NULL,
  "league" varchar NOT NULL,
  "country" varchar NOT NULL
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "watchlists" ("account_id");

CREATE INDEX ON "watchlist_coins" ("symbol");

CREATE INDEX ON "portfolios" ("account_id");

CREATE INDEX ON "transactions" ("coin_id");

CREATE INDEX ON "transactions" ("type");

ALTER TABLE "watchlists" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "watchlist_coins" ADD FOREIGN KEY ("watchlist_id") REFERENCES "watchlists" ("id");

CREATE TABLE "coin_portfolios" (
  "coin_id" bigserial,
  "portfolios_id" bigserial,
  PRIMARY KEY ("coin_id", "portfolios_id")
);

ALTER TABLE "coin_portfolios" ADD FOREIGN KEY ("coin_id") REFERENCES "coin" ("id");

ALTER TABLE "coin_portfolios" ADD FOREIGN KEY ("portfolios_id") REFERENCES "portfolios" ("id");


ALTER TABLE "portfolios" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("portfolio_id") REFERENCES "portfolios" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("coin_id") REFERENCES "coin" ("id");

ALTER TABLE "football" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");
