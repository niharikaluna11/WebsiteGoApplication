package config

import (
	"OrderProcessingService/models"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
	mysqlGorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: environment variable %s is not set", key)
	}
	return value
}

func ConnectToDB() (*gorm.DB, error) {
	dbUser := mustGetenv("DB_USER")
	dbPwd := mustGetenv("DB_PASSWORD")
	dbName := mustGetenv("DB_NAME")
	instanceConnName := mustGetenv("INSTANCE_CONNECTION_NAME")

	dialer, err := cloudsqlconn.NewDialer(context.Background(), cloudsqlconn.WithRefreshTimeout(30*time.Second))
	if err != nil {
		return nil, fmt.Errorf("failed to create Cloud SQL dialer: %w", err)
	}

	mysql.RegisterDialContext("cloudsqlconn", func(ctx context.Context, _ string) (net.Conn, error) {
		opts := []cloudsqlconn.DialOption{}
		return dialer.Dial(ctx, instanceConnName, opts...)
	})

	dsnConfig := mysql.Config{
		User:                 dbUser,
		Passwd:               dbPwd,
		Net:                  "cloudsqlconn",
		Addr:                 "localhost:3306",
		DBName:               dbName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	dsn := dsnConfig.FormatDSN()

	db, err := gorm.Open(mysqlGorm.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Cloud SQL DB: %w", err)
	}

	if err := db.AutoMigrate(&models.Order{}, &models.Payment{}, &models.User{}); err != nil {
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}

	return db, nil
}