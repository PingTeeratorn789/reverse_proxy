package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/PingTeeratorn789/reverse_proxy/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logFileMutex sync.Mutex
	logger       *zap.Logger
	file         *os.File
)

type Logger struct{}
type Dependencies struct {
	Configs *configs.Config
}

func (l *Logger) Close() {
	logger.Sync()
	file.Close()
}

func (l *Logger) GetLogger() *zap.Logger {
	return logger
}

func GetConfig() *zap.Logger {
	return logger
}

func NewLogger(d Dependencies) *Logger {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "",
		CallerKey:      "caller",
		FunctionKey:    "",
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName: func(message string, enc zapcore.PrimitiveArrayEncoder) {
		},
	}

	// สร้าง output ไฟล์ log และ console
	file = createLogFile(d.Configs.App.Log)
	fileOutput := zapcore.AddSync(file)
	consoleOutput := zapcore.Lock(os.Stdout)

	// สร้าง Core สำหรับ logger ด้วย encoder และ output
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(fileOutput, consoleOutput),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	).With(
		[]zapcore.Field{
			zap.String("service-name", d.Configs.App.Name),
		},
	)

	// สร้าง global logger
	logger = zap.New(core, zap.AddCaller())
	return &Logger{}
}

func createLogFile(pathFolder string) *os.File {

	// // ตรวจสอบว่า folder "logs" มีอยู่หรือไม่
	if _, err := os.Stat(pathFolder); os.IsNotExist(err) {
		// ถ้าไม่มี ให้สร้าง folder "logs"
		err := os.Mkdir(pathFolder, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}

	logfile := path.Join(pathFolder, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
