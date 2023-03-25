BEGIN;

CREATE TABLE users  (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email varchar(255) NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NULL,
    password varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NULL
);

COMMIT