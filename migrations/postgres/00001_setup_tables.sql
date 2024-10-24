-- +goose Up
SELECT
    'up SQL query';

CREATE TABLE wizards (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    specialization TEXT
);

CREATE TABLE wizard_stats (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    wizard_id INT REFERENCES wizards(id) UNIQUE,
    power INT NOT NULL DEFAULT 0,
    mana INT NOT NULL DEFAULT 0,
    intelligence INT NOT NULL DEFAULT 0,
    luck INT NOT NULL DEFAULT 0
);

-- +goose Down
SELECT
    'down SQL query';

DROP TABLE wizard_stats;

DROP TABLE wizards;
