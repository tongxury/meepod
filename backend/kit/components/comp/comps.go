package comp

import (
	"gitee.com/meepo/backend/kit/components/auth"
	"gitee.com/meepo/backend/kit/components/auth/verification"
)

type PreComp struct {
	comps *Components
}

var preCompIns = &PreComp{
	comps: &Components{},
}

func Comps() *Components {
	return preCompIns.comps
}

type Components struct {
	auth auth_comp.IAuthComp
}

func (c *Components) Preparing() *PreComp {
	return preCompIns
}

func (p *PreComp) Auth(authVerifyComp verification_comp.IAuthVerifyComp, authStore auth_comp.IAuthStore) *PreComp {
	p.comps.auth = auth_comp.Assemble(authVerifyComp, authStore)
	return p
}

func (c *Components) Auth() auth_comp.IAuthComp {
	if c.auth == nil {
		panic("auth comp is not initialized")
	}

	return c.auth
}
