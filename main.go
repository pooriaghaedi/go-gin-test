package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Mobilenumber string `json:"mobilenumber"`
	Nationalcode string `json:"nationalcode"`
	UUID         string `json:"uuid"`
}

var users = []user{
	{ID: "1", Name: "Pooria Ghaedi", Mobilenumber: "02133990496", Nationalcode: "2980510743", UUID: ""},
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
			// max := findmax(a.Mobilenumber)
			// countOfMax := countmax(max, a.Mobilenumber)
			// fmt.Printf("number of %v is %v .", max, countOfMax)
			sliced := strings.Split(a.Mobilenumber, "")
			slicedNational := strings.Split(a.Nationalcode, "")
			ints := make([]int, len(sliced))

			for i, s := range sliced {
				ints[i], _ = strconv.Atoi(s)
			}
			// ints = sort.Ints(ints)

			sort.Ints(ints)
			intsnational := make([]int, len(slicedNational))

			for i, s := range slicedNational {
				intsnational[i], _ = strconv.Atoi(s)
			}
			// ints = sort.Ints(ints)

			sort.Ints(intsnational)
			// var slice []int
			// slice[len(ints)-1:][0]
			var uuid1 []int
			uuid1 = append(uuid1, ints[len(ints)-1])
			uuid1 = append(uuid1, ints[len(ints)-2])
			uuid1 = append(uuid1, ints[len(ints)-3])
			uuid1 = append(uuid1, intsnational[0])
			uuid1 = append(uuid1, intsnational[1])
			uuid1 = append(uuid1, intsnational[2])
			// arrayToString := strings.Trim(strings.Replace(fmt.Sprint(uuid1), " ", " ", -1), "[]")
			arrayToString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(uuid1), " "), " "), "[]")
			arrayToString = strings.ReplaceAll(arrayToString, " ", "")
			fmt.Println(arrayToString)
			// fmt.Println(uuid1)
			// sorted := strings.Fields(a.Mobilenumber)
			// fmt.Println(sorted, len(sorted))
			// if countOfMax < 3 {
			// a.UUID=ints
			// }
			a.UUID = arrayToString
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// func sortNumbers (c string) string {
// 	sort.Strings(c)
// }
func findmin(c string) int {
	a := 0
	for _, b := range c {
		value, _ := strconv.Atoi(string(b))
		if a > value {
			a = value
		}
	}
	return a

}

func findmax(c string) int {
	var a int
	for _, b := range c {
		value, _ := strconv.Atoi(string(b))

		if a < value {
			a = value
		}
	}
	return a

}
func countmax(a int, c string) int {
	count := 0
	for _, b := range c {
		if fmt.Sprint(a) == string(b) {
			count++
		}
	}

	return count

}

// func changeUUID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of users, looking for
// 	// an user whose ID value matches the parameter.
// 	for _, a := range users {
// 		if a.ID == id {
// 			a.UUID=
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
// }

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
