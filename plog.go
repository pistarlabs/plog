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
		log *log.Logger
	}

	// Options is list of options used for Logger
	Options struct {
		Debug  bool
		Prefix string
	}
)

// New return logger instance with empty options
func New(opt Options) *Logger {
	l := &Logger{}

	if opt.Debug {
		l.log = log.New(os.Stdout, opt.Prefix, log.LstdFlags)
	}

	return l
}

// Default return logger instance with default options
func Default() *Logger {
	return New(Options{
		Debug:  true,
		Prefix: "[log] ",
	})
}

// Handler is a function to chain handler and log the activity
func (l *Logger) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Begin of time
		start := time.Now()
		// URL accessed
		path := r.URL

		// Wrapping default response writer to get status code
		w = NewResponseWriter(w)
		// Next handler
		h.ServeHTTP(w, r)
		res := w.(ResponseWriter)

		// End of time
		end := time.Now()
		// Latency
		latency := end.Sub(start)
		// Client IP Address
		clientIP := r.RemoteAddr
		// Method access
		method := r.Method

		l.log.Printf("|%d|%13v|%s|%s %s", res.Status(), latency, clientIP, method, path)
	})
}
