package logs

import (
	"fmt"
	"log"
	"os"
)

const (
	//Trace constant is for Trace logging inside logs/trace.txt
	Trace = "Trace"

	//Info constant is for Info logging inside logs/info.txt
	Info = "Info"

	//Warning constant is for Warning logging inside logs/warning.txt
	Warning = "Warning"

	//Error constant is for Error logging inside logs/error.txt
	Error = "Error"
)

var (
	//LogTrace is for trace logging
	LogTrace *log.Logger

	//LogInfo is for info logging
	LogInfo *log.Logger

	//LogWarning is for warning logging
	LogWarning *log.Logger

	//LogError is for error logging
	LogError *log.Logger
)

func init() {

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModeDir)
		fmt.Println("go_log pkg: logs folder created, logs files will be created")
	} else {
		fmt.Println("go_log pkg: logs folder exist, logs will be appended")
	}

	traceFile, err := os.OpenFile("logs/trace.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	infoFile, err := os.OpenFile("logs/info.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	warningFile, err := os.OpenFile("logs/warning.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errorFile, err := os.OpenFile("logs/error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Unable to create|open file")
		return
	}

	LogTrace = log.New(traceFile, "Trace: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	LogInfo = log.New(infoFile, "Info: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	LogWarning = log.New(warningFile, "Warning: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	LogError = log.New(errorFile, "Error: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
}

//Log is used for logging with level
//For simple logging use Log* methods
func Log(logLevel string, message string) {
	switch logLevel {
	case Trace:
		LogTrace.Printf(message)
	case Info:
		LogInfo.Printf(message)
	case Warning:
		LogWarning.Printf(message)
	case Error:
		LogError.Printf(message)
	default:
		fmt.Printf(message)
	}
}
