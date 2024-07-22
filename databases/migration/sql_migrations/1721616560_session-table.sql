-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE auth_user(
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(256) NOT NULL,
    password        VARCHAR(256) NOT NULL,
    full_name       VARCHAR(256),
    email           VARCHAR(256),
    user_status     INTEGER,
    user_role       INTEGER,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

ALTER TABLE auth_user
    ADD CONSTRAINT unique_username UNIQUE (username);

CREATE TABLE session_history(
    user_id         INTEGER NOT NULL,
    token           TEXT,
    session_data    TEXT,
    created_at      TIMESTAMP DEFAULT NOW(),
    expired_at      TIMESTAMP
);

-- +migrate StatementEnd