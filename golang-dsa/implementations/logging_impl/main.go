package main

import (
	"log"
	"logging_impl/logger"
)

func main() {
	slog := logger.Constructor("output.txt")
	slog.Info("This is Naveen")
	slog.Warn("There is something wrong here")
	slog.Error("Everything is breaking")
	log.Println("some error")
}
