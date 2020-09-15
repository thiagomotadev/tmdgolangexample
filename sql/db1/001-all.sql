CREATE DATABASE tmdgolangexample;

\connect tmdgolangexample

CREATE SCHEMA example;

CREATE TABLE example.random_text (
    id serial PRIMARY KEY,
    random_text TEXT, 
    random_text_sum CHAR(64)
);
