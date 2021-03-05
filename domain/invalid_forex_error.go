package domain

import "fmt"

type InvalidForexError struct {
	Line int
	Col  int
}

func (e *InvalidForexError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}
