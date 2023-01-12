package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main()  {
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

	for _, osys := range []string{"darwin", "windows", "linux"} {
		for _, arch := range []string{"amd64", "arm64"} {
			if osys == "windows" && arch == "arm64" {
				continue
			}
			cmd := exec.Command(fmt.Sprintf("GOOS=%s GOARCH=%s go build -ldflags \"-s -w\" -o %s/cmd/%s/%s/ %s/...", osys, arch, os.Args[1], osys, arch, os.Args[1]))
			_, err := cmd.Output()

			if err != nil {
				fmt.Printf("OS: %s, ARCH: %s -> compilation failed:\n%s\n", osys, arch, err.Error())
				continue
			}
			fmt.Printf("OS: %s, ARCH: %s -> compilation done.\n", osys, arch)
		}
	}
}