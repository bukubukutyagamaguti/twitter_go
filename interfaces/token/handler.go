package token

type TokenHandler interface {
	Generate(int, string) (string, error)
	VerityToken(string) error
}
