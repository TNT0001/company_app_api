package wraperror

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// Fields are attached to error
type Fields map[string]interface{}

// DetailTrace wraps error with detail
type DetailTrace interface {
	GetTraces() string
	GetTraceSource() string
	GetStackTrace() []string
	GetFields() Fields
}

type detailTrace struct {
	tracer     TraceHistory
	stackTrace []string
	fields     Fields
}

func (e *detailTrace) GetTraces() string {
	result := ""

	if e.tracer != nil {
		result = e.tracer.Traces()
	}

	return result
}

func (e *detailTrace) GetTraceSource() string {
	result := ""

	if e.tracer != nil {
		result = e.tracer.TraceSource()
	}

	return result
}

func (e *detailTrace) GetStackTrace() []string {
	return e.stackTrace
}

func (e *detailTrace) GetFields() Fields {
	return e.fields
}

type withTrace struct {
	error
	DetailTrace
}

// TraceHistory is the interface that can get trace data with.
type TraceHistory interface {
	Traces() string
	AddTraces(traces ...interface{})
	TraceSource() string
	SetTraceSource(source string)
}

type stack []uintptr

// WithTrace binds error and trace history into an error.
func WithTrace(err error, fields Fields, tracer TraceHistory) error {
	if err == nil {
		return nil
	}

	stackTrace := []string{}
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])
	var st stack = pcs[0:n]

	for _, pc := range st {
		f := errors.Frame(pc)
		fStr := fmt.Sprintf("%+v", f)
		fStr = strings.Replace(fStr, "\n\t", " in ", -1)
		stackTrace = append(stackTrace, fStr)
	}

	dt := &detailTrace{
		tracer:     tracer,
		fields:     fields,
		stackTrace: stackTrace,
	}

	return &withTrace{err, dt}
}

// GetStatusCode of error
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	default:
		return http.StatusInternalServerError
	}
}
