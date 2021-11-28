package solution

const NotImplemented = "not implemented"

type Provider interface {
	SolveA(input string) string
	SolveB(input string) string
}
