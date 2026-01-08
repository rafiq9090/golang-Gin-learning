package handler

import (
	"go_project_Gin/internal/notification"
	"go_project_Gin/internal/service"
	"go_project_Gin/internal/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetUint("user_id")
	caption := c.PostForm("caption")

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		utils.JSONError(c, "Image required", http.StatusBadRequest, nil)
		return
	}
	defer file.Close()
	fileName := "upload/" + strconv.Itoa(int(time.Now().Unix())) + filepath.Ext(header.Filename)
	os.MkdirAll("upload", 0755)
	if err != nil {
		utils.JSONError(c, "Failed to create upload directory", http.StatusInternalServerError, nil)
		return
	}

	defer file.Close()
	out, err := os.Create(fileName)
	if err != nil {
		utils.JSONError(c, "Failed to create file", http.StatusInternalServerError, nil)
		return
	}
	defer out.Close()
	io.Copy(out, file)

	imageURL := "http://localhost:8080/" + fileName
	post, err := service.Post.CreatePost(ctx, userID, caption, imageURL)
	if err != nil {
		utils.JSONError(c, "Failed to create post", http.StatusInternalServerError, nil)
		return
	}
	// Fetch user to get email for notification
	userObj, err := service.Auth.GetUserById(ctx, userID)
	if err == nil {
		go notification.SendTaskNotification(userObj.Email, userID, post.ID, "post_created")
	}
	c.JSON(http.StatusOK, post)
}

func GetPostsByUserId(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetUint("user_id")

	posts, err := service.Post.GetPostsByUserId(ctx, userID)
	if err != nil {
		utils.JSONError(c, "Failed to get posts", http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()
	posts, err := service.Post.GetAllPost(ctx)
	if err != nil {
		utils.JSONError(c, "Failed to get posts", http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, posts)
}
