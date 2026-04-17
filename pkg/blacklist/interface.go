package blacklist

type Interface interface {
	Good(string) bool
}
