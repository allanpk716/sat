package sat

type Dicter interface {
	Init(opts ...Option) error
	Read(string) string
	ReadReverse(string) string
	IsChs(s string, threshold float32) bool
}
