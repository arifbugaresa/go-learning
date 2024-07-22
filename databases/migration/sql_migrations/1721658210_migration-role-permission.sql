-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE roles (
    id SERIAL       PRIMARY KEY,
    name            VARCHAR(256),
    permission      INTEGER,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

CREATE TABLE permissions (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(256),
    permission      TEXT,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
)

-- +migrate StatementEnd