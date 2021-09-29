package entity

const (
	WatchedTableName = "watched"
)

// ArticleModel is a model for entity.Article
type Watched struct {
	id       *Detailed `gorm:"foreignKey:id" json:"id"`
	season   int       `gorm:"type:int;not_null" json:"season"`
	episodes int       `gorm:"type:int;not_null" json:"episodes"`
}

func NewWatched(id *Detailed, season, episodes int) *Watched {
	return &Watched{
		id:       id,
		season:   season,
		episodes: episodes,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Watched) TableName() string {
	return WatchedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
