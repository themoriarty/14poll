package src

import(
	"net/http"
	"appengine"
	"appengine/user"
	)

type httpFn func (http.ResponseWriter, *http.Request, *appengine.Context, *user.User)
type context struct{
	callback httpFn
}

func (this context) fn(w http.ResponseWriter, r *http.Request){
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
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	this.callback(w, r, &ctx, usr)
}

func AuthRequired(oldHandler httpFn) http.HandlerFunc{
	return context{oldHandler}.fn
}