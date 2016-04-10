\set ON_ERROR_STOP
\encoding utf8

SET search_path TO users;


CREATE TABLE user_account (
    -- table columns with their constraints
    id                    BIGSERIAL,
    username              TEXT                        NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT name_uq UNIQUE (username),
    CONSTRAINT user_account_pk PRIMARY KEY (id)
);

CREATE TABLE user_password (
    -- table columns with their constraints
    id                    BIGSERIAL,
    user_account_id       BIGINT                      NOT NULL,
    password              TEXT                        NOT NULL,
    salt                  TEXT                        NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT password_pk PRIMARY KEY (id),
    CONSTRAINT password_uq UNIQUE (id, user_account_id), -- TODO:CHECK_BEST_WAY
    CONSTRAINT user_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id)
);

CREATE TABLE user_email (
    -- table columns with their constraints
    user_account_id       BIGINT                      NOT NULL,
    address               VARCHAR(254)                NOT NULL,
    verified              BOOLEAN                              DEFAULT FALSE,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT user_email_address_pk PRIMARY KEY (user_account_id, address)
);

CREATE TABLE user_personal_information (
    -- table columns with their constraints
    id                    BIGSERIAL                   NOT NULL,
    user_account_id       BIGINT                      NOT NULL,
    given_name            TEXT                        NOT NULL,
    last_name             TEXT                        NOT NULL,
    mother_name           TEXT                                 DEFAULT NULL,
    father_name           TEXT                                 DEFAULT NULL,
    nationality           TEXT                                 DEFAULT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    PRIMARY KEY (id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id)
);

CREATE TABLE user_current_emails (
    user_account_id    BIGINT NOT NULL,
    user_email_address VARCHAR(254),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT user_email_addresses_uq UNIQUE (user_email_address),
    CONSTRAINT user_email_addresses_fk FOREIGN KEY (user_email_address, user_account_id) REFERENCES user_email (address, user_account_id)
);

-- table defines wich user_account is the current information;
CREATE TABLE user_current_information (
    -- table columns with their constraints
    user_account_id              BIGINT NOT NULL,
    user_password_id             BIGINT DEFAULT NULL,
    user_personal_information_id BIGINT DEFAULT NULL,
    -- table constraints
    CONSTRAINT user_account_uq PRIMARY KEY (user_account_id),
    CONSTRAINT user_email_address_uq UNIQUE (user_personal_information_id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT user_personal_information_fk FOREIGN KEY (user_personal_information_id) REFERENCES user_personal_information (id),
    CONSTRAINT user_password_fk FOREIGN KEY (user_password_id, user_account_id) REFERENCES user_password (id, user_account_id)
    --     CONSTRAINT user_email_address_fk FOREIGN KEY (user_account_id, user_email_address) REFERENCES user_email (user_account_id, address)
);
