package regmodel

type RegOutput struct {
	IsRegistered bool
}

func NewRegOutput(r bool) *RegOutput {
	return &RegOutput{
		IsRegistered: r,
	}
}
