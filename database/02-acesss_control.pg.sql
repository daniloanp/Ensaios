\set ON_ERROR_STOP
\encoding utf8

SET search_path TO access_control;

DROP TABLE IF EXISTS module;

CREATE TABLE module (
    id     BIGSERIAL, --
    name   VARCHAR(255) NOT NULL, -- TODO:CHECK_IF_IS_URL_SAFE
    parent BIGINT,
    -- table constraints
    CHECK (parent != id),
    UNIQUE (parent, name),
    CONSTRAINT module_pk PRIMARY KEY (id),
    CONSTRAINT parent_module_fk FOREIGN KEY (parent) REFERENCES module (id)
);

DROP TABLE IF EXISTS operation;
CREATE TABLE operation (
    id     BIGSERIAL,
    name   VARCHAR(255), -- TODO:CHECK_IF_IS_URL_SAFE
    module BIGINT,
    -- table constraints
    UNIQUE (module, name),
    CONSTRAINT operation_pk PRIMARY KEY (id),
    CONSTRAINT module_fk FOREIGN KEY (module) REFERENCES module (id)
);

DROP TABLE IF EXISTS permission;
CREATE TABLE permission (
    id         BIGSERIAL,
    descriptor VARCHAR(255),
    -- table constraints
    CONSTRAINT permission_pk PRIMARY KEY (id)
);

DROP TABLE IF EXISTS operation_permission;
CREATE TABLE operation_permission (
    operation  BIGINT NOT NULL,
    permission BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT operation_permission_pk PRIMARY KEY (operation, permission),
    CONSTRAINT permission_fk FOREIGN KEY (permission) REFERENCES permission (id),
    CONSTRAINT operation_fk FOREIGN KEY (operation) REFERENCES operation (id)
);

DROP TABLE IF EXISTS "role";
CREATE TABLE "role" (
    id         BIGSERIAL,
    descriptor VARCHAR(255), -- just a mneumonic TODO:CHECK IF IS NECESSARY
    parent     BIGINT,
    -- table constraints
    CONSTRAINT role_pk PRIMARY KEY (id),
    CONSTRAINT parent_role_fk FOREIGN KEY (parent) REFERENCES "role" (id)
);

DROP TABLE IF EXISTS "role_permission";
CREATE TABLE "role_permission" (
    permission BIGINT NOT NULL,
    role       BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT role_permission_pk PRIMARY KEY (role, permission),
    CONSTRAINT permission_fk FOREIGN KEY (permission) REFERENCES permission (id),
    CONSTRAINT role_fk FOREIGN KEY (role) REFERENCES role (id)
);








