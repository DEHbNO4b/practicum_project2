CREATE TABLE IF NOT EXISTS users (
    id serial,
    login varchar(150) NOT NULL UNIQUE,
    pass_hash bytea NOT NULL
);
CREATE INDEX if NOT EXISTS idx_login ON users(login);

CREATE TABLE IF NOT EXISTS logpassdata (
    user_id integer,
    login varchar(150) NOT NULL,
    pass varchar(150) NOT NULL,
    meta text
);
