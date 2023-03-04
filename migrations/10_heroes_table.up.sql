CREATE TABLE IF NOT EXISTS heroes (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    level INTEGER NOT NULL,
    hit_points INTEGER NOT NULL,
    armor_class INTEGER NOT NULL,
    iniative INTEGER NOT NULL,
    attack INTEGER NOT NULL,
    damage VARCHAR NOT NULL,
    class_id INTEGER REFERENCES classes(id),
    attribute_id INTEGER REFERENCES attributes(id)
);