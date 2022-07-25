package transformer

type Transformer interface {
	Transform(model any) any
}

type Manager struct {
	Serializer Serializer
}

// NewManager /**
func NewManager() *Manager {
	return &Manager{
		Serializer: NewJSONSerializer(),
	}
}

// CreateData /**
func (m *Manager) CreateData(data any) any {
	return m.serializeResource(data)
}

// SerializeResource /**
func (m *Manager) serializeResource(resource any) any {
	switch v := resource.(type) {
	case Collection:
		return m.Serializer.SerializeCollection(v)
	case Item:
		return m.Serializer.SerializeItem(v)
	case Error:
		return m.Serializer.SerializeError(v)
	default:
		return m.Serializer.SerializeNil()
	}
}
