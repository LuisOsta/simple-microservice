package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandleLoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialRetriever = func(username string, password string) (AuthDocument, error)
type Auth struct {
	GetUserByCredentials CredentialRetriever
}

// gets the username and password from the request
// get the collection instance
// atempt to search for a document in the collection with a matching username and password
// if the document is found, return the document
// if the document is not found, return an error
func (a *Auth) HandleLogin(c *gin.Context) {

	var requestBody HandleLoginRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	credentials, err := a.GetUserByCredentials(requestBody.Username, requestBody.Password)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := createAuthToken(credentials.ID.Hex())

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Header(AUTH_HEADER, "Bearer "+token)
	c.JSON(http.StatusOK, credentials)
}
