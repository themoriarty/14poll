package core


import (
	"path"
	"bytes"
	"html/template"
	)

const templates_dir = "src/templates"

var templates = template.Must(template.ParseFiles(path.Join(templates_dir, "user"), path.Join(templates_dir, "option_and_votes"), path.Join(templates_dir, "poll")))


type Context struct{
	Token string
	Data interface{}
	Show_results bool
	Show_vote bool
}

func (this *Context) Prepare(data interface{}) *Context{
	return &Context{this.Token, data, this.Show_results, this.Show_vote}
}

func (this *Context) OptionAndVotes() OptionAndVotes{
	return this.Data.(OptionAndVotes)
}
func itos(v int) string{
	return map[int]string{VoteFor: "VoteFor",
		VoteAgainst: "VoteAgainst",
	VoteNeutral: "VoteNeutral"}[v]
}

func (this OptionAndVotes) GetResult() map[string]int{
	ret := map[string]int {}
	for i := 0; i < VoteLast; i++{
		ret[itos(i)] = 0
	}
	for _, v := range(this.Votes){
		ret[itos(v.Value)]++
	}
	return ret
}
func (this OptionAndVotes) ToString(key string) string{
	return key + "!"
}


func (poll *Poll) Render(token string, show_results bool, show_vote bool) (string, error) {
	writer := bytes.NewBuffer(nil)
	if err := templates.ExecuteTemplate(writer, "poll", &Context{token, poll, show_results, show_vote}); err != nil{
		return "", err
	}
	return writer.String(), nil
}
