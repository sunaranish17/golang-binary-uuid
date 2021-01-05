package main

import (
	"golang-binary-uuid/models"
	"log"
)

func main() {

	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Post{}, &models.Comment{})

	post := &models.Post{Name: "Test Post"}

	if db.Create(&post).Error != nil {
		log.Panic("Unable to create Post")
	}

	log.Print("Post's id: ", post.ID)

	db.Where("id=?", post.ID).First(&post)
	log.Print("where ", post)

	comment := &models.Comment{Name: "First Post", PostID: post.ID}
	if db.Create(&comment).Error != nil {
		log.Panic("Unable to add comment")
	}

	fetchedPost := &models.Post{}
	if err := db.Where("id=?", comment.PostID).Preload("Comment").First(&fetchedPost).Error; err != nil {
		log.Panic("Unable to find created Post")
	}
	// fmt.Println("Post: %+v\n", fetchedPost)
}
