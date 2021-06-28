package cmdutils

import "fmt"

// Version information.
var (
	ReleaseVersion = "None"
	BuildTS        = "None"
	GitHash        = "None"
	GitBranch      = "None"
	GoVersion      = "None"
)

func PrintVersionInfo() {
	fmt.Printf("Release Version: %s\n", ReleaseVersion)
	fmt.Printf("Git Commit Hash: %s\n", GitHash)
	fmt.Printf("Git Branch: %s\n", GitBranch)
	fmt.Printf("UTC Build Time: %s\n", BuildTS)
	fmt.Printf("Go Version: %s\n", GoVersion)
}
