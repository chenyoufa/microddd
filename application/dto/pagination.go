package dto

type PaginationParam struct {
	Pagination bool `form:"-"`
	OnlyCount  bool `form:"-"`
	Current    int  `form:"current,default=1"`
	PageSize   int  `form:"pageSize,default=10" binding:"max=100"`
}

func (a PaginationParam) GetCurrent() int {
	return a.Current
}

func (a PaginationParam) GetPageSize() int {
	pageSize := a.PageSize
	if a.PageSize <= 0 {
		pageSize = 100
	}
	return pageSize
}

type OrderDirection int

const (
	OrderByASC OrderDirection = iota + 1
	OrderByDESC
)

// Create order fields key and define key index direction
func NewOrderFieldWithKeys(keys []string, directions ...map[int]OrderDirection) []*OrderField {
	m := make(map[int]OrderDirection)
	if len(directions) > 0 {
		m = directions[0]
	}

	fields := make([]*OrderField, len(keys))
	for i, key := range keys {
		d := OrderByASC
		if v, ok := m[i]; ok {
			d = v
		}

		fields[i] = NewOrderField(key, d)
	}

	return fields
}

func NewOrderFields(orderFields ...*OrderField) []*OrderField {
	return orderFields
}

func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}

type OrderField struct {
	Key       string
	Direction OrderDirection
}

// UserQueryOptions 查询可选参数项
type QueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}
