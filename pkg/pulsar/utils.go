package pulsar

import (
	"fmt"
)

func makeHTTPPath(apiVersion string, componentPath string) string {
	return fmt.Sprintf("/admin/%s%s", apiVersion, componentPath)
}
