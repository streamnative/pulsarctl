package pulsar

import (
	"fmt"
	"reflect"
)

func makeHttpPath(apiVersion string, componentPath string) string {
	return fmt.Sprintf("/admin/%s%s", apiVersion, componentPath)
}

// IsNil check if the interface is nil
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
