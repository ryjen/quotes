package storage

import (
	"context"
)

type MemoryDB struct {
	quotes map[string]*Quote
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		quotes: map[string]*Quote{},
	}
}

func (db *MemoryDB) ListQuotes(ctx context.Context) ([]*Quote, error) {
	values := []*Quote{}
	for _, val := range db.quotes {
		values = append(values, val)
	}
	return values, nil
}

func (db *MemoryDB) SaveQuote(ctx context.Context, quote BaseQuote) (string, error) {
	newQuote := &Quote{quote, newId()}
	db.quotes[newQuote.ID] = newQuote
	return newQuote.ID, nil
}

func (db *MemoryDB) GetQuote(ctx context.Context, id string) (*Quote, error) {
	return db.quotes[id], nil
}
