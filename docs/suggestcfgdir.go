package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// suggestCfgDir() takes a stab at return the recommended location
// of the directory that stores user configuration data. It does
// so by sniffing the operating system. It returns the suggested
// directory in the dir variable and the OS it detected in the os variable.
// Go Playground:
// https://play.golang.org/p/BeXM5iS66X3
// Gist:
// https://gist.github.com/tomcam/508f7a95a269b0d39781590ad47e6e75
func suggestCfgDir() (dir, system string) {
	system = runtime.GOOS
	// Try to determine where user application home dir would be
	switch system {
	case "windows":
		return os.Getenv("%APPDATA%"), system
	case "darwin":
		u, _ := user.Current()
		return filepath.Join(u.Username, u.HomeDir, "Library", "Preferences"), system
	case "linux":
		return os.Getenv("HOME"), system
	default:
		return os.Getenv("HOME"), system
	}
}

func main() {
	configDir, opsys := suggestCfgDir()
	fmt.Printf("Suggested place to store config files on %s: %s\n", opsys, configDir)
}
