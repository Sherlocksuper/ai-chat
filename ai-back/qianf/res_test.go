package qianf

import "testing"

func TestGenerateImage(t *testing.T) {
	token, err := Token(APIKey, APISecretKey)
	if err != nil {
		t.Errorf("TestToken failed")
	}
	GenerateImage(token.AccessToken)
}
