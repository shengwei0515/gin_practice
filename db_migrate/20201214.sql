CREATE TABLE "users" (
   id bigserial primary key,
   "username" varchar(255) default NULL,
   "password" varchar(255) default NULL
);