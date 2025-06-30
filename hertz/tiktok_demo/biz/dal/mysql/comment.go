package mysql

import (
	"time"

	"github.com/njupt-sakura/hertz/tiktok_demo/pkg/constants"
	"gorm.io/gorm"
)

type Comment struct {
	ID          int            `json:"id"`
	UserId      int64          `json:"user_id"`
	VideoId     int64          `json:"video_id"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return constants.MysqlDefaultDsn
}

func AddNewComment(comment *Comment) error {
	panic("implement me")
}
