\set ON_ERROR_STOP
\encoding utf8

SET search_path TO content;

CREATE TABLE public_profile (
    id                      BIGSERIAL                   NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT public_profile_pk PRIMARY KEY (id),
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id)
);

CREATE TABLE public_profile_administrator (
    public_profile_id             BIGINT NOT NULL,
    administrator_user_account_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT public_profile_administrator_pk PRIMARY KEY (public_profile_id, administrator_user_account_id),
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile (id),
    CONSTRAINT administrator_user_account_fk FOREIGN KEY (administrator_user_account_id) REFERENCES users.user_account (id)
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
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile (id),
    CONSTRAINT public_profile_account_pk PRIMARY KEY (id)
);

CREATE TABLE public_profile_current_revision (
    public_profile_revision_id BIGINT NOT NULL,
    public_profile_id          BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT public_profile_current_revision_pk UNIQUE (public_profile_id, public_profile_revision_id), -- just one current revision;
    CONSTRAINT public_profile_revision_uq UNIQUE (public_profile_revision_id), -- just one current revision;
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile (id),
    CONSTRAINT public_profile_revision_fk FOREIGN KEY (public_profile_revision_id) REFERENCES public_profile_revision (id)
);

CREATE TABLE article (
    id                      BIGSERIAL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id),
    CONSTRAINT article_pk PRIMARY KEY (id)
);

CREATE TABLE article_revision (
    id                      BIGSERIAL,
    article_id              BIGINT,
    title                   TEXT,
    intro                   TEXT,
    content                 TEXT, --
    html_content            BYTEA, --
    language_code           TEXT, -- should be a table.
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT article_revision_pk PRIMARY KEY (id),
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id),
    CONSTRAINT article_fk FOREIGN KEY (article_id) REFERENCES article (id)
);

CREATE TABLE article_current_revision (
    article_id          BIGINT NOT NULL,
    article_revision_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT article_revision_uq UNIQUE (article_id), -- just one article;
    CONSTRAINT article_current_revision_pk PRIMARY KEY (article_id, article_revision_id),
    CONSTRAINT article_revision_fk FOREIGN KEY (article_revision_id) REFERENCES article_revision (id),
    CONSTRAINT article_fk FOREIGN KEY (article_id) REFERENCES article (id)
);

CREATE TABLE article_revision_author_mapping (-- authors
    article_revision_id BIGINT NOT NULL,
    public_profile_id   BIGINT NOT NULL,
    -- table constraints
    PRIMARY KEY (article_revision_id, public_profile_id),
    CONSTRAINT article_revision_fk FOREIGN KEY (article_revision_id) REFERENCES article_revision (id),
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile (id)
);


CREATE TABLE form (
    id    BIGSERIAL,
    title TEXT NOT NULL,
    -- table constraints
    CONSTRAINT form_pk PRIMARY KEY (id)
);

CREATE TYPE form_field_type AS ENUM ('combobox', 'radiolist', 'checkbox', 'date', 'datetime', 'textfield', 'textarea');
CREATE TABLE form_field (
    id      BIGSERIAL,
    form_id BIGINT,
    label   TEXT            NOT NULL,
    type    form_field_type NOT NULL,
    mask    TEXT    DEFAULT NULL,
    options TEXT [] DEFAULT NULL, -- should be a
    -- table constraints
    CONSTRAINT form_field_pk PRIMARY KEY (id),
    CONSTRAINT form_field_pk FOREIGN KEY (form_id) REFERENCES form (id)
);

CREATE TABLE pool (
    id                      BIGSERIAL,
    title                   TEXT                        NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT pool_pk PRIMARY KEY (id)
);

CREATE TABLE pool_step (
    id                    BIGSERIAL,
    title                 TEXT   NOT NULL,
    pool_id               BIGINT NOT NULL,
    form_id               BIGINT NOT NULL,
    previous_pool_step_id BIGINT DEFAULT NULL,
    -- table constraints
    CHECK (id != previous_pool_step_id),
    UNIQUE (pool_id, previous_pool_step_id),
    UNIQUE (pool_id, form_id),
    CONSTRAINT pool_step_pk PRIMARY KEY (id),
    CONSTRAINT form_fk FOREIGN KEY (form_id) REFERENCES form (id),
    CONSTRAINT pool_fk FOREIGN KEY (pool_id) REFERENCES pool_step (id),
    CONSTRAINT previous_pool_step_fk FOREIGN KEY (previous_pool_step_id, pool_id) REFERENCES pool_step (id, pool_id)
);


CREATE TABLE file_node (
    id                      BIGSERIAL,
    name                    TEXT                        NOT NULL,
    mimetype                TEXT                        NOT NULL,
    size                    BIGINT                      NOT NULL, -- in bytes
    content                 BYTEA                       NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id),
    CONSTRAINT file_node_pk PRIMARY KEY (id)
);


CREATE TABLE file_link (

);




