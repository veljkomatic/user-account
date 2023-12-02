package log

type Fielder interface {
	LogFields() Fields
}

var _ Fielder = (*Fields)(nil)

type Fields map[string]any

func (f Fields) LogFields() Fields { return f }
