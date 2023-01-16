package util

import "testing"

func Test_UniqueValue(t *testing.T) {
	id, err := UniqueValue()
	if err != nil {
		t.Fatalf("Unexpected error while generating unique id %v", err)
	}

	if len(id) != 64 {
		t.Fatalf("Should have generated exactly a 64 byte string ID")
	}
}
