\set ON_ERROR_STOP
\encoding utf8

SET search_path TO content;

CREATE TABLE profile (
    id                      BIGSERIAL                   NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT pk_profile PRIMARY KEY (id),
    CONSTRAINT fk_profile__user_account FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id)
);

CREATE TABLE profile_admin (-- is a N-N table
    profile_id     BIGINT NOT NULL,
    admin_user_account_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT pk_profile_admin PRIMARY KEY (profile_id, admin_user_account_id),
    CONSTRAINT fk_profile_admin__profile FOREIGN KEY (profile_id) REFERENCES profile (id),
    CONSTRAINT fk_profile_admin__user_account FOREIGN KEY (admin_user_account_id) REFERENCES users.user_account (id)
);

CREATE TABLE profile_revision (
    -- table columns with their constraints
    id                      BIGSERIAL,
    profile_id       BIGINT                      NOT NULL,
    name                    TEXT                        NOT NULL,
    location                TEXT                                 DEFAULT NULL,
    short_description       VARCHAR(255)                         DEFAULT NULL,
    about                   VARCHAR(2048)               NOT NULL,
    picture_id              BIGINT, -- missing picture image,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT pk_profile_revision PRIMARY KEY (id),
    CONSTRAINT fk_profile_revision__profile FOREIGN KEY (profile_id) REFERENCES profile (id)
);

CREATE TABLE profile_current_revision (
    profile_revision_id BIGINT NOT NULL,
    profile_id          BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT uq_profile_current_revision__profile_id__profile_revision_id UNIQUE (profile_id, profile_revision_id), -- just one current revision;
    CONSTRAINT uq_profile_current_revision__profile_revision_id UNIQUE (profile_revision_id), -- just one current revision;
    CONSTRAINT fk_profile_current_revision__profile FOREIGN KEY (profile_id) REFERENCES profile (id),
    CONSTRAINT fk_profile_current_revision__profile_revision FOREIGN KEY (profile_revision_id) REFERENCES profile_revision (id)
);


CREATE TABLE content_set (
    id                      BIGSERIAL,
    name                    TEXT                        NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT pk_content_set PRIMARY KEY (id)
);

CREATE TABLE content_set_admin (
    user_account_id BIGINT NOT NULL,
    content_set_id  BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT pk_content_set_admin PRIMARY KEY (user_account_id, content_set_id),
    CONSTRAINT fk_content_set_admin__user_account FOREIGN KEY (user_account_id) REFERENCES users.user_account (id),
    CONSTRAINT fk_content_set_admin__content_set FOREIGN KEY (content_set_id) REFERENCES content_set (id)
);

CREATE TABLE content (
    id                      BIGSERIAL,
    content_set_id          BIGINT                      NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT pk_content PRIMARY KEY (id),
    CONSTRAINT fk_content__user_account FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id),
    CONSTRAINT fk_content__content_set FOREIGN KEY (content_set_id) REFERENCES content_set (id)
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
    CONSTRAINT pk_content_revision PRIMARY KEY (id),
    CONSTRAINT fk_content_revision__user_account FOREIGN KEY (creator_user_account_id) REFERENCES users.user_account (id),
    CONSTRAINT fk_content_revision__content FOREIGN KEY (content_id) REFERENCES content (id)
);

CREATE TABLE content_current_revision (
    content_id          BIGINT NOT NULL,
    content_revision_id BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT pk_content_current_revision PRIMARY KEY (content_id, content_revision_id),
    CONSTRAINT uq_content_current_revision__content_id UNIQUE (content_id), -- just one content;
    CONSTRAINT fk_content_current_revision__content_revision FOREIGN KEY (content_revision_id) REFERENCES content_revision (id),
    CONSTRAINT fk_content_current_revision__content FOREIGN KEY (content_id) REFERENCES content (id)
);

CREATE TABLE content_revision_v_profile (-- authors
    content_revision_id BIGINT NOT NULL,
    profile_id   BIGINT NOT NULL,
    -- table constraints
    CONSTRAINT pk_content_revision_v_profile PRIMARY KEY (content_revision_id, profile_id),
    CONSTRAINT fk_content_revision_v_profile__content_revision FOREIGN KEY (content_revision_id) REFERENCES content_revision (id),
    CONSTRAINT fk_content_revision_v_profile__profile FOREIGN KEY (profile_id) REFERENCES profile (id)
);

CREATE TABLE form (
    id    BIGSERIAL,
    title TEXT NOT NULL,
    -- table constraints
    CONSTRAINT form_pk PRIMARY KEY (id)
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
    CONSTRAINT pk_form_field PRIMARY KEY (id),
    CONSTRAINT fk_form_field__form FOREIGN KEY (form_id) REFERENCES form (id)
);

CREATE TABLE pool (
    id                      BIGSERIAL,
    title                   TEXT                        NOT NULL,
    content_set_id          BIGINT                      NOT NULL,
    registration_datetime   TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT current_timestamp(2),
    creator_user_account_id BIGINT                      NOT NULL, -- user who created the node.
    -- table constraints
    CONSTRAINT pk_pool PRIMARY KEY (id),
    CONSTRAINT fk_pool__content_set FOREIGN KEY (content_set_id) REFERENCES content_set (id)
);

CREATE TABLE pool_step (
    id                    BIGSERIAL,
    title                 TEXT   NOT NULL,
    pool_id               BIGINT NOT NULL,
    form_id               BIGINT NOT NULL,
    previous_pool_step_id BIGINT DEFAULT NULL,
    -- table constraints
    CONSTRAINT pk_pool_step PRIMARY KEY (id),
    CONSTRAINT ck_pool_step_id__neq__previous_pool_step_id CHECK (id != previous_pool_step_id),
    CONSTRAINT uq_pool_step_id__pool_id UNIQUE (id, pool_id),
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
    FOREIGN KEY (html_template_id) REFERENCES html_template (id)
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
    pool_id      BIGINT NOT NULL,
    page_view_id BIGINT NOT NULL,
    name         TEXT   NOT NULL,
    -- table constraints
    PRIMARY KEY (pool_id, page_view_id),
    UNIQUE (page_view_id, name),
    FOREIGN KEY (pool_id) REFERENCES pool (id),
    FOREIGN KEY (page_view_id) REFERENCES page_view (id)
);

CREATE TABLE menu (
    id   BIGSERIAL NOT NULL,
    name TEXT      NOT NULL,
    -- table constraints,
    PRIMARY KEY (id)
);

CREATE TABLE menu_item (
    id           BIGSERIAL,
    menu_id      BIGINT NOT NULL,
    operation_id BIGINT DEFAULT NULL, -- giver us a url.
    label        TEXT   NOT NULL,
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (menu_id) REFERENCES menu (id),
    FOREIGN KEY (operation_id) REFERENCES controller.operation (id)
);

CREATE TABLE menu_item_sub_menu (
    menu_item_id BIGINT NOT NULL,
    menu_id      BIGINT NOT NULL,
    -- table constraints
    PRIMARY KEY (menu_item_id),
    FOREIGN KEY (menu_item_id) REFERENCES menu_item (id),
    FOREIGN KEY (menu_id) REFERENCES menu (id)
);

CREATE TABLE menu_page_view_mapping (
    menu_id      BIGINT NOT NULL,
    page_view_id BIGINT NOT NULL,
    name         TEXT   NOT NULL,
    -- table constraints
    PRIMARY KEY (menu_id, page_view_id),
    UNIQUE (page_view_id, name),
    FOREIGN KEY (menu_id) REFERENCES menu (id),
    FOREIGN KEY (page_view_id) REFERENCES page_view (id)
);

CREATE TABLE comment_section (
    id BIGSERIAL NOT NULL,
    -- table constraints
    PRIMARY KEY (id)
);

CREATE TABLE comment (
    id                       BIGSERIAL NOT NULL,
    author_profile_id BIGINT    NOT NULL,
    comment_section_id       BIGINT DEFAULT NULL,
    parent_comment_id        BIGINT DEFAULT NULL,
    -- table constraints
    PRIMARY KEY (id),
    FOREIGN KEY (author_profile_id) REFERENCES profile (id),
    FOREIGN KEY (comment_section_id) REFERENCES comment_section (id),
    FOREIGN KEY (parent_comment_id) REFERENCES comment_section (id)
);

CREATE TABLE page_view_comment_section_mapping (
    comment_section_id BIGINT DEFAULT NULL,
    page_view_id       BIGINT NOT NULL,
    --table constraints
    PRIMARY KEY (comment_section_id, page_view_id),
    FOREIGN KEY (comment_section_id) REFERENCES comment_section (id),
    FOREIGN KEY (page_view_id) REFERENCES page_view (id)

);