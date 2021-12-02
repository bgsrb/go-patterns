package main

import "fmt"

//RedisLogger
type RedisLogger struct {
}

func (c *RedisLogger) Log(msg string) {
	fmt.Printf("RedisLogger Log : %s \r\n", msg)
}

func NewRedisLogger() *RedisLogger {
	return &RedisLogger{}
}

//ConsoleLogger
type ConsoleLogger struct {
}

func (c *ConsoleLogger) Log(msg string) {
	fmt.Printf("ConsoleLogger Log : %s \r\n", msg)
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

//DBLogger
type DBLogger struct {
}

func (d *DBLogger) Log(msg string) {
	fmt.Printf("DBLogger Log : %s \r\n", msg)
}

func NewDBLogger() *DBLogger {
	return &DBLogger{}
}

//Logger
type Logger interface {
	Log(msg string)
}

type logger struct {
	loggers []Logger
}

func (l *logger) Log(msg string) {
	for _, l := range l.loggers {
		l.Log(msg)
	}
}

func (l *logger) AddLoger(logger Logger) {
	l.loggers = append(l.loggers, logger)
}

func NewLogger(loggers ...Logger) *logger {
	return &logger{
		loggers: loggers,
	}
}

func main() {
	var l = NewLogger(NewConsoleLogger(), NewDBLogger())
	l.AddLoger(NewRedisLogger())
	l.Log("123")
}
