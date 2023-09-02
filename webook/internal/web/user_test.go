package web

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncrypto(t *testing.T) {
	password := "123456"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}

func TestNil(t *testing.T) {
	testTypeAssert(nil)

}
func testTypeAssert(c any) {
	claims := c.(*UserClaims)
	println(claims.Uid)
}

func TestUserHandler_SignUp(t *testing.T) {
	testCases := []struct {
		name string
	}{}
	req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer([]byte(
		`{
"email":"123@qq.com",
"password":"12345"
}`)))
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	println(resp, req)
	//resp.Code

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//handler := NewUserHandler(nil)
			//ctx := &gin.Context{}
			//handler.SignUp(ctx)
		})
	}
}
