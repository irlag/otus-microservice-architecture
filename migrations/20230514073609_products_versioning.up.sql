CREATE EXTENSION IF NOT EXISTS pgcrypto;

ALTER TABLE products
    ADD COLUMN version uuid NOT NULL DEFAULT gen_random_uuid()
;