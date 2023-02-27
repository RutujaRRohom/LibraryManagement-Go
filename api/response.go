package api

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

type Message struct {
	Msg string `json:"message"`
}

func Response(w http.ResponseWriter, status int, response interface{}) {

	respBytes, err := json.Marshal(response)
	if err != nil {
		logger.WithField("err", err.Error()).Error()
		status = http.StatusInternalServerError
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBytes)
}
