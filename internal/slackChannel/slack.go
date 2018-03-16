package slackChannel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/Neoin/ReminderBot/internal/interfaces"
)

type Slack struct {
}

const url = ""

func (s *Slack) Send(m interfaces.Message) error {
	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(
		struct {
			Text string `json:"text"`
		}{
			Text: m.String(),
		},
	)
	resp, err := http.Post(url, "application/json", buf)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return err
}
