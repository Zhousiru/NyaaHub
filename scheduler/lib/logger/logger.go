package logger

import "fmt"

func Error(collection, msg string) {
	// TODO
	fmt.Println("[ ERR! ] [", collection, "]", msg)
}

func Info(collection, msg string) {
	// TODO
	fmt.Println("[ INFO ] [", collection, "]", msg)
}
