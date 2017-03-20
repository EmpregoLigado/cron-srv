package runner

import (
	"gopkg.in/h2non/gock.v1"
	"testing"
	"time"
)

func TestRunnerRun(t *testing.T) {
	url := "https://github.com"

	defer gock.Off()
	gock.New(url).
		Get("/").
		Reply(200).
		BodyString("foo")

	c := &Config{Url: url}
	r := New()
	r.Run() <- c

	time.Sleep(time.Second * 1)

	if !gock.IsDone() {
		t.Errorf("Expected to call %s", url)
	}
}
