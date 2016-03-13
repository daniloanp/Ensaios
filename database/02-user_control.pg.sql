\set ON_ERROR_STOP
\encoding utf8

SET search_path TO user_control;

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

CREATE TABLE operation_permission (
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

CREATE TABLE role_permission (
    -- table columns with their constraints
    permission_id BIGINT NOT NULL,
    role_id       BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT role_permission_pk PRIMARY KEY (role_id, permission_id),
    CONSTRAINT permission_fk FOREIGN KEY (permission_id) REFERENCES permission (id),
    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES role (id)
);

CREATE TABLE user_account (
    -- table columns with their constraints
    id              BIGSERIAL,
    name            TEXT                        NOT NULL,
    create_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL,
    -- table constraints
    CONSTRAINT name_unique UNIQUE (name),
    CONSTRAINT user_account_pk PRIMARY KEY (id)
);

CREATE TABLE role_account (
    -- table columns with their constraints
    user_account_id BIGINT NOT NULL,
    role_id         BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT role_account_pk PRIMARY KEY (user_account_id, role_id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES role (id)
);


CREATE TABLE user_session (
    -- table columns with their constraints
    id              TEXT   NOT NULL,
    user_account_id BIGINT NOT NULL,
    role_id         BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT user_session_pk PRIMARY KEY (id),
    CONSTRAINT role_account_fk FOREIGN KEY (user_account_id, role_id) REFERENCES role_account (user_account_id, role_id)
);

CREATE TABLE user_password (
    -- table columns with their constraints
    id              BIGSERIAL,
    user_account_id BIGINT                      NOT NULL,
    password        TEXT                        NOT NULL,
    salt            TEXT                        NOT NULL,
    create_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL,
    -- table constraints
    CONSTRAINT password_pk PRIMARY KEY (id),
    CONSTRAINT password_unique UNIQUE (id, user_account_id), -- TODO:CHECK_BEST_WAY
    CONSTRAINT user_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id)
);

-- ALTER TABLE user_account ADD CONSTRAINT current_password_id FOREIGN KEY (current_password_id) REFERENCES user_password(id);
CREATE TABLE user_current_password (
    -- table columns with their constraints
    user_account_id  BIGINT NOT NULL,
    user_password_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT user_unique UNIQUE (user_account_id),
    CONSTRAINT current_password_pk PRIMARY KEY (user_password_id, user_account_id),
    -- TODO:CHECK_BEST_WAY
    CONSTRAINT current_account_password_fk FOREIGN KEY (user_password_id, user_account_id) REFERENCES user_password (id, user_account_id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id)
);

CREATE TABLE user_email (
    -- table columns with their constraints
    user_account_id   BIGINT                      NOT NULL,
    address           VARCHAR(254)                NOT NULL,
    verified          BOOLEAN                              DEFAULT FALSE,
    register_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT user_email_address_pk PRIMARY KEY (user_account_id, address)
);

CREATE TABLE user_personal_information (
    -- table columns with their constraints
    id                BIGSERIAL                   NOT NULL,
    user_account_id   BIGINT                      NOT NULL,
    given_name        TEXT                                 DEFAULT NULL,
    last_name         TEXT                                 DEFAULT NULL,
    mother_name       TEXT                                 DEFAULT NULL,
    country           TEXT                                 DEFAULT NULL, -- NO REFERENCE TABLE;
    nationality       TEXT                                 DEFAULT NULL, -- NO REFERENCE TABLE;
    register_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    PRIMARY KEY (id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id)
);

CREATE TABLE user_current_emails(
    user_account_id   BIGINT                      NOT NULL,
    user_email_addresses VARCHAR(254),
    CONSTRAINT user_email_addresses_unique UNIQUE (user_email_addresses),
    CONSTRAINT user_email_addresses_fk FOREIGN KEY (user_email_addresses) REFERENCES ()
);

-- table defines wich user_account is the current information;
CREATE TABLE user_current_information (
    -- table columns with their constraints
    user_account_id              BIGINT       NOT NULL,
    user_password_id             BIGINT       NOT NULL,
    user_personal_information_id BIGINT DEFAULT NULL,
    -- table constraints
    CONSTRAINT user_account_unique PRIMARY KEY (user_account_id),
    CONSTRAINT user_email_address_unique UNIQUE(user_personal_information_id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT user_password_fk FOREIGN KEY (user_password_id, user_account_id) REFERENCES user_password (id, user_account_id),
    CONSTRAINT user_email_address_fk FOREIGN KEY (user_account_id, user_email_address) REFERENCES user_email (user_account_id, address)
);












