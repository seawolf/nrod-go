package messages

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"bitbucket.org/sea_wolf/nrod-go/v2/nrod-go/data"
)

type MovementHeader struct {
	Header
}

type MovementBody struct {
	Body

	TimestampGBTT          TimeFromUnix `json:"gbtt_timestamp"`
	PlannedTimestamp       TimeFromUnix `json:"planned_timestamp"`
	ActualTimestamp        TimeFromUnix `json:"actual_timestamp"`
	OriginalLocationStanox string       `json:"original_loc_stanox"`
	OriginalLOCTimestamp   string       `json:"original_loc_timestamp"`
	TimetableVariation     string       `json:"timetable_variation"`
	CurrentTrainID         string       `json:"current_train_id"`
	DelayMonitoringPoint   string       `json:"delay_monitoring_point"`
	NextReportRunTime      string       `json:"next_report_run_time"`
	ReportingStanox        string       `json:"reporting_stanox"`
	CorrectionInd          string       `json:"correction_ind"`
	EventSource            string       `json:"event_source"`
	TrainFileAddress       string       `json:"train_file_address"`
	Platform               string       `json:"platform"`
	DivisionCode           string       `json:"division_code"`
	TrainTerminated        string       `json:"train_terminated"`
	TrainID                string       `json:"train_id"`
	Offroute               string       `json:"offroute_ind"`
	VariationStatus        string       `json:"variation_status"`
	TrainServiceCode       string       `json:"train_service_code"`
	OperatorID             string       `json:"toc_id"`
	LocationStanox         string       `json:"loc_stanox"`
	AutoExpected           string       `json:"auto_expected"`
	Direction              string       `json:"direction_ind"`
	Route                  string       `json:"route"`
	PlannedEventType       string       `json:"planned_event_type"`
	NextReportStanox       string       `json:"next_report_stanox"`
	Line                   string       `json:"line_ind"`
}

func NewMovementMessage(jsonBody []byte) *MovementMessage {
	var message *MovementMessage
	err := json.Unmarshal(jsonBody, &message)

	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Detected a MovementMessage!")
	}

	return message
}

func (message *MovementMessage) TOCName() (name string) {
	id, err := strconv.Atoi(message.Body.OperatorID)

	if err == nil {
		name = data.TOCsByID[id]
	}

	return
}

func (message *MovementMessage) LocationName(stanoxNumber string) (name string) {
	id, err := strconv.Atoi(stanoxNumber)

	if err == nil {
		name = data.StanoxByID[id]
	}

	return
}

func (message *MovementMessage) Platform() string {
	return strings.Trim(message.Body.Platform, " ")
}

func (message *MovementMessage) Arrival() bool {
	return message.Body.EventType == "ARRIVAL"
}

func (message *MovementMessage) Departure() bool {
	return message.Body.EventType == "DEPARTURE"
}

func (message *MovementMessage) Headcode() string {
	if len(message.Body.TrainID) >= 6 {
		return message.Body.TrainID[2:6]
	}
	return "????"
}

func (message *MovementMessage) ToString() string {
	locationName := message.LocationName(message.Body.LocationStanox)
	if locationName == "" {
		return ""
	}

	var locationInfo string
	if message.Arrival() {
		locationInfo = fmt.Sprintf("arrived at %s", locationName)
	} else if message.Departure() {
		locationInfo = fmt.Sprintf("departed from %s", locationName)
	} else {
		return ""
	}

	if message.Body.Platform != "" {
		locationInfo = fmt.Sprintf("%s platform %s", locationInfo, message.Platform())
	}

	if message.Body.Line != "" {
		locationInfo = fmt.Sprintf(`%s on line "%s"`, locationInfo, message.Body.Line)
	}

	if message.Body.Direction != "" {
		locationInfo = fmt.Sprintf(`%s travelling %s`, locationInfo, message.Body.Direction)
	}

	if !message.Arrival() && message.Body.NextReportStanox != "" {
		nextReportStanox := message.LocationName(message.Body.NextReportStanox)
		if nextReportStanox != "" {
			locationInfo = fmt.Sprintf("%s towards %s", locationInfo, nextReportStanox)
		}
	}

	return fmt.Sprintf(
		"train %s by %s has %s",
		message.Headcode(),
		message.TOCName(),
		locationInfo,
	)
}
