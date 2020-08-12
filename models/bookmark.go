package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Link defines user object structure
type Link struct {
	ID              bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string        `json:"name" bson:"name"`
	MetaImage       string        `json:"meta_image" bson:"meta_image"`
	MetaDescription string        `json:"meta_description" bson:"meta_description"`
	MetaSite        string        `json:"meta_site" bson:"meta_site"`
	MetaURL         string        `json:"meta_url" bson:"meta_url"`
	MetaTitle       string        `json:"meta_title" bson:"meta_title"`
	Link            string        `json:"link" bson:"link"`
	Owner           bson.ObjectId `json:"owner" bson:"owner"`
	CreateAt        time.Time     `json:"created_at" bson:"created_at"`
}

// BookmarkModel defines the model structure
type BookmarkModel struct{}

// CreateBookmark handles creating a bookmark by the user
func (l *BookmarkModel) CreateBookmark(data Link) error {
	// Connect to the bookmark collection
	collection := dbConnect.Use(databaseName, "bookmark")
	// Assign result to error object while saving bookmark
	err := collection.Insert(bson.M{
		"name":             data.Name,
		"meta_image":       data.MetaImage,
		"meta_description": data.MetaDescription,
		"meta_site":        data.MetaSite,
		"meta_url":         data.MetaURL,
		"meta_title":       data.MetaTitle,
		"link":             data.Link,
		"owner":            data.Owner,
		"created_at":       data.CreateAt,
	})

	return err
}

// FetchBookmarks handles fetching bookmarks by a user
func (l *BookmarkModel) FetchBookmarks(id bson.ObjectId) (bookmarks []Link, err error) {
	collection := dbConnect.Use(databaseName, "bookmark")

	err = collection.Find(bson.M{"owner": id}).Sort("-$natural").All(&bookmarks)

	return bookmarks, err
}

// DeleteBookmark handles deleting a bookmark
func (l *BookmarkModel) DeleteBookmark(id string) error {
	collection := dbConnect.Use(databaseName, "bookmark")

	err := collection.RemoveId(id)

	return err
}
