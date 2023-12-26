CREATE TABLE IF NOT EXISTS users (
    id serial,
    login varchar(150) NOT NULL UNIQUE,
    pass_hash bytea NOT NULL
);
CREATE INDEX if NOT EXISTS idx_login ON users(login);

CREATE TABLE IF NOT EXISTS logpass_data (
    user_id integer,
    login varchar(150) NOT NULL,
    pass varchar(150) NOT NULL,
    meta text
);
CREATE TABLE IF NOT EXISTS text_data (
    user_id integer,
    text text,
    meta text
);
CREATE TABLE IF NOT EXISTS binary_data (
    user_id integer,
    data bytea ,
    meta text
);
CREATE TABLE IF NOT EXISTS card_data (
    user_id integer,
    card_id varchar(16),
    pass varchar(3),
    date timestamp    ,
    meta text
);
