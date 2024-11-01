-- +goose Up
SELECT
    'up SQL query';

CREATE TABLE wizards (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    specialization TEXT NOT NULL DEFAULT ''
);

CREATE INDEX ON idx_wizards_id ON wizards(id)

CREATE TABLE wizard_stats (
    wizard_id INT REFERENCES wizards(id) ON DELETE CASCADE PRIMARY KEY,
    power INT NOT NULL DEFAULT 0,
    mana INT NOT NULL DEFAULT 0,
    intelligence INT NOT NULL DEFAULT 0,
    luck INT NOT NULL DEFAULT 0
);

CREATE INDEX ON idx_wizard_stats_wizard_id ON wizard_stats(wizard_id)

-- +goose Down
SELECT
    'down SQL query';

DROP TABLE wizard_stats;

DROP TABLE wizards;
