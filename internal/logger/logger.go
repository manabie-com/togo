package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// const prefix = "src/git.begroup.team/platform-core/"

// ConsoleEncoderName ...
const ConsoleEncoderName = "custom_console"

var (
	ll Logger
	xl Logger
)

// Logger wraps zap.Logger
type Logger struct {
	*zap.Logger
	S *zap.SugaredLogger
}

// PrintError prints all error with all meta data and line number.
// It's prefered to be used at top level function.
//
//     func DoSomething() (_err error) {
//         defer ll.PrintError("DoSomething", &_err)
//
func (logger Logger) PrintError(msg string, err *error) {
	if *err != nil {
		ll.S.Errorf("%v: %+v", msg, *err)
	}
}

// Short-hand functions for logging.
var (
	Any        = zap.Any
	Bool       = zap.Bool
	Duration   = zap.Duration
	Float64    = zap.Float64
	Int        = zap.Int
	Int64      = zap.Int64
	Skip       = zap.Skip
	String     = zap.String
	Stringer   = zap.Stringer
	Time       = zap.Time
	Uint       = zap.Uint
	Uint32     = zap.Uint32
	Uint64     = zap.Uint64
	Uintptr    = zap.Uintptr
	ByteString = zap.ByteString
)

// DefaultConsoleEncoderConfig ...
var DefaultConsoleEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalColorLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   ShortColorCallerEncoder,
}

// Error wraps error for zap.Error.
func Error(err error) zapcore.Field {
	if err == nil {
		return Skip()
	}
	return String("error", err.Error())
}

// Interface ...
func Interface(key string, val interface{}) zapcore.Field {
	if val, ok := val.(fmt.Stringer); ok {
		return zap.Stringer(key, val)
	}
	return zap.Reflect(key, val)
}

// Stack ...
func Stack() zapcore.Field {
	return zap.Stack("stack")
}

// Int32 ...
func Int32(key string, val int32) zapcore.Field {
	return zap.Int(key, int(val))
}

// Object ...
func Object(key string, val interface{}) zapcore.Field {
	//return zap.Any(key, val)
	return zap.Stringer(key, Dump(val))
}

type dd struct {
	v interface{}
}

func (d dd) String() string {
	return pp.Sprint(d.v)
}

// Dump renders object for debugging
func Dump(v interface{}) fmt.Stringer {
	return dd{v}
}

func trimPath(c zapcore.EntryCaller) string {
	// index := strings.Index(c.File, prefix)
	// if index < 0 {
	// 	return c.TrimmedPath()
	// }
	// return c.File[index+len(prefix):]
	return c.TrimmedPath()
}

// ShortColorCallerEncoder encodes caller information with sort path filename and enable color.
func ShortColorCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	callerStr := trimPath(caller) + ":" + strconv.Itoa(caller.Line)
	enc.AppendString(callerStr)
}

func NewWithName(name string, opts ...zap.Option) Logger {
	return newWithName(name, opts...)
}

func newWithName(name string, opts ...zap.Option) Logger {
	if name == "" {
		_, filename, _, _ := runtime.Caller(1)
		name = filepath.Dir(truncFilename(filename))
	}

	var enabler zap.AtomicLevel
	if e, ok := enablers[name]; ok {
		enabler = e
	} else {
		enabler = zap.NewAtomicLevel()
		enablers[name] = enabler
	}

	setLogLevelFromEnv(name, enabler)
	loggerConfig := zap.Config{
		Level:            enabler,
		Development:      false,
		Encoding:         ConsoleEncoderName, //"json",
		EncoderConfig:    DefaultConsoleEncoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	stacktraceLevel := zap.NewAtomicLevelAt(zapcore.PanicLevel)

	opts = append(opts, zap.AddStacktrace(stacktraceLevel))
	logger, err := loggerConfig.Build(opts...)
	if err != nil {
		panic(err)
	}
	return Logger{logger, logger.Sugar()}
}

// New returns new zap.Logger
func New(opts ...zap.Option) Logger {
	return newWithName("", opts...)
}

// ServeHTTP supports logging level with an HTTP request.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	type payload struct {
		Name  string `json:"name"`
		Level string `json:"level"`
	}

	enc := json.NewEncoder(w)

	switch r.Method {
	case "GET":
		var payloads []payload
		for k, e := range enablers {
			lvl := e.Level()
			payloads = append(payloads, payload{
				Name:  k,
				Level: lvl.String(),
			})
		}
		err := enc.Encode(payloads)
		if err != nil {
			panic(err)
		}

	case "PUT":
		var req payload
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = enc.Encode(errorResponse{
				Error: fmt.Sprintf("Request body must be valid JSON: %v", err),
			})
			if err != nil {
				panic(err)
			}
			return
		}

		if req.Level == "" {
			w.WriteHeader(http.StatusBadRequest)
			err := enc.Encode(errorResponse{
				Error: errLevelNil.Error(),
			})
			if err != nil {
				panic(err)
			}
			return
		}

		var lv zapcore.Level
		err := lv.UnmarshalText([]byte(req.Level))
		if err != nil {
			panic(err)
		}

		if req.Name == "" {
			for _, enabler := range enablers {
				enabler.SetLevel(lv)
			}

		} else {
			enabler, ok := enablers[req.Name]
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				err := enc.Encode(errorResponse{
					Error: errEnablerNotFound.Error(),
				})
				if err != nil {
					panic(err)
				}
				return
			}

			enabler.SetLevel(lv)
		}

		err = enc.Encode(req)
		if err != nil {
			panic(err)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := enc.Encode(errorResponse{
			Error: "Only GET and PUT are supported.",
		})
		if err != nil {
			panic(err)
		}
	}
}

var (
	errEnablerNotFound = errors.New("enabler not found")
	errLevelNil        = errors.New("must specify a logging level")

	enablers = make(map[string]zap.AtomicLevel)
)

func truncFilename(filename string) string {
	// index := strings.Index(filename, prefix)
	// return filename[index+len(prefix):]
	return filename
}

// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

var envPatterns []*regexp.Regexp

func init() {
	err := zap.RegisterEncoder(ConsoleEncoderName, func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return NewConsoleEncoder(cfg), nil
	})
	if err != nil {
		panic(err)
	}

	ll = New()
	xl = New(zap.AddCallerSkip(1))

	envLog := os.Getenv("LOG_LEVEL")
	if envLog == "" {
		return
	}

	var lv zapcore.Level
	err = lv.UnmarshalText([]byte(envLog))
	if err != nil {
		panic(err)
	}

	for _, enabler := range enablers {
		enabler.SetLevel(lv)
	}

	var errPattern string
	envPatterns, errPattern = initPatterns(envLog)
	if errPattern != "" {
		ll.Fatal("Unable to parse LOG_LEVEL. Please set it to a proper value.", String("invalid", errPattern))
	}

	ll.Info("Enable debug log", String("LOG_LEVEL", envLog))
}

func initPatterns(envLog string) ([]*regexp.Regexp, string) {
	patterns := strings.Split(envLog, ",")
	result := make([]*regexp.Regexp, len(patterns))
	for i, p := range patterns {
		r, err := parsePattern(p)
		if err != nil {
			return nil, p
		}

		result[i] = r
	}
	return result, ""
}

func parsePattern(p string) (*regexp.Regexp, error) {
	p = strings.Replace(strings.Trim(p, " "), "*", ".*", -1)
	return regexp.Compile(p)
}

func setLogLevelFromEnv(name string, enabler zap.AtomicLevel) {
	for _, r := range envPatterns {
		if r.MatchString(name) {
			enabler.SetLevel(zap.DebugLevel)
		}
	}
}
