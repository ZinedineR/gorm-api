package entity

import "github.com/google/uuid"

const (
	UserTableName = "user"
)

// ArticleModel is a model for entity.Article
type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Username string    `gorm:"type:varchar;not_null" json:"username"`
	Password string    `gorm:"type:varchar;not_null" json:"password"`
	Admin    bool      `gorm:"type:bool;default:false;not_null" json:"admin"`
}

func NewUser(id uuid.UUID, username, password string, admin bool) *User {
	return &User{
		Id:       id,
		Username: username,
		Password: password,
		Admin:    admin,
	}
}
func NewAdmin(admin bool) *User {
	return &User{
		Admin: admin,
	}
}

// TableName specifies table name for ArticleModel.
func (model *User) TableName() string {
	return UserTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
