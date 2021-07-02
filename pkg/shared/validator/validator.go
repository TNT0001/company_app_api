package validator

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	v "github.com/go-playground/validator/v10"
)

const (
	enumImportantTypeTag        = "enumImportantType"
	enumPeriodTag               = "enumPeriod"
	enumLearningPurposeTag      = "enumLearningPurpose"
	enumSensitivityRecordingTag = "enumSensitivityRecording"
	enumSpeechRateTag           = "enumSpeechRate"
	enumTypeTag                 = "enumType"
	enumPlanTypeTag             = "enumPlanType"
	timeFormat                  = "timeFormat"
	password                    = "password"
	urlcustom                   = "urlcustom"
)

var tagFuncMaps = map[string]func(v.FieldLevel) bool{
	enumImportantTypeTag:        IsEnumIntType,
	enumPeriodTag:               IsEnumIntType,
	enumLearningPurposeTag:      IsEnumIntType,
	enumSensitivityRecordingTag: IsEnumIntType,
	enumSpeechRateTag:           IsEnumIntType,
	enumTypeTag:                 IsEnumIntType,
	enumPlanTypeTag:             IsEnumIntType,
	timeFormat:                  DateTimeValidator,
	password:                    PasswordValidator,
	urlcustom:                   UrlValidator,
}

var tagNameFunc = []func(reflect.StructField) string{
	FormTagName,
}

// New func: calls validator.New and add custom validators
func New() *v.Validate {
	validate, ok := binding.Validator.Engine().(*v.Validate)
	if !ok {
		return nil
	}

	for key, value := range tagFuncMaps {
		switch key {
		case timeFormat:
			_ = validate.RegisterValidation(key, value, true)
		default:
			_ = validate.RegisterValidation(key, value)
		}
	}

	for _, value := range tagNameFunc {
		validate.RegisterTagNameFunc(value)
	}

	return validate
}
