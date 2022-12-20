package impl

import "fmt"

type DBMemory struct {
	data map[string]string
}

func (s *DBMemory) Create(key, value string) error {
	s.data[key] = value
	return nil
}

func (s *DBMemory) Read(key string) (string, error) {
	value, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return value, nil
}

func (s *DBMemory) Update(key, value string) error {
	_, ok := s.data[key]
	if !ok {
		return fmt.Errorf("key not found: %s", key)
	}
	s.data[key] = value
	return nil
}

func (s *DBMemory) Delete(key string) error {
	_, ok := s.data[key]
	if !ok {
		return fmt.Errorf("key not found: %s", key)
	}
	delete(s.data, key)
	return nil
}
