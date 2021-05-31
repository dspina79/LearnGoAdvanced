package main

import (
	"fmt"
	"io/ioutil"
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

	grepCommand := exec.Command("grep", "hello")
	gIn, _ := grepCommand.StdinPipe()
	gOut, _ := grepCommand.StdoutPipe()

	grepCommand.Start()
	gIn.Write([]byte("hello grep\ngoodbye grep\nthere are other things than hello"))
	gIn.Close()
	grepBytes, _ := ioutil.ReadAll(gOut)
	fmt.Println(">grep hello")
	fmt.Println(string(grepBytes))
}
