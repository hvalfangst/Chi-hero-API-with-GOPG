package models

type Feat struct {
	tableName struct{} `pg:"feats,alias:f"`
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
}
