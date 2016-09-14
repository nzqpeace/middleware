package logger

import (
	"encoding/json"
	"strconv"
	"strings"
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

	if !l.config.Body || len(body) > l.config.MaxLenToPrint {
		body = body[:l.config.MaxLenToPrint]
	} else {
		b := ctx.RequestCtx.Request.Body()
		contentType := string(ctx.GetRequestCtx().Request.Header.ContentType())
		if len(b) > 0 && strings.ToLower(contentType) == "application/json" {
			var p interface{}
			err := json.Unmarshal(b, &p)
			if err == nil {
				b, err = json.Marshal(p)
				if err == nil {
					body = string(b)
				}
			}
		}
	}
	//finally print the logs
	ctx.Log("[%s] %s %v %4v %s %s %s %s\n", requestID, date, status, latency, ip, method, path, body)

}

// New returns the logger middleware
// receives optional configs(logger.Config)
func New(cfg ...Config) iris.HandlerFunc {
	c := DefaultConfig().Merge(cfg)
	l := &loggerMiddleware{config: c}

	return l.Serve
}
