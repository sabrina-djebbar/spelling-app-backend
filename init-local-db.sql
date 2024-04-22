-- Create the 'postgres' role
CREATE ROLE postgres WITH SUPERUSER LOGIN PASSWORD 'secret';

-- Create the 'user' database
CREATE USER users WITH ENCRYPTED PASSWORD 'secret' IN ROLE postgres;
CREATE DATABASE users OWNER postgres;

CREATE DATABASE spelling OWNER postgres;
