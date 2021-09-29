package entity

const (
	StreamedTableName = "streamed"
)

// ArticleModel is a model for entity.Article
type Streamed struct {
	id       *TV    `gorm:"foreignKey:id" json:"id"`
	platform string `gorm:"type:varchar;not_null" json:"platform"`
}

func NewStreamed(id *TV, platform string) *Streamed {
	return &Streamed{
		id:       id,
		platform: platform,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Streamed) TableName() string {
	return StreamedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
