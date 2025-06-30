package slack

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/censoredplanet/cp-api/internal/constants"
)

const (
	SlackUnknown = iota
	SlackInfo
	SlackError
	SlackFatal
)

type SlackPort interface {
	Info(details ...string)  // fileName, parentfunc, childfunc, info
	Error(details ...string) // fileName, parentfunc, childfunc, error
	Fatal(details ...string) // fileName, parentfunc, childfunc, fatal
}

type slack struct {
	SendToSlack bool
	Environment string
	WebHookURL  string
}

func NewSlack() *slack {
	environment := strings.ToUpper(os.Getenv("ENV"))
	if environment == "" {
		log.Fatal("ENV is not set")
		return nil
	}

	sendToSlack := true

	if environment == constants.EnvLOCAL {
		sendToSlack = false
	}

	webHookURL := os.Getenv("SLACK_WEBHOOK_URL")

	if environment != constants.EnvDEV && environment != constants.EnvPROD && environment != constants.EnvLOCAL {
		log.Fatal("ENV is not correctly set")
		return nil
	}

	return &slack{
		SendToSlack: sendToSlack,
		Environment: environment,
		WebHookURL:  webHookURL,
	}
}

func (s *slack) Info(details ...string) {
	text := s.Environment + ": " + strings.Join(details, " | ")

	if s.SendToSlack {
		d, err := messageData(text, constants.InfoEmoji)
		if err != nil {
			log.Printf("slack_error.go | Info | messageData | %s", err.Error())
			return
		}

		err = s.sendMessage(d)
		if err != nil {
			log.Printf("slack_error.go | Info | sendMessage: %s", err.Error())
			return
		}
	}

	log.Printf("Info: %s\n", text)
}

func (s *slack) Error(details ...string) {

	text := s.Environment + ": " + strings.Join(details, " | ")

	if s.SendToSlack {
		d, err := messageData(text, constants.ErrorEmoji)
		if err != nil {
			log.Printf("slack_error.go | Error | messageData | %s", err.Error())
			return
		}

		err = s.sendMessage(d)
		if err != nil {
			log.Printf("slack_error.go | Error | sendMessage: %s", err.Error())
			return
		}
	}

	log.Printf("Error: %s\n", text)
}

func (s *slack) Fatal(details ...string) {
	text := s.Environment + ": " + strings.Join(details, " | ")

	if s.SendToSlack {
		d, err := messageData(text, constants.FatalEmoji)
		if err != nil {
			log.Printf("slack_error.go | Fatal | messageData | %s", err.Error())
			return
		}

		err = s.sendMessage(d)
		if err != nil {
			log.Printf("slack_error.go | Fatal | sendMessage: %s", err.Error())
			return
		}
	}

	log.Fatalf("Fatal: %s\n", text)
}

// helper function to send a message to Slack
func (s *slack) sendMessage(msgData []byte) error {
	resp, err := http.Post(s.WebHookURL, "application/json", bytes.NewBuffer(msgData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send message to Slack")
	}

	return nil
}
