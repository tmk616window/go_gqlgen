CREATE TABLE "public"."users" ("id" serial NOT NULL, "name" text, "created_at" timestamptz, "updated_at" timestamptz, PRIMARY KEY ("id") , UNIQUE ("id"));
alter table "public"."users" alter column "created_at" set default now();
alter table "public"."users" alter column "updated_at" set default now();


CREATE TABLE "public"."todos" ("id" serial NOT NULL, "title" text, "user_id" integer, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE restrict);




insert into users (id, name) values 
(1, 'user1');


insert into todos (id, title, user_id) values 
(1, 'title1',1);
