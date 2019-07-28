package db

import (
	"database/sql"
)

// Config 数据库配置
type Config struct {

}
// NewDB 连接数据库,返回sql.DB和错误
func NewDB(c *Config) (*sql.DB,error) {
	return &sql.DB{},nil
}


