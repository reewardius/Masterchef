package modules

type Module struct{}

type Scheme interface {
	Cook(string, []byte) (string, error)
	CookShh(string) (string, error)
	ToHTML() string
}
