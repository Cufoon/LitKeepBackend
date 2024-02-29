package schema

type BillKind struct {
	UserID        string
	KindID        string
	Name          string
	Description   string
	ReplaceSystem string
	Owner         string
}

type BillKindTree struct {
	BillKind
	Children []BillKind
}
