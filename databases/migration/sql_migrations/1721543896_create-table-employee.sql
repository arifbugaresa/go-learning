-- +migrate Up
-- +migrate StatementBegin

create table employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    division VARCHAR(20) NOT NULL
);

insert into employees (full_name, email, age, division)
VALUES (
        'Arif Yuniarto Fajar Bugaresa', 'arifbugaresa@ymail.com', '12', 'IT'
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS employees;

-- +migrate StatementEnd