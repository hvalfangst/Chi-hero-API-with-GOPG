CREATE TABLE IF NOT EXISTS spells (
id SERIAL PRIMARY KEY,
spell_name VARCHAR NOT NULL,
spell_level INTEGER NOT NULL,
spell_damage VARCHAR NOT NULL,
class_id INT NOT NULL,
FOREIGN KEY (class_id) REFERENCES classes(id)
);