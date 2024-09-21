package protocol

import (
	"testing"
)

func TestValidate(t *testing.T) {
	if IsUsernameValid("你好世界") {
		t.Errorf("fail")
	}

	if !IsUsernameValid("1123123123") {
		t.Errorf("fail")
	}
}
