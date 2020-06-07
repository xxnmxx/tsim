package tsim

type ObjectType string

const (
	CORP_OBJ = "CORP"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
