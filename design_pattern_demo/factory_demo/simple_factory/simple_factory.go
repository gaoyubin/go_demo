package simple_factory

import "fmt"

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}
func (g *Gun) GetName() string {
	return g.name
}
func (g *Gun) GetPower() int {
	return g.power
}

type AK47 struct {
	Gun
}

func NewAk47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "musket gun",
			power: 1,
		},
	}
}

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
