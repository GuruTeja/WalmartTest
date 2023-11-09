package models

import (
	"github.com/gorilla/mux"
	"time"
)

type App struct {
	Router *mux.Router
}

type Alert struct {
	AlertId          string    `json:"alert_id"`
	ServiceId        string    `json:"service_id"`
	ServiceName      string    `json:"service_name"`
	Model            string    `json:"model"`
	AlertType        string    `json:"alert_type"`
	AlertTs          string    `json:"alert_ts"`
	Severity         string    `json:"severity"`
	TeamSlack        string    `json:"team_slack"`
	CreatedTimeStamp time.Time `json:"created_time_stamp,omitempty"`
}

type CreateAlertsResponse struct {
	AlertId string `json:"alert_id"`
	Error   string `json:"error"`
}

type GetAlertsRequest struct {
	ServiceId string
	StartTs   string
	EndTs     string
}

type GetAlertsErrorResponse struct {
	ServiceID string `json:"service_id"`
	Error     string `json:"error"`
}

type GetAlertsResponse struct {
	ServiceID   string   `json:"service_id"`
	ServiceName string   `json:"service_name"`
	Alerts      []*Alert `json:"alerts"`
}
