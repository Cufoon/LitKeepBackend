package vo

type BillKind struct {
	UserID      string `json:"userID" validate:"required"`
	KindID      string `json:"kindID" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
