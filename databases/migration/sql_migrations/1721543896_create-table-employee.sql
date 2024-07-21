-- +migrate Up
-- +migrate StatementBegin

create table employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    division VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    modified_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(256),
    modified_by VARCHAR(256)
);

insert into employees (full_name, email, age, division, modified_by, modified_at, created_at, created_by)
VALUES (
        'Arif Yuniarto Fajar Bugaresa', 'arifbugaresa@ymail.com', '12', 'IT', 'admin', '2024-07-21 15:30:00', '2024-07-21 15:30:00', 'admin'
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS employees;

-- +migrate StatementEnd