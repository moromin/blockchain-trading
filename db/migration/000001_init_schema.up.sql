CREATE TABLE "ohlcs" (
  "id" bigserial PRIMARY KEY,
  "opentime" timestamptz NOT NULL,
  "open" numeric NOT NULL,
  "high" numeric NOT NULL,
  "low" numeric NOT NULL,
  "close" numeric NOT NULL,
  "volume" numeric NOT NULL,
  "closetime" timestamptz NOT NULL,
  "quote_asset_volume" numeric NOT NULL,
  "numer_of_trades" bigint NOT NULL,
  "taker_buy_base_asset_volume" numeric NOT NULL,
  "taker_buy_quote_asset_volume" numeric NOT NULL,
  "symbol_id" bigint NOT NULL,
  "interval_id" bigint NOT NULL
);

CREATE TABLE "symbols" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "base_id" bigint NOT NULL,
  "quote_id" bigint NOT NULL
);

CREATE TABLE "currencies" (
  "id" bigserial PRIMARY KEY,
  "coin" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "intervals" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

ALTER TABLE "ohlcs" ADD FOREIGN KEY ("symbol_id") REFERENCES "symbols" ("id");

ALTER TABLE "ohlcs" ADD FOREIGN KEY ("interval_id") REFERENCES "intervals" ("id");

ALTER TABLE "symbols" ADD FOREIGN KEY ("base_id") REFERENCES "currencies" ("id");

ALTER TABLE "symbols" ADD FOREIGN KEY ("quote_id") REFERENCES "currencies" ("id");
