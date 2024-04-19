CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS credentials (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id uuid NOT NULL,
	password VARCHAR(50) NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
	);

CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(50) NOT NULL,
	parent_code VARCHAR(20) NOT NULL,
	date_of_birth DATE,
	created timestamp DEFAULT NOW()
);

