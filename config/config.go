package config

type database struct {
	User     string
	Password string
	Address  string
	DBName   string
}

//Database cointains the configuration and credentialls to connect to the database
var Database = database{
	User:     "admin",
	Password: "password",
	Address:  "127.0.0.1:3306",
	DBName:   "the_mirror",
}
