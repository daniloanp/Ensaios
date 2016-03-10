
SET SEARCH_PATH TO application_structure;

DROP TABLE IF EXISTS application;
CREATE TABLE application
(
    id bigserial,
    name text,
    -- constraints
    CONSTRAINT application_pk PRIMARY KEY (id),
    CONSTRAINT application_name_unique UNIQUE(name)
);

DROP TABLE IF EXISTS operation;
CREATE TABLE operation (
    id bigserial,
    name text,
     -- constraints
    CONSTRAINT operation_pk PRIMARY KEY (id),
    CONSTRAINT operation_name_unique UNIQUE(name)
);