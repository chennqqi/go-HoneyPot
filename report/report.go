package report

import (
	"github.com/chennqqi/goutils/utime"
)

type HoneypotRecord struct {
	Src      string     `json:"src"`
	Dst      string     `json:"dst"`
	Srcport  int64      `json:"srcport"`
	Dstport  int64      `json:"dstport"`
	Time     utime.Time `json:"time"`
	Protocol string     `json:"protocol"`

	Payload string `json:"payload"`
	Raw     []byte `json:"raw"`
}

type Reporter interface {
	Pub(record *HoneypotRecord) error
	Close() error
}
