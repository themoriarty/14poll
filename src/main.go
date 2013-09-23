package src

import (
	"fmt"
	"time"
	"net/http"
	"appengine"
	"appengine/user"
	"src/core"
)

func init(){
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	usr := user.Current(ctx)

	if usr == nil{
		url, err := user.LoginURL(ctx, r.URL.String())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	user1 := core.User{"testName1", "testEmail1"};
	user2 := core.User{"testName2", "testEmail2"};
	poll := core.Poll{
		[]core.OptionAndVotes{core.OptionAndVotes{core.Option{"op1", user1, time.Now()}, map[int]core.Vote{0: core.Vote{core.VoteNeutral, time.Now()}, 1: core.Vote{core.VoteAgainst, time.Now()}}}},
		[]core.User{user1, user2}};
	if out, err := poll.Render(true, true); err != nil{
		fmt.Fprintf(w, "Can't render poll: %s", err)
	} else{
		fmt.Fprintf(w, "hello, %v! %s", usr, out)
	}	
}