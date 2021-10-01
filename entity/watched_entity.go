package entity

import "gorm.io/gorm"

const (
	WatchedTableName = "watched"
)

// ArticleModel is a model for entity.Article
type Watched struct {
	gorm.Model
	Id       Detailed `gorm:"foreignKey:Id" json:"id"`
	Season   int      `gorm:"type:int;not_null" json:"season"`
	Episodes int      `gorm:"type:int;not_null" json:"episodes"`
}

func NewWatched(id Detailed, season, episodes int) *Watched {
	return &Watched{
		Id:       id,
		Season:   season,
		Episodes: episodes,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Watched) TableName() string {
	return WatchedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
