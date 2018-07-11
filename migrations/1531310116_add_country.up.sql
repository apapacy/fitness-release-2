BEGIN;

CREATE TABLE countries (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	code bigint NOT NULL,
	a2 text NOT NULL,
	a3 text NOT NULL
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


COMMIT;
