-- Create the 'postgres' role
CREATE ROLE postgres WITH SUPERUSER LOGIN PASSWORD 'secret';

DROP DATABASE IF EXISTS "user";

-- Create the 'user' database
CREATE DATABASE "user" OWNER postgres;
