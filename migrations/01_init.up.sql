CREATE TABLE users (
                       id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       name TEXT NOT NULL,
                       phone TEXT UNIQUE NOT NULL,
                       role INT2 NOT NULL,
                       login TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE markets (
                         id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                         name TEXT NOT NULL,
                         address TEXT NOT NULL,
                         phone TEXT NOT NULL,
                         director_id BIGINT REFERENCES users (id) NOT NULL
);


CREATE TABLE users_discounts (
                                 id BIGINT REFERENCES users (id) PRIMARY KEY,
                                 sale int2 NOT NULL
)