package model

import (
	"time"

	"github.com/google/uuid"
)

type SeatMap struct {
	ID        uuid.UUID `json:"id"`
	Aircraft  string    `json:"aircraft"`
	Cabins    []Cabin   `json:"cabins"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Cabin struct {
	ID          uuid.UUID `json:"id"`
	SeatMapID   uuid.UUID `json:"seatMapId"`
	Deck        string    `json:"deck"`
	SeatColumns []string  `json:"seatColumns"`
	FirstRow    int       `json:"firstRow"`
	LastRow     int       `json:"lastRow"`
	SeatRows    []SeatRow `json:"seatRows"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type SeatRow struct {
	ID        uuid.UUID `json:"id"`
	CabinID   uuid.UUID `json:"cabinId"`
	RowNumber int       `json:"rowNumber"`
	Seats     []Seat    `json:"seats"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Seat struct {
	ID                 uuid.UUID `json:"id"`
	SeatRowID          uuid.UUID `json:"seatRowId"`
	StorefrontSlotCode string    `json:"storefrontSlotCode"`
	Available          bool      `json:"available"`
	Code               string    `json:"code"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
