package storage

type Storage interface {
	CreateUser(username *string, email *string, password *string) (int64, error)
}
