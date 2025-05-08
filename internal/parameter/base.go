package parameter

type Parameters interface {
	SetDefaults()
}

type BaseParameter struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

func (p *BaseParameter) SetDefaults() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 || p.Limit > 100 {
		p.Limit = 10
	}
}
