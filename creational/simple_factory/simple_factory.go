package simple_factory

// ProductType is used as a enum for product type
type ProductType int

// Names of products as enum
const (
	ProductA ProductType = iota
	ProductB
)

// Product describes abstract product
type Product interface {
	show() string
}

// ConcreteProductA represents concrete product of A
type ConcreteProductA struct {
}

func (p ConcreteProductA) show() string {
	return "ProductA"
}

// ConcreteProductB represents concrete product of B
type ConcreteProductB struct {
}

func (p ConcreteProductB) show() string {
	return "ProductB"
}

// Factory represents product factory
type Factory struct {
}

// CreateProduct creates product
func (f Factory) CreateProduct(t ProductType) Product {
	if t == ProductA {
		return ConcreteProductA{}
	} else if t == ProductB {
		return ConcreteProductB{}
	}
	return nil
}
