package tictacgo

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

type ArbitraryToken rune

func (at ArbitraryToken) Generate(rand *rand.Rand, size int) reflect.Value {
	minCodePoint := int('A')
	maxCodePoint := int('Z')
	allowedCodePoints := make([]int, maxCodePoint-minCodePoint+1)
	for i := range allowedCodePoints {
		allowedCodePoints[i] = minCodePoint + i
	}
	randCodePoint := allowedCodePoints[rand.Intn(len(allowedCodePoints)%size)]
	arbitraryRune := rune(byte(randCodePoint))
	return reflect.ValueOf(arbitraryRune)
}

type ArbitraryTokenPair [2]ArbitraryToken

func (atp ArbitraryTokenPair) Generate(rand *rand.Rand, size int) reflect.Value {
	writeIndex := 0
	for i := 0; i < 100; i++ {
		at := ArbitraryToken('_')
		newToken := ArbitraryToken(at.GenerateRune(rand, size))
		if writeIndex == 0 {
			atp[0] = newToken
			writeIndex++
		} else if newToken != atp[0] {
			atp[1] = newToken
			return reflect.ValueOf(atp)
		}
	}
	panic("Failed to generate pair of arbitrary tokens in 100 tries")
}

func (at ArbitraryToken) GenerateRune(rand *rand.Rand, size int) rune {
	return at.Generate(rand, size).Interface().(rune)
}

type ArbitraryBoard struct {
	Board
	Player1Token rune
	Player2Token rune
}

func (ab ArbitraryBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	ab.Board = EmptyBoard()
	atp := ArbitraryTokenPair{}.Generate(rand, size).Interface().(ArbitraryTokenPair)
	ab.Player1Token = rune(atp[0])
	ab.Player2Token = rune(atp[1])
	if ab.Player1Token == ab.Player2Token {
		panic("Winning and losing tokens must not be the same")
	}
	return reflect.ValueOf(ab)
}

type ArbitraryPendingBoard struct{ ArbitraryBoard }

func (apb ArbitraryPendingBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	apb.ArbitraryBoard = apb.ArbitraryBoard.Generate(rand, size).Interface().(ArbitraryBoard)

	shuffledSpaces := rand.Perm(apb.Board.SpacesLen())

	for i, space := range shuffledSpaces[:size%6] {
		var currentToken rune
		if i%2 == 0 {
			currentToken = apb.Player1Token
		} else {
			currentToken = apb.Player2Token
		}
		apb.Board = apb.Board.AssignSpace(space, &currentToken)
	}
	return reflect.ValueOf(apb)
}

func TestArbitraryPendingBoard(t *testing.T) {
	assert := assert.New(t)

	qcErr := quick.Check(func(apb ArbitraryPendingBoard) bool {
		state, winner := apb.Board.GameState()
		return state == Pending && winner == nil
	}, nil)
	assert.Nil(qcErr)
}

// ArbitraryVictoryBoard has spaces assigned to an arbitrary winning player which should result in a victory
// and arbitrary other spaces assigned to the losing player.
//
// The number of moves is always 5, since that guarantees that the losing player can't chose other spots which
// would also potentially qualify as a victory (e.g. both players filling separate rows).
type ArbitraryVictoryBoard struct {
	ArbitraryBoard
	WinningToken rune
	LosingToken  rune
}

func (avb ArbitraryVictoryBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	avb.ArbitraryBoard = avb.ArbitraryBoard.Generate(rand, size).Interface().(ArbitraryBoard)

	// choose who will win
	if rand.Int31n(2) == 0 {
		avb.WinningToken = avb.Player1Token
		avb.LosingToken = avb.Player2Token
	} else {
		avb.WinningToken = avb.Player2Token
		avb.LosingToken = avb.Player1Token
	}

	// choose how they win
	possibleWinningVectors := [][]int{}
	possibleWinningVectors = append(possibleWinningVectors, avb.rowIndexVectors()...)
	possibleWinningVectors = append(possibleWinningVectors, avb.columnIndexVectors()...)
	possibleWinningVectors = append(possibleWinningVectors, avb.diagonalIndexVectors()...)

	winningVectorIndex := rand.Intn(len(possibleWinningVectors)) % size

	winningVector := possibleWinningVectors[winningVectorIndex]

	// assign winning spaces
	for _, spaceIndex := range winningVector {
		avb.Board = avb.Board.AssignSpace(spaceIndex, &avb.WinningToken)
	}

	// assign losing spaces
	possibleLosingSpaces := avb.Board.AvailableSpaces()
	rand.Shuffle(len(possibleLosingSpaces), func(i, j int) {
		possibleLosingSpaces[i], possibleLosingSpaces[j] = possibleLosingSpaces[j], possibleLosingSpaces[i]
	})
	for _, spaceIndex := range possibleLosingSpaces[:2] {
		avb.Board = avb.Board.AssignSpace(spaceIndex, &avb.LosingToken)
	}

	return reflect.ValueOf(avb)
}

func TestArbitraryVictoryGameState(t *testing.T) {
	assert := assert.New(t)

	qcErr := quick.Check(func(avb ArbitraryVictoryBoard) bool {
		state, winner := avb.GameState()
		// fmt.Printf("state %s, winner %#v, board: \n%s", state, spaceToString(winner, "null"), avb.Board.String())
		return state == Victory && winner != nil && *winner == avb.WinningToken
	}, nil)
	assert.Nil(qcErr)
}

type ArbitraryFullBoard struct {
	ArbitraryBoard
}

func (afb ArbitraryFullBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	afb.ArbitraryBoard = afb.ArbitraryBoard.Generate(rand, size).Interface().(ArbitraryBoard)

	shuffledSpaces := rand.Perm(afb.Board.SpacesLen())

	for i, space := range shuffledSpaces {
		var currentToken rune
		if i%2 == 0 {
			currentToken = afb.Player1Token
		} else {
			currentToken = afb.Player2Token
		}
		afb.Board = afb.Board.AssignSpace(space, &currentToken)
	}
	if len(afb.Board.AvailableSpaces()) > 0 {
		panic("Board should have been filled")
	}
	return reflect.ValueOf(afb)
}

func TestFullBoardEitherTieOrVictory(t *testing.T) {
	assert := assert.New(t)

	qcErr := quick.Check(func(afb ArbitraryFullBoard) bool {
		// fmt.Println(fmt.Sprintf("checking state of board : \n%s", afb.Board.String()))
		state, winner := afb.GameState()
		// fmt.Println(fmt.Sprintf("and the result was: %s, %s", state, spaceToString(winner, "null")))
		switch state {
		case Tie:
			if winner == nil {
				// Ties don't have a winner
				return true
			}
			return false
		case Victory:
			if winner != nil {
				// If the board is full and it's not a tie, must be a victory (with a winner)
				return true
			}
			return false
		case Pending:
		default:
			break
		}
		// Full boards should never be pending
		return false
	}, nil)
	assert.Nil(qcErr)
}
