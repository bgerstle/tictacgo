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
}

func (ab ArbitraryBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	atp := ArbitraryTokenPair{}.Generate(rand, size).Interface().(ArbitraryTokenPair)

	ab.Board = NewBoard([2]PlayerInfo{
		PlayerInfo{rune(atp[0])},
		PlayerInfo{rune(atp[1])},
	})
	return reflect.ValueOf(ab)
}

type ArbitraryPendingBoard struct{ ArbitraryBoard }

func (apb ArbitraryPendingBoard) Generate(rand *rand.Rand, size int) reflect.Value {
	apb.ArbitraryBoard = apb.ArbitraryBoard.Generate(rand, size).Interface().(ArbitraryBoard)

	shuffledSpaces := rand.Perm(apb.Board.SpacesLen())

	for _, space := range shuffledSpaces[:size%6] {
		apb.Board, _, _ = apb.Board.AssignSpace(space)
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

	// first player always wins for now
	avb.WinningToken = avb.Board.ActivePlayerToken()
	avb.LosingToken = avb.Board.NextPlayerToken()

	// choose how they win
	possibleWinningVectors := [][]int{}
	possibleWinningVectors = append(possibleWinningVectors, avb.rowIndexVectors()...)
	possibleWinningVectors = append(possibleWinningVectors, avb.columnIndexVectors()...)
	possibleWinningVectors = append(possibleWinningVectors, avb.diagonalIndexVectors()...)

	winningVectorIndex := rand.Intn(len(possibleWinningVectors)) % size

	winningVector := possibleWinningVectors[winningVectorIndex]

	// choose 2 losing spaces that aren't in the winning vector
	possibleLosingSpaces := avb.Board.AvailableSpaces()
	losingSpaces := []int{}
	for _, losingSpace := range possibleLosingSpaces {
		taken := false
		for _, winningSpace := range winningVector {
			if losingSpace == winningSpace {
				taken = true
				break
			}
		}
		if !taken {
			losingSpaces = append(losingSpaces, losingSpace)
		}
	}
	rand.Shuffle(len(losingSpaces), func(i, j int) {
		losingSpaces[i], losingSpaces[j] = losingSpaces[j], losingSpaces[i]
	})
	losingSpaces = losingSpaces[:2]

	allSpaces := interleaveInts(winningVector, losingSpaces)
	for _, space := range allSpaces {
		avb.Board, _, _ = avb.Board.AssignSpace(space)
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
