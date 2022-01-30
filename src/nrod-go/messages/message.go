package messages

import (
	"encoding/json"
	"fmt"
	"strconv"

	"bitbucket.org/sea_wolf/nrod-go/v2/nrod-go/data"
)

type Message struct {
	Header Header `json:"header"`
	Body   Body   `json:"body"`
}

type Header struct {
	MsgType            string `json:"msg_type"`
	SourceDevID        string `json:"source_dev_id"`
	UserID             string `json:"user_id"`
	OriginalDataSource string `json:"original_data_source"`
	MsgQueueTimestamp  string `json:"msg_queue_timestamp"`
	SourceSystemID     string `json:"source_system_id"`
}

type Body struct {
	EventType  string `json:"event_type"`
	TrainID    string `json:"train_id"`
	OperatorID string `json:"toc_id"`
}

type GenericMessage interface {
	ToString() string
}

func Detect(subscriptionMessage []byte) []GenericMessage {
	var rawMessages []*json.RawMessage
	_ = json.Unmarshal(subscriptionMessage, &rawMessages)

	var messages []GenericMessage
	for _, obj := range rawMessages {
		var msg *Message
		_ = json.Unmarshal(*obj, &msg)

		if msg.IsActivationMessage() {
			var m *ActivationMessage
			_ = json.Unmarshal(*obj, &m)
			messages = append(messages, m)
		} else if msg.IsCancellationMessage() {
			var m *CancellationMessage
			_ = json.Unmarshal(*obj, &m)
			messages = append(messages, m)
		} else if msg.IsMovementMessage() {
			var m *MovementMessage
			_ = json.Unmarshal(*obj, &m)
			messages = append(messages, m)
		}
	}

	return messages
}

func (message *Message) IsActivationMessage() bool {
	return message.Header.MsgType == "0001"
}

func (message *Message) IsCancellationMessage() bool {
	return message.Header.MsgType == "0002"
}

func (message *Message) IsMovementMessage() bool {
	return message.Header.MsgType == "0003"
}

func (message *Message) ToString() string {
	return fmt.Sprintf("Message%v", message)
}

func (body *Body) TOCName() (name string) {
	id, err := strconv.Atoi(body.OperatorID)

	if err == nil {
		name = data.TOCsByID[id]
	}

	return
}

func (body *Body) Headcode() string {
	if len(body.TrainID) >= 6 {
		return body.TrainID[2:6]
	}
	return "????"
}

func (*Body) LocationName(stanoxNumber string) (name string) {
	id, err := strconv.Atoi(stanoxNumber)

	if err == nil {
		name = data.StanoxByID[id]
	}

	return
}
