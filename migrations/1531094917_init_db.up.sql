BEGIN;

ALTER TABLE registrations ADD COLUMN created_at timestamptz NOT NULL;

ALTER TABLE registrations ADD COLUMN updated_at timestamptz NOT NULL;

COMMIT;
