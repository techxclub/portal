package domain

type OTPGeneration struct {
	Type  string
	Value string
}

type OTPVerification struct {
	Value string
	Code  string
}

type AuthDetails struct {
	Status string
}
