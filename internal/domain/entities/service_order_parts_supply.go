package entities

// ServiceOrderPartsSupply represents the quantity of a parts supply allocated
// to a specific service order. It maps the many-to-many relationship while
// keeping the use case layer free from persistence details.
type ServiceOrderPartsSupply struct {
	PartsSupplyID  string
	ServiceOrderID string
	Quantity       int
}
