
SET SEARCH_PATH TO session_management;


CREATE TABLE user_session (
    -- table columns with their constraints
    id              TEXT   NOT NULL,
    user_account_id BIGINT DEFAULT NULL,
    role_id         BIGINT NOT NULL,
    -- missing browser and referer.
    -- table constraints
    CONSTRAINT user_session_pk PRIMARY KEY (id),
    CONSTRAINT user_account_fk FOREIGN KEY (user_account_id) REFERENCES users.user_account(id),
    CONSTRAINT role_account_fk FOREIGN KEY (user_account_id, role_id) REFERENCES permissions.role_account_mapping (user_account_id, role_id)
);
