CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS credentials (
	id text PRIMARY KEY,
	user_id text NOT NULL,
	password text NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
	);

CREATE TABLE IF NOT EXISTS users (
	id text PRIMARY KEY,
	username text NOT NULL,
	parent_code text NOT NULL,
	date_of_birth DATE,
	created timestamp DEFAULT NOW()
);

