package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type API struct {
	UserService Service
}

func ProvideUserAPI(us Service) API {
	return API{UserService: us}
}

func (api *API) Add(c *gin.Context) {
	var userDTO UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api.UserService.Add(userDTO.Username, userDTO.FirstName, userDTO.LastName, userDTO.Email)
}
