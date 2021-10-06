package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email            string   `json:"email" gorm:"type:varchar(200);"`
	Name             string   `json:"name" gorm:"type:varchar(200);"`
	Username         string   `json:"username" gorm:"type:varchar(200);unique;"`
	Password         string   `json:"password" gorm:"type:varchar(200);"`
	Avatar           string   `json:"avatar" gorm:"type:varchar(200);"`
	About            string   `json:"about" gorm:"type:varchar(1000);"`
	HasDeployedToken bool     `json:"has_deployed_token gorm:boolean"`
	HasVerifiedEmail bool     `json:"has_verified_email gorm:boolean"`
	IsAlphaUser 	 bool     `json:"is_alpha_user gorm:boolean"`
	Token            Token    `json:"token" gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	Wallets          []Wallet `json:"wallets" gorm:"ForeignKey:userId;AssociationForeignKey:user_id"`
}

type UserProfile struct {
	ID               uint   `json:"id"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Username         string `json:"username"`
	Avatar           string `json:"avatar"`
	About            string `json:"about"`
	HasDeployedToken bool   `json:"hasDeployedToken"`
	HasVerifiedEmail bool   `json:"hasVerifiedEmail"`
	IsAlphaUser      bool   `json:"isAlphaUser"`
	Token            Token  `json:"token"`
}

type UserPublicProfile struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Username         string `json:"username"`
	Avatar           string `json:"avatar"`
	About            string `json:"about"`
	HasDeployedToken bool   `json:"hasDeployedToken"`
	Token            Token  `json:"token"`
}

func GetUserProfile(user *User) UserProfile {
	return UserProfile{
		user.ID,
		user.Email,
		user.Name,
		user.Username,
		user.Avatar,
		user.About,
		user.HasDeployedToken,
		user.HasVerifiedEmail,
		user.IsAlphaUser,
		user.Token,
	}
}

func GetUserPublicProfile(user *User) UserPublicProfile {
	return UserPublicProfile{
		user.ID,
		user.Name,
		user.Username,
		user.Avatar,
		user.About,
		user.HasDeployedToken,
		user.Token,
	}
}

func GetUserByID(id uint, db gorm.DB) (User, error) {
	var m User
	if err := db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, nil
		}
		return User{}, err
	}
	return m, nil
}

func GetUserByEmail(e string, db gorm.DB) (User, error) {
	var m User
	if err := db.Where(&User{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, nil
		}
		return User{}, err
	}
	return m, nil
}

func GetUserByUsername(username string, db gorm.DB) (User, error) {
	var m User
	if err := db.Preload("Token").Where("username = ?", username).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, err
		}
		return User{}, err
	}
	return m, nil
}