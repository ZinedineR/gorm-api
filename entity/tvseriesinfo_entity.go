package entity

import (
	"github.com/gofrs/uuid"
)

const (
	TVTableName = "tvseries_info"
)

// ArticleModel is a model for entity.Article
type TV struct {
	id       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	title    string    `gorm:"type:varchar;not_null" json:"title"`
	producer string    `gorm:"type:varchar;null" json:"producer"`
}

func NewTV(id uuid.UUID, title, producer string) *TV {
	return &TV{
		id:       id,
		title:    title,
		producer: producer,
	}
}

// TableName specifies table name for ArticleModel.
func (model *TV) TableName() string {
	return TVTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
