package core

type BaseSystem struct {
	Codename        string
	PreviousSystems []string
	NextSystems     []string
}

func (s BaseSystem) GetCodename() string {
	return s.Codename
}

func (s BaseSystem) GetPreviousSystems() []string {
	return s.PreviousSystems
}

func (s BaseSystem) GetNextSystems() []string {
	return s.NextSystems
}
