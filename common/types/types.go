package types

// KeyValueStr is  a string tuple
type KeyValueStr struct {
	Key   string
	Value string
}

// NewKeyValueStr returns a new KeyValueStr
func NewKeyValueStr(key, value string) KeyValueStr {
	return KeyValueStr{key, value}
}
