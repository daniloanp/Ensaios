\set ON_ERROR_STOP
\encoding utf8

SET search_path TO content;

CREATE TABLE public_profile (
    id                      BIGSERIAL                   NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id)
);

CREATE TABLE public_profile_administrator (
    public_profile_id             BIGINT NOT NULL,
    administrator_user_account_id BIGINT NOT NULL,
    -- table constraints
    PRIMARY KEY (public_profile_id, administrator_user_account_id),
    FOREIGN KEY (public_profile_id) REFERENCES public_profile (id),
    FOREIGN KEY (administrator_user_account_id) REFERENCES users.user_account (id)
);

CREATE TABLE public_profile_revision (
    -- table columns with their constraints
    id                      BIGSERIAL,
    public_profile_id       BIGINT                      NOT NULL,
    name                    TEXT                        NOT NULL,
    location                TEXT                                 DEFAULT NULL,
    short_description       VARCHAR(255)                         DEFAULT NULL,
    about                   VARCHAR(2048)               NOT NULL,
    picture_id              BIGINT, -- missing picture image,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    FOREIGN KEY (public_profile_id) REFERENCES public_profile (id),
    PRIMARY KEY (id)
);

CREATE TABLE public_profile_current_revision (
    public_profile_revision_id BIGINT NOT NULL,
    public_profile_id          BIGINT NOT NULL,
    -- table constraints
    UNIQUE (public_profile_id, public_profile_revision_id), -- just one current revision;
    UNIQUE (public_profile_revision_id), -- just one current revision;
    FOREIGN KEY (public_profile_id) REFERENCES public_profile (id),
    FOREIGN KEY (public_profile_revision_id) REFERENCES public_profile_revision (id)
);


CREATE TABLE content_set (
    id                      BIGSERIAL,
    name                    TEXT                        NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    PRIMARY KEY (id)
);

CREATE TABLE content_set_administrator (
    user_account_id BIGINT NOT NULL,
    content_set_id  BIGINT NOT NULL,
    -- table constraints
    PRIMARY KEY (user_account_id, content_set_id),
    FOREIGN KEY (user_account_id) REFERENCES users.user_account (id),
    FOREIGN KEY (content_set_id) REFERENCES content_set (id)
);

CREATE TABLE content (
    id                      BIGSERIAL,
    content_set_id          BIGINT                      NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id),
    FOREIGN KEY (content_set_id) REFERENCES content_set (id),
    PRIMARY KEY (id)
);

CREATE TABLE content_revision (
    id                      BIGSERIAL,
    content_id              BIGINT,
    filename                TEXT                        NOT NULL,
    mimetype                TEXT                        NOT NULL,
    numbytes                BIGINT                      NOT NULL, -- in bytes
    content                 BYTEA                       NOT NULL,
    language_code           TEXT, -- should be a table.
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id),
    FOREIGN KEY (content_id) REFERENCES content (id)
);

CREATE TABLE content_current_revision (
    content_id          BIGINT NOT NULL,
    content_revision_id BIGINT NOT NULL,
    -- table constraints
    UNIQUE (content_id), -- just one content;
    PRIMARY KEY (content_id, content_revision_id),
    FOREIGN KEY (content_revision_id) REFERENCES content_revision (id),
    FOREIGN KEY (content_id) REFERENCES content (id)
);

CREATE TABLE content_revision_author_mapping (-- authors
    content_revision_id BIGINT NOT NULL,
    public_profile_id   BIGINT NOT NULL,
    -- table constraints
    PRIMARY KEY (content_revision_id, public_profile_id),
    FOREIGN KEY (content_revision_id) REFERENCES content_revision (id),
    FOREIGN KEY (public_profile_id) REFERENCES public_profile (id)
);

CREATE TABLE form (
    id    BIGSERIAL,
    title TEXT NOT NULL,
    -- table constraints
    PRIMARY KEY (id)
);

CREATE TYPE form_field_input_type AS ENUM ('combobox', 'radiolist', 'checkbox', 'date', 'datetime', 'textfield', 'textarea');
CREATE TABLE form_field (
    id         BIGSERIAL,
    form_id    BIGINT,
    label      TEXT                  NOT NULL,
    input_type form_field_input_type NOT NULL,
    mask       TEXT    DEFAULT NULL,
    options    TEXT [] DEFAULT NULL, -- should be a
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (form_id) REFERENCES form (id)
);

CREATE TABLE pool (
    id                      BIGSERIAL,
    title                   TEXT                        NOT NULL,
    content_set_id          BIGINT                      NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (content_set_id) REFERENCES content_set (id)
);

CREATE TABLE pool_step (
    id                    BIGSERIAL,
    title                 TEXT   NOT NULL,
    pool_id               BIGINT NOT NULL,
    form_id               BIGINT NOT NULL,
    previous_pool_step_id BIGINT DEFAULT NULL,
    -- table constraints
    CHECK (id != previous_pool_step_id),
    PRIMARY KEY (id),
    UNIQUE (id, pool_id),
    UNIQUE (pool_id, previous_pool_step_id),
    UNIQUE (pool_id, form_id),
    FOREIGN KEY (form_id) REFERENCES form (id),
    FOREIGN KEY (pool_id) REFERENCES pool (id),
    FOREIGN KEY (previous_pool_step_id, pool_id) REFERENCES pool_step (id, pool_id)
);

CREATE TABLE pool_step_record (
    id           BIGSERIAL,
    pool_step_id BIGINT NOT NULL,
    data         JSON DEFAULT NULL,
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (pool_step_id) REFERENCES pool_step (id)
);


CREATE TABLE page_view (
    id                      BIGSERIAL,
    title                   TEXT,
    description             TEXT,
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id)
);


CREATE TABLE page_view_page_view_mapping (-- a page can be composed by others
    parent_view_id BIGINT,
    child_view_id  BIGINT,
    -- table constraints
    CHECK (parent_view_id != child_view_id),
    PRIMARY KEY (parent_view_id, child_view_id),
    FOREIGN KEY (child_view_id) REFERENCES page_view (id),
    FOREIGN KEY (parent_view_id) REFERENCES page_view (id)
);

CREATE TABLE html_template (
    id                      BIGSERIAL,
    template                BYTEA, -- HTML(?)
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id)
);

CREATE TABLE page_view_operation_mapping (
    page_view_id     BIGINT NOT NULL,
    operation_id     BIGINT NOT NULL,
    html_template_id BIGINT DEFAULT NULL,
    -- table constraints
    PRIMARY KEY (page_view_id, operation_id),
    FOREIGN KEY (operation_id) REFERENCES controller.operation (id),
    FOREIGN KEY (page_view_id) REFERENCES page_view (id),
    FOREIGN KEY (html_template_id) REFERENCES html_template(id)
);

CREATE TABLE page_view_content_mapping (
    content_id   BIGINT NOT NULL,
    page_view_id BIGINT NOT NULL,
    name         TEXT   NOT NULL,
    -- table constraints
    PRIMARY KEY (content_id, page_view_id),
    UNIQUE (page_view_id, name),
    FOREIGN KEY (content_id) REFERENCES content (id),
    FOREIGN KEY (page_view_id) REFERENCES page_view (id)
);

CREATE TABLE page_view_pool_mapping (
    pool_id   BIGINT NOT NULL,
    page_view_id BIGINT NOT NULL,
    name         TEXT   NOT NULL,
    -- table constraints
    PRIMARY KEY (pool_id, page_view_id),
    UNIQUE (page_view_id, name),
    FOREIGN KEY (pool_id) REFERENCES pool (id),
    FOREIGN KEY (page_view_id) REFERENCES page_view (id)
);


