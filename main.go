package main

import (
	"fmt"
	"os"
	"os/exec"
)

func prepareCmd(sourcePath, osys, arch string) *exec.Cmd {
	destPath := fmt.Sprintf("%s/cmd/%s/%s/", os.Args[1], osys, arch)
	cmd := exec.Command("go", "build", "-ldflags", "\"-s -w\"", "-o", destPath, sourcePath)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%s", osys))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%s", arch))
	return cmd
}

func main() {
	argsLen := len(os.Args)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Wrong number of arguments, 1 wanted, %d given.\n", argsLen)
		os.Exit(1)
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		//path does not exist
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	sourcePath := fmt.Sprintf("%s/...", os.Args[1])
	for _, osys := range []string{"darwin", "windows", "linux"} {
		for _, arch := range []string{"amd64", "arm64"} {
			if osys == "windows" && arch == "arm64" {
				continue
			}

			_, err := prepareCmd(sourcePath, osys, arch).Output()
			if err != nil {
				fmt.Printf("OS: %s, ARCH: %s -> compilation failed:\n%s\n", osys, arch, err.Error())
				continue
			}
			fmt.Printf("OS: %s, ARCH: %s -> compilation done.\n", osys, arch)
		}
	}
}
