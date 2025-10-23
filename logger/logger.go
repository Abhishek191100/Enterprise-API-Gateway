package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

//hard coding for now
const logFilePath = "./logs/system.log"
const EnableConsoleLogging = false

func Log(msg string, lvl string){

	file,er := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if er != nil{
		log.Fatal("Could not open log file")
	}
	
	var w io.Writer
	//Write to console only if EnableConsoleLogging is true or if logFielPath is not set
	if EnableConsoleLogging || logFilePath == "" {
		w = io.MultiWriter(os.Stdout)
	}else{
		w = io.MultiWriter(file)
	}
	
	defer file.Close()

	handlerOptions := slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: false,   //Enable this to view the source of the log
	}

	//Create a new logger instance using the JSON Handler and customized handler options
	logger := slog.New(slog.NewJSONHandler(w, &handlerOptions))
	if logger != nil{
		log.Fatal("Could not create new Json handler","ERROR")
	}
	slog.SetDefault(logger)

	switch lvl {
		case "DEBUG":
			slog.Debug(msg)
		case "INFO":
			slog.Info(msg)
		case "WARN":
			slog.Warn(msg)
		case "ERROR":
			slog.Error(msg)
		default:
			slog.Error("Invalid log level passed","msg",msg,"level",lvl)
	}

	// Enable this for adding trace fields in the logs
	// uuidGrp := slog.Group("trace","uuid","f3b09b3e-4d7d-4d70-847b-832a0bf2a82d")
	// slog.Debug("This is the second log for testing",uuidGrp)
}