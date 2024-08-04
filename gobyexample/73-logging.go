package main

import (
	"log"
	"os"
	"bytes"
	"fmt"
	"log/slog"
)

func main() {
	log.Println("Standard logging.")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetPrefix("AuthService:")
	log.Println("With microseconds")
	
	
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("With file:line")


	redisLog := log.New(os.Stdout, "redis:", log.LstdFlags)
	redisLog.Println("Service up at port 3002.")
	
	redisLog.SetPrefix("memcache:")
	redisLog.Println("Service up at port 3002.")


	var buf bytes.Buffer
	bufLog := log.New(&buf, "bufLog:", log.LstdFlags)
	bufLog.Println("WHat's good buf gang?") // this will log into var buf and not os.Stdout ie terminal here

	fmt.Println(buf.String())


	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	slogLogger := slog.New(jsonHandler)

	slogLogger.Info("hi there!")
	slogLogger.Info("hello again!", "key", "val", "age", "name", "aditya") // this will give key name for aditya val as !BADKEY since odd number of values were passed
}