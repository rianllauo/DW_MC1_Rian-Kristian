package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(nama string) *Profile {
	p := Profile{Name: nama}

	if p.Name == "sasuke" {
		p.Name = nama
		p.Health = 200
		p.Power = 100
		p.Exp = 0
	}

	if p.Name == "goku" {
		p.Name = nama
		p.Health = 400
		p.Power = 300
		p.Exp = 100
	}

	if p.Name == "naruto" {
		p.Name = nama
		p.Health = 150
		p.Power = 200
		p.Exp = 50
	}

	return &p
}

func PowerUp(p *Profile, m int) *Profile {
	pa := p

	if pa.Name == p.Name {
		pa.Health = pa.Health + (pa.Health * m)
		pa.Power = pa.Power + (pa.Power * m)
		pa.Exp = pa.Exp + (pa.Exp * m)
	}

	return pa
}

func main() {
	profile := MakeProfile("naruto")
	fmt.Println("Name :", profile.Name)
	fmt.Println("Health :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 5)
	fmt.Println("Name :", powerUp.Name)
	fmt.Println("Health :", powerUp.Health)
	fmt.Println("Power :", powerUp.Power)
	fmt.Println("Exp :", powerUp.Exp)
}
