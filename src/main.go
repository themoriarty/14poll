package src

import (
	"fmt"
//	"time"
	"net/http"
	"appengine"
	"appengine/user"
//	"src/core"
	"src/nosurf"
)

func init(){
	http.Handle("/poll/", nosurf.New(AuthRequired(pollHandler)))
	http.Handle("/", nosurf.New(AuthRequired(handler)))
}

func handler(w http.ResponseWriter, r *http.Request, ctx *appengine.Context, usr *user.User){/*
	token := nosurf.Token(r)
	user1 := core.User{"testName1", "testEmail1"};
	user2 := core.User{"testName2", "testEmail2"};
	poll := core.Poll{
		"testPoll",
		[]core.OptionAndVotes{core.OptionAndVotes{core.Option{"op1", user1.UserId, time.Now()}, []core.UserVote{core.UserVote{user1.UserId, core.Vote{core.VoteNeutral, time.Now()}}, core.UserVote{user2.UserId, core.Vote{core.VoteAgainst, time.Now()}}}, true}, 
			core.OptionAndVotes{core.Option{"op2", user1.UserId, time.Now()}, []core.UserVote{core.UserVote{user2.UserId, core.Vote{core.VoteNeutral, time.Now()}}, core.UserVote{user1.UserId, core.Vote{core.VoteAgainst, time.Now()}}}, true}}, true};
	if _, err := FindPoll(ctx, "testPoll"); err != nil{
		err := StorePoll(ctx, &poll)
		if (err != nil){
			fmt.Fprintf(w, "Can't save poll: %s", err)
			return
		}
	}
	if pollV, err := FindPoll(ctx, "testPoll"); err != nil{
		fmt.Fprintf(w, "Can't load poll: %s", err)
		return
	} else{
		poll = *pollV
	}
	if out, err := poll.Render(token, core.UserId(usr.Email)); err != nil{
		fmt.Fprintf(w, "Can't render poll: %s", err)
	} else{
		fmt.Fprintf(w, "%s", out)
	}*/
	fmt.Fprintf(w, "%s", "Nothing to see there")
}

