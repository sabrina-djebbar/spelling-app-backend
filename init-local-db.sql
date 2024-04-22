-- Create the 'postgres' role
CREATE ROLE postgres WITH SUPERUSER LOGIN PASSWORD 'secret';

-- Create the 'user' database
-- CREATE DATABASE  "user" OWNER postgres;
CREATE USER "user" WITH ENCRYPTED PASSWORD 'secret' IN ROLE postgres;
CREATE DATABASE "user";

CREATE DATABASE spelling OWNER postgres;
