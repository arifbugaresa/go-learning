package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var Logger *logrus.Logger

const maxStack = 5

type ErrorField struct {
	File string `json:"file"`
	Line int    `json:"line"`
	Func string `json:"func"`
}

func Initiator() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "function.name",
		},
	})

	// check log directory
	panicDir := viper.GetString("storage.log.panic")
	if panicDir == "" {
		log.Fatal("panicDir is not set in the configuration")
	}
	if _, err := os.Stat(panicDir); os.IsNotExist(err) {
		err = os.MkdirAll(panicDir, 0777)
		if err != nil {
			panic(err)
		}
	}

	errorDir := viper.GetString("storage.log.error")
	if _, err := os.Stat(errorDir); os.IsNotExist(err) {
		err = os.MkdirAll(errorDir, 0777)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(viper.GetString("storage.log.error")+"/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(fmt.Errorf("error open file : %v", err))
		return
	}

	Logger.Out = f
}

func ErrorWithCtx(ctx *gin.Context, fields map[string]interface{}, args ...interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}

	fields["trace_id"], _ = ctx.Get("trace_id")
	fields["mode"] = viper.GetString("app.mode")
	fields["at"] = time.Now()
	fields["stack"] = getStackTrace()

	entry := Logger.WithFields(fields)
	entry.Error(args...)
}

func getStackTrace() []ErrorField {
	var (
		trace []ErrorField
		field ErrorField
	)

	stackBuf := make([]uintptr, maxStack)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			field.File = frame.File
			field.Line = frame.Line
			field.Func = frame.Function
			trace = append(trace, field)
		}
		if !more {
			break
		}
	}
	return trace
}
