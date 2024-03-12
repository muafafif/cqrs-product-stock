package database

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	source string
	db     *gorm.DB
}

func New(source string) *PostgreSQL {
	return &PostgreSQL{source: source}
}

func (p *PostgreSQL) Connect() error {
	db, err := gorm.Open(postgres.Open(p.source), &gorm.Config{})
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *PostgreSQL) Create(ctx context.Context, model, value interface{}) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		// Transaction to save the value
		if err := tx.Model(model).Save(value).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Rollback needed when commiting the value get error
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func (p *PostgreSQL) Update(ctx context.Context, column string, model, value interface{}) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		// Transaction to update the value
		if err := tx.Model(model).Update(column, value).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Rollback needed when commiting the value get error
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func (p *PostgreSQL) Delete(ctx context.Context, model, value interface{}) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		// Transaction to delete the value
		if err := tx.Model(model).Delete(value).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Rollback needed when commiting the value get error
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}
