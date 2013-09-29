package src

import (
	"fmt"
	"regexp"
	"net/http"
	"appengine"
	"appengine/user"
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
func pollVoteHandler(w http.ResponseWriter, r *http.Request, ctx *appengine.Context, usr *user.User, pollId string){
	if (r.Method != "POST"){
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "%s: %s", pollId, "vote")
}
func pollViewHandler(w http.ResponseWriter, r *http.Request, ctx *appengine.Context, usr *user.User, pollId string){
	fmt.Fprintf(w, "%s: %s", pollId, "view")
}
