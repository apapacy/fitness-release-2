BEGIN;

CREATE TABLE registrations (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
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
