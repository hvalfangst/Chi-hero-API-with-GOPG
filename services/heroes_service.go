package services

import (
	"Chi-hero-API-with-GOPG/models"
	"github.com/go-pg/pg/v10"
)

func GetHeroes(db *pg.DB) ([]*models.Hero, error) {
	heroes := make([]*models.Hero, 0)

	err := db.Model(&heroes).
		Relation("Attribute").
		Relation("Class").
		Relation("Class.Feat").
		Relation("Class.Spell").
		Select()

	return heroes, err
}

func GetHero(db *pg.DB, heroID string) (*models.Hero, error) {
	hero := &models.Hero{}

	err := db.Model(hero).
		Where("h.id = ?", heroID).
		Relation("Attribute").
		Relation("Class").
		Relation("Class.Feat").
		Relation("Class.Spell").
		Select()

	return hero, err
}
