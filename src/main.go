package src

import (
	"fmt"
	"time"
	"net/http"
	"appengine"
	"appengine/user"
	"src/core"
	"src/nosurf"
)

func init(){
	http.Handle("/poll/", nosurf.New(AuthRequired(pollHandler)))
	http.Handle("/", nosurf.New(AuthRequired(handler)))
}

func handler(w http.ResponseWriter, r *http.Request, ctx *appengine.Context, usr *user.User){
	token := nosurf.Token(r)
	w.Header().Set("Content-type", "text/html; encoding=utf-8")
	user1 := core.User{"testName1", "testEmail1"};
	user2 := core.User{"testName2", "testEmail2"};
	poll := core.Poll{
		"testPoll",
		[]core.OptionAndVotes{core.OptionAndVotes{core.Option{"op1", user1, time.Now()}, map[int]core.Vote{0: core.Vote{core.VoteNeutral, time.Now()}, 1: core.Vote{core.VoteAgainst, time.Now()}}}, core.OptionAndVotes{core.Option{"op2", user1, time.Now()}, map[int]core.Vote{0: core.Vote{core.VoteNeutral, time.Now()}, 1: core.Vote{core.VoteAgainst, time.Now()}}}},
		[]core.User{user1, user2}};
	if out, err := poll.Render(token, true, true); err != nil{
		fmt.Fprintf(w, "Can't render poll: %s", err)
	} else{
		fmt.Fprintf(w, "%s", out)
	}	
}

