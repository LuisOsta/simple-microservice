package profile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createProfileTest struct {
	body               createProfileBody
	profileCreator     ProfileCreator
	expectedStatusCode int
}

func mockSuccessProfileCreator(address string, phone string, uid string) (ProfileDocument, error) {
	oid, _ := primitive.ObjectIDFromHex(uid)

	return ProfileDocument{
		Address: address,
		Phone:   phone,
		UserID:  oid,
	}, nil
}

func mockFailedProfileCreator(address string, phone string, uid string) (ProfileDocument, error) {
	return ProfileDocument{}, fmt.Errorf("Failed to create profile")
}

func TestHandleCreateProfile(t *testing.T) {
	tests := []createProfileTest{
		{
			body: createProfileBody{
				Address: "address",
				Phone:   "phone",
				UserId:  "userId",
			},
			profileCreator:     mockSuccessProfileCreator,
			expectedStatusCode: 200,
		},
		{
			body: createProfileBody{
				Address: "address",
				Phone:   "phone",
				UserId:  "userId",
			},
			profileCreator:     mockFailedProfileCreator,
			expectedStatusCode: 500,
		},
		// should fail bind json
		{
			body:               createProfileBody{},
			profileCreator:     mockSuccessProfileCreator,
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		body, _ := json.Marshal(test.body)

		if (test.body == createProfileBody{}) {
			body = []byte("")
		}
		c.Request, _ = http.NewRequest("POST", "/", io.Reader(bytes.NewBuffer(body)))
		c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")
		p := Profile{
			CreateProfile: test.profileCreator,
		}
		p.HandleCreateProfile(c)

		if w.Code != test.expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", test.expectedStatusCode, w.Code)
		}
	}

}

type updateProfileTest struct {
	body               updateProfileBody
	profileUpdator     ProfileUpdator
	expectedStatusCode int
}

func mockSuccessProfileUpdator(uid string, p updatePayload) (ProfileDocument, error) {
	oid, _ := primitive.ObjectIDFromHex(uid)

	return ProfileDocument{
		Address: p.Address,
		Phone:   p.Phone,
		UserID:  oid,
	}, nil
}

func mockFailedProfileUpdator(uid string, p updatePayload) (ProfileDocument, error) {
	return ProfileDocument{}, fmt.Errorf("Failed to update profile")
}

func TestHandleUpdateProfile(t *testing.T) {
	tests := []updateProfileTest{
		{
			body: updateProfileBody{
				Address: "address",
				Phone:   "phone",
			},
			profileUpdator:     mockSuccessProfileUpdator,
			expectedStatusCode: 200,
		},
		{
			body: updateProfileBody{
				Address: "address",
				Phone:   "phone",
			},
			profileUpdator:     mockFailedProfileUpdator,
			expectedStatusCode: 500,
		},
		{
			body:               updateProfileBody{},
			profileUpdator:     mockSuccessProfileUpdator,
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		body, _ := json.Marshal(test.body)
		if (test.body == updateProfileBody{}) {
			body = []byte("")
		}
		c.Request, _ = http.NewRequest("PUT", "/profile/1234", io.Reader(bytes.NewBuffer(body)))
		c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")
		p := Profile{
			UpdateProfile: test.profileUpdator,
		}
		p.HandleUpdateProfile(c)

		if w.Code != test.expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", test.expectedStatusCode, w.Code)
		}
	}
}
