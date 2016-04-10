\set ON_ERROR_STOP
\encoding utf8

SET search_path TO permissions;

CREATE TABLE module (
    -- table columns with their constraints
    id               BIGSERIAL, --
    name             VARCHAR(255) NOT NULL, -- TODO:CHECK_IF_IS_URL_SAFE
    parent_module_id BIGINT,
    -- table constraints
    CHECK (parent_module_id != id),
    UNIQUE (parent_module_id, name),
    CONSTRAINT module_pk PRIMARY KEY (id),
    CONSTRAINT parent_module_fk FOREIGN KEY (parent_module_id) REFERENCES module (id)
);

CREATE TABLE operation (
    -- table columns with their constraints
    id        BIGSERIAL,
    name      VARCHAR(255), -- TODO:CHECK_IF_IS_URL_SAFE
    module_id BIGINT,
    -- table constraints
    UNIQUE (module_id, name),
    CONSTRAINT operation_pk PRIMARY KEY (id),
    CONSTRAINT module_fk FOREIGN KEY (module_id) REFERENCES module (id)
);

CREATE TABLE permission (
    -- table columns with their constraints
    id          BIGSERIAL,
    description VARCHAR(255),
    -- table constraints
    CONSTRAINT permission_pk PRIMARY KEY (id)
);

CREATE TABLE operation_permission_mapping (
    -- table columns with their constraints
    operation_id  BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT operation_permission_pk PRIMARY KEY (operation_id, permission_id),
    CONSTRAINT permission_fk FOREIGN KEY (permission_id) REFERENCES permission (id),
    CONSTRAINT operation_fk FOREIGN KEY (operation_id) REFERENCES operation (id)
);


CREATE TABLE "role" (
    -- table columns with their constraints
    id             BIGSERIAL,
    description    VARCHAR(255), -- just a mneumonic TODO:CHECK IF IS NECESSARY
    parent_role_id BIGINT,
    -- table constraints
    CHECK (parent_role_id != id),
    CONSTRAINT role_pk PRIMARY KEY (id),
    CONSTRAINT parent_role_fk FOREIGN KEY (parent_role_id) REFERENCES "role" (id)
);

CREATE TABLE role_permission_mapping (
    -- table columns with their constraints
    permission_id BIGINT NOT NULL,
    role_id       BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT role_permission_pk PRIMARY KEY (role_id, permission_id),
    CONSTRAINT permission_fk FOREIGN KEY (permission_id) REFERENCES permission (id),
    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES role (id)
);


CREATE TABLE role_account_mapping (
    -- table columns with their constraints
    user_account_id BIGINT NOT NULL,
    role_id         BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT role_account_pk PRIMARY KEY (user_account_id, role_id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES users.user_account (id),
    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES role (id)
);

