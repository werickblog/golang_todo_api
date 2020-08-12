package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tesh254/golang_todo_api/forms"
	"github.com/tesh254/golang_todo_api/helpers"
	"github.com/tesh254/golang_todo_api/models"
	"github.com/tesh254/golang_todo_api/services"
)

// Import bookmark model from the models file
var bookmarkModel = new(models.BookmarkModel)

// BookmarkController defines the bookmark controller
type BookmarkController struct{}

func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

// FetchBookmarks controller handles fetching all bookmarks of a specific user
func (b *BookmarkController) FetchBookmarks(c *gin.Context) {
	user := c.MustGet("User").(models.User)

	if user.Email == "" {
		responseWithError(c, 404, "Please login")
		return
	}

	var linkModel models.BookmarkModel

	results, err := linkModel.FetchBookmarks(user.ID)

	if err != nil {
		responseWithError(c, 500, "Problem fetching your articles")
		return
	}

	if results != nil {
		c.JSON(200, gin.H{"bookmarks": results})
	} else {
		c.JSON(200, gin.H{"bookmarks": []string{}})
	}

}

// DeleteBookmark controller handles deleting a single bookmark
func (b *BookmarkController) DeleteBookmark(c *gin.Context) {
	user := c.MustGet("User").(models.User)

	if user.Email == "" {
		responseWithError(c, 404, "Please login")
		return
	}

	bookmarkID, found := c.GetQuery("bookmark_id")

	if !found {
		responseWithError(c, 400, "Please provide bookmark id")
		return
	}

	var linkModel models.BookmarkModel

	err := linkModel.DeleteBookmark(bookmarkID)

	if err != nil {
		responseWithError(c, 500, "Problem deleting bookmark")
		return
	}

	c.JSON(204, gin.H{"message": "Deleted bookmark successfully"})
}

// CreateBookmak controller handles creating a bookmark of a specifi user
func (b *BookmarkController) CreateBookmak(c *gin.Context) {
	user := c.MustGet("User").(models.User)

	if user.Email == "" {
		responseWithError(c, 404, "Please login")
		return
	}

	var data forms.BookmarkPayload

	// Check if required fields are provided
	if c.BindJSON(&data) != nil {
		log.Fatal(c.BindJSON(&data))
		responseWithError(c, 406, "Please provide link, and name")
		return
	}

	var linkModel models.BookmarkModel

	if !helpers.IsValidURL(data.Link) {
		responseWithError(c, 400, "Link is invalid")
	}

	var scrapper services.Scrapper

	var meta services.Meta = scrapper.CallWebsite(data.Link, c)

	var bookmarkPayload models.Link = models.Link{
		Name:            data.Name,
		MetaImage:       meta.Image,
		MetaDescription: meta.Description,
		MetaSite:        meta.Site,
		MetaURL:         meta.URL,
		Link:            data.Link,
		Owner:           user.ID,
		CreateAt:        time.Now(),
	}

	err := linkModel.CreateBookmark(bookmarkPayload)

	if err != nil {
		responseWithError(c, 500, "Problem saving your bookmark")
		log.Fatal(err)
		return
	}

	c.JSON(201, gin.H{"message": "Bookmark saved", "bookmark": bookmarkPayload})
}
