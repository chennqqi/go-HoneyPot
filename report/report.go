package report

import (
	"errors"
	"time"
)

var (
	timeFormat = time.RFC3339Nano
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte(`"null"`), nil
	}
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 1 {
		return errors.New("No data")
	}
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	if string(data) == "null" {
		*t = Time(time.Unix(0, 0))
		return nil
	}

	now, err := time.ParseInLocation(timeFormat, string(data), time.Local)
	*t = Time(now)
	return err
}

func (t Time) ToTime() time.Time {
	return time.Time(t)
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

type HoneypotRecord struct {
	Id      string `json:"id"`
	Src     string `json:"src"`
	Dst     string `json:"dst"`
	Srcport int64  `json:"srcport"`
	Dstport int64  `json:"dstport"`
	Payload string `json:"payload"`
	Raw     []byte `json:"raw"`
	Time    Time   `json:"time"`
}

type Reporter interface {
	Pub(record *HoneypotRecord) error
	Close() error
}
