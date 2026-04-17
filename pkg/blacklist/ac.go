package blacklist

import ac "github.com/cloudflare/ahocorasick"

type acBlacklist struct {
	matcher *ac.Matcher
}

func (a acBlacklist) Good(s string) bool {
	return !a.matcher.Contains([]byte(s))
}

type acBlacklistBuilder []string

func NewACBuilder() *acBlacklistBuilder {
	var builder acBlacklistBuilder
	return &builder
}

func (b *acBlacklistBuilder) Add(word string) *acBlacklistBuilder {
	*b = append(*b, word)
	return b
}

func (b acBlacklistBuilder) Build() Interface {
	if len(b) == 0 {
		return pseudoBlacklist{}
	}
	return &acBlacklist{ac.NewStringMatcher(b)}
}
