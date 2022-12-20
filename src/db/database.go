package db

type IDatabase interface {
	Create(string, string) error
	Read(string) (string, error)
	Update(string, string) error
	Delete(string) error
}

type Database struct {
	service IDatabase
}

func NewDBService(database IDatabase) *Database {
	return &Database{service: database}
}

func (s *Database) Create(key, value string) error {
	return s.service.Create(key, value)
}

func (s *Database) Read(key string) (string, error) {
	return s.service.Read(key)
}

func (s *Database) Update(key, value string) error {
	return s.service.Update(key, value)
}

func (s *Database) Delete(key string) error {
	return s.service.Delete(key)
}
