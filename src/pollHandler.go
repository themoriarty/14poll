package src

import (
	"fmt"
	"regexp"
	"net/http"
//	"encoding/json"
//	"strings"
	"appengine"
	"appengine/user"
	"appengine/datastore"
	"src/core"
)

var pollView = regexp.MustCompile("^/poll/(\\w+)/?$")
var pollVote = regexp.MustCompile("^/poll/(\\w+)/vote$")
func pollHandler(w http.ResponseWriter, r *http.Request, ctx *appengine.Context, usr *user.User){
	if match := pollVote.FindStringSubmatch(r.URL.Path); match != nil{
		pollVoteHandler(w, r, ctx, usr, match[1]);
		return
	}
	if match := pollView.FindStringSubmatch(r.URL.Path); match != nil{
		pollViewHandler(w, r, ctx, usr, match[1]);
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s", "unknown url")
}

type vote struct{
	Poll string
	Option string
	Choice string
}

type response struct{
	Error string
	Data interface{}
}

func pollVoteHandler(w http.ResponseWriter, r *http.Request, ctx *appengine.Context, usr *user.User, pollId string){
	if (r.Method != "POST"){
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//voteDecoder := json.NewDecoder(strings.NewReader(r.FormValue("data")))
	//v := vote{r.FormValue("poll"), , };
	if err := datastore.RunInTransaction(*ctx, func(tc appengine.Context) error{
		poll, err := FindPoll(ctx, pollId)
		if err != nil{
			return err
		}
		if err := poll.CastVote(core.UserId(usr.String()), r.FormValue("option"), r.FormValue("choice")); err != nil{
			(*ctx).Errorf("can't save poll %s: %s", poll, err)
			return err
		}
		(*ctx).Infof("new poll: %s", poll)
		return StorePoll(ctx, poll)
	}, nil); err != nil{
		fmt.Fprintf(w, "{error: '%s'}", err); // TODO json
		return
	}
	fmt.Fprintf(w, "{ok: true}")
}
func pollViewHandler(w http.ResponseWriter, r *http.Request, ctx *appengine.Context, usr *user.User, pollId string){
	fmt.Fprintf(w, "%s: %s", pollId, "view")
}
