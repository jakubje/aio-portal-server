CREATE TABLE "sessions" (
      "id" uuid PRIMARY KEY,
     "email" varchar NOT NULL,
     "account_id" bigint NOT NULL,
     "refresh_token" varchar NOT NULL,
     "user_agent" varchar NOT NULL,
     "client_ip" varchar NOT NULL,
     "is_blocked" boolean NOT NULL DEFAULT false,
     "expires_at" timestamptz NOT NULL
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");