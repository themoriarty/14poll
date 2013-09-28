package core

import (
	"time"
	)

type User struct{
	Name string
	Email string
}

type Option struct{
	Name string
	ByUser User
	Created time.Time
}

const (
	VoteFor = iota
	VoteAgainst
	VoteNeutral
	VoteLast
)

type Vote struct{
	Value int
	Created time.Time
}

type OptionAndVotes struct{
	Option Option
	Votes map[int]Vote
}

type Poll struct{
	Id string
	Options []OptionAndVotes
	Users []User
}
