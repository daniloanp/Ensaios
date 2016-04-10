\set ON_ERROR_STOP
\encoding utf8

SET search_path TO content;

CREATE TABLE public_profile (
    id                    BIGSERIAL NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT public_profile_pk PRIMARY KEY (id),
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account(id)
);

CREATE TABLE public_profile_administrator (
    public_profile_id BIGINT NOT NULL,
    administrator_user_account_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT public_profile_administrator_pk PRIMARY KEY (public_profile_id, administrator_user_account_id),
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile(id),
    CONSTRAINT administrator_user_account_fk FOREIGN KEY (administrator_user_account_id) REFERENCES users.user_account(id)
);

CREATE TABLE public_profile_revision (
    -- table columns with their constraints
    id                    BIGSERIAL,
    public_profile_id     BIGINT                      NOT NULL,
    name                  TEXT                        NOT NULL,
    location              TEXT                                 DEFAULT NULL,
    short_description     VARCHAR(255)                         DEFAULT NULL,
    about                 VARCHAR(2048)               NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    picture_id            BIGINT, -- missing picture image,
    -- table constraints
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile(id),
    CONSTRAINT public_profile_account_pk PRIMARY KEY (id)
);

CREATE TABLE public_profile_current_revision(
    public_profile_revision_id BIGINT NOT NULL,
    public_profile_id          BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT public_profile_current_revision_pk UNIQUE (public_profile_id, public_profile_revision_id), -- just one current revision;
    CONSTRAINT public_profile_revision_uq UNIQUE (public_profile_revision_id), -- just one current revision;
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile(id),
    CONSTRAINT public_profile_revision_fk FOREIGN KEY (public_profile_revision_id) REFERENCES public_profile_revision(id)
);

CREATE TABLE file_node (
    id                    BIGSERIAL,
    name                  TEXT                        NOT NULL,
    mimetype              TEXT                        NOT NULL,
    size                  BIGINT                      NOT NULL, -- in bytes
    content               BYTEA                       NOT NULL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account(id),
    CONSTRAINT file_node_pk PRIMARY KEY (id)
);

CREATE TABLE article (
    id                    BIGSERIAL,
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account(id),
    CONSTRAINT article_pk PRIMARY KEY (id)
);

CREATE TABLE article_revision (
    id                    BIGSERIAL,
    article_id            BIGINT,
    title                 TEXT,
    content               TEXT, --
    html_content          BYTEA, --
    language_code         TEXT, -- should be a table.
    registration_datetime TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT creator_user_account_fk FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account(id),
    CONSTRAINT article_revision_pk PRIMARY KEY (id),
    CONSTRAINT article_fk FOREIGN KEY (article_id) REFERENCES article (id)
);

CREATE TABLE article_revision_author_mapping ( -- authors
    article_revision_id BIGINT NOT NULL,
    public_profile_id BIGINT NOT NULL,
    --
    PRIMARY KEY (article_revision_id, public_profile_id),
    CONSTRAINT article_revision_fk FOREIGN KEY (article_revision_id) REFERENCES article_revision(id),
    CONSTRAINT public_profile_fk FOREIGN KEY (public_profile_id) REFERENCES public_profile(id)
);
