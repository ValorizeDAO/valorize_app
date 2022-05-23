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
	IsAlphaUser      bool     `json:"is_alpha_user gorm:boolean"`
	Token            Token    `json:"token" gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	Wallets          []Wallet `json:"wallets" gorm:"ForeignKey:userId;AssociationForeignKey:user_id"`
	Links            []Link   `json:"links" gorm:"ForeignKey:userId;AssociationForeignKey:user_id"`
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
	Links            []Link `json:"links"`
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

func (m *Model) GetUserProfile(user *User) UserProfile {
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
		user.Links,
	}
}

func (m *Model) GetUserPublicProfile(user *User) UserPublicProfile {
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

func (m *Model) GetUserByID(id uint) (User, error) {
	var u User
	if err := m.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, nil
		}
		return User{}, err
	}
	return u, nil
}

func (m *Model) GetUserByEmail(e string) (User, error) {
	var u User
	if err := m.db.Where(&User{Email: e}).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, nil
		}
		return User{}, err
	}
	return u, nil
}

func (m *Model) GetUserByUsername(username string) (User, error) {
	var u User
	if err := m.db.Preload("Token", "token_type= ?", "creator").Where("username = ?", username).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, err
		}
		return User{}, err
	}
	return u, nil
}

func (m *Model) GetUserProfileByUsername(username string) (User, error) {
	var u User
	if err := m.db.Preload("Token").Where("username = ?", username).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, err
		}
		return User{}, err
	}
	links, err := m.GetUserLinks(&u)
	if err != nil {
		return User{}, err
	}
	u.Links = links
	return u, nil
}
