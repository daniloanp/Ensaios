\set ON_ERROR_STOP
\encoding utf8

SET search_path TO user_control;
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

CREATE TABLE operation (
    id     BIGSERIAL,
    name   VARCHAR(255), -- TODO:CHECK_IF_IS_URL_SAFE
    module BIGINT,
    -- table constraints
    UNIQUE (module, name),
    CONSTRAINT operation_pk PRIMARY KEY (id),
    CONSTRAINT module_fk FOREIGN KEY (module) REFERENCES module (id)
);

CREATE TABLE permission (
    id         BIGSERIAL,
    description VARCHAR(255),
    -- table constraints
    CONSTRAINT permission_pk PRIMARY KEY (id)
);

CREATE TABLE operation_permission (
    operation  BIGINT NOT NULL,
    permission BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT operation_permission_pk PRIMARY KEY (operation, permission),
    CONSTRAINT permission_fk FOREIGN KEY (permission) REFERENCES permission (id),
    CONSTRAINT operation_fk FOREIGN KEY (operation) REFERENCES operation (id)
);


CREATE TABLE "role" (
    id         BIGSERIAL,
    description VARCHAR(255), -- just a mneumonic TODO:CHECK IF IS NECESSARY
    parent     BIGINT,
    -- table constraints
    CONSTRAINT role_pk PRIMARY KEY (id),
    CONSTRAINT parent_role_fk FOREIGN KEY (parent) REFERENCES "role" (id)
);


CREATE TABLE role_permission (
    permission BIGINT NOT NULL,
    role       BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT role_permission_pk PRIMARY KEY (role, permission),
    CONSTRAINT permission_fk FOREIGN KEY (permission) REFERENCES permission (id),
    CONSTRAINT role_fk FOREIGN KEY (role) REFERENCES role (id)
);

CREATE TABLE user_account (
    id              BIGSERIAL,
    name            TEXT NOT NULL,
    password        TEXT,
    salt            TEXT,
    create_datetime TEXT,
    -- table constraints
    CONSTRAINT name_unique UNIQUE (name),
    CONSTRAINT user_account_pk PRIMARY KEY (id)
);

CREATE TABLE role_account (
    user_account BIGINT NOT NULL,
    role BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT role_account_pk PRIMARY KEY (user_account, "role"),
    CONSTRAINT user_account_fk FOREIGN KEY(user_account) REFERENCES user_account(id),
    CONSTRAINT role_fk FOREIGN KEY(role) REFERENCES role(id)
);

CREATE TABLE user_session (
    id TEXT NOT NULL,
    user_account BIGINT NOT NULL,
    role BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT user_session_pk PRIMARY KEY (id),
    CONSTRAINT role_account_fk FOREIGN KEY(user_account, role) REFERENCES role_account(user_account, role)
);

