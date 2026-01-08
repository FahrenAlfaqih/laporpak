package handler

import (
	"laporpak/model"
	"laporpak/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := service.CreateUser(req.Name)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id harus angka",
		})
		return
	}

	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, ok := service.UpdateUser(id, req.Name)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user tidak ditemukan",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id harus angka",
		})

		return
	}

	ok := service.DeleteUser(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user tidak ditemukan",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user berhasil dihapus",
	})
}
