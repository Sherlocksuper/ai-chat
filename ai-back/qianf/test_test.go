package qianf

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := Token("vhyR5VgbKVSf6OsuxCdzQiSv", "7gocmq6fdGaeM4dXLfnJOq02jzDkiWPV")
	fmt.Println(token.AccessToken)
	if err != nil {
		t.Errorf("TestToken failed")
	}
}
