package systemssorter

type ISystem interface {
	GetCodename() string

	GetPreviousSystems() []string
	GetNextSystems() []string
}
