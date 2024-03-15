CREATE TABLE "eur_quotes" (
    "id" VARCHAR(15) PRIMARY KEY,
    "usd" FLOAT,
    "mxn" FLOAT,
    "refresh_time" TIMESTAMP
);

CREATE TABLE "usd_quotes" (
    "id" VARCHAR(15) PRIMARY KEY,
    "eur" FLOAT,
    "mxn" FLOAT,
    "refresh_time" TIMESTAMP
);

CREATE TABLE "mxn_quotes" (
    "id" VARCHAR(15) PRIMARY KEY,
    "eur" FLOAT,
    "usd" FLOAT,
    "refresh_time" TIMESTAMP
);
