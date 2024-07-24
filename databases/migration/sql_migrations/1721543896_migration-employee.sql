-- +migrate Up
-- +migrate StatementBegin

create table employees (
    id              SERIAL PRIMARY KEY,
    full_name       VARCHAR(50) NOT NULL,
    email           TEXT UNIQUE NOT NULL,
    age             INT NOT NULL,
    division        VARCHAR(20) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS employees;

-- +migrate StatementEnd