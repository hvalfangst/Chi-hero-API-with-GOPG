package models

type Spell struct {
	tableName   struct{} `pg:"spells,alias:s"`
	ID          int64    `json:"id"`
	SpellName   string   `json:"spell_name"`
	SpellLevel  int      `json:"spell_level"`
	SpellDamage string   `json:"spell_damage"`
}
