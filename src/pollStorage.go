package src

import (
//	"fmt"
	"src/core"
	"appengine"
	"appengine/datastore"
)

func FindPoll(ctx *appengine.Context, id string) (*core.Poll, error){
	pollKey := datastore.NewKey(*ctx, "poll", id, 0, nil)
	poll := new(core.Poll)
	err := datastore.Get(*ctx, pollKey, poll)
	if (err != nil){
		(*ctx).Errorf("Can't load poll: %s", err)
		return nil, err
	}
	optionsQuery := datastore.NewQuery("option").Ancestor(pollKey)
	options := make([]core.OptionAndVotes, 0, 0)
	if _, err := optionsQuery.GetAll(*ctx, &options); err != nil{
		(*ctx).Errorf("Can't load options for poll %s: %s", poll, err)
		return nil, err
	}
	poll.Options = options
	return poll, nil
}

func StorePoll(ctx *appengine.Context, poll *core.Poll) error{	
	pollKey := datastore.NewKey(*ctx, "poll", poll.Id, 0, nil)
	if poll.Dirty{
		if _, err := datastore.Put(*ctx, pollKey, poll); err != nil{
			(*ctx).Errorf("Can't store poll %s: %s", poll, err)
			return err
		}
	}
	for _, option := range(poll.Options){
		key := datastore.NewKey(*ctx, "option", option.Option.Name, 0, pollKey)
		if option.Dirty{
			if _, err := datastore.Put(*ctx, key, &option); err != nil{
				(*ctx).Errorf("Can't store option %s: %s", option, err)
				return err
			}
		}
	}

	for i := range(poll.Options){
		poll.Options[i].Dirty = false
	}
	poll.Dirty = false
	return nil
}

