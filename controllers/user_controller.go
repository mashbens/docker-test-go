package controllers

import (
	"app-test/config"
	"app-test/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserReq struct {
	ID    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers(c echo.Context) error {
	var users []models.User

	db := config.DBManager()

	db = db.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func CreateUser(c echo.Context) (err error) {
	var userReq UserReq

	if err = c.Bind(&userReq); err != nil {
		return
	}

	db := config.DBManager()

	user := models.User{
		Name:  userReq.Name,
		Email: userReq.Email,
	}

	db = db.Create(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}

func GetUser(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	db := config.DBManager()

	db = db.First(&user, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}

func UpdateUser(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id := c.Param("id")

	NewUser := models.User{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	}

	db := config.DBManager()

	db = db.Model(&models.User{}).Where("id = ?", id).Updates(NewUser)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   NewUser,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	db := config.DBManager()

	db = db.Where("id = ?", id).Delete(&models.User{})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success Delleted",
		"id":     id,
	})
}
