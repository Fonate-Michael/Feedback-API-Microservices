package controllers

import (
	"app/db"
	"app/models"
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostFeedBack(context *gin.Context) {
	UserId := context.MustGet("user_id").(int)

	var feeedback models.Feedback

	err := context.BindJSON(&feeedback)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to Bind JSON"})
		return
	}

	_, err = db.DB.Exec("INSERT INTO feedbacks(user_id, sentiment, message) VALUES($1, $2, $3)", UserId, feeedback.Sentiment, feeedback.Message)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to Insert Feedback"})
		return
	}

	context.IndentedJSON(200, gin.H{"message": "Feedback Submitted Successfully"})

}

func GetFeedBack(context *gin.Context) {

	row, err := db.DB.Query("SELECT * FROM feedbacks")
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to Query Feedbacks"})
		return
	}
	defer row.Close()

	var feedbacks []models.Feedback

	for row.Next() {
		var feedback models.Feedback
		err := row.Scan(&feedback.Id, &feedback.UserId, &feedback.Sentiment, &feedback.Message)
		if err != nil {
			context.JSON(500, gin.H{"error": "Failed to Scan Feedbacks"})
			return
		}
		feedbacks = append(feedbacks, feedback)
	}

	context.IndentedJSON(200, feedbacks)
}

func GetFeedBackById(context *gin.Context) {
	idStr := context.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {

		context.JSON(500, gin.H{"error": "Failed to Convert ID to Integer"})
		return
	}
	var feedback models.Feedback
	err = db.DB.QueryRow("SELECT * FROM feedbacks WHERE id = $1", id).Scan(&feedback.Id, &feedback.UserId, &feedback.Sentiment, &feedback.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(404, gin.H{"error": "Feedback Not Found"})
			return
		}
		context.JSON(500, gin.H{"error": "Failed to Scan Feedback"})
		return
	}
	context.IndentedJSON(200, feedback)

}

func SearchFeedBacks(context *gin.Context) {
	searchQuery := context.Query("q")

	searchTerm := "%" + searchQuery + "%"

	row, err := db.DB.Query("SELECT * FROM feedbacks WHERE sentiment ILIKE $1", searchTerm)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to Query Feedbacks"})
		return
	}
	defer row.Close()

	var feedbacks []models.Feedback

	for row.Next() {
		var feedback models.Feedback
		err := row.Scan(&feedback.Id, &feedback.UserId, &feedback.Sentiment, &feedback.Message)
		if err != nil {
			context.JSON(500, gin.H{"error": "Failed to Scan Feedbacks"})
			return
		}
		feedbacks = append(feedbacks, feedback)
	}

	context.IndentedJSON(200, feedbacks)
}

func DeleteFeedBack(context *gin.Context) {
	idStr := context.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to Convert ID to Integer"})
		return
	}

	_, err = db.DB.Exec("DELETE FROM feedbacks WHERE id = $1", id)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to Delete Feedback"})
		return
	}
	context.IndentedJSON(200, gin.H{"message": "Feedback Deleted Successfully"})

}
