package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// book represents data about a book entity.
type book struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Author    string  `json:"artist"`
	Publisher string  `json:"publisher"`
	Price     float64 `json:"price"`
}

// book slice data.
var books = []book{}

func main() {

	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.GET("/books", getBooks)
	router.POST("/book", addBooks)
	router.POST("/books/:id", getBookByID)
	router.PUT("/removeBook/:id", removeBookByID)

	router.Run("localhost:8089")

}

/*
GET api to get all books
*/
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

/*
addBooks adds an book from JSON received in the request body.
*/
func addBooks(c *gin.Context) {
	var newBook book

	/*   Call BindJSON to bind the received JSON to
	     newBook. */
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

/*
getBookByID locates the book whose ID value matches the id
parameter sent by the client, then returns that book as a response.
*/
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	/*   Loop over the list of books, looking for
	     an book whose ID value matches the parameter. */
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

/*
getBookByID locates the book whose ID value matches the id
parameter sent by the client, then return true after deletion or message if id not found.
*/
func removeBookByID(c *gin.Context) {
	id := c.Param("id")

	/*   Loop over the list of books, looking for
	     an book whose ID value matches the parameter. */
	for i, a := range books {
		if a.ID == id {
			books[i] = books[len(books)-1]
			books[len(books)-1] = book{}
			books = books[:len(books)-1]
			fmt.Println(i)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted sucessfully."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
