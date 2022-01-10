package storage

import "context"

type SavedQuote = map[string]interface{}

type quoteFields struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type NewQuote struct {
	quoteFields
	ID *string `json:"id,omitempty"`
}

type Quote struct {
	quoteFields
	ID string `json:"id"`
}

type Storage interface {
	AddQuote(context.Context, NewQuote) (string, error)
	ListQuotes(context.Context) ([]*Quote, error)
	GetQuote(context.Context, string) (*Quote, error)
	SaveQuote(context.Context, *Quote) error
	RemoveQuote(context.Context, string) error
}
