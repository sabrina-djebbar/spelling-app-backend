CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id text PRIMARY KEY,
    username text NOT NULL,
    parent_code text NOT NULL,
    date_of_birth DATE,
    created timestamp DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS credentials (
	id text PRIMARY KEY,
	user_id text NOT NULL,
	password text NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE users
    OWNER TO postgres;


ALTER TABLE credentials OWNER TO postgres;


