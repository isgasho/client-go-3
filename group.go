package client_go

type Group struct {
	tranType TranType
	protocol Protocol
	branches []*Branch
}

func NewTccGroup(protocol Protocol, tryUri, confirmUri, cancelUri string) *Group {
	g := &Group{
		tranType: TCC,
		protocol: protocol,
	}
	g.branches = []*Branch{
		NewBranch(tryUri, Try),
		NewBranch(confirmUri, Confirm),
		NewBranch(cancelUri, Cancel),
	}
	g.SetProtocol(g.protocol)
	return g
}

func NewSagaGroup(protocol Protocol, normalUri, compensation string) *Group {
	g := &Group{
		tranType: SAGA,
		protocol: protocol,
	}
	g.branches = []*Branch{
		NewBranch(normalUri, Normal),
		NewBranch(compensation, Compensation),
	}
	g.SetProtocol(protocol)
	return g
}

func (g *Group) GetTranType() TranType {
	return g.tranType
}

func (g *Group) SetData(data []byte) *Group {
	g.set(func(branch *Branch) {
		branch.SetData(data)
	})
	return g
}

func (g *Group) SetProtocol(protocol Protocol) *Group {
	g.set(func(branch *Branch) {
		branch.SetProtocol(protocol)
	})
	return g
}

// todo should rename to  metadata?
func (g *Group) SetHeader(data []byte) *Group {
	g.set(func(branch *Branch) {
		branch.SetHeader(data)
	})
	return g
}

func (g *Group) SetLevel(level Level) *Group {
	g.set(func(branch *Branch) {
		branch.SetLevel(level)
	})
	return g
}

func (g *Group) set(fn func(branch *Branch)) {
	for _, branch := range g.branches {
		fn(branch)
	}
}
