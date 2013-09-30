package core


import (
	"path"
	"bytes"
	"html/template"
	)

const templates_dir = "src/templates"

var templates = template.Must(template.ParseFiles(path.Join(templates_dir, "user"), path.Join(templates_dir, "option_and_votes"), path.Join(templates_dir, "poll")))


type Context struct{
	UserId UserId
	Token string
	Data interface{}
	ShowResults bool
}

func (this *Context) Prepare(data interface{}) *Context{
	return &Context{this.UserId, this.Token, data, this.ShowResults}
}

func (this *Context) OptionAndVotes() OptionAndVotes{
	return this.Data.(OptionAndVotes)
}
func itos(v int) string{
	return map[int]string{VoteFor: "VoteFor",
		VoteAgainst: "VoteAgainst",
	VoteNeutral: "VoteNeutral"}[v]
}
func VoteStoi(v string) int{
	return map[string]int{"VoteFor": VoteFor,
		"VoteAgainst": VoteAgainst,
	"VoteNeutral": VoteNeutral}[v]
}


func (this OptionAndVotes) GetResult() map[string]int{
	ret := map[string]int {}
	for i := 0; i < VoteLast; i++{
		ret[itos(i)] = 0
	}
	for _, v := range(this.Votes){
		ret[itos(v.Vote.Value)]++
	}
	return ret
}
func (this OptionAndVotes) ToString(key string) string{
	return key + "!"
}


func (poll *Poll) Render(token string, usr UserId) (string, error) {
	writer := bytes.NewBuffer(nil)
	if err := templates.ExecuteTemplate(writer, "poll", &Context{usr, token, poll, poll.DoneFor(usr)}); err != nil{
		return "", err
	}
	return writer.String(), nil
}
