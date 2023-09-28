package main

func convertToNumber(rock string) int {
	switch rock {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	default:
		return -1
	}
}

func convertToString(number int) string {
	switch number {
	case 1:
		return ROCK
	case 2:
		return PAPER
	case 3:
		return SCISSORS
	default:
		return ""
	}
}