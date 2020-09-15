package types

import "fmt"

// SuperSlice type ...
type SuperSlice []interface{}

// Find func ...
func (s SuperSlice) Find(value interface{}) (int, bool) {
	for index, item := range s {
		if value == item {
			return index, true
		}
	}

	return -1, false
}

// Join func ...
func (s SuperSlice) Join(sep string) string {
	var joined string

	for index, item := range s {
		if index != 0 {
			joined += sep
		}

		joined += fmt.Sprint(item)
	}

	return joined
}
