package keygen

type Generator interface {
	Generate() (string, error)
}
