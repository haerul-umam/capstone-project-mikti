CREATE TYPE role AS ENUM (
  'ADMIN',
  'BUYER'
);

CREATE TYPE status AS ENUM (
  'MENUNGGU',
  'DITERIMA',
  'DITOLAK'
);

CREATE TYPE payment AS ENUM (
  'CREDIT',
  'DEBIT',
  'VA',
  'EWALLET'
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "email" varchar(255) UNIQUE,
  "password" varchar(255) NOT NULL,
  "role" role NOT NULL DEFAULT 'BUYER',
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp
);

CREATE TABLE "category_event" (
  "id" serial PRIMARY KEY,
  "name" varchar(255),
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp
);

CREATE TABLE "event" (
  "id" serial PRIMARY KEY,
  "category_id" integer,
  "name" varchar(255) NOT NULL,
  "date" date NOT NULL,
  "price" integer NOT NULL,
  "is_free" bool DEFAULT false,
  "city" varchar(255) NOT NULL,
  "description" text,
  "quota" integer NOT NULL DEFAULT 1,
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "order" (
  "id" uuid PRIMARY KEY,
  "event_id" integer,
  "user_id" uuid,
  "name_event" varchar(255) NOT NULL,
  "date_event" date NOT NULL,
  "price_event" integer NOT NULL,
  "is_free" bool NOT NULL,
  "description" text,
  "city" varchar(255) NOT NULL,
  "quantity" int,
  "payment_method" payment NOT NULL,
  "amount" integer,
  "status" status DEFAULT 'MENUNGGU',
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp
);

ALTER TABLE "event" ADD FOREIGN KEY ("category_id") REFERENCES "category_event" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");