package handler

import (
	"net/http"
	"seat-map/internal/pkg/httpx"
	"seat-map/internal/service"

	"github.com/google/uuid"
)

type SeatMapHandler struct {
	SeatMapService *service.SeatMapService
}

func (handler *SeatMapHandler) GetSeatMapByID(w http.ResponseWriter, req *http.Request) {
	seatMapID := req.PathValue("seatMapID")
	if seatMapID == "" {
		httpx.ResponseJson(w, http.StatusBadRequest, "ERROR", "Missing seat map id", nil)
		return
	}

	parsedSeatMapID, err := uuid.Parse(seatMapID)
	if err != nil {
		httpx.ResponseJson(w, http.StatusBadRequest, "ERROR", "Invalid format seat map id", nil)
		return
	}

	seatMap, err := handler.SeatMapService.GetSeatMapByID(parsedSeatMapID)
	if err != nil {
		httpx.ResponseJson(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error(), nil)
		return
	}

	if seatMap == nil {
		httpx.ResponseJson(w, http.StatusOK, "DATA_NOT_FOUND", "Seat map not found", nil)
		return
	}

	httpx.ResponseJson(w, http.StatusOK, "SUCCESS", "SUCCESS", seatMap)
}
