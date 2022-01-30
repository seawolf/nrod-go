package messages

import (
	"fmt"
)

type ActivationMessage struct {
	Message

	Header ActivationHeader `json:"header"`
	Body   ActivationBody   `json:"body"`
}

type ActivationHeader struct {
	Header
}

type ActivationBody struct {
	Body

	ScheduleSource              string `json:"schedule_source"`
	TrainFileAddress            string `json:"train_file_address"`
	TrainUID                    string `json:"train_uid"`
	CreationTimestamp           string `json:"creation_timestamp"`
	TrainPlannedOriginTimestamp string `json:"tp_origin_timestamp"`
	TrainPlannedOriginStanox    string `json:"tp_origin_stanox"`
	OriginDepartureTimestamp    string `json:"origin_dep_timestamp"`
	TrainServiceCode            string `json:"train_service_code"`
	D1266RecordNumber           string `json:"d1266_record_number"`
	TrainCallType               string `json:"train_call_type"`
	TrainCallMode               string `json:"train_call_mode"`
	ScheduleType                string `json:"schedule_type"`
	ScheduleOriginStanox        string `json:"sched_origin_stanox"`
	ScheduleWorkingTimetableID  string `json:"schedule_wtt_id"`
	ScheduleStartDate           string `json:"schedule_start_date"`
	ScheduleEndDate             string `json:"schedule_end_date"`
}

func (message *ActivationMessage) ToString() string {
	locationInfo := message.Body.LocationName(message.Body.ScheduleOriginStanox)
	if locationInfo != "" {
		locationInfo = fmt.Sprintf(" originating at %s", locationInfo)
	}

	return fmt.Sprintf(
		"Activation of train %s for %s%s",
		message.Body.Headcode(),
		message.Body.TOCName(),
		locationInfo,
	)
}
