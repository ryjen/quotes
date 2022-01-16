package storage

import (
	"context"
	"errors"
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

func (db *MemoryDB) PageQuotes(ctx context.Context, key string, count int) (*PagedQuotes, error) {
	values := []*Quote{}
	found := -1
	pos := 0
	var last *string
	for _, val := range db.quotes {
		if found == -1 {
			if val.ID != key {
				continue
			}
			found = pos
		}
		pos++
		if pos-count >= found {
			last = &val.ID
			break
		}
		values = append(values, val)
	}
	return &PagedQuotes{values, last}, nil
}

func (db *MemoryDB) AddQuote(ctx context.Context, quote QuoteWithOptionalID) (string, error) {
	id := quote.GetOrGenerateID()
	db.quotes[id] = &Quote{quote.QuoteWithoutID, id}
	return id, nil
}

func (db *MemoryDB) AddQuotes(ctx context.Context, quotes []QuoteWithOptionalID) ([]string, error) {
	var ids []string
	for _, quote := range quotes {
		id, err := db.AddQuote(ctx, quote)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (db *MemoryDB) SaveQuote(ctx context.Context, quote *Quote) error {
	if len(quote.ID) == 0 {
		return errors.New("not found")
	}
	db.quotes[quote.ID] = quote
	return nil
}

func (db *MemoryDB) GetQuote(ctx context.Context, id string) (*Quote, error) {
	q, ok := db.quotes[id]

	if !ok {
		return q, errors.New("not found")
	}

	return q, nil
}

func (db *MemoryDB) RemoveQuote(ctx context.Context, id string) error {
	_, ok := db.quotes[id]
	if !ok {
		return errors.New("not found")
	}
	db.quotes[id] = nil
	return nil
}
