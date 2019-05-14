package userdb

type localDBHandler struct {
	userHandler
}

func (l *localDBHandler) ConnectDB(connectString string) error {
	// Local DB Handler use local DB, so there are no process here.
	return nil
}
