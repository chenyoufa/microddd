// Package aggregate holds aggregates that combines many entities into a full object
package aggregate

import (
	"errors"

	"gomicroddd/domain/entity"
	"gomicroddd/domain/valueobject"

	"github.com/google/uuid"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	Person *entity.Person
	// a customer can hold many products
	products []*entity.Item
	// a customer can perform many transactions
	transactions []valueobject.Transaction
}

//外面的聚合不能访问内部聚合,除了聚合根 所以聚合根是大写
//只能通过聚合根来导航
//实体是可修改的所以是指针
//值对象是只读的所以是结构体的值
//聚合与聚合之间通过聚合根的id引用，所以聚合根应该是可访问的
func (p *Customer) GetProducts() []*entity.Item {
	return p.products
}

func (p *Customer) GetTransactions() []valueobject.Transaction {
	return p.transactions
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Create a new person and generate ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}
	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		Person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.Person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.Person == nil {
		c.Person = &entity.Person{}
	}
	c.Person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.Person == nil {
		c.Person = &entity.Person{}
	}
	c.Person.Name = name
}

// SetName changes the name of the Customer
func (c *Customer) GetName() string {
	return c.Person.Name
}
