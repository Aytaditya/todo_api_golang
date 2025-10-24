package storage

// all methods which will be implemented by different storage types (like sqlite, postgres etc) will be defined here (not necesarry step)
type Storage interface {
	CreateUser(username *string, email *string, password *string) (int64, error)
}
