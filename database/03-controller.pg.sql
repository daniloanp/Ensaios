\set ON_ERROR_STOP
\encoding utf8

SET search_path TO controller;

CREATE TABLE module (
    -- table columns with their constraints
    id               BIGSERIAL, --
    name             VARCHAR(255) NOT NULL, -- TODO:CHECK_IF_IS_URL_SAFE
    parent_module_id BIGINT,
    -- table constraints
    CONSTRAINT chk_module_id_uint CHECK (id > 0),
    CONSTRAINT chk_module_name_rule CHECK(name::text ~ '^[a-z0-9][a-z0-9\\-]*$'),
    CONSTRAINT chk_module_no_circular_parent CHECK (parent_module_id != id),
    CONSTRAINT chk_module_needs_parent CHECK(name = '' OR parent_module_id is distinct from null),
    CONSTRAINT uq_parent_module_name UNIQUE (parent_module_id, name),
    CONSTRAINT pk_module PRIMARY KEY (id),
    CONSTRAINT fk_module_parent_module FOREIGN KEY (parent_module_id) REFERENCES module (id)
);

CREATE TABLE operation (
    -- table columns with their constraints
    id        BIGSERIAL,
    name      VARCHAR(255), -- TODO:CHECK_IF_IS_URL_SAFE
    module_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT chk_operation_name_rule CHECK(name::text ~ '^([a-z0-9][a-z0-9\\-]*|)$'),
    CONSTRAINT uq_module_name UNIQUE (module_id, name),
    CONSTRAINT pk_operation PRIMARY KEY (id),
    CONSTRAINT fk_module FOREIGN KEY (module_id) REFERENCES module (id)
);

CREATE TABLE permission (
    -- table columns with their constraints
    id          BIGSERIAL,
    description VARCHAR(255),
    -- table constraints
     PRIMARY KEY (id)
);

CREATE TABLE operation_permission_mapping (
    -- table columns with their constraints
    operation_id  BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    -- table constraints
     PRIMARY KEY (operation_id, permission_id),
     FOREIGN KEY (permission_id) REFERENCES permission (id),
     FOREIGN KEY (operation_id) REFERENCES operation (id)
);


CREATE TABLE "role" (
    -- table columns with their constraints
    id             BIGSERIAL,
    description    VARCHAR(255), -- just a mneumonic TODO:CHECK IF IS NECESSARY
    parent_role_id BIGINT,
    -- table constraints
    CHECK (parent_role_id != id),
    PRIMARY KEY (id),
    FOREIGN KEY (parent_role_id) REFERENCES "role" (id)
);

CREATE TABLE permission_role_mapping (
    -- table columns with their constraints
    permission_id BIGINT NOT NULL,
    role_id       BIGINT NOT NULL,
    -- table constraints
     PRIMARY KEY (role_id, permission_id),
     FOREIGN KEY (permission_id) REFERENCES permission (id),
     FOREIGN KEY (role_id) REFERENCES role (id)
);

CREATE TABLE role_account_mapping (
    -- table columns with their constraints
    user_account_id BIGINT NOT NULL,
    role_id         BIGINT NOT NULL,
    -- table constraints
     PRIMARY KEY (user_account_id, role_id),
     FOREIGN KEY (user_account_id) REFERENCES users.user_account (id),
     FOREIGN KEY (role_id) REFERENCES role (id)
);
