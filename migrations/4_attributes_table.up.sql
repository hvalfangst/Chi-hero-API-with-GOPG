CREATE TABLE IF NOT EXISTS attributes(
    id SERIAL PRIMARY KEY,
    dexterity INTEGER not null,
    strength INTEGER not null,
    wisdom INTEGER not null,
    intelligence INTEGER not null,
    constitution INTEGER not null,
    charisma INTEGER not null
);