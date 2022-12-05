package main

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

type Config struct {
	Host     string
	User     string
	Password string

	Schema string
	Port   string
}

func InitConf() Config {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	scheme := os.Getenv("DB_SCHEME")
	return Config{
		Host:     host,
		User:     user,
		Password: password,
		Schema:   scheme,
		Port:     port,
	}
}

func NewClient(conf *Config) (*Client, error) {
	db, err := connect(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database error: %w", err)
	}
	return &Client{
		db: db,
	}, nil
}

func connect(conf *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp([%s]:%s)/%s?charset=utf8mb4&interpolateParams=true&parseTime=true&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Schema)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open DB. err: %+v", err)
	}

	return db, nil
}

func (c *Client) InsertResult(ctx context.Context, data *Result) error {
	var retData Result
	if err := c.db.WithContext(ctx).Where("id = ?", data.ID).Assign(data).FirstOrCreate(&retData).Error; err != nil {
		return err
	}
	return nil
}
