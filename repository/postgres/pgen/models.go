// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgen

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Game struct {
	ID          pgtype.UUID
	Creator     int32
	CorrectWord string
	CreatedAt   pgtype.Timestamptz
	StartedAt   pgtype.Timestamptz
	EndedAt     pgtype.Timestamptz
}

type GamePlayer struct {
	GameID        pgtype.UUID
	PlayerID      int32
	PlayedWords   []byte
	BestGuess     pgtype.Text
	BestGuessTime pgtype.Timestamptz
	Finished      pgtype.Timestamptz
	Rank          pgtype.Int4
}

type Player struct {
	ID        int32
	Username  string
	Password  string
	SessionTs pgtype.Int8
}
