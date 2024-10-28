-- +goose Up
SELECT
    'up SQL query';

CREATE TABLE wizards (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    specialization TEXT NOT NULL DEFAULT ''
);

CREATE TABLE wizard_stats (
    wizard_id INT REFERENCES wizards(id) PRIMARY KEY,
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
