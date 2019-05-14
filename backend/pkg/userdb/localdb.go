package userdb

type localDBHandler struct {
	userHandler
}

func (l *localDBHandler) Initialize(connectString string) error {
	// Local DB Handler use local DB, so there are no process here.
	return nil
}
