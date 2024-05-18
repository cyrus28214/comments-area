package server

import (
	"net/http"

	"github.com/cyrus28214/comments-area/model"
	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type CreateCommentResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (s *Server) CreateComment(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		s.HandleBindingError(c, err)
		return
	}

	var comment = model.Comment{
		Author:  req.Name,
		Content: req.Content,
	}

	result := s.DB.Create(&comment)

	if result.Error != nil {
		s.HandleDatabaseError(c, result.Error)
		return
	}

	res := CreateCommentResponse{
		ID:      comment.ID,
		Name:    comment.Author,
		Content: comment.Content,
	}

	s.Log.Debugf("Comment created with ID: %d", comment.ID)
	c.JSON(http.StatusCreated, res)
}
