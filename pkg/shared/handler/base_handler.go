package handler

import (
	"bytes"
	"encoding/json"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/validator"
	"go-api/pkg/shared/wraperror"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/form"
	v "github.com/go-playground/validator/v10"
)

// BaseHTTPHandler base handler struct.
type BaseHTTPHandler struct {
	Logger infrastructure.Logger
}

// ParseJSON form struct.
// https://github.com/go-playground/form
func (h *BaseHTTPHandler) ParseJSON(c *gin.Context, i interface{}) ([]string, error) {
	// mapping post to struct.
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		return nil, wraperror.WithTrace(err, nil, nil)
	}
	s := buf.String()
	mesages, err := parseJSON(s, i)
	return mesages, wraperror.WithTrace(err, nil, nil)
}

// parseJSON function
func parseJSON(jsonString string, i interface{}) (messages []string, err error) {
	if !isJSON(jsonString) {
		messages = append(messages, "JSON input invalid.")
		return messages, wraperror.WithTrace(err, nil, nil)
	}
	var swapErr error
OUTER:
	swapErr = json.NewDecoder(strings.NewReader(jsonString)).Decode(&i)
	if swapErr != nil {
		err = swapErr
		if terr, ok := err.(*json.UnmarshalTypeError); ok {
			jsonString = strings.Replace(jsonString, terr.Field, "error"+terr.Field, -1)
			messages = append(messages, terr.Field+" is error.")
			goto OUTER
		}
	}
	return messages, err
}

// isJSON check a string data is JSON or not
func isJSON(input string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(input), &js) == nil
}

// ParseForm  form struct.
// https://github.com/go-playground/form
func (h *BaseHTTPHandler) ParseForm(c *gin.Context, i interface{}) error {
	// mapping post to struct.
	err := c.Request.ParseForm()
	if err != nil {
		return wraperror.WithTrace(err, nil, nil)
	}
	decoder := form.NewDecoder()
	err = decoder.Decode(&i, c.Request.Form)
	if err != nil {
		return wraperror.WithTrace(err, nil, nil)
	}
	return nil
}

// Validate func
func (h *BaseHTTPHandler) Validate(req interface{}) ([]string, error) {
	var ret []string
	mValidator := validator.New()
	err := mValidator.Struct(req)

	if err != nil {
		if _, ok := err.(*v.InvalidValidationError); ok {
			return ret, wraperror.WithTrace(err, nil, nil)
		}

		for _, errV := range err.(v.ValidationErrors) {
			ret = append(ret, errV.Field()+" is error.")
		}
	}
	return ret, err
}
