
SET SEARCH_PATH TO application_structure;

CREATE TABLE application
(
    id bigserial,
    name text,
    -- constraints
    CONSTRAINT application_pk PRIMARY KEY (id),
    CONSTRAINT application_name_unique UNIQUE(name)
);
