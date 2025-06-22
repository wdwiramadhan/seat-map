package handler

import (
	"encoding/json"
	"net/http"
	"seat-map/internal/dto"
	"seat-map/internal/service"

	"github.com/google/uuid"
)

type SeatMapHandler struct {
	SeatMapService *service.SeatMapService
}

func (handler *SeatMapHandler) GetSeatMapByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	seatMapID := req.PathValue("seatMapID")

	if seatMapID == "" {
		response := dto.APIResponse{
			Code:    "ERROR",
			Message: "Missing seat map id",
			Data:    nil,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	parsedSeatMapID, err := uuid.Parse(seatMapID)
	if err != nil {
		response := dto.APIResponse{
			Code:    "ERROR",
			Message: "Invalid format seat map id",
			Data:    nil,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	seatMap, err := handler.SeatMapService.GetSeatMapByID(parsedSeatMapID)
	if err != nil {
		response := dto.APIResponse{
			Code:    "ERROR",
			Message: "ERROR",
			Data:    nil,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := dto.APIResponse{
		Code:    "SUCCESS",
		Message: "SUCCESS",
		Data:    seatMap,
	}

	json.NewEncoder(w).Encode(response)
}
