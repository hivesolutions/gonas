package gonas

import "io"
import "log"

type Logger interface {
    Trace(msg string)
    Info(msg string)
    Warning(msg string)
    Error(msg string)
    Tracef(format string, v ...interface{})
    Infof(format string, v ...interface{})
    Warningf(format string, v ...interface{})
    Errorf(format string, v ...interface{})
}

type AbstractLogger struct {
    TraceLog *log.Logger
    InfoLog *log.Logger
    WarningLog *log.Logger
    ErrorLog *log.Logger
}

func (srv *AbstractLogger) InitLog(
    traceH io.Writer,
    infoH io.Writer,
    warningH io.Writer,
    errorH io.Writer) error {
    srv.TraceLog = log.New(traceH,  "[TRACE] ", log.Ldate|log.Ltime)
    srv.InfoLog = log.New(infoH, "[INFO] ", log.Ldate|log.Ltime)
    srv.WarningLog = log.New(warningH, "[WARNING] ", log.Ldate|log.Ltime)
    srv.ErrorLog = log.New(errorH, "[ERROR] ", log.Ldate|log.Ltime)
    return nil
}

func (logger *AbstractLogger) Trace(msg string) {
    logger.TraceLog.Println(msg)
}

func (logger *AbstractLogger) Tracef(format string, v ...interface{}) {
    logger.TraceLog.Printf(format, v)
}

func (logger *AbstractLogger) Info(msg string) {
    logger.InfoLog.Println(msg)
}

func (logger *AbstractLogger) Infof(format string, v ...interface{}) {
    logger.InfoLog.Printf(format, v)
}

func (logger *AbstractLogger) Warning(msg string) {
    logger.WarningLog.Println(msg)
}

func (logger *AbstractLogger) Warningf(format string, v ...interface{}) {
    logger.WarningLog.Printf(format, v)
}

func (logger *AbstractLogger) Error(msg string) {
    logger.ErrorLog.Println(msg)
}

func (logger *AbstractLogger) Errorf(format string, v ...interface{}) {
    logger.ErrorLog.Printf(format, v)
}
