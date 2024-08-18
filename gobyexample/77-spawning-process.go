package main

import (
	"os/exec"
	"fmt"
	"io"
)


func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("date>>>", string(dateOut))


	// _, err1 := exec.Command("date", "-x").Output()
	// if err1 != nil {
	// 	// fmt.Println(">>>", err1)
	// 	// panic(err1)
	// 	switch e := err1.(type) {
	// 	case *exec.Error:
	// 		fmt.Println("Exec Error", err1)
	// 	case *exec.ExitError:
	// 		fmt.Println("Exec Exit Error", e.ExitCode())
	// 	default:
	// 		panic(err1)
	// 	}
	// }

	grepCommand := exec.Command("grep", "hello")

	grepIn, _ := grepCommand.StdinPipe()
	grepOut, _ := grepCommand.StdoutPipe()

	grepCommand.Start()
	grepIn.Write([]byte("hello there!\nTrying out spawning processes"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCommand.Wait()
	fmt.Println(">>>", string(grepBytes))



	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err2 := lsCmd.Output()
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("ls -a -l -h>>>", string(lsOut))
}