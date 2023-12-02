package mock

// EnvAccessor is a mock implementation of the EnvAccessor interface for testing.
type EnvAccessor struct {
	EnvVars map[string]string
}

func (m *EnvAccessor) GetEnv(key string) string {
	return m.EnvVars[key]
}
