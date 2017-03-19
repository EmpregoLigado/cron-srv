package runner

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rubyist/circuitbreaker"
	"time"
)

type Runner interface {
	Run() chan *Config
}

type Config struct {
	Url     string
	Retries int64
	Timeout int64
}

type runner struct {
	runChannel chan *Config
}

func New() Runner {
	r := &runner{
		runChannel: make(chan *Config),
	}

	go r.register()

	return r
}

func (r *runner) Run() chan *Config {
	return r.runChannel
}

func (r *runner) register() {
	for {
		select {
		case event := <-r.runChannel:
			r.run(event)
		}
	}
}

func (r *runner) run(c *Config) {
	timeout := time.Second * time.Duration(c.Timeout)
	client := circuit.NewHTTPClient(timeout, c.Retries, nil)

	_, err := client.Get(c.Url)
	if err == nil {
		log.WithField("url", c.Url).Info("Event job event sent")
		return
	}

	log.WithField("url", c.Url).Info("Failed to send event")
}
