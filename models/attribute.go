package models

type Attribute struct {
	tableName    struct{} `pg:"attributes,alias:a"`
	ID           int64    `json:"id"`
	Dexterity    int      `json:"dexterity"`
	Strength     int      `json:"strength"`
	Wisdom       int      `json:"wisdom"`
	Intelligence int      `json:"intelligence"`
	Constitution int      `json:"constitution"`
	Charisma     int      `json:"charisma"`
}
