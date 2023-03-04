package models

type Hero struct {
	tableName   struct{}   `pg:"heroes,alias:h"`
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Level       int        `json:"level"`
	HitPoints   int        `json:"hit_points"`
	ArmorClass  int        `json:"armor_class"`
	Iniative    int        `json:"iniative"`
	Attack      int        `json:"attack"`
	Damage      string     `json:"damage"`
	ClassID     int64      `json:"class_id"`
	Class       *Class     `pg:"rel:has-one" json:"class"`
	AttributeID int64      `json:"attribute_id"`
	Attribute   *Attribute `pg:"rel:has-one" json:"attribute"`
}
