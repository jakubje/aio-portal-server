CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "email" varchar UNIQUE NOT NULL,
                         "name" varchar NOT NULL,
                         "last_name" varchar NOT NULL,
                         "password" varchar NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "watchlists" (
                              "id" BIGSERIAL PRIMARY KEY,
                              "name" varchar NOT NULL,
                              "account_id" bigint NOT NULL
);

CREATE TABLE "coins" (
                         "coin_id" varchar(10) PRIMARY KEY NOT NULL,
                         "name" varchar NOT NULL,
                         "price" varchar NOT NULL,
                         "market_cap" varchar NOT NULL,
                         "circulating_supply" varchar NOT NULL,
                         "total_supply" varchar NOT NULL,
                         "max_supply" varchar NOT NULL,
                         "rank" varchar NOT NULL,
                         "volume" varchar NOT NULL,
                         "image_url" varchar NOT NULL,
                         "description" varchar NOT NULL,
                         "website" varchar NOT NULL,
                         "social_media_links" varchar [],
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "watchlist_coins" (
                                   "watchlist_id" bigint NOT NULL,
                                   "coin_id" varchar NOT NULL
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
                                "id" uuid DEFAULT uuid_generate_v4 (),
                                "account_id" bigint NOT NULL,
                                "portfolio_id" bigint NOT NULL,
                                "type" int NOT NULL,
                                "symbol" varchar NOT NULL,
                                "price_per_coin" float NOT NULL,
                                "quantity" float NOT NULL,
                                "time_transacted" timestamptz NOT NULL,
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

CREATE INDEX ON "portfolios" ("account_id");

CREATE INDEX ON "transactions" ("type");

ALTER TABLE "watchlists" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "watchlist_coins" ADD FOREIGN KEY ("watchlist_id") REFERENCES "watchlists" ("id");

ALTER TABLE "watchlist_coins" ADD FOREIGN KEY ("coin_id") REFERENCES "coins" ("coin_id");

ALTER TABLE "portfolios" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("portfolio_id") REFERENCES "portfolios" ("id");

ALTER TABLE "football" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");