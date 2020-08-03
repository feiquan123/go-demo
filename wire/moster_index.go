package main

type MosterParam string

type Moster struct {
	Name string
}

func NewMoster(name MosterParam) Moster {
	return Moster{
		Name: string(name),
	}
}
