package entity

import "gorm.io/gorm"

const (
	DetailedTableName = "detailed"
)

// ArticleModel is a model for entity.Article
type Detailed struct {
	gorm.Model
	Id       TV  `gorm:"foreignKey:Id" json:"id"`
	Season   int `gorm:"type:int;not_null" json:"season"`
	Episodes int `gorm:"type:int;not_null" json:"episodes"`
	Year     int `gorm:"type:int;not_null" json:"year"`
}

func NewDetailed(id TV, season, episodes, year int) *Detailed {
	return &Detailed{
		Id:       id,
		Season:   season,
		Episodes: episodes,
		Year:     year,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Detailed) TableName() string {
	return DetailedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
