CREATE TABLE IF NOT EXISTS service_user
(
    id         SERIAL PRIMARY KEY NOT NULL,
    username   VARCHAR(32) UNIQUE NOT NULL,
    password   VARCHAR(128)       NOT NULL,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS service_admin
(
    id         SERIAL PRIMARY KEY NOT NULL,
    username   VARCHAR(32) UNIQUE NOT NULL,
    password   VARCHAR(128)       NOT NULL,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);


INSERT INTO service_user(username, password)
VALUES ('test', 'test');
INSERT INTO service_admin(username, password)
VALUES ('test', 'test');