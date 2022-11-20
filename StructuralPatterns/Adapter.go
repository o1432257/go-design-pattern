package StructuralPatterns

const (
	knife    = "砍"
	shutDown = "電腦關機"
)

type Weapon interface {
	Attack() string
}

type Knife struct {
}

func (k *Knife) Attack() string {
	return knife
}

type Hero struct {
	Name   string
	weapon Weapon
}

func (h *Hero) Skill() string {
	return h.weapon.Attack()
}

// PowerOff 不適配的技能
type PowerOff struct{}

func (p *PowerOff) ShutDown() string {
	return shutDown
}

// Adapter 適配器
type Adapter struct {
	powerOff *PowerOff
}

func (a *Adapter) Attack() string {
	return a.powerOff.ShutDown()
}

func NewAdapter(p *PowerOff) *Adapter {
	return &Adapter{p}
}
