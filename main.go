package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Mobilenumber string `json:"mobilenumber"`
	Nationalcode string `json:"nationalcode"`
}

var users = []user{
	{ID: "1", Name: "Pooria Ghaedi", Mobilenumber: "09133990496", Nationalcode: "2980510743"},
}

// postusers adds an user from JSON received in the request body.
func postUser(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to
	// newuser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new user to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUser)
	router.GET("/users/:id", getUserByID)

	router.Run("0.0.0.0:8080")
}

// getUsers responds with the list of all Users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
