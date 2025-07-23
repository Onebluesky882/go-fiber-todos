mkdir -p name && cd name && touch name.controller.go name.model.go name.routes.go name.service.go

```go
สร้าง template
mkdir -p name && cd name && touch name.controller.go name.model.go name.routes.go name.service.go


ตัวอย่าง sehema
type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey; not null;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null"`
	NickName  string    `json:"nickname"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}





```
