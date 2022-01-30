package messages

import (
	"fmt"

	"bitbucket.org/sea_wolf/nrod-go/v2/nrod-go/data"
)

type CancellationMessage struct {
	Message

	Header CancellationHeader `json:"header"`
	Body   CancellationBody   `json:"body"`
}

type CancellationHeader struct {
	Header
}

type CancellationBody struct {
	Body

	TrainFileAddress        string `json:"train_file_address"`
	TrainServiceCode        string `json:"train_service_code"`
	DivisionCode            string `json:"division_code"`
	LocationStanox          string `json:"loc_stanox"`
	DepartureTimestamp      string `json:"dep_timestamp"`
	CancellationType        string `json:"canx_type"`
	CancellationTimestamp   string `json:"canx_timestamp"`
	OriginLocationStanox    string `json:"orig_loc_stanox"`
	OriginLocationTimestamp string `json:"orig_loc_timestamp"`
	CancellationReasonCode  string `json:"canx_reason_code"`
}

func (body *CancellationBody) Reason() string {
	return data.DelayReasons[body.CancellationReasonCode]
}

func (message *CancellationMessage) ToString() string {
	cancellationLocation := message.Body.LocationName(message.Body.LocationStanox)
	if cancellationLocation != "" {
		cancellationLocation = fmt.Sprintf(" effective from %s", cancellationLocation)
	}

	originLocation := message.Body.LocationName(message.Body.OriginLocationStanox)
	if originLocation != "" {
		originLocation = fmt.Sprintf(" (originated at %s)", originLocation)
	}

	reason := message.Body.Reason()
	if reason != "" {
		reason = fmt.Sprintf(" due to %s", reason)
	}

	return fmt.Sprintf(
		"Cancellation of train %s for %s%s%s%s",
		message.Body.Headcode(),
		message.Body.TOCName(),
		cancellationLocation,
		originLocation,
		reason,
	)
}
