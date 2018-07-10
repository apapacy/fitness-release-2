BEGIN;

ALTER TABLE registrations DROP COLUMN created_at;

ALTER TABLE registrations DROP COLUMN updated_at;

COMMIT;
