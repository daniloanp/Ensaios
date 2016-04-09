\set ON_ERROR_STOP
\encoding utf8

-- ----------------------------------------------------------------------------
-- Disconnect every meconsultedb's session
-- ----------------------------------------------------------------------------
SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'ensaios'
      AND pid <> pg_backend_pid();


DROP DATABASE IF EXISTS ensaios;
CREATE DATABASE ensaios
WITH OWNER = postgres
    template = template0
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    LC_COLLATE = 'pt_BR.utf8'
    LC_CTYPE = 'pt_BR.utf8'
CONNECTION LIMIT = -1;

\connect ensaios postgres

SET client_min_messages TO WARNING;


CREATE SCHEMA IF NOT EXISTS permissions;
CREATE SCHEMA IF NOT EXISTS users;


CREATE SCHEMA IF NOT EXISTS user_control; -- store permission over
CREATE SCHEMA IF NOT EXISTS application_structure; -- store applications structures
CREATE SCHEMA IF NOT EXISTS application_content;   -- content produced by users




