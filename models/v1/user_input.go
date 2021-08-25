package v1

func NewUserInput() *UserInput {
	i := &UserInput{}
	return i
}

type UserInput struct {
	Direction
}

type Direction int

const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionLeft
	DirectionRight
	DirectionDown
)
