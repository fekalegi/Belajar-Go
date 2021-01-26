package tweet

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Tweet struct {
	TweetID           uuid.UUID `gorm:"type:uuid;primary_key;not null; json:"tweet_id"`
	Username          string    `sql:"type:json" json:"username"`
	ContentText       string    `sql:"type:json" json:"content_text"`
	ContentAttachment string    `sql:"type:json" json:"content_attachment"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         sql.NullTime `gorm:"index"`
}
