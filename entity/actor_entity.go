package entity

import "gorm.io/gorm"

const (
	ActorTableName = "actor"
)

// ArticleModel is a model for entity.Article
type Actor struct {
	gorm.Model
	Id   TV     `gorm:"foreignKey:Id" json:"id"`
	Name string `gorm:"type:varchar;not_null" json:"name"`
}

func NewActor(id TV, name string) *Actor {
	return &Actor{
		Id:   id,
		Name: name,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Actor) TableName() string {
	return ActorTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
