package auth

import (
	"github.com/gin-gonic/gin"
)

const COLLECTION_NAME = "auth"

type HandleLoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// gets the username and password from the request
// get the collection instance
// atempt to search for a document in the collection with a matching username and password
// if the document is found, return the document
// if the document is not found, return an error
func HandleLogin(c *gin.Context) {

	var requestBody HandleLoginRequestBody
	c.BindJSON(&requestBody)

	credentials, err := getUserByCredentials(requestBody.Username, requestBody.Password)

	if err != nil {
		println(err.Error())
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, credentials)

}
