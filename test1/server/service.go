package main

type ServiceA struct {
}

type Aras struct {
	X, Y int
}

func (s *ServiceA) Add(aras *Aras, replay *int) error {
	*replay = aras.Y + aras.X
	return nil
}
