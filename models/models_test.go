package models

import "testing"

func TestNamePrefixes(t *testing.T) {
	prefix := ItemNamePrefix(1)
	if len(prefix) < 1 {
		t.Errorf("Prefix of 0 length: %s", prefix)
	}
}

func TestPrefixedName(t *testing.T) {
	for i := 0; i < 1000; i++ {
		suffix := "Stabber"
		name := PrefixedItemName(suffix, uint16(i))
		if len(name) <= (len(suffix) + 1) {
			t.Errorf("Prefixed name '%s' is no longer than suffix '%s'", name, suffix)
		}
	}
}
