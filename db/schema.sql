DROP DATABASE IF EXISTS lbc_db;

CREATE DATABASE lbc_db;

\c lbc_db

CREATE TABLE request (
    id  BIGSERIAL PRIMARY KEY NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE  DEFAULT NOW() NOT NULL,
    last_update TIMESTAMP WITH TIME ZONE  DEFAULT NOW() NOT NULL,
    int1   integer   DEFAULT 0 NOT NULL,
    int2   integer   DEFAULT 0 NOT NULL,
    limite   integer   DEFAULT 0 NOT NULL,
    str1   VARCHAR(50)   DEFAULT '' NOT NULL,
    str2   VARCHAR(50)   DEFAULT '' NOT NULL
);