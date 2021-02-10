package database

import (
	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our db based on the models passed in
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
