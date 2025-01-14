package models

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}
