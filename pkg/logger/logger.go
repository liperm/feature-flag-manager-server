package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

type CustomLogger struct {
	logger *log.Logger
}

var Logger CustomLogger

func Init() {
	generalLog, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	multiWritter := io.MultiWriter(generalLog, os.Stdout)
	Logger.logger = log.New(multiWritter, "Log:\t", log.Ldate|log.Ltime)
}

func (l *CustomLogger) Request(r interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		method := getCallerFunctionName(details)
		l.write("Request", method, r)
		return
	}

	l.write("Request", "UKNOWN", r)
}

func (l *CustomLogger) Response(r interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		method := getCallerFunctionName(details)
		l.write("Response", method, r)
		return
	}

	l.write("Response", "UKNOWN", r)
}

func (l *CustomLogger) Error(e error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		method := getCallerFunctionName(details)
		formattedMethod := fmt.Sprintf("[%s]:", method)
		l.logger.Println(formattedMethod, "ERROR", e.Error())
	}

	formattedMethod := fmt.Sprintf("[%s]:", "UNKNOWN")
	l.logger.Println(formattedMethod, "ERROR", e.Error())
}

func (l *CustomLogger) write(ctx string, method string, r interface{}) {
	formattedMethod := fmt.Sprintf("[%s]:", method)
	requestJson, err := json.Marshal(r)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error logging %s", ctx), err)
		return
	}
	l.logger.Println(formattedMethod, fmt.Sprintf("%s Received", ctx), string(requestJson))
}

func getCallerFunctionName(details *runtime.Func) string {
	splitDetails := strings.Split(details.Name(), ".")
	functionName := splitDetails[len(splitDetails)-1]
	return functionName
}
