CREATE TYPE role AS ENUM (
  'BUYER',
  'SELLER'
);

CREATE TABLE "user" (
  "user_id" serial PRIMARY KEY,
  "name" varchar(255) not null,
  "password" varchar(255) not null,
  "email" varchar(255) UNIQUE,
  "role" role not null,
  "created_at" timestamp default current_timestamp,
  "updated_at" timestamp default current_timestamp
);

CREATE TABLE "item" (
  "item_id" serial PRIMARY KEY,
  "seller_id" integer not null,
  "name" varchar(255) not null,
  "quantity" integer not null default 0,
  "price" integer not null,
  "description" text,
  "created_at" timestamp default current_timestamp,
  "updated_at" timestamp default current_timestamp
);

CREATE TABLE "order" (
  "order_id" serial PRIMARY KEY,
  "buyer_id" integer not null,
  "seller_id" integer not null,
  "sub_total" integer not null,
  "order_date" date not null,
  "created_at" timestamp default current_timestamp,
  "updated_at" timestamp default current_timestamp
);

CREATE TABLE "order_detail" (
  "detail_id" serial PRIMARY KEY,
  "order_id" integer not null,
  "item_id" integer not null,
  "quantity" integer not null,
  "item_name" varchar(255) not null,
  "price" integer not null,
  "sub_total" integer not null,
  "created_at" timestamp default current_timestamp,
  "updated_at" timestamp default current_timestamp
);

ALTER TABLE "item" ADD FOREIGN KEY ("seller_id") REFERENCES "user" ("user_id");

ALTER TABLE "order" ADD FOREIGN KEY ("buyer_id") REFERENCES "user" ("user_id");

ALTER TABLE "order" ADD FOREIGN KEY ("seller_id") REFERENCES "user" ("user_id");

ALTER TABLE "order_detail" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("order_id");

ALTER TABLE "order_detail" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("item_id");