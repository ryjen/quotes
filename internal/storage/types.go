package storage

import (
	"context"
	"fmt"
)

type SavedQuote = map[string]interface{}

type QuoteWithoutID struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type Quote struct {
	QuoteWithoutID
	ID string `json:"id"`
}

type QuoteWithOptionalID struct {
	QuoteWithoutID
	ID *string `json:"id,omitempty"`
}

type PagedQuotes struct {
	Quotes []*Quote
	LastID *string
}

type Storage interface {
	AddQuote(context.Context, QuoteWithOptionalID) (string, error)
	AddQuotes(context.Context, []QuoteWithOptionalID) ([]string, error)
	ListQuotes(context.Context) ([]*Quote, error)
	PageQuotes(context.Context, string, int) (*PagedQuotes, error)
	GetQuote(context.Context, string) (*Quote, error)
	SaveQuote(context.Context, *Quote) error
	RemoveQuote(context.Context, string) error
}

func (q *Quote) String() string {
	return fmt.Sprintf("%s %v", q.ID, q.QuoteWithoutID)
}

func (q QuoteWithoutID) String() string {
	return fmt.Sprintf("%s - %s", q.Text, q.Author)
}

func (q *QuoteWithOptionalID) GetOrGenerateID() string {

	if q.ID == nil || len(*q.ID) == 0 {
		return NewID()
	}
	return *q.ID
}
