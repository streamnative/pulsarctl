// Copyright (c) 2020 StreamNative, Inc.. All Rights Reserved.

package os

import (
	"os/exec"
	"runtime"
)

// Interactor abstracts opening a url on the users OS
type Interactor interface {
	OpenURL(url string) error
}

// DefaultInteractor is the default OS interactor that wraps go standard os
// package functionality to satisfy different interfaces
type DefaultInteractor struct{}

// OpenURL opens the passed in url in the default OS browser
func (i *DefaultInteractor) OpenURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
