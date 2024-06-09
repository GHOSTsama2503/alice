package middlewares

import (
	"alice/modules"

	"gopkg.in/telebot.v3"
)

func UserState(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) (err error) {

		userId := c.Sender().ID
		if userId == c.Bot().Me.ID {
			return nil
		}

		inState := modules.UserInState(userId)
		if !inState {
			return next(c)
		}

		state, err := modules.GetState(userId)
		if err != nil {
			return err
		}

		return state(c)
	}
}
