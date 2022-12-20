package factory

import (
	serv "msg-app/src/db"
	impl "msg-app/src/db/impl"
)

func MemoryStorageFactory() *serv.Database {
	return serv.NewDBService(&impl.DBMemory{})
}
