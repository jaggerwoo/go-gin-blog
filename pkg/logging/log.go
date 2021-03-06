package logging

import (
    "log"
    "os"
    "runtime"
    "path/filepath"
    "fmt"
)
type Level int

var (
    F *os.File
    DefaultPrefix = ""
    DefaultCallerDepth = 2
    logger *log.Logger
    logPrefix = ""
    levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
    DEBUG Level = iota
    INFO
    WARNING
    ERROR
    FATAL
)

func Setup() {
    filePath := getLogFileFullPath()
    F = openLogFile(filePath)

    // func New(out io.Writer, prefix string, flag int) *Logger {
	// 		return &Logger{out: out, prefix: prefix, flag: flag}
	// }

    logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
    setPrefix(DEBUG)
    logger.Println(v)
}

func Info(v ...interface{}) {
    setPrefix(INFO)
    logger.Println(v)
}

func Warn(v ...interface{}) {
    setPrefix(WARNING)
    logger.Println(v)
}

func Error(v ...interface{}) {
    setPrefix(ERROR)
    logger.Println(v)
}

func Fatal(v ...interface{}) {
    setPrefix(FATAL)
    logger.Fatalln(v)
}

func setPrefix(level Level) {
    _, file, line, ok := runtime.Caller(DefaultCallerDepth)
    if ok {
        logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
    } else {
        logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
    }

    logger.SetPrefix(logPrefix)
}