CREATE TABLE IF NOT EXISTS scenarios (
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    chat_id       BIGINT PRIMARY KEY,
    access_token  VARCHAR(255) UNIQUE,
    refresh_token VARCHAR(255) UNIQUE,
    scenario_id   BIGINT NOT NULL REFERENCES scenarios(id)
);
