package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisshihi/golang-lessons/models"
	"github.com/mattn/go-sqlite3"
)

func signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sqliteErr sqlite3.Error
	if err := user.Save(); err != nil {
		status, msg := handleDBError(err)
		if errors.As(err, &sqliteErr) {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				msg = "Email already in use"
			}
		}
		c.JSON(status, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

type getUserByEmailReq struct {
	Email string `json:"email" binding:"required,email"`
}

func getUserByEmail(c *gin.Context) {
	var req getUserByEmailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		status, msg := handleDBError(err)
		c.JSON(status, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
	})
}

// Обработка ошибок sqlite3 и sql
func handleDBError(err error) (int, string) {
	var sqliteErr sqlite3.Error
	switch {
	case errors.As(err, &sqliteErr):
		switch sqliteErr.ExtendedCode {
		case sqlite3.ErrConstraintUnique:
			return http.StatusConflict, "Email already in use"
		default:
			return http.StatusInternalServerError, sqliteErr.Error()
		}
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound, "User not found"
	default:
		return http.StatusInternalServerError, err.Error()
	}
}
