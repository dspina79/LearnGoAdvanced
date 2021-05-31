package main

import (
	"fmt"
	"os/exec"
)

func checkError(err error, cmd string) {
	if err != nil {
		panic(cmd + " " + err.Error())
	}
}

func main() {
	dateCmd := exec.Command("date")

	dateOutput, err := dateCmd.Output()
	checkError(err, "date")
	fmt.Println("> date")
	fmt.Print(string(dateOutput))

	dirCommand := exec.Command("bash", "-c", "ls -a -l -h")
	dirOutput, err := dirCommand.Output()
	checkError(err, "ls")
	fmt.Println("$> ls -a -l -h")
	fmt.Println(string(dirOutput))
}
