package server

import (
	"errors"
	"net/http"

	"github.com/cyrus28214/comments-area/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentView struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type getCommentsRequest struct {
	Page uint `form:"page" binding:"required"` //page从1开始，size表示每页显示的数量
	Size int  `form:"size" binding:"required"` //如果size=-1，则返回全部评论
}

type getCommentsResponse struct {
	Comments []commentView `json:"comments"`
	Total    uint          `json:"total"`
}

func (s *Server) GetComments(c *gin.Context) {
	var req getCommentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		s.HandleBindingError(c, err)
		return
	}

	var total int64
	if result := s.DB.Model(&model.Comment{}).Count(&total); result.Error != nil {
		s.HandleDatabaseError(c, result.Error)
		return
	}

	var comments []commentView
	tx := s.DB.Model(&model.Comment{}).Select("id", "author AS name", "content").Order("created_at DESC") //按创建时间倒序排列
	if req.Size == -1 {
		result := tx.Find(&comments)
		if result.Error != nil {
			s.HandleDatabaseError(c, result.Error)
			return
		}
	} else {
		offset := (int(req.Page) - 1) * req.Size
		result := tx.Limit(req.Size).Offset(offset).Find(&comments)
		if result.Error != nil {
			s.HandleDatabaseError(c, result.Error)
			return
		}
	}

	c.JSON(http.StatusOK, Success(getCommentsResponse{
		Comments: comments,
		Total:    uint(total),
	}))
}

type CreateCommentRequest struct {
	Name    string `json:"name" binding:"required"`    //必填，非空
	Content string `json:"content" binding:"required"` //必填，非空
}

const (
	codeCommentNotFound = 10001
)

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

	res := commentView{
		ID:      comment.ID,
		Name:    comment.Author,
		Content: comment.Content,
	}

	s.Log.Debugf("Comment created with ID: %d", comment.ID)
	c.JSON(http.StatusCreated, res)
}

type deleteCommentRequest struct {
	ID uint `form:"id" binding:"required"`
}

func (s *Server) DeleteComment(c *gin.Context) {
	var req deleteCommentRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		s.HandleBindingError(c, err)
		return
	}

	var comment model.Comment
	result := s.DB.First(&comment, req.ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, &Response{
				Code: codeCommentNotFound,
				Msg:  "Comment not found",
				Data: nil,
			})
		} else {
			s.HandleDatabaseError(c, result.Error)
		}
	}

	result = s.DB.Delete(&comment)
	if result.Error != nil {
		s.HandleDatabaseError(c, result.Error)
		return
	}

	s.Log.Debugf("Comment with ID: %d deleted", req.ID)
	c.JSON(http.StatusOK, Success(nil))
}
