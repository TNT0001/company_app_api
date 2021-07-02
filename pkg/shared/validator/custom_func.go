package validator

import (
	"go-api/pkg/shared/utils"
	"net/url"
	"reflect"
	"time"

	v "github.com/go-playground/validator/v10"
)

// IsEnumIntType func
func IsEnumIntType(fl v.FieldLevel) bool {
	value := fl.Field().Int()
	nameEnumType := fl.GetTag()[len("enum"):]

	switch nameEnumType {
	case "Period":
		return utils.PeriodType(value).CheckPeriod()
	case "ImportantType":
		return utils.ImportantType(value).CheckImportantType()
	case "LearningPurpose":
		return utils.LearningPurpose(value).CheckLearningPurpose()
	case "SensitivityRecording":
		return utils.SensitivityRecording(value).CheckSensitivityRecording()
	case "SpeechRate":
		return utils.SpeechRate(value).CheckSpeechRate()
	case "Type":
		return utils.ChangeTargetType(value).CheckChangeTargetType()
	case "PlanType":
		return utils.PlanType(value).CheckPlanType()
	default:
		return false
	}
}

// DateTimeValidator func
func DateTimeValidator(fl v.FieldLevel) bool {
	timeType := reflect.TypeOf(time.Time{})
	val := fl.Field()

BEGIN:
	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return true
		}
		val = val.Elem()

		goto BEGIN
	case reflect.String:
		if val.String() == "" {
			return true
		}

		_, err := time.Parse(utils.FormatTime, val.String())
		return err == nil
	default:
		return val.Type() == timeType
	}
}

func PasswordValidator(fl v.FieldLevel) bool {
	value := fl.Field().String()
	return HasDigit(value) && HasLowwerCaseLetter(value) && HasUpperCaseLetter(value)
}

func UrlValidator(fl v.FieldLevel) bool {
	value := fl.Field().String()
	_, err := url.ParseRequestURI(value)
	if err != nil {
		return false
	}

	u, err := url.Parse(value)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
