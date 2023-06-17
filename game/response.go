package game

import (
	"time"

	"github.com/Chat-Map/wordle-server/game/word"
	"github.com/google/uuid"
	"github.com/lordvidex/x/ptr"
)

type Response struct {
	CreatedAt       time.Time             `json:"created_at"`
	StartedAt       *time.Time            `json:"started_at"`
	EndedAt         *time.Time            `json:"ended_at"`
	Creator         string                `json:"creator"`
	CorrectWord     *string               `json:"correct_word,omitempty"` // returned only if game has ended
	Guesses         []GuessResponse       `json:"guesses"`                // contains the guesses of the current player
	GamePerformance []PlayerGuessResponse `json:"game_performance"`       // contains the best guesses of all players
	ID              uuid.UUID             `json:"id"`
}

type GuessResponse struct {
	// Word can be nil if the word was not played by this user
	Word     *string   `json:"word,omitempty"`
	PlayedAt time.Time `json:"played_at"`
	Status   []int     `json:"status,omitempty"`
}

type PlayerGuessResponse struct {
	Username      string        `json:"username,omitempty"`
	GuessResponse GuessResponse `json:"guess_response,omitempty"`
	RankOffset    *int          `json:"rank_offset,omitempty"`
}

func ToResponse(g Game, username string) Response {
	setWord := func(w string) *string {
		if g.EndedAt == nil {
			return nil
		}
		return ptr.String(w)
	}
	perf := make([]PlayerGuessResponse, 0, len(g.Sessions))
	for name, s := range g.Sessions {
		perf = append(perf, PlayerGuessResponse{
			Username:      name,
			GuessResponse: ToGuess(s.BestGuess(), false),
		})
	}
	var guesses []GuessResponse
	userSession, ok := g.Sessions[username]
	if ok {
		guesses = make([]GuessResponse, len(userSession.Guesses))
		for i, guess := range userSession.Guesses {
			guesses[i] = ToGuess(guess, true)
		}
	}
	return Response{
		CreatedAt:       g.CreatedAt,
		StartedAt:       g.StartedAt,
		EndedAt:         g.EndedAt,
		Creator:         g.Creator,
		CorrectWord:     setWord(g.CorrectWord.Word),
		Guesses:         guesses,
		GamePerformance: perf,
		ID:              g.ID,
	}
}

// ToGuess converts a word.Word to a guessResponse.
// If showWord is true, the word is returned, otherwise it is nil.
func ToGuess(w word.Word, showWord bool) GuessResponse {
	guessed := func() *string {
		if showWord {
			return ptr.String(w.Word)
		}
		return nil
	}
	return GuessResponse{
		Word:     guessed(),
		PlayedAt: w.PlayedAt.Time,
		Status:   w.Stats.Ints(),
	}
}