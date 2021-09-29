package entity

const (
	ActorTableName = "actor"
)

// ArticleModel is a model for entity.Article
type Actor struct {
	id   *TV    `gorm:"foreignKey:id" json:"id"`
	name string `gorm:"type:varchar;not_null" json:"name"`
}

func NewActor(id *TV, name string) *Actor {
	return &Actor{
		id:   id,
		name: name,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Actor) TableName() string {
	return ActorTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
