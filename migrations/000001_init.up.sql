CREATE TABLE users
(
    id serial NOT NULL UNIQUE ,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE ,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE wallet
(
    id serial NOT NULL UNIQUE ,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    balance INT
);