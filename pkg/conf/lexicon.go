package conf

// Chars building blocks composing lexicons
type Chars struct {
	Name    string `yaml:"char"`
	Unicode string `yaml:"unicode"`
}

// LexiconSet set of lexicons
type LexiconSet struct {
	Item []Chars `yaml:"lexicon"`
}

var lexicon map[string][]Chars

// GetLexicon prepares the lexicon to be used by downstream packages
func (confRcv *ConfigRcv) GetLexicon() map[string][]Chars {
	return lexicon
}
