package models

type Class struct {
	tableName struct{} `pg:"classes,alias:c"`
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Spell     []Spell  `pg:"rel:has-many" json:"spell"`
	Feat      *Feat    `pg:"rel:has-one" json:"feat"`
}
