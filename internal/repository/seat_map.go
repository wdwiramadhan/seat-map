package repository

import (
	"context"
	"database/sql"
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
	c.id AS cabin_id, c.seat_map_id AS cabin_seat_map_id, c.deck, c.first_column, c.last_column, c.created_at AS cabin_created_at, c.updated_at AS cabin_updated_at,
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

			cID        sql.NullString
			cSeatMapID sql.NullString
			deck       sql.NullString
			// seatColumns             pq.StringArray
			firstColumn, lastColumn sql.NullInt64
			cCreatedAt, cUpdatedAt  sql.NullTime

			srID                     sql.NullString
			srCabinID                sql.NullString
			rowNum                   sql.NullInt64
			srCreatedAt, srUpdatedAt sql.NullTime

			sID        sql.NullString
			sSeatRowID sql.NullString
			// slotChars              pq.StringArray
			storefrontCode         sql.NullString
			available              sql.NullBool
			code                   sql.NullString
			sCreatedAt, sUpdatedAt sql.NullTime
		)

		err := rows.Scan(
			&smID, &smAircraft, &smCreatedAt, &smUpdatedAt,
			&cID, &cSeatMapID, &deck, &firstColumn, &lastColumn, &cCreatedAt, &cUpdatedAt,
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
		if cID.Valid {
			parsedCabinID, _ := uuid.Parse(cID.String)
			if existing, ok := cabinMap[parsedCabinID]; ok {
				cabin = existing
			} else {
				cabin = &model.Cabin{
					ID:          parsedCabinID,
					Deck:        deck.String,
					SeatColumns: make([]string, 0),
					FirstColumn: int(firstColumn.Int64),
					LastColumn:  int(lastColumn.Int64),
					CreatedAt:   cCreatedAt.Time,
					UpdatedAt:   cUpdatedAt.Time,
				}

				cabinMap[parsedCabinID] = cabin
				seatMap.Cabins = append(seatMap.Cabins, *cabin)
			}
		}

		var seatRow *model.SeatRow
		if srID.Valid {
			parsedSeatRowID, _ := uuid.Parse(srID.String)
			if existing, ok := seatRowMap[parsedSeatRowID]; ok {
				seatRow = existing
			} else {
				seatRow = &model.SeatRow{
					ID:        parsedSeatRowID,
					RowNumber: int(rowNum.Int64),
					CreatedAt: srCreatedAt.Time,
					UpdatedAt: srUpdatedAt.Time,
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
		if sID.Valid && seatRow != nil {
			parsedSeatID, _ := uuid.Parse(sID.String)
			seat = &model.Seat{
				ID: parsedSeatID,
				// SlotCharacteristics: slotChars,
				StorefrontSlotCode: storefrontCode.String,
				Available:          available.Bool,
				Code:               code.String,
				CreatedAt:          sCreatedAt.Time,
				UpdatedAt:          sUpdatedAt.Time,
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
