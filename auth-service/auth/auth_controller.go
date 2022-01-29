package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

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
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	credentials, err := getUserByCredentials(requestBody.Username, requestBody.Password)

	if err != nil {
		log.Println(err.Error())
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := createAuthToken(credentials.ID.Hex())

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.Header(AUTH_HEADER, "Bearer "+token)
	c.JSON(200, credentials)
}
