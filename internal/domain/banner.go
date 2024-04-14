package domain

import (
	"time"
)

type Banner struct {
	Id        int       `db:"id" json:"banner_id"`
	Content   string    `db:"content" json:"content"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	TagsIds   []int `db:"tag_ids" json:"tag_ids"`
	FeatureId int   `db:"feature_id" json:"feature_id"`
}
