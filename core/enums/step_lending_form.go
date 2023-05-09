package enums

type StepLendingForm string

const (
	StepBasicInfo    StepLendingForm = "base_info"
	StepBusinessInfo StepLendingForm = "business_info"
	StepFileInfo     StepLendingForm = "file_info"
	StepSubmit       StepLendingForm = "submit"
)
