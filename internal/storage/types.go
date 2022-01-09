package storage

import "context"

type SavedQuote = map[string]interface{}

type NewQuote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type Quote struct {
	NewQuote
	ID string `json:"id"`
}

type Storage interface {
	ListQuotes(context.Context) ([]*Quote, error)
	SaveQuote(context.Context, NewQuote) (string, error)
	GetQuote(context.Context, string) (*Quote, error)
}
