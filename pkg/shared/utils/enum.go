package utils

// PeriodType int
type PeriodType int

// ImportantType int
type ImportantType int

// LearningPurpose int
type LearningPurpose int

// SensitivityRecording int
type SensitivityRecording int

// SpeechRate int
type SpeechRate int

// ChangeTargetType int
type ChangeTargetType int

// PlanType int
type PlanType int

const (
	// AllLevel Type
	AllLevel ImportantType = iota
	// One Type
	One
	// Two Type
	Two
	// Three Type
	Three
	// Four Type
	Four
	// Five Type
	Five
)

const (
	// All Type
	All PeriodType = iota
	// Day Type
	Day
	// Week Type
	Week
	// Month Type
	Month
)

const (
	// Business Type
	Business LearningPurpose = iota + 1

	// DailyConversation Type
	DailyConversation

	// Travel Type
	Travel
)

const (
	// High value
	High SensitivityRecording = iota + 1

	// Normal value
	Normal

	// Low value
	Low
)

const (
	// Fast Type
	Fast SpeechRate = iota + 1

	// Usually Type
	Usually

	// Slowly Type
	Slowly
)

const (
	// Master Type
	Master ChangeTargetType = iota + 1

	// Custom Type
	Custom
)

const (
	// MonthPT Type
	MonthPT PlanType = iota + 1

	// YearPT Type
	YearPT
)

// CheckPeriod func
func (p PeriodType) CheckPeriod() bool {
	switch p {
	case Day, Week, Month, All:
		return true
	default:
		return false
	}
}

// CheckImportantType func
func (i ImportantType) CheckImportantType() bool {
	switch i {
	case AllLevel, One, Two, Three, Four, Five:
		return true
	default:
		return false
	}
}

// CheckLearningPurpose func
func (l LearningPurpose) CheckLearningPurpose() bool {
	switch l {
	case Business, DailyConversation, Travel:
		return true
	default:
		return false
	}
}

// CheckSensitivityRecording func
func (s SensitivityRecording) CheckSensitivityRecording() bool {
	switch s {
	case High, Normal, Low:
		return true
	default:
		return false
	}
}

// CheckSpeechRate func
func (r SpeechRate) CheckSpeechRate() bool {
	switch r {
	case Fast, Usually, Slowly:
		return true
	default:
		return false
	}
}

// CheckChangeTargetType func
func (c ChangeTargetType) CheckChangeTargetType() bool {
	switch c {
	case Master, Custom:
		return true
	default:
		return false
	}
}

// CheckPlanType func
func (p PlanType) CheckPlanType() bool {
	switch p {
	case MonthPT, YearPT:
		return true
	default:
		return false
	}
}
