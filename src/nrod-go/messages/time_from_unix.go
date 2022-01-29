package messages

import (
	"strconv"
	"strings"
	"time"
)

type TimeFromUnix struct {
	time.Time

	Parsed bool
}

func (t *TimeFromUnix) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")

	if len(str) > 10 {
		str = str[0:10]
	}

	i, _ := strconv.ParseInt(str, 10, 64)
	value := time.Unix(i, 0)

	if !value.IsZero() {
		t.Time = value
	}

	return nil
}

func (t *TimeFromUnix) Format(format string) string {
	if t.Time.IsZero() {
		return "─"
	}

	return t.Time.Format(format)
}
