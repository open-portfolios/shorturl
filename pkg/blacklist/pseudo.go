package blacklist

type pseudoBlacklist struct{}

func NewPseudo() Interface               { return pseudoBlacklist{} }
func (pseudoBlacklist) Good(string) bool { return true }
