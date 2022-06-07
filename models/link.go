package models

import (
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

func (m *Model) GetUserLinks(user *User) ([]Link, error) {
	var links []Link
	if err := m.db.Where("user_id = ?", user.ID).Find(&links).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return make([]Link, 0), err
		}
		return make([]Link, 0), err
	}
	return links, nil
}

func (m *Model) SaveLink(user *User, link Link) error {
	link.UserId = user.ID
	if err := m.db.Save(&link).Error; err != nil {
		return err
	}
	return nil
}

func (m *Model) CreateLink(user *User, link Link) (Link, error) {
	link.UserId = user.ID
	result := m.db.Create(&link)
	if result.Error != nil {
		return link, result.Error
	}
	return link, nil
}

func (m *Model) DeleteLink(link Link) error {
	if err := m.db.Delete(&link).Error; err != nil {
		return err
	}
	return nil
}
