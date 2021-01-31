CREATE TABLE "Accounts" (
   id bigserial ,
   "username" varchar(255) default NULL,
   "password" varchar(255) default NULL,
	PRIMARY KEY (id, username),
   UNIQUE (username)
);