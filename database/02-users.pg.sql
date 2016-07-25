\set ON_ERROR_STOP
\encoding utf8

SET search_path TO users;

CREATE TABLE nickname (
    id                    BIGSERIAL                   NOT NULL,
    name                  VARCHAR(255)                NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT pk_nickname PRIMARY KEY (id),
    CONSTRAINT ck_nickname__name__rule CHECK (name :: TEXT ~ '^[a-z][a-z0-9\\-_.]*$' AND length(name) >= 3),
    CONSTRAINT uq_nickname__name UNIQUE (name)
);

CREATE TABLE user_account (
    -- table columns with their constraints
    id                    BIGSERIAL                   NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    --    CONSTRAINT uq_username UNIQUE (username),
    CONSTRAINT pk_user_account PRIMARY KEY (id)
);

CREATE TABLE user_nickname (
    id                    BIGSERIAL                   NOT NULL,
    user_account_id       BIGINT                      NOT NULL,
    nickname_id           BIGINT                      NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- Table constraint
    CONSTRAINT pk_user_nickname PRIMARY KEY (id),
    CONSTRAINT uq_user_nickname__id_user_nickname UNIQUE (id, user_account_id, nickname_id),
    CONSTRAINT fk_user_nickname__user_account_id FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT fk_user_nickname__nickname_id FOREIGN KEY (nickname_id) REFERENCES nickname (id)
);

CREATE TABLE user_password (
    -- table columns with their constraints
    id                    BIGSERIAL                   NOT NULL,
    user_account_id       BIGINT                      NOT NULL,
    password              TEXT                        NOT NULL,
    salt                  TEXT                        NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT pk_user_password PRIMARY KEY (id),
    CONSTRAINT uq_user_password__id_user_account UNIQUE (id, user_account_id),
    CONSTRAINT fk_user_password__user_account FOREIGN KEY (user_account_id) REFERENCES user_account (id)
);

CREATE TABLE email (
    id                    BIGSERIAL                   NOT NULL,
    address               VARCHAR(255)                NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT pk_email PRIMARY KEY (id),
    CONSTRAINT uq_email__address UNIQUE (address)
);

CREATE TABLE user_email (
    -- table columns with their constraints
    user_account_id       BIGINT                      NOT NULL,
    email_id              BIGINT                      NOT NULL,
    verified              BOOLEAN                              DEFAULT FALSE,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    CONSTRAINT pk_user_email PRIMARY KEY (user_account_id, email_id),
    CONSTRAINT fk_user_email__email FOREIGN KEY (email_id) REFERENCES email (id),
    CONSTRAINT fk_user_email__user_account FOREIGN KEY (user_account_id) REFERENCES user_account (id)
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
    CONSTRAINT pk_user_personal_information PRIMARY KEY (id),
    CONSTRAINT uq_personal_information__id_user_account_id UNIQUE (id, user_account_id),
    CONSTRAINT fk_user_personal_information__user_account FOREIGN KEY (user_account_id) REFERENCES user_account (id)
);

CREATE TABLE user_current_email (-- several
    user_account_id BIGINT NOT NULL,
    user_email_id   BIGINT NOT NULL,
    CONSTRAINT uq_user_current_email__user_email_id UNIQUE (user_email_id),
    CONSTRAINT fk_user_current_email__user_account FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT fk_user_current_email__user_email FOREIGN KEY (user_email_id, user_account_id) REFERENCES user_email (email_id, user_account_id)
);

CREATE TABLE user_current_nickname (
    user_nickname_id BIGINT NOT NULL,
    nickname_id      BIGINT NOT NULL,
    user_account_id  BIGINT NOT NULL,
    CONSTRAINT pk_user_current_nickname PRIMARY KEY (user_nickname_id),
    CONSTRAINT uq_user_current_nickname__account_id UNIQUE (user_account_id),
    CONSTRAINT uq_nickname_id UNIQUE (nickname_id),
    CONSTRAINT fk_user_current_nickname__user_nickname FOREIGN KEY (user_nickname_id, nickname_id, user_account_id) REFERENCES user_nickname (id, nickname_id, user_account_id)
);
-- table defines wich user_account is the current information;
CREATE TABLE user_current_information (
    -- table columns with their constraints
    user_account_id              BIGINT NOT NULL,
    user_password_id             BIGINT DEFAULT NULL,
    user_personal_information_id BIGINT DEFAULT NULL, -- table constraints
    CONSTRAINT pk_user_current_information PRIMARY KEY (user_account_id),
    CONSTRAINT uq_user_current_information UNIQUE (user_personal_information_id),
    CONSTRAINT fk_user_current_information__user_account FOREIGN KEY (user_account_id) REFERENCES user_account (id),
    CONSTRAINT fk_user_current_information__user_personal_information FOREIGN KEY (user_personal_information_id, user_account_id) REFERENCES user_personal_information (id, user_account_id),
    CONSTRAINT fk_user_current_information__user_password FOREIGN KEY (user_password_id, user_account_id) REFERENCES user_password (id, user_account_id)
);