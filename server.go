package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}
	_name := "Jon Snow"
	_isAdmin := true
	_expiration := time.Now().Add(time.Hour * 72).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = _name
	claims["admin"] = _isAdmin
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
		"user":    _name,
		"expires": strconv.Itoa(int(_expiration)),
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func writeCookie(c echo.Context) error {
	return c.String(http.StatusOK, "write a cookie")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://valorize.local:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// Login route

	e.Static("/*", "app/dist")
	e.GET("/public", accessible)
	e.POST("/login", login)
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/test", restricted)

	api := e.Group("/api/v0")
	api.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "All systems GO")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
