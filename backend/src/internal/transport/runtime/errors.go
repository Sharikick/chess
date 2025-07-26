package runtime

type RuntimeError string

func (e RuntimeError) Error() string {
	return string(e)
}

const (
	ErrMainOmitted RuntimeError = "main func not found"
)
