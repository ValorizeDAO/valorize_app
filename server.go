package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/db/models"
)

func (auth *AuthHandler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.User{}

	err := auth.db.First(&user, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": username + " does not exist",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return echo.ErrUnauthorized
	}

	_user := user.Username
	_expiration := time.Now().Add(time.Hour * 144).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = _user
	claims["exp"] = _expiration

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(144 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"token":   t,
		"user":    _user,
		"expires": strconv.Itoa(int(_expiration)),
	})
}

type AuthHandler struct {
	db *gorm.DB
}

func (auth *AuthHandler) register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := models.User{
		Email:    email,
		Username: username,
		Password: string(hash),
	}
	checkIfUser := user
	if auth.db.First(&checkIfUser).Error != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": username + " already exists",
		})
	}
	fmt.Print(user.CreatedAt)
	if auth.db.Create(&user).Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not create user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"email": email,
		"user":  username,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func validateJwtInCookie(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Valid)
		//TODO: get better key management for app
		return []byte("secret"), nil
	})
	if token.Valid {
		return c.String(http.StatusOK, "hello "+claims.Username)
	} else {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
}

func main() {
	cfg := config.NewConfig()
	dbInstance := db.Init(cfg)
	auth := AuthHandler{db: dbInstance}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://valorize.local:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Static("/*", "app/dist")
	e.GET("/public", accessible)
	e.POST("/login", auth.login)
	e.POST("/register", auth.register)

	r := e.Group("/restricted")
	r.GET("/test", validateJwtInCookie)

	api := e.Group("/api/v0")
	api.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "All systems GO")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
