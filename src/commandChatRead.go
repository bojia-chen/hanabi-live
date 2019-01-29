/*
	Sent when the user opens the in-game chat
	"data" is empty
*/

package main

import (
	"strconv"
)

func commandChatRead(s *Session, d *CommandData) {
	/*
		Validate
	*/

	// Validate that the game exists
	gameID := s.CurrentGame()
	var g *Game
	if v, ok := games[gameID]; !ok {
		s.Warning("Game " + strconv.Itoa(gameID) + " does not exist.")
		return
	} else {
		g = v
	}

	// Validate that the game has started
	if !g.Running {
		s.Warning("Game " + strconv.Itoa(gameID) + " has not started yet.")
		return
	}

	// Validate that they are in the game
	i := g.GetPlayerIndex(s.UserID())
	if i == -1 {
		s.Warning("You are in not game " + strconv.Itoa(gameID) + ", so you cannot send an action.")
		return
	}

	/*
		Mark that they have read all of the in-game chat
	*/

	p := g.Players[i]
	p.ChatReadIndex = len(g.Chat)
}