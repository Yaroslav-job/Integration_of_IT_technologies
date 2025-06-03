package generator

import "testing"

func TestRandomString(t *testing.T) {
	str := randomString()
	if len(str) != StrLen {
		t.Errorf("Expected length %d, got %d", StrLen, len(str))
	}
}
