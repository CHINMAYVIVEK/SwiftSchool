package serversync

func NewServerSync() *ServerSync {
	return &ServerSync{}
}

type ServerSync struct {
}

func (s *ServerSync) Sync() error {
	return nil
}
