package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lunovoy/test-project-verba/internal/models"
	"github.com/gin-gonic/gin"
)

var (
	ErrorFetchData  = "error fetching user data"
	ErrorDecodeData = "error decoding user data"
	ErrorAddUser    = "error trying to add user to db"
)

func (h *Handler) addUser(c *gin.Context) {

	response, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorFetchData: err.Error()})
	}
	defer response.Body.Close()

	var user models.User
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorDecodeData: err.Error()})
		return
	}

	fmt.Println(user)

	userEmail, err := h.services.User.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorAddUser: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"added user's email": userEmail,
	})
}
