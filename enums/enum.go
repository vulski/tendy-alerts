package enums

type Enum interface {
	Value() interface{}
	Is(interface{}) bool
}
