CREATE TABLE files
(
    id         TEXT PRIMARY KEY,
    name       TEXT   NOT NULL,
    bucket     TEXT   NOT NULL,
    size       BIGINT NOT NULL,
    user_token TEXT,
    url        TEXT   NOT NULL
);
