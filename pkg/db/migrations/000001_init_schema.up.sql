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
  "name" varchar NOT NULL,
  "account_id" bigint NOT NULL
);

CREATE TABLE "coins" (
  "id" bigserial PRIMARY KEY,
  "portfolio_id" bigint NOT NULL,
  "coin_name" varchar NOT NULL,
  "coin_symbol" varchar NOT NULL,
  "amount" integer NOT NULL,
  "time_created" timestamptz NOT NULL DEFAULT (now()),
  "no_of_coins" numeric NOT NULL
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "coin_id" bigint NOT NULL,
  "coin_name" varchar NOT NULL,
  "symbol" varchar NOT NULL,
  "type" int NOT NULL,
  "amount" integer NOT NULL,
  "time_transacted" timestamp NOT NULL,
  "time_created" timestamptz NOT NULL DEFAULT (now()),
  "price_purchased_at" numeric NOT NULL,
  "no_of_coins" numeric NOT NULL
);

CREATE TABLE "football" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "team" varchar,
  "league" varchar,
  "country" varchar
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "watchlists" ("account_id");

CREATE INDEX ON "watchlist_coins" ("symbol");

CREATE INDEX ON "portfolios" ("account_id");

CREATE INDEX ON "coins" ("portfolio_id");

CREATE INDEX ON "transactions" ("coin_id");

CREATE INDEX ON "transactions" ("type");

ALTER TABLE "watchlists" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "watchlist_coins" ADD FOREIGN KEY ("watchlist_id") REFERENCES "watchlists" ("id");

ALTER TABLE "portfolios" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "coins" ADD FOREIGN KEY ("portfolio_id") REFERENCES "portfolios" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("coin_id") REFERENCES "coins" ("id");

ALTER TABLE "football" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");
