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
	method := getCallerFunctionName()
	l.write("Request", method, r)
	return
}

func (l *CustomLogger) RequestToMethod(method string, r interface{}) {
	l.write("Request", method, r)
}

func (l *CustomLogger) Response(r interface{}) {
	method := getCallerFunctionName()
	l.write("Response", method, r)
}

func (l *CustomLogger) ResponseFromMethod(method string, r interface{}) {
	l.write("Response", method, r)
}

func (l *CustomLogger) Error(e error) {
	method := getCallerFunctionName()
	formattedMethod := l.formatMethod(method)
	l.logger.Println(formattedMethod, "ERROR", e.Error())
}

func (l *CustomLogger) ErrorFromMethod(method string, e error) {
	formattedMethod := l.formatMethod(method)
	l.logger.Println(formattedMethod, "ERROR", e.Error())
}
func (l *CustomLogger) formatMethod(method string) string {
	formattedMethod := fmt.Sprintf("[%s]:", method)
	return formattedMethod
}

func (l *CustomLogger) write(ctx string, method string, r interface{}) {
	formattedMethod := l.formatMethod(method)
	requestJson, err := json.Marshal(r)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error logging %s", ctx), err)
		return
	}
	l.logger.Println(formattedMethod, fmt.Sprintf("%s Received", ctx), string(requestJson))
}

func getCallerFunctionName() string {
	functionName := "UNKNOWN"

	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		splitDetails := strings.Split(details.Name(), ".")
		functionName = splitDetails[len(splitDetails)-1]
	}

	return functionName
}
