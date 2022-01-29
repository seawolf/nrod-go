package messages

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
	EventType string `json:"event_type"`
}

type MovementMessage struct {
	Header MovementHeader `json:"header"`
	Body   MovementBody   `json:"body"`
}
