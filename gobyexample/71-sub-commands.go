package main

import (
	"os"
	"flag"
	"fmt"
)


func main() {

	jetCmd := flag.NewFlagSet("jet", flag.ExitOnError) 
	jetPort := jetCmd.Int("port", 3000, "port number")

	lagCmd := flag.NewFlagSet("lag", flag.ExitOnError)
	lagEnv := lagCmd.String("env", "development", "env of script running")

	if len(os.Args) < 2 {
		fmt.Println(os.Args)
		fmt.Println("Expected jet or lag commands")
		os.Exit(1)
	}

	fmt.Println(">>>", os.Args[1])

	switch os.Args[1] {
	case "jet":
		jetCmd.Parse(os.Args[2:])
		fmt.Println(*jetPort)
		fmt.Println(jetCmd.Args())
		
	case "lag":
		lagCmd.Parse(os.Args[2:])
		fmt.Println(*lagEnv)
		fmt.Println(lagCmd.Args())
	}
}