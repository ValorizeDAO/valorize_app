package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Link struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Label       string `json:"label" gorm:"type:varchar(200);"`
	Icon        string `json:"icon" gorm:"type:varchar(200);"`
	Url         string `json:"url" gorm:"type:varchar(200);"`
	UserId      uint   `json:"user_id"`
	Description string `json:"description" gorm:"type:varchar(200);"`
}

func GetUserLinks(user *User, db gorm.DB) ([]Link, error) {
	var links []Link
	if err := db.Where("user_id = ?", user.ID).Find(&links).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return make([]Link, 0), err
		}
		return make([]Link, 0), err
	}
	return links, nil
}

func SaveLink(user *User, link Link, db gorm.DB) error {
	link.UserId = user.ID
	fmt.Println("\n\n SAVING \n\n")
	if err := db.Save(&link).Error; err != nil {
		return err
	}
	return nil
}

func CreateLink(user *User, link Link, db gorm.DB) (Link, error) {
	link.UserId = user.ID
	fmt.Println("\n\n CREATING \n\n")
	result := db.Create(&link)
	if result.Error != nil {
		return link, result.Error
	}
	return link, nil
}

func DeleteLink(link Link, db gorm.DB) error {
	if err := db.Delete(&link).Error; err != nil {
		return err
	}
	return nil
}
