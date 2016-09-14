package logger

import (
	"strconv"
	"time"

	"github.com/kataras/iris"
)

type loggerMiddleware struct {
	config Config
}

// Serve serves the middleware
func (l *loggerMiddleware) Serve(ctx *iris.Context) {
	//all except latency to string
	var date, status, ip, method, path, body, requestID string
	var latency time.Duration
	var startTime, endTime time.Time
	path = ctx.RequestPath(false)
	method = ctx.MethodString()
	requestID = strconv.FormatUint(ctx.GetRequestCtx().ID(), 10)
	body = bytes.NewBuffer(ctx.RequestCtx.Request.Body()).String()

	startTime = time.Now()

	ctx.Next()
	//no time.Since in order to format it well after
	endTime = time.Now()
	date = endTime.Format("01/02 - 15:04:05")
	latency = endTime.Sub(startTime)

	if l.config.Status {
		status = strconv.Itoa(ctx.Response.StatusCode())
	}

	if l.config.IP {
		ip = ctx.RemoteAddr()
	}

	if !l.config.Method {
		method = ""
	}

	if !l.config.Path {
		path = ""
	}

	if !l.config.Date {
		date = ""
	}

	if !l.config.RequestID {
		requestID = ""
	}

	if !l.config.Body {
		body = ""
	}
	//finally print the logs
	ctx.Log("[%d] %s %v %4v %s %s %s %s\n", requestID, date, status, latency, ip, method, path, body)

}

// New returns the logger middleware
// receives optional configs(logger.Config)
func New(cfg ...Config) iris.HandlerFunc {
	c := DefaultConfig().Merge(cfg)
	l := &loggerMiddleware{config: c}

	return l.Serve
}
