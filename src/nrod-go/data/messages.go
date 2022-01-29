package data

var MessageTypes = map[string]string{
	"0001": "Train Activation",
	"0002": "Train Cancellation",
	"0003": "Train Movement",
	"0004": "Unidentified Train", // not used in production
	"0005": "Train Reinstatement",
	"0006": "Change of Origin",
	"0007": "Change of Identity",
	"0008": "Change of Location",
}
