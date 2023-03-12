BEGIN;

CREATE TABLE products (
      id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
      name varchar(255) NOT NULL,
      article character(12) NOT NULL,
      brand varchar(125) NOT NULL,
      country_of_origin varchar(125) NOT NULL,
      description text,
      price numeric(10, 2) NOT NULL,
      rating numeric(2,1) NOT NULL DEFAULT '0',
      vote_count integer NOT NULL DEFAULT '0',
      vote_sum integer NOT NULL DEFAULT '0'
);

COMMIT