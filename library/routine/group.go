package routine

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

type Group struct {
	mac *MaxAmountCtrl
	egp *errgroup.Group
}

func NewGroup(max int) *Group {
	return &Group{
		mac: NewMaxAmountCtrl(max),
		egp: &errgroup.Group{},
	}
}

func NewGroupWithContext(max int, ctx context.Context) (*Group, context.Context) {
	egp, ctx := errgroup.WithContext(ctx)
	return &Group{
		mac: NewMaxAmountCtrl(max),
		egp: egp,
	}, ctx
}

func (g *Group) Go(f func() error) {
	g.mac.Incr()
	g.egp.Go(func() error {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("panic err:", err)
			}
			g.mac.Decr()
		}()
		return f()
	})
}

func (g *Group) Wait() error {
	return g.egp.Wait()
}
