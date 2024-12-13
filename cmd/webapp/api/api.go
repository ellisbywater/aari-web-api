package webapp

import "context"

type API struct {
	Users interface {
		Register(context.Context)
		Update(context.Context)
		Delete(context.Context)
	}
	Positions interface {
		Create(context.Context)
		GetById(context.Context)
		Update(context.Context)
		Delete(context.Context)
	}
}
