package databases

import (
	"fmt"
	"log"
	"os"
	"time"

	"Github.com/Yobubble/go-clean-boilerplate/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresqlDatabase struct {
	db *gorm.DB
}

// GetDB implements Database.
func (p *postgresqlDatabase) GetDB() *gorm.DB {
	return p.db
}

func NewPostgresqlDatabase(cfg configs.Config) Database {
	pgConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSLMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Enable color
		},
	)

	db, dbErr := gorm.Open(postgres.Open(pgConn), &gorm.Config{
		Logger: newLogger,
	})

	if dbErr != nil {
		panic(dbErr)
	}

	fmt.Printf("Connect to postgres with port : %d", cfg.DB.Port)
	return &postgresqlDatabase{
		db: db,
	}
}
