package partition

import "testing"

func TestPartition(t *testing.T) {
	var (
		key   = "a_key"
		value = "a_value"
	)
	Set(key, value)
	if Get(key) != value {
		t.Fatal("no Equal")
	}
}
