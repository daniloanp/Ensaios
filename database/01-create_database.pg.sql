\set ON_ERROR_STOP
\encoding utf8

DROP DATABASE IF EXISTS ensaios;
CREATE DATABASE ensaios
WITH OWNER = postgres
    TEMPLATE template0
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    LC_COLLATE = 'pt_BR.UTF-8'
    LC_CTYPE = 'pt_BR.UTF-8'
CONNECTION LIMIT = 63;

\connect ensaios postgres
SET client_min_messages TO WARNING;



CREATE SCHEMA IF NOT EXISTS access_control; -- store permission over
SET SEARCH_PATH TO access_control;

CREATE SCHEMA IF NOT EXISTS user_data; -- users information and login
SET SEARCH_PATH TO user_data;

CREATE SCHEMA IF NOT EXISTS application_content;   -- content produced by users
SET SEARCH_PATH TO application_content;




