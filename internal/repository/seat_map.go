package repository

import (
	"context"
	"seat-map/internal/model"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SeatMapRepository struct {
	DB *pgxpool.Pool
}

func (repo *SeatMapRepository) GetSeatMapByID(seatMapID uuid.UUID) (*model.SeatMap, error) {
	query := `SELECT  sm.id AS seat_map_id, sm.aircraft, sm.created_at AS seat_map_created_at, sm.updated_at AS seat_map_updated_at,
			c.id AS cabin_id, c.seat_map_id AS cabin_seat_map_id, c.deck, c.seat_columns, c.first_row, c.last_row, c.created_at AS cabin_created_at, c.updated_at AS cabin_updated_at,
			sr.id AS seat_row_id, sr.cabin_id AS seat_row_cabin_id, sr.row_number, sr.created_at AS seat_row_created_at, sr.updated_at AS seat_row_updated_at,
			s.id AS seat_id, s.seat_row_id AS seat_seat_row_id, s.storefront_slot_code, s.available, s.code, s.created_at AS seat_created_at, s.updated_at AS seat_updated_at
			FROM seat_maps AS sm LEFT JOIN cabins as c ON sm.id = c.seat_map_id LEFT JOIN seat_rows AS sr ON c.id = sr.cabin_id LEFT JOIN seats AS s ON s.seat_row_id = sr.id
			WHERE sm.id = $1`

	rows, err := repo.DB.Query(context.Background(), query, seatMapID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var seatMap *model.SeatMap
	cabinMap := make(map[uuid.UUID]*model.Cabin)
	seatRowMap := make(map[uuid.UUID]*model.SeatRow)

	for rows.Next() {
		var (
			smID                     uuid.UUID
			smAircraft               string
			smCreatedAt, smUpdatedAt time.Time

			cID                    *string
			cSeatMapID             *string
			deck                   *string
			seatColumns            []string
			firstRow, lastRow      *int64
			cCreatedAt, cUpdatedAt *time.Time

			srID                     *string
			srCabinID                *string
			rowNum                   *int64
			srCreatedAt, srUpdatedAt *time.Time

			sID                    *string
			sSeatRowID             *string
			storefrontCode         *string
			available              *bool
			code                   *string
			sCreatedAt, sUpdatedAt *time.Time
		)

		err := rows.Scan(
			&smID, &smAircraft, &smCreatedAt, &smUpdatedAt,
			&cID, &cSeatMapID, &deck, &seatColumns, &firstRow, &lastRow, &cCreatedAt, &cUpdatedAt,
			&srID, &srCabinID, &rowNum, &srCreatedAt, &srUpdatedAt,
			&sID, &sSeatRowID, &storefrontCode, &available, &code, &sCreatedAt, &sUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if seatMap == nil {
			seatMap = &model.SeatMap{
				ID: smID, Aircraft: smAircraft,
				CreatedAt: smCreatedAt, UpdatedAt: smUpdatedAt,
			}
		}

		var cabin *model.Cabin
		if cID != nil {
			parsedCabinID, _ := uuid.Parse(*cID)
			if existing, ok := cabinMap[parsedCabinID]; ok {
				cabin = existing
			} else {
				parsedSeatMapID, _ := uuid.Parse(*cSeatMapID)
				cabin = &model.Cabin{
					ID:          parsedCabinID,
					SeatMapID:   parsedSeatMapID,
					Deck:        *deck,
					SeatColumns: seatColumns,
					FirstRow:    int(*firstRow),
					LastRow:     int(*lastRow),
					CreatedAt:   *cCreatedAt,
					UpdatedAt:   *cUpdatedAt,
				}

				cabinMap[parsedCabinID] = cabin
				seatMap.Cabins = append(seatMap.Cabins, *cabin)
			}
		}

		var seatRow *model.SeatRow
		if srID != nil {
			parsedSeatRowID, _ := uuid.Parse(*srID)
			if existing, ok := seatRowMap[parsedSeatRowID]; ok {
				seatRow = existing
			} else {
				parsedCabinID, _ := uuid.Parse(*srCabinID)
				seatRow = &model.SeatRow{
					ID:        parsedSeatRowID,
					CabinID:   parsedCabinID,
					RowNumber: int(*rowNum),
					CreatedAt: *srCreatedAt,
					UpdatedAt: *srUpdatedAt,
				}

				seatRowMap[parsedSeatRowID] = seatRow
				for i := range seatMap.Cabins {
					if seatMap.Cabins[i].ID == cabin.ID {
						seatMap.Cabins[i].SeatRows = append(seatMap.Cabins[i].SeatRows, *seatRow)
						break
					}
				}
			}
		}

		var seat *model.Seat
		if sID != nil && seatRow != nil {
			parsedSeatID, _ := uuid.Parse(*sID)
			parsedSeatRowID, _ := uuid.Parse(*sSeatRowID)
			seat = &model.Seat{
				ID:                 parsedSeatID,
				SeatRowID:          parsedSeatRowID,
				StorefrontSlotCode: *storefrontCode,
				Available:          *available,
				Code:               *code,
				CreatedAt:          *sCreatedAt,
				UpdatedAt:          *sUpdatedAt,
			}

			for i := range seatMap.Cabins {
				if seatMap.Cabins[i].ID == cabin.ID {
					for j := range seatMap.Cabins[i].SeatRows {
						if seatMap.Cabins[i].SeatRows[j].ID == seatRow.ID {
							seatMap.Cabins[i].SeatRows[j].Seats = append(seatMap.Cabins[i].SeatRows[j].Seats, *seat)
						}
					}
				}
			}

		}

	}

	return seatMap, nil
}
