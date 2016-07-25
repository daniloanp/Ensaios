\set ON_ERROR_STOP
\encoding utf8

SET search_path TO controller;

CREATE TABLE module (
    -- table columns with their constraints
    id               BIGSERIAL, --
    name             VARCHAR(255) NOT NULL, -- TODO:CHECK_IF_IS_URL_SAFE
    parent_module_id BIGINT,
    -- table constraints
    CONSTRAINT pk_module PRIMARY KEY (id),
    CONSTRAINT ck_module__name CHECK (name :: TEXT ~ '^[a-z0-9][a-z0-9\\-]*$'),
    CONSTRAINT ck_module__no_circular_parent CHECK (parent_module_id != id),
    CONSTRAINT ck_module__needs_parent CHECK (name = '' OR parent_module_id IS DISTINCT FROM NULL),
    CONSTRAINT uq_module__parent_module_name UNIQUE (parent_module_id, name),
    CONSTRAINT fk_module__module FOREIGN KEY (parent_module_id) REFERENCES module (id)
);

CREATE TABLE operation (
    -- table columns with their constraints
    id        BIGSERIAL,
    name      VARCHAR(255), -- TODO:CHECK_IF_IS_URL_SAFE
    module_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT pk_operation PRIMARY KEY (id),
    CONSTRAINT ck_operation__name CHECK(name::text ~ '^([a-z0-9][a-z0-9\\-]*|)$'),
    CONSTRAINT uq_operation__module_id__name UNIQUE (module_id, name),
    CONSTRAINT fk_operation__module FOREIGN KEY (module_id) REFERENCES module (id)
);

CREATE TABLE permission (
    -- table columns with their constraints
    id          BIGSERIAL,
    description VARCHAR(255),
    -- table constraints
    CONSTRAINT pk_permission PRIMARY KEY (id)
);

CREATE TABLE operation_v_permission (
    -- table columns with their constraints
    operation_id  BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT pk_operation_v_permission PRIMARY KEY (operation_id, permission_id),
    CONSTRAINT fk_operation_v_permission__permission FOREIGN KEY (permission_id) REFERENCES permission (id),
    CONSTRAINT fk_operation_v_permission__operation FOREIGN KEY (operation_id) REFERENCES operation (id)
);


CREATE TABLE "role" (
    -- table columns with their constraints
    id             BIGSERIAL,
    description    VARCHAR(255), -- just a mneumonic TODO:CHECK IF IS NECESSARY
    parent_role_id BIGINT,
    -- table constraints
    CONSTRAINT pk_role PRIMARY KEY (id),
    CONSTRAINT ck_role__parent__neq__parent_role_id CHECK (parent_role_id != id),
    CONSTRAINT fk_role__role FOREIGN KEY (parent_role_id) REFERENCES "role" (id)
);

CREATE TABLE permission_v_role (
    -- table columns with their constraints
    permission_id BIGINT NOT NULL,
    role_id       BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT pk_permission_v_role PRIMARY KEY (role_id, permission_id),
    CONSTRAINT fk_permission_v_role__permission FOREIGN KEY (permission_id) REFERENCES permission (id),
    CONSTRAINT fk_permission_v_role__role FOREIGN KEY (role_id) REFERENCES role (id)
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
