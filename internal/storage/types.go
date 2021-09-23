package storage

import "context"

type SavedQuote = map[string]interface{}

type BaseQuote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type Quote struct {
	BaseQuote
	ID string `json:"id"`
}

type Storage interface {
	ListQuotes(context.Context) ([]*Quote, error)
	SaveQuote(context.Context, BaseQuote) (string, error)
	GetQuote(context.Context, string) (*Quote, error)
}
