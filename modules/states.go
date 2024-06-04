package modules

import (
	"errors"

	"gopkg.in/telebot.v3"
)

type StateType string

const (
	FreeState    StateType = "free"
	WaitingState StateType = "waiting"
)

var states map[int64]telebot.HandlerFunc

func InitStates() {
	states = make(map[int64]telebot.HandlerFunc)
}

func SetState(userId int64, state telebot.HandlerFunc) {
	states[userId] = state
}

func UserInState(userId int64) (inState bool) {
	_, inState = states[userId]
	return
}

func GetState(userId int64) (telebot.HandlerFunc, error) {
	state, ok := states[userId]
	if !ok {
		return state, errors.New("user state not found")
	}

	return state, nil
}

func DelState(userId int64) (err error) {
	_, err = GetState(userId)
	if err != nil {
		return
	}

	delete(states, userId)

	return
}
