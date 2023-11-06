package handlers

type PaginationParams struct {
	Page  int
	Limit int
}

const (
	defaultPage  = 1
	defaultLimit = 10
)

func (p PaginationParams) isValid() bool {
	return p.Page > 0 && p.Limit > 0
}
