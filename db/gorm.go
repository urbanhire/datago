package db

import (
	"context"
	"errors"
	grm "github.com/jinzhu/gorm"
)

type Gorm interface {
	Ping() error
	GetConnection() (*grm.DB, error)
}

type gorm struct {
	vendor string
	uri      string
	ctx      context.Context
	db       *grm.DB
}

//NewGorm : instance of Golang Object Relational Mapping
func NewGorm(ctx context.Context, vendor string, uri string) Gorm {
	return &gorm{ctx: ctx, vendor: vendor, uri: uri}
}

func (g *gorm) validate() error {
	if len(g.vendor) == 0 {
		return errors.New("Vendor is empty")
	}
	if len(g.uri) == 0 {
		return errors.New("uri is empty")
	}
	return nil
}

func (g *gorm) GetConnection() (*grm.DB, error) {
	var err error

	if err = g.Ping(); err == nil {
		return g.db, err
	}

	if err = g.validate(); err != nil {
		return nil, err
	}

	g.db, err = grm.Open(g.vendor, g.uri)
	if err != nil {
		return nil, err
	}

	return g.db, nil
}

func (g *gorm) Ping() error {
	if g.db == nil {
		return errors.New("gorm instance is nil")
	}
	return g.db.DB().Ping()
}
