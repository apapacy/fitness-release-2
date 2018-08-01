BEGIN;

CREATE TABLE country (
	id uuid PRIMARY KEY,
	code bigint NOT NULL UNIQUE,
	a2 text NOT NULL UNIQUE,
	a3 text NOT NULL UNIQUE,
	capital_id uuid NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL
);

CREATE TABLE country_translations (
	id uuid REFERENCES country ON DELETE CASCADE,
	locale text,
	name text NOT NULL,
	fullname text NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	PRIMARY KEY (id, locale)
);

CREATE TABLE city (
	id uuid PRIMARY KEY,
	country_id uuid NOT NULL REFERENCES country ON DELETE RESTRICT,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL
);

CREATE TABLE city_translations (
	id uuid REFERENCES city ON DELETE CASCADE,
	locale text,
	name text NOT NULL,
	fullname text NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	PRIMARY KEY (id, locale)
);

ALTER TABLE country 
	ADD CONSTRAINT capital_id_fk FOREIGN KEY (capital_id) REFERENCES city ON DELETE SET NULL;

CREATE TABLE firma (
	id uuid PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL
);

CREATE TABLE firma_translations (
	id uuid REFERENCES firma ON DELETE CASCADE,
	locale text,
	name text NOT NULL,
	fullname text NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	PRIMARY KEY (id, locale)
);

CREATE TABLE club (
	id uuid PRIMARY KEY,
	firma_id uuid NOT NULL REFERENCES firma ON DELETE RESTRICT,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL
);

CREATE TABLE club_translations (
	id uuid REFERENCES firma(id) ON DELETE CASCADE,
	locale text,
	name text NOT NULL,
	fullname text NOT NULL,
	address text NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	PRIMARY KEY (id, locale)
);

COMMIT;
