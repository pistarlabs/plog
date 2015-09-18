package plog

import (
	"log"
	"net/http"
	"os"
	"time"
)

type (
	// Logger is logger
	Logger struct {
		// Debug logger
		log *log.Logger
	}

	// LoggerOptions is list of options used for Logger
	LoggerOptions struct {
		// Debug state if logger should output log in server side
		Debug bool
	}
)

func New(options LoggerOptions) *Logger {
	l := &Logger{}

	if options.Debug {
		l.log = log.New(os.Stdout, "[log] ", log.LstdFlags)
	}

	return l
}

func Default() *Logger {
	return New(LoggerOptions{
		Debug: true,
	})
}

func (l *Logger) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL

		w = NewResponseWriter(w)

		h.ServeHTTP(w, r)

		res := w.(ResponseWriter)
		end := time.Now()
		latency := end.Sub(start)
		clientIP := r.RemoteAddr
		method := r.Method

		l.log.Printf("|%d|%13v|%s|%s %s", res.Status(), latency, clientIP, method, path)
	})
}
