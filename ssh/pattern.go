package ssh

import (
	"bytes"
	"errors"
	"regexp"
)

// Pattern bulabula
type Pattern struct {
	str   string // Its appearance in the file, not the value that gets compiled.
	regex *regexp.Regexp
	not   bool // True if this is a negated match
}

// Copied from regexp.go with * and ? removed.
var specialBytes = []byte(`\.+()|[]{}^$`)

func special(b byte) bool {
	return bytes.IndexByte(specialBytes, b) >= 0
}

// NewPattern bulabula
func NewPattern(s string) (pattern *Pattern, err error) {
	if s == "" {
		err = errors.New("ssh_config: empty pattern")
		return
	}
	negated := false
	if s[0] == '!' {
		negated = true
		s = s[1:]
	}
	var buf bytes.Buffer
	buf.WriteByte('^')
	for i := 0; i < len(s); i++ {
		// A byte loop is correct because all metacharacters are ASCII.
		switch b := s[i]; b {
		case '*':
			buf.WriteString(".*")
		case '?':
			buf.WriteString(".?")
		default:
			// borrowing from QuoteMeta here.
			if special(b) {
				buf.WriteByte('\\')
			}
			buf.WriteByte(b)
		}
	}
	buf.WriteByte('$')
	r, err := regexp.Compile(buf.String())
	if err != nil {
		return
	}
	pattern = &Pattern{str: s, regex: r, not: negated}
	return
}

// String bulabula
func (p *Pattern) String() string {
	return p.str
}

// Match bulabula
func (p *Pattern) Match(s string, strict bool) bool {
	if strict {
		return p.str == s
	}
	return p.regex.MatchString(s)
}
