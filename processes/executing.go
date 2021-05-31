package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Executing ls")

	binaryLs, lsError := exec.LookPath("ls")
	check(lsError)

	args := []string{"ls", "-l", "-h", "-a"}
	env := os.Environ()

	execError := syscall.Exec(binaryLs, args, env)
	check(execError)

}
