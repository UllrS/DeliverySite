package settings

import "fmt"

type Conf struct {
	DATABASE       string
	DATABASE_LOGIN string
	DATABASE_PWD   string
	DATABASE_TYPE  string
	DATABASE_ADDR  string
	DATABASE_NAME  string
}

func (c *Conf) GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", c.DATABASE_LOGIN, c.DATABASE_PWD, c.DATABASE_TYPE, c.DATABASE_ADDR, c.DATABASE_NAME)
}

var DB_config = Conf{
	DATABASE:       "mysql",
	DATABASE_LOGIN: "root",
	DATABASE_PWD:   "root",
	DATABASE_TYPE:  "tcp",
	DATABASE_ADDR:  "127.0.0.1:3306",
	DATABASE_NAME:  "golang"}
