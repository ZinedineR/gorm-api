package entity

const (
	DetailedTableName = "detailed"
)

// ArticleModel is a model for entity.Article
type Detailed struct {
	id       *TV `gorm:"foreignKey:id" json:"id"`
	season   int `gorm:"type:int;not_null" json:"season"`
	episodes int `gorm:"type:int;not_null" json:"episodes"`
	year     int `gorm:"type:int;not_null" json:"year"`
}

func NewDetailed(id *TV, season, episodes, year int) *Detailed {
	return &Detailed{
		id:       id,
		season:   season,
		episodes: episodes,
		year:     year,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Detailed) TableName() string {
	return DetailedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
