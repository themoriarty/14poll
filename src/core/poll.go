package core

import (
	"time"
	)

type UserId string

type User struct{
	UserId
	Email string
}

type Option struct{
	Name string
	ByUser UserId
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

type UserVote struct{
	UserId UserId
	Vote Vote
}

type OptionAndVotes struct{
	Option Option
	Votes []UserVote
	Dirty bool `datastore:"-"`
}

type Poll struct{
	Id string
	Options []OptionAndVotes `datastore:"-"`
	Dirty bool `datastore:"-"`
}

func (this OptionAndVotes) DoneFor(userId UserId) bool{
	for _, vote := range(this.Votes){
		if vote.UserId == userId{
			return true
		}
	}
	return false
}

func (this *Poll) DoneFor(userId UserId) bool{
	for _, option := range(this.Options){
		if !option.DoneFor(userId){
			return false
		}
	}
	return true
}

func (this *Poll) CastVote(userId UserId, optionName string, choiceString string) error{
	vote := Vote{VoteStoi(choiceString), time.Now()}
	for i, option := range(this.Options){
		if option.Option.Name == optionName{
			newVotes := []UserVote{}
			for _, vote := range(option.Votes){
				if vote.UserId != userId{
					newVotes = append(newVotes, vote)
				}
			}
			this.Options[i].Votes = append(newVotes, UserVote{userId, vote})
			this.Options[i].Dirty = true
			return nil
		}
	}
	newOption := OptionAndVotes{Option{optionName, userId, time.Now()}, []UserVote{UserVote{userId, vote}}, true}
	this.Options = append(this.Options, newOption)
	return nil
}