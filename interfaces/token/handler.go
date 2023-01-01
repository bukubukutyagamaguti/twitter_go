//go:generate mockgen -source=handler.go -package=mock_token -destination=./mock/handler.go
package token

type TokenHandler interface {
	Generate(int, string) (string, error)
	VerityToken(string) error
}
