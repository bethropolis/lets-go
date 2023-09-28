# Rock, Paper, Scissors Game in Go

This Go program is a simple implementation of the classic game "Rock, Paper, Scissors." It allows you to play against the computer and provides a winner based on your choice and the computer's choice.

## How to Play

1. Clone or download this repository to your local machine.

2. Navigate to the project directory in your terminal.

3. Run the program using the `go run` command:

   ```shell
   go run .
   ```

4. Follow the on-screen instructions to make your choice (rock, paper, or scissors). You can type your choice and press Enter.

5. The program will reveal the computer's choice and determine the winner.

6. After each game, you will have the option to play again by typing 'y' for yes or 'n' for no.

## Rules

- Rock beats scissors.
- Scissors beats paper.
- Paper beats rock.

## Code Explanation

- The program uses constants (`ROCK`, `PAPER`, and `SCISSORS`) to represent the possible choices.

```go
  //LOGIC
  rock = 1
  paper = 2
  scissors = 3

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

winner(rock, rock) // "draw"
winner(rock, paper) // "lose"
winner(rock, scissors) // "win"
winner(paper, rock) // "win"
winner(paper, paper) // "draw"
winner(paper, scissors) // "lose"
winner(scissors, rock) // "lose"
winner(scissors, paper) // "win"
winner(scissors, scissors) // "draw"
```
- computer choice = `rand.Intn(3) + 1`
- It uses a switch statement to determine the winner based on the user's choice and the computer's choice.
- It returns the winner as a string.


## Notes

- This is a basic implementation of Rock, Paper, Scissors for educational purposes.
- Feel free to modify and expand upon this code to add more features or improve the user experience.
