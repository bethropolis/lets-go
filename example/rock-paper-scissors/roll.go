package main



func winner(user, computer int) string {
	sub := user - computer

	if sub == 1 || sub == -2 {
		return "win"
	}
	if sub == -1 || sub == 2 {
		return "lose"
	}
	return "draw"
}
