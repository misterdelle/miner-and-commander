package logger

import (
	"fmt"
	"log/slog"
	"os"
)

// var Logger *zap.Logger
var Logger *slog.Logger
var f *os.File

func InitializeLogger(logFile string) {

	// w := zapcore.AddSync(&lumberjack.Logger{
	// 	//		Filename:   "DEV_CaricaProvvigioniGGR.log",
	// 	Filename:   logFile,
	// 	MaxSize:    10, // megabytes
	// 	MaxBackups: 5,
	// 	MaxAge:     10,   // days
	// 	Compress:   true, // disabled by default
	// })

	// encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// core := zapcore.NewCore(
	// 	zapcore.NewConsoleEncoder(encoderConfig),
	// 	w,
	// 	zap.InfoLevel,
	// )

	// Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	var err error
	f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening logFile: %v", err)
	}

	Logger = slog.New(slog.NewTextHandler(f, nil))
	slog.SetDefault(Logger)
}

func CloseLoggerFile() {
	f.Close()
}
