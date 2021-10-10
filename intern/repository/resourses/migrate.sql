CREATE TABLE IF NOT EXISTS models
(
    id         serial primary key,
    urn        varchar(255) default '',
    struct     jsonb,
    name       varchar(255),
    is_dataset boolean
);

CREATE TABLE IF NOT EXISTS users
(
    id serial primary key
);

CREATE TABLE IF NOT EXISTS user_access
(
    user_id  varchar(255) references users,
    model_id integer references models,
    unique (user_id, model_id)
);
