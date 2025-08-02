package facade

import "fmt"

type Projector struct{}

func (p *Projector) TurnOn() {
	fmt.Println("Projector's on")
}

func (p *Projector) TurnOff() {
	fmt.Println("Projector's off")
}

type SoundSystem struct{}

func (s *SoundSystem) TurnOn() {
	fmt.Println("SoundSystem's on")
}

func (s *SoundSystem) TurnOff() {
	fmt.Println("SoundSystem's off")
}

type Screen struct{}

func (s *Screen) Lower() {
	fmt.Println("Screen's lowered")
}

func (s *Screen) Raise() {
	fmt.Println("Screen's raised")
}

type LightingSystem struct{}

func (l *LightingSystem) Dim() {
	fmt.Println("LightingSystem's dimmed")
}

func (l *LightingSystem) Brighten() {
	fmt.Println("LightingSystem's brightened")
}

type HomeTheaterFacade struct {
	Projector      *Projector
	SoundSystem    *SoundSystem
	Screen         *Screen
	LightingSystem *LightingSystem
}

func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("Get ready to watch a movie")
	h.LightingSystem.Dim()
	h.Screen.Lower()
	h.Projector.TurnOn()
	h.SoundSystem.TurnOn()
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down")
	h.Projector.TurnOff()
	h.SoundSystem.TurnOff()
	h.Screen.Raise()
	h.LightingSystem.Brighten()
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		Projector:      &Projector{},
		SoundSystem:    &SoundSystem{},
		Screen:         &Screen{},
		LightingSystem: &LightingSystem{},
	}
}
