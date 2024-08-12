-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE roles (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(256),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

INSERT INTO roles (name) VALUES ('Role Admin');

CREATE TABLE permissions (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(256),
    module          VARCHAR(256),
    access_code     TEXT,
    grant_code      VARCHAR(1),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

ALTER TABLE permissions
    ADD CONSTRAINT uq_accesscode_grantcode UNIQUE (access_code,grant_code);

insert into permissions (name, module, access_code, grant_code) values
    ('Permission Read Employee', 'employee', 'employee', 'r'),
    ('Permission Update Employee', 'employee', 'employee', 'u'),
    ('Permission Delete Employee', 'employee', 'employee', 'd'),
    ('Permission Create Employee', 'employee', 'employee', 'c');

CREATE TABLE role_permissions (
    id              SERIAL PRIMARY KEY,
    role_id         INTEGER NOT NULL,
    permission_id   INTEGER NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(356) DEFAULT 'SYSTEM',
    modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

ALTER TABLE role_permissions
    ADD CONSTRAINT uq_roleid_permissionid UNIQUE (role_id,permission_id);

INSERT INTO role_permissions (role_id, permission_id) VALUES ((select id from roles where name = 'Role Admin'),(select id from permissions where name = 'Permission Create Employee'));

-- +migrate StatementEnd