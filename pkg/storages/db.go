package storage

import (
	"log"

	"github.com/danik-tro/usd-rate/pkg/core"
	"github.com/danik-tro/usd-rate/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func initDb(config core.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBUri), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Subscriber{})

	return db
}

func NewStorage(config core.Config) Storage {
	db := initDb(config)

	return Storage{db}
}

func (s *Storage) IsEmailSubscribed(email string) (bool, error) {
	var es models.Subscriber
	if err := s.db.First(&es, "email = ?", email).Error; err == nil {
		return true, nil
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	} else {
		return false, err
	}
}

func (s *Storage) SubscribeEmail(sub *models.Subscriber) error {
	if err := s.db.Create(sub).Error; err != nil {
		return err
	}

	return nil
}

func (s *Storage) TotalSubscribers() int64 {
	var totalSubscribers int64
	s.db.Model(&models.Subscriber{}).Count(&totalSubscribers)
	return totalSubscribers
}

func (s *Storage) FetchBatchSubscribers(batchSize, offset int64) []models.Subscriber {
	var subscribers []models.Subscriber
	if r := s.db.Limit(int(batchSize)).Offset(int(offset)).Find(&subscribers); r.Error != nil {
		log.Printf("Failed to fetch batch of subscribers: %s", r.Error)
	}

	return subscribers
}
