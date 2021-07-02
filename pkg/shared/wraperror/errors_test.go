package wraperror

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	err = errors.New("Have error")
)

//TestWithTraceReturnNil func
func TestWithTraceReturnNil(t *testing.T) {
	err := WithTrace(nil, nil, nil)
	assert.Nil(t, err)
}

//TestWithTraceReturnError func
func TestWithTraceReturnError(t *testing.T) {
	err := WithTrace(err, nil, nil)
	assert.Error(t, err)
}

//TestGetStatusCode func
func TestGetStatusCode(t *testing.T) {
	testCases := []error{
		nil,
		err,
	}
	expectError := []int{
		http.StatusOK,
		http.StatusInternalServerError,
	}
	for i, testCase := range testCases {
		t.Run(strconv.Itoa(int(i)), func(t *testing.T) {
			err := GetStatusCode(testCase)
			assert.Equal(t, expectError[i], err)
		})
	}
}
