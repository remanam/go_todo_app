CREATE TABLE users (
    id serial UNIQUE,
    firstname varchar(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash varchar(255) NOT NULL

)