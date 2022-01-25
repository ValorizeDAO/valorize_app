package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"
	"valorize-app/services/stringsUtil"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	server *Server
}

func NewAuthHandler(s *Server) *AuthHandler {
	return &AuthHandler{
		server: s,
	}
}

func (auth *AuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user := models.User{}

	err := auth.server.DB.First(&user, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": username + " does not exist",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "incorrect password for " + username,
		})

	}

	token, _, err := services.NewToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not create token",
		})
	}

	cookie := services.CreateTokenCookie(token)
	c.SetCookie(cookie)

	userStruct, err := json.Marshal(models.GetUserProfile(&user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not find logged in user information",
		})
	}
	return c.JSON(http.StatusOK, json.RawMessage(userStruct))

}

func (auth *AuthHandler) Logout(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "not logged in",
		})
	}
	cookie.Value = ""
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "goodbye",
	})
}

func (auth *AuthHandler) Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	if stringsUtil.StringInSlice(username, []string{
		"admin",
		"root",
		"register",
		"login",
		"edit-profile",
		"dashboard",
		"logout",
	}) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "sorry, username is reserved",
		})
	}
	if username == "" || email == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "username, email, and password are required",
		})
	}

	if len(username) < 3 || len(username) > 20 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "username must be between 3 and 20 characters",
		})
	}

	invalidChars := "@#$%^&*()=+[]{}|;:'\",<.>/?`~"
	if strings.ContainsAny(username, invalidChars) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "username cannot contain any of: " + invalidChars,
		})
	}
	user := models.User{
		Email:    email,
		Username: username,
		Password: string(hash),
		Avatar:   "https://res.cloudinary.com/valorize/image/upload/v1628386107/default_avatar.png",
	}

	if auth.server.DB.Create(&user).Error != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": user.Username + " already exists",
		})
	}

	go ethereum.StoreUserKeystore(password, user.ID, auth.server.DB)

	token, _, err := services.NewToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not create token",
		})
	}

	cookie := services.CreateTokenCookie(token)
	c.SetCookie(cookie)

	userStruct, err := json.Marshal(models.GetUserProfile(&user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not find logged in user information",
		})
	}
	return c.JSON(http.StatusCreated, json.RawMessage(userStruct))

}

func (auth *AuthHandler) ShowUser(c echo.Context) error {
	user, err := services.AuthUser(c, *auth.server.DB)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find logged in user information",
		})
	}

	userStruct, err := json.Marshal(models.GetUserProfile(&user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not find logged in user information",
		})
	}
	return c.JSON(http.StatusOK, json.RawMessage(userStruct))
}

func (auth *AuthHandler) UpdatePicture(c echo.Context) error {
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	userData, err := services.AuthUser(c, *auth.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find " + userData.Username,
		})
	}

	file, err := c.FormFile("picture")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not process file",
		})
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filename := os.Getenv("ENVIRONMENT")[0:1] + "/" + strconv.FormatUint(uint64(userData.ID), 10) + "_avatar"
	resp, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{PublicID: filename})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "server storage error",
		})
	}

	cachedUrl := resp.SecureURL + "?t=" + strconv.FormatInt(time.Now().Unix(), 10)
	userData.Avatar = cachedUrl
	if auth.server.DB.Save(&userData).Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not save user data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"image": cachedUrl,
	})
}

func (auth *AuthHandler) UpdateProfile(c echo.Context) error {
	userData, err := services.AuthUser(c, *auth.server.DB)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find logged in user",
		})
	}

	name := c.FormValue("name")
	about := c.FormValue("about")

	userData.Name = name
	userData.About = about

	err = auth.server.DB.Save(&userData).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "database error",
		})
	}
	userStruct, err := json.Marshal(models.GetUserProfile(&userData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not parse user data",
		})
	}
	return c.JSON(http.StatusOK, json.RawMessage(userStruct))
}

func (auth *AuthHandler) UpdateTokenData(c echo.Context) error {
	user, err := services.AuthUser(c, *auth.server.DB)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find logged in user",
		})
	}
	tokenType := c.FormValue("tokenType")
	contractVersion := c.FormValue("contractVersion")
	vaultAddress := c.FormValue("vaultAddress")
	tokenName := c.FormValue("tokenName")
	tokenTicker := c.FormValue("tokenTicker")
	adminAddresses := c.FormValue("adminAddresses")
	chainId := c.FormValue("chainId")
	txHash := c.FormValue("txHash")
	contractAddress := c.FormValue("contractAddress")
	if tokenType != "simple" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error, invalid token type parameter",
		})
	}

	adminAccounts := []*models.Wallet{}
	addresses := strings.Split(adminAddresses, ",")
	for i := range addresses {
		wallet := models.Wallet{}

		auth.server.DB.FirstOrCreate(&wallet, models.Wallet{Address: addresses[i]})
		fmt.Printf("==================\n\n\n\nid: %v  wallet: %v \n\n\n\n========================", wallet.ID, wallet.Address)
		adminAccounts = append(adminAccounts, &wallet)
	}

	token := models.Token{
		UserId:          user.ID,
		TokenType:       tokenType,
		ContractVersion: contractVersion,
		Name:            tokenName,
		Symbol:          tokenTicker,
		ChainId:         chainId,
		VaultAddress:    vaultAddress,
		AdminAddresses:  adminAccounts,
		Address:         contractAddress,
		TxHash:          txHash,
	}

	err = auth.server.DB.Create(&token).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not store contract information: " + err.Error(),
		})
	}

	user.HasDeployedToken = true
	err = auth.server.DB.Save(&user).Error
	tokenStruct, err := json.Marshal(token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not update user info: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json.RawMessage(tokenStruct))
}

type response struct {
	Success bool          `json:"success"`
	Links   []models.Link `json:"links"`
}

var jsonRequest map[string][]models.Link

func (auth *AuthHandler) UpdateLinks(c echo.Context) error {

	if err := c.Bind(&jsonRequest); err != nil {
		return err
	}
	userData, err := services.AuthUser(c, *auth.server.DB)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find logged in user",
		})
	}

	links := jsonRequest["links"]
	userLinks, err := models.GetUserLinks(&userData, *auth.server.DB)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not get user links",
		})
	}

	addedLinks := []models.Link{}
	for _, link := range links {
		var err error
		if link.ID != 0 {
			isUsersLink := checkIfUserLinkExists(userLinks, link)
			if isUsersLink {
				err = models.SaveLink(&userData, link, *auth.server.DB)
				addedLinks = append(addedLinks, link)
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "could not save '" + link.Label + "': " + err.Error(),
				})
			}
		} else {
			link.UserId = userData.ID
			link, err = models.CreateLink(&userData, link, *auth.server.DB)
			addedLinks = append(addedLinks, link)
		}
	}
	return c.JSON(http.StatusOK, response{
		Success: true,
		Links:   addedLinks,
	})
}

func checkIfUserLinkExists(userLinks []models.Link, link models.Link) bool {
	linkExists := false

	for _, userLink := range userLinks {
		if userLink.ID == link.ID && userLink.UserId == link.UserId {
			return true
		}
	}
	return linkExists
}

func (auth *AuthHandler) DeleteLinks(c echo.Context) error {
	linkId := c.FormValue("id")
	linkIdInt, err := strconv.Atoi(linkId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "id parameter invalid",
		})
	}

	userData, err := services.AuthUser(c, *auth.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find " + userData.Username,
		})
	}

	links, err := models.GetUserLinks(&userData, *auth.server.DB)

	if err != nil || len(links) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find user links",
		})
	}

	for _, link := range links {
		if link.ID != uint(linkIdInt) {
			continue
		} //boolean gate to check if link is associated with user
		err = models.DeleteLink(link, *auth.server.DB)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not delete link",
			})
		} else {
			return c.JSON(http.StatusOK, map[string]string{
				"success": "link deleted",
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}
