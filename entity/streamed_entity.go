package entity

const (
	StreamedTableName = "streamed"
)

// ArticleModel is a model for entity.Article
type Streamed struct {
	Id       *TV    `gorm:"foreignKey:id" json:"id"`
	Platform string `gorm:"type:varchar;not_null" json:"platform"`
}

func NewStreamed(id *TV, platform string) *Streamed {
	return &Streamed{
		Id:       id,
		Platform: platform,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Streamed) TableName() string {
	return StreamedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
