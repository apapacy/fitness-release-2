BEGIN;

CREATE TABLE countries (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	code bigint NOT NULL UNIQUE,
	a2 text NOT NULL UNIQUE,
	a3 text NOT NULL UNIQUE
);


CREATE TABLE cities (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	country_id uuid NOT NULL REFERENCES countries(id)
);


CREATE TABLE cities_translations (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	locale text NOT NULL,
	name text NOT NULL,
	fullname text NOT NULL,
	city_id uuid NOT NULL REFERENCES cities(id)
);


CREATE TABLE countries_translations (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	locale text NOT NULL,
	name text NOT NULL,
	fullname text NOT NULL,
	country_id uuid NOT NULL REFERENCES countries(id)
);


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
