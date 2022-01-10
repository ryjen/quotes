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

func (db *MemoryDB) AddQuote(ctx context.Context, quote NewQuote) (string, error) {
	var id string
	if quote.ID == nil {
		id = newId()
		quote.ID = &id
	} else {
		id = *quote.ID
	}
	db.quotes[id] = &Quote{quote.quoteFields, id}
	return id, nil
}

func (db *MemoryDB) SaveQuote(ctx context.Context, quote *Quote) error {
	db.quotes[quote.ID] = quote
	return nil
}
func (db *MemoryDB) GetQuote(ctx context.Context, id string) (*Quote, error) {
	return db.quotes[id], nil
}

func (db *MemoryDB) RemoveQuote(ctx context.Context, id string) error {
	db.quotes[id] = nil
	return nil
}
