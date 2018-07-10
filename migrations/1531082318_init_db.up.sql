BEGIN;


CREATE TABLE registrations (
	id uuid NOT NULL PRIMARY KEY,
	username text NOT NULL,
	email text NOT NULL,
	password text NOT NULL
);


CREATE TABLE users (
	id uuid NOT NULL PRIMARY KEY,
	username text NOT NULL,
	email text NOT NULL,
	password text NOT NULL
);


COMMIT;
