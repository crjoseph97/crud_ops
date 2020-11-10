package app

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"path"
// 	"runtime"
// 	"strconv"
// 	"time"

// 	"git.hifx.in/crud_ops/domain"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/gommon/color"
// 	glog "github.com/labstack/gommon/log"
// )

// type logger struct {
// 	appName     string
// 	level       glog.Lvl
// 	internelLog *log.Logger
// 	output      io.Writer
// 	logChan     chan combinedLog
// }

// // Lvls holds array of log levels
// var Lvls = [8]string{"-", "debug", "info", "warn", "error", "off", "panic", "fatal"}

// const defaultStackSkip = 0

// // NewLogger inits the logger function
// func NewLogger(appName string, outputFile string) (echo.Logger, error) {
// 	w, err := OpenLogFile(outputFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	l := &logger{
// 		level:   glog.INFO,
// 		appName: appName,
// 		output:  w,
// 		logChan: make(chan combinedLog),
// 	}

// 	go logWriter(l.logChan, l.output)
// 	return l, nil
// }

// func (l *logger) DisableColor() {}

// func (l *logger) EnableColor() {}

// func (l *logger) Prefix() string {
// 	return l.appName
// }

// func (l *logger) SetPrefix(p string) {
// 	l.appName = p
// }

// func (l *logger) Level() glog.Lvl {
// 	return l.level
// }

// func (l *logger) SetLevel(level glog.Lvl) {
// 	l.level = level
// }

// func (l *logger) Output() io.Writer {
// 	return l.output
// }

// func (l *logger) SetOutput(w io.Writer) {
// 	l.output = w
// }

// func (l *logger) Color() *color.Color {
// 	return nil
// }

// func (l *logger) SetHeader(h string) {}

// func (l *logger) Print(i ...interface{}) {
// 	l.log(glog.INFO, "", i...)
// }

// func (l *logger) Printf(format string, args ...interface{}) {
// 	l.log(glog.INFO, format, args...)
// }

// func (l *logger) Printj(j glog.JSON) {
// 	l.log(glog.INFO, "json", j)
// }

// func (l *logger) Debug(i ...interface{}) {
// 	l.log(glog.DEBUG, "", i...)
// }

// func (l *logger) Debugf(format string, args ...interface{}) {
// 	l.log(glog.DEBUG, format, args...)
// }

// func (l *logger) Debugj(j glog.JSON) {
// 	l.log(glog.DEBUG, "json", j)
// }

// func (l *logger) Info(i ...interface{}) {
// 	l.log(glog.INFO, "", i...)
// }

// func (l *logger) Infof(format string, args ...interface{}) {
// 	l.log(glog.INFO, format, args...)
// }

// func (l *logger) Infoj(j glog.JSON) {
// 	l.log(glog.INFO, "json", j)
// }

// func (l *logger) Warn(i ...interface{}) {
// 	l.log(glog.WARN, "", i...)
// }

// func (l *logger) Warnf(format string, args ...interface{}) {
// 	l.log(glog.WARN, format, args...)
// }

// func (l *logger) Warnj(j glog.JSON) {
// 	l.log(glog.WARN, "json", j)
// }

// func (l *logger) Error(i ...interface{}) {
// 	l.log(glog.ERROR, "", i...)
// }

// func (l *logger) Errorf(format string, args ...interface{}) {
// 	l.log(glog.ERROR, format, args...)
// }

// func (l *logger) Errorj(j glog.JSON) {
// 	l.log(glog.ERROR, "json", j)
// }

// // Panic is used to report panic. there is no need to use it in our code except the Recover middleware.
// func (l *logger) Panic(i ...interface{}) {
// 	l.log(6, "", i...)
// }

// func (l *logger) Panicf(format string, args ...interface{}) {
// 	l.log(6, format, args...)
// }

// func (l *logger) Panicj(j glog.JSON) {
// 	l.log(6, "json", j)
// }

// func (l *logger) Fatal(i ...interface{}) {
// 	l.log(7, "", i...)
// 	os.Exit(1)
// }

// func (l *logger) Fatalf(format string, args ...interface{}) {
// 	l.log(7, format, args...)
// 	os.Exit(1)
// }

// func (l *logger) Fatalj(j glog.JSON) {
// 	l.log(7, "json", j)
// 	os.Exit(1)
// }

// type combinedLog struct {
// 	Time      string      `json:"time,omitempty"`
// 	AppName   string      `json:"app_name,omitempty"`
// 	Level     string      `json:"level,omitempty"`
// 	AdminID   interface{} `json:"admin_id,omitempty"`
// 	AdminName interface{} `json:"admin_name,omitempty"`

// 	RequestID    interface{} `json:"request_id,omitempty"`
// 	RemoteIP     string      `json:"remote_ip,omitempty"`
// 	Host         string      `json:"host,omitempty"`
// 	URI          string      `json:"uri,omitempty"`
// 	Method       string      `json:"method,omitempty"`
// 	Path         string      `json:"path,omitempty"`
// 	Protocol     string      `json:"protocol,omitempty"`
// 	Referer      string      `json:"referer,omitempty"`
// 	UserAgent    string      `json:"user_agent,omitempty"`
// 	Status       int         `json:"status,omitempty"`
// 	Latency      string      `json:"latency,omitempty"`
// 	LatencyHuman string      `json:"latency_human,omitempty"`
// 	BytesIn      int64       `json:"bytes_in,omitempty"`
// 	BytesOut     int64       `json:"bytes_out,omitempty"`

// 	Query    interface{} `json:"query,omitempty"`
// 	ESIndexs interface{} `json:"es_indexs,omitempty"`

// 	FileName   string      `json:"file_name,omitempty"`
// 	Error      string      `json:"error,omitempty"`
// 	ErrorStack interface{} `json:"error_stack,omitempty"`
// 	Message    string      `json:"message,omitempty"`
// }

// // HACK: try to send context as the first arugument, to print the context details in logging.
// func (l *logger) log(level glog.Lvl, format string, args ...interface{}) {
// 	var (
// 		index                         = 0
// 		errStack                      interface{}
// 		message, errStr               string
// 		fileName                      string
// 		AdminID, AdminName, RequestID interface{}
// 	)

// 	// if the first element is echo
// 	c, ok := args[index].(echo.Context)
// 	if ok {
// 		AdminID = c.Get(domain.ADMINID)
// 		AdminName = c.Get(domain.ADMINNAME)
// 		RequestID = c.Get(domain.REQUESTID)
// 		index = 1
// 	}

// 	if level == glog.ERROR && len(args[index:]) > 0 {
// 		errStr, errStack = l.buildErr(format, args[index])
// 		message = "reporting an unexpected error."
// 	} else if level == 6 && len(args[index:]) > 0 {
// 		errStr, errStack = l.buildPanicErr(format, args[index:]...)
// 		message = "reporting a crash error."
// 	} else {
// 		switch format {
// 		case "":
// 			if len(args[index:]) > 0 {
// 				message = fmt.Sprintf("%v", args[index:]...)
// 			}

// 		case "json":
// 			b, err := json.Marshal(args[index])
// 			l.logChan <- combinedLog{
// 				Time:      time.Now().Format(time.RFC3339),
// 				Level:     Lvls[6],
// 				AdminID:   AdminID,
// 				AdminName: AdminName,
// 				RequestID: RequestID,
// 				Error:     err.Error(),
// 				Message:   "unexpected error while logging.",
// 			}

// 			message = string(b)

// 		default:
// 			if len(args[index:]) > 0 {
// 				message = fmt.Sprintf("%v", args[index:]...)
// 			}
// 		}
// 	}

// 	_, file, line, _ := runtime.Caller(defaultStackSkip)
// 	fileName = path.Clean(file) + ":" + strconv.Itoa(line)
// 	l.logChan <- combinedLog{
// 		Time:       time.Now().Format(time.RFC3339),
// 		AdminID:    AdminID,
// 		AdminName:  AdminName,
// 		RequestID:  RequestID,
// 		Level:      Lvls[level],
// 		AppName:    l.appName,
// 		Error:      errStr,
// 		ErrorStack: errStack,
// 		FileName:   fileName,
// 		Message:    message,
// 	}
// }

// func (l *logger) buildPanicErr(format string, args ...interface{}) (errStr string, errTrace interface{}) {
// 	if format != "" {
// 		errStr = fmt.Sprintf(format, args...)
// 	}

// 	if format == "" {
// 		errStr = fmt.Sprint(args...)
// 	}

// 	return errStr, getStackTrace(defaultStackSkip, 10)
// }

// func (l *logger) buildErr(format string, args ...interface{}) (string, interface{}) {
// 	var (
// 		errtrace      string
// 		errStr        string
// 		errtraceArray = []string{}
// 	)

// 	if format != "" {
// 		if format == "json" && args[0] != nil {
// 			args[0], _ = json.Marshal(args[0])
// 		}

// 		errStr = fmt.Sprintf(format, args...)
// 	} else {
// 		for _, arg := range args {
// 			// the argument contains the error, then get message and stacktrace from it.
// 			if err, ok := arg.(error); ok {
// 				errStr += err.Error()
// 				if ego, ok := err.(*errgo.Err); ok {
// 					if conf.Config.Mode == conf.MODEDEV {
// 						errtraceArray = append(errtraceArray, ego.Stack(), "")
// 					} else {
// 						errtrace += errgo.ErrorStack(ego) + "\n\n"
// 					}
// 				}

// 				continue
// 			}

// 			errStr += fmt.Sprint(arg)
// 		}
// 	}

// 	if conf.Config.Mode == conf.MODEDEV {
// 		return errStr, errtraceArray
// 	}

// 	return errStr, errtrace
// }

// func getStackTrace(skip, limit int) (stackTrace interface{}) {
// 	// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
// 	pc := make([]uintptr, limit)
// 	n := runtime.Callers(skip, pc)
// 	if n == 0 {
// 		// No pcs available. Stop now.
// 		// This can happen if the first argument to runtime.Callers is large.
// 		return "no stack traces available."
// 	}

// 	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
// 	frames := runtime.CallersFrames(pc)

// 	// Loop to get frames.
// 	// A fixed number of pcs can expand to an indefinite number of Frames.

// 	// we are skipping the first frame here..
// 	frame, more := frames.Next()
// 	var stacktrace string
// 	stackArray := make([]string, 0, n)
// 	for more {
// 		frame, more = frames.Next()

// 		// In DevMode will will send the stack traces as an  array to make it looks pretty
// 		if conf.Config.Mode == conf.MODEDEV {
// 			stackArray = append(stackArray, fmt.Sprintf("%v:%d", frame.File, frame.Line))
// 		} else {
// 			stacktrace += fmt.Sprintf("%s \n\t %v:%d \n", frame.Function, frame.File, frame.Line)
// 		}
// 	}

// 	if conf.Config.Mode == conf.MODEDEV {
// 		return stackArray
// 	}

// 	return stackTrace
// }

// // CombinedLogger combines both the access and error logging functionalities.
// func CombinedLogger() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			var (
// 				req       = c.Request()
// 				res       = c.Response()
// 				level     = Lvls[c.Logger().Level()]
// 				errStr    string
// 				stacktace interface{}
// 				message   string = "request has been processed successfully."
// 				err       error
// 				l         = c.Logger().(*logger)
// 			)

// 			start := time.Now()
// 			err = next(c)
// 			latency := time.Since(start)
// 			if err != nil {
// 				level = Lvls[glog.ERROR]
// 				errStr, stacktace = l.buildErr("%v", err)
// 				message = "request has failed because of an error."
// 			}

// 			var cl int64
// 			if req.Header.Get(echo.HeaderContentLength) != "" {
// 				cl, err = strconv.ParseInt(req.Header.Get(echo.HeaderContentLength), 10, 64)
// 				if err != nil {
// 					l.logChan <- combinedLog{
// 						Time:      time.Now().Format(time.RFC3339),
// 						Level:     Lvls[6],
// 						AdminID:   c.Get(domain.ADMINID),
// 						AdminName: c.Get(domain.ADMINNAME),
// 						RequestID: c.Get(domain.REQUESTID),
// 						Error:     err.Error(),
// 						Message:   "unexpected error while logging. unable to read HeaderContentLength",
// 					}
// 				}
// 			}
// 			l.logChan <- combinedLog{
// 				Time:         time.Now().Format(time.RFC3339),
// 				Level:        level,
// 				AdminID:      c.Get(domain.ADMINID),
// 				AdminName:    c.Get(domain.ADMINNAME),
// 				Query:        c.Get(domain.QUERY),
// 				ESIndexs:     c.Get(domain.ESINDEX),
// 				RequestID:    c.Get(domain.REQUESTID),
// 				AppName:      c.Logger().Prefix(),
// 				Error:        errStr,
// 				ErrorStack:   stacktace,
// 				RemoteIP:     c.RealIP(),
// 				Host:         req.Host,
// 				URI:          req.RequestURI,
// 				Method:       req.Method,
// 				Path:         req.URL.Path,
// 				Protocol:     req.Proto,
// 				Referer:      req.Referer(),
// 				UserAgent:    req.UserAgent(),
// 				Status:       res.Status,
// 				Latency:      strconv.FormatInt(int64(latency), 10),
// 				LatencyHuman: latency.String(),
// 				BytesIn:      cl,
// 				BytesOut:     res.Size,
// 				Message:      message,
// 			}

// 			return err
// 		}
// 	}
// }

// func logWriter(logchan chan combinedLog, w io.Writer) {
// 	for log := range logchan {
// 		var (
// 			d   []byte
// 			err error
// 		)

// 		if conf.Config.Mode == conf.MODEDEV && log.Error != "" {
// 			d, err = json.MarshalIndent(log, "", "")
// 		} else {
// 			d, err = json.Marshal(log)
// 		}

// 		if err != nil {
// 			logchan <- combinedLog{
// 				Time:      time.Now().Format(time.RFC3339),
// 				Level:     Lvls[6],
// 				AdminID:   log.AdminID,
// 				AdminName: log.AdminName,
// 				RequestID: log.RequestID,
// 				Error:     err.Error(),
// 				Message:   "unexpected error while logging. unable to log provided details.",
// 			}

// 			return
// 		}

// 		// // TODO: need to write a good implementation for text logging.
// 		if !conf.Config.HasJSONLogging() {
// 			d = d[1 : len(d)-1]
// 		}

// 		d = append(d, []byte("\n")...)
// 		w.Write(d)
// 	}
// }
