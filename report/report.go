package report

type HoneypotRecord struct {
	Id      string `json:"id"`
	Src     string `json:"src"`
	Dst     string `json:"dst"`
	Port    int64  `json:"port"`
	Payload string `json:"payload"`
	Raw     []byte `json:"raw"`
}

type Reporter interface {
	Pub(record *HoneypotRecord) error
}
