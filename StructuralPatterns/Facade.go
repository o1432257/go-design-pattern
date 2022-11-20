package StructuralPatterns

type TV struct {
	status bool //開關
}

func (t *TV) On() {
	t.status = true
}

func (t *TV) Off() {
	t.status = false
}

type Light struct {
	status bool //開關
}

func (l *Light) On() {
	l.status = true
}

func (l *Light) Off() {
	l.status = false
}

type Audio struct {
	status bool //開關
}

func (a *Audio) On() {
	a.status = true
}

func (a *Audio) Off() {
	a.status = false
}

type LivingRoom struct {
	tv    TV
	light Light
	audio Audio
}

func (hp *LivingRoom) DoMovie() {
	hp.tv.On()
	hp.audio.On()
	hp.light.Off()
}

func (hp *LivingRoom) DoDinner() {
	hp.tv.Off()
	hp.audio.Off()
	hp.light.On()
}
