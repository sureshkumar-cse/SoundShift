package utils

import (
	"log"
	"os"
)

// InitLogger initializes a log file
func InitLogger(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)
	log.Println("Logger initialized.")
	return file, nil
}
