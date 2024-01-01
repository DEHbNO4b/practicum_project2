CREATE TABLE IF NOT EXISTS users (
    id serial,
    login varchar(150) NOT NULL UNIQUE,
    pass_hash bytea NOT NULL
);
CREATE INDEX if NOT EXISTS idx_login ON users(login);

CREATE TABLE IF NOT EXISTS logpass_data (
    user_id integer NOT NULL,
    login varchar(150) NOT NULL,
    pass varchar(150) NOT NULL,
    meta text
);
CREATE TABLE IF NOT EXISTS text_data (
    user_id integer NOT NULL,
    text text NOT NULL,
    meta text
);
CREATE TABLE IF NOT EXISTS binary_data (
    user_id integer NOT NULL,
    data bytea NOT NULL,
    meta text
);
CREATE TABLE IF NOT EXISTS card_data (
    user_id integer NOT NULL,
    card_id varchar(16) NOT NULL,
    pass varchar(3) NOT NULL,
    date varchar(10) NOT NULL,
    meta text
);
