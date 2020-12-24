package team

import "gorm.io/gorm"

type Repository interface {
	Save(team Team) (Team, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Save(team Team) (Team, error) {
	err := r.db.Create(&team).Error

	if err != nil {
		return team, err
	}
	return team, nil
}
