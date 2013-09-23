package core

import (
	"testing"
)

func TestUser(t* testing.T){
	user := User{"testName", "testEmail"};
	if out, err := user.Render(); err != nil{
		t.Error("Can't render user", user, err)
	} else{
		t.Log(out)
	}
}