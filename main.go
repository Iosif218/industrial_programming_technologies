package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type tea struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
}

var teas = []tea{
	{ID: "1", Name: "Lapsang Souchong Black Tea", Type: "Black", Country: "China", Description: "From the Fujian province of China comes this very distinctive tea with its exotic smoky flavor. After plucking, the leaves are withered over cypress or pine wood fires. After the rolling process, they are placed into wooden barrels until they begin to emit their own pleasant aroma. As a final step they are placed in bamboo baskets and hung on racks over smoky pine fires where they dry and absorb the essence of the smoke.", Cost: "$7.95"},
	{ID: "2", Name: "Glenburn Second Flush Darjeeling", Type: "Black", Country: "India", Description: "A rare, fragrant second flush Darjeeling from the foothills of the Himalayas in India brews to a lovely amber color. The taste is exquisite and complex, with a refined malty and muscatel character, and a fruity, floral finish.", Cost: "$14.95"},
	{ID: "3", Name: "Dragon Phoenix Pearl Jasmine Tea", Type: "Green", Country: "China", Description: "Finest downy green leaves are hand-rolled into small spheres and scented with fragrant jasmine blossoms. It brews up an exquisite cup with an alluring aroma and delicate sweet floral taste.", Cost: "$18.95"},
}

var users = []Credentials{
	{Username: "user", Password: "password"},
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
	{Username: "user3", Password: "password3"},
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var validUser *Credentials
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			validUser = &user
			break
		}
	}

	if validUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func refreshToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	newExpirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = newExpirationTime.Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := newToken.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func getTeas(c *gin.Context) {
	c.JSON(http.StatusOK, teas)
}

func getTeaByID(c *gin.Context) {
	id := c.Param("id")

	for _, tea := range teas {
		if tea.ID == id {
			c.JSON(http.StatusOK, tea)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "tea not found"})
}

func createTea(c *gin.Context) {
	var newTea tea

	if err := c.BindJSON(&newTea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	teas = append(teas, newTea)
	c.JSON(http.StatusCreated, newTea)
}

func updateTea(c *gin.Context) {
	id := c.Param("id")
	var updatedTea tea

	if err := c.BindJSON(&updatedTea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, tea := range teas {
		if tea.ID == id {
			teas[i] = updatedTea
			c.JSON(http.StatusOK, updatedTea)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "tea not found"})
}

func deleteTea(c *gin.Context) {
	id := c.Param("id")

	for i, tea := range teas {
		if tea.ID == id {
			teas = append(teas[:i], teas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "tea deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "tea not found"})
}

func main() {
	router := gin.Default()

	router.POST("/login", login)
	router.POST("/refresh", refreshToken)

	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		protected.GET("/teas", getTeas)
		protected.POST("/teas", createTea)
		router.GET("/teas/:id", getTeaByID)
		router.PUT("/teas/:id", updateTea)
		router.DELETE("/teas/:id", deleteTea)
	}

	router.Run(":8080")
}
