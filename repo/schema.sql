CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS credentials (
	id SERIAL PRIMARY KEY,
	user_id SERIAL,
	password VARCHAR(50),
	FOREIGN KEY (user_id) REFERENCES users(id)
	);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) NOT NULL,
	parent_code VARCHAR(4),
	date_of_birth DATE,
	created timestamp DEFAULT NOW()
);

