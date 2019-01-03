package report

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/chennqqi/go-HoneyPot/config"
)

type httpReport struct {
	client *http.Client
	Url    string
}

func NewHttpReport(remote *config.RemoteHttp) (Reporter, error) {
	var rpt httpReport
	rpt.client = &http.Client{}
	rpt.Url = remote.Uri
	return &rpt, nil
}

func (r *httpReport) Pub(p *HoneypotRecord) error {
	txt, _ := json.Marshal(p)
	rd := bytes.NewBuffer(txt)
	req, _ := http.NewRequest("POST", r.Url, rd)
	req.Header.Set("Content-Type", "application/json")

	client := r.client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	return nil
}

func (r *httpReport) Close() error {
	return nil
}
