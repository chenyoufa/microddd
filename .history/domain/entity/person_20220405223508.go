package entity

import (
	"gomicroddd/domain/valueobject"

	"github.com/google/uuid"
)

// Person is a entity that represents a person in all Domains
type Person struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID
	// Name is the name of the person
	Name string
	// Age is the age of the person
	Age int
}

var products []*Item

// a customer can perform many transactions
var transactions []valueobject.Transaction

func (p *Person) GetProducts() []*Item {
	return products
}
