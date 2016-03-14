--the script to remove all tables in the database
/*
DROP TABLE IF EXISTS cats CASCADE;
*/

create table cats
(
	id uuid,

	name character varying(1000) not null,
	gender character varying(1000) not null,

	create_time timestamp with time zone not null default current_timestamp,
	update_time timestamp with time zone not null default current_timestamp,
	CONSTRAINT "cats_pk" PRIMARY KEY (id)
);

