package main

import "sync"

type DbConnection struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}

var connectionPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return &DbConnection{
			Host:     "localhost",
			Port:     3306,
			DbName:   "test",
			User:     "root",
			Password: "123456",
		}
	},
}

func main() {
	connection := connectionPool.Get().(*DbConnection)
	println(connection.Host, connection.Port, connection.DbName, connection.User, connection.Password)
}
