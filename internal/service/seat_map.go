package service

import (
	"seat-map/internal/model"
	"seat-map/internal/repository"

	"github.com/google/uuid"
)

type SeatMapService struct {
	SeatMapRepository *repository.SeatMapRepository
}

func (service *SeatMapService) GetSeatMapByID(seatMapID uuid.UUID) (*model.SeatMap, error) {
	return service.SeatMapRepository.GetSeatMapByID(seatMapID)
}
