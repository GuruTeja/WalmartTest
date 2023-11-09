package controller

import (
	"encoding/json"
	"net/http"
	"walmartTest/pkg/models"
	"walmartTest/pkg/repository"
)

var alarmRepository repository.AlarmRepository

func SaveAlert(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var alarm models.Alert

	if err := json.NewDecoder(req.Body).Decode(&alarm); err != nil {
		respondWithJson(w, http.StatusBadRequest, models.CreateAlertsResponse{AlertId: alarm.AlertId, Error: err.Error()})
		return
	}
	//Save alarm
	_, err := alarmRepository.Insert(alarm)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, models.CreateAlertsResponse{AlertId: alarm.AlertId, Error: err.Error()})
		return
	}
	response := models.CreateAlertsResponse{AlertId: alarm.AlertId, Error: ""}
	respondWithJson(w, http.StatusCreated, response)
}

func GetAlerts(w http.ResponseWriter, req *http.Request) {
	serviceId := req.URL.Query().Get("service_id")
	startTs := req.URL.Query().Get("start_ts")
	endTs := req.URL.Query().Get("end_ts")

	if serviceId == "" {
		respondWithJson(w, http.StatusBadRequest, "service_id should be specified")
		return
	}

	if startTs == "" {
		respondWithJson(w, http.StatusBadRequest, "start_ts should be specified")
		return
	}

	if endTs == "" {
		respondWithJson(w, http.StatusBadRequest, "end_ts should be specified")
		return
	}

	var getAlertRequest = models.GetAlertsRequest{
		ServiceId: serviceId,
		StartTs:   startTs,
		EndTs:     endTs,
	}
	//Get alarm
	alerts, err := alarmRepository.Find(getAlertRequest)

	if err != nil {
		respondWithJson(w, http.StatusNotFound, models.GetAlertsErrorResponse{
			ServiceID: serviceId,
			Error:     "Failed to get Alerts",
		})
		return
	}

	if len(alerts) == 0 {
		respondWithJson(w, http.StatusNotFound, models.GetAlertsErrorResponse{
			ServiceID: serviceId,
			Error:     "No Alerts found",
		})
		return
	}

	response := models.GetAlertsResponse{
		ServiceID:   alerts[0].ServiceId,
		ServiceName: alerts[0].ServiceName,
		Alerts:      alerts,
	}

	respondWithJson(w, http.StatusCreated, response)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
