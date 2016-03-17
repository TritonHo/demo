--the script to remove all tables in the database
/*
DROP TABLE IF EXISTS cats CASCADE;
DROP TABLE IF EXISTS users CASCADE;
*/

create table cats
(
	id uuid,
	user_id uuid not null,

	name character varying(1000) not null,
	gender character varying(1000) not null,

	create_time timestamp with time zone not null default current_timestamp,
	update_time timestamp with time zone not null default current_timestamp,
	CONSTRAINT "cats_pk" PRIMARY KEY (id)
);

create table users
(
	id uuid,

	email character varying(200) not null,
	password_digest character varying(1000) not null,

	first_name character varying(255) null,
	last_name character varying(255) null,

	create_time timestamp with time zone not null default current_timestamp,
	update_time timestamp with time zone not null default current_timestamp,
	CONSTRAINT "users_pk" PRIMARY KEY (id)
);
ALTER TABLE users ADD CONSTRAINT users_u1 UNIQUE (email);

