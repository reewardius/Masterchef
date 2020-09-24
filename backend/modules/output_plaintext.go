package modules

var OutputPlainText = &Module{}

type options struct {
	Output string `json:"Output"`
}

func (Module) Cook(_ string, opts []byte) (string, error) {
	return "", nil
}

func (Module) CookShh(string) (string, error) {
	return "", nil
}

func (Module) ToHTML() string {
	return ""
}
