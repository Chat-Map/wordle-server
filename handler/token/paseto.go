package token

import (
	"context"
	"errors"
	"time"

	"github.com/lordvidex/x/auth"
	"github.com/o1egl/paseto/v2"

	"github.com/kodekulture/wordle-server/game"
)

const (
	sessionTsKey = "xts"
)

var (
	defaultFooter = "kodekulture"
)

type Paseto struct {
	footer       string
	symmetricKey []byte
}

func New(key []byte, footer string) (*Paseto, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length, key must be 32 bytes long")
	}
	if footer == "" {
		footer = defaultFooter
	}
	pas := Paseto{
		symmetricKey: key,
		footer:       footer,
	}
	return &pas, nil
}

func (p *Paseto) Generate(ctx context.Context, player game.Player, period time.Duration) (auth.Token, error) {
	payload := p.fromPlayer(player, period)
	str, err := paseto.Encrypt(p.symmetricKey, payload, p.footer)
	if err != nil {
		return "", err
	}
	return auth.Token(str), nil
}

func (p *Paseto) Validate(ctx context.Context, token auth.Token) (game.Player, error) {
	var payload paseto.JSONToken
	if err := paseto.Decrypt(string(token), p.symmetricKey, &payload, &p.footer); err != nil {
		return game.Player{}, err
	}
	if err := payload.Validate(paseto.IssuedBy(p.footer), paseto.ValidAt(time.Now())); err != nil {
		return game.Player{}, err
	}
	return p.toPlayer(payload)
}

func (p *Paseto) fromPlayer(player game.Player, period time.Duration) paseto.JSONToken {
	now := time.Now()
	payload := paseto.JSONToken{
		IssuedAt:   now,
		NotBefore:  now,
		Expiration: now.Add(period),
		Issuer:     p.footer,
	}
	if player.Password != "" {
		player.Password = ""
	}
	payload.Set("player", player)
	payload.Set(sessionTsKey, player.SessionTs)
	return payload
}

func (p *Paseto) toPlayer(t paseto.JSONToken) (game.Player, error) {
	var player game.Player
	err := t.Get("player", &player)
	if err != nil {
		return game.Player{}, err
	}
	return player, nil
}
