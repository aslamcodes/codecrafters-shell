package executable

import (
	"os"
	"os/exec"
	"strings"
)

func RunExecutable(program string, args []string) {
	cmd := exec.Command(program, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func IsInPath(exec string) (string, bool) {
	path := os.Getenv("PATH")
	for location := range strings.SplitSeq(path, string(os.PathListSeparator)) {
		file := location + "/" + exec
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}
	return "", false
}
