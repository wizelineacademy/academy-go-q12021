package repository

const (
	// simple statements

	AndCondition = "AND"
	OrCondition  = "OR"

	// advanced

	EqualsCondition         = "="
	EqualOrGreaterCondition = ">="
	EqualOrLessCondition    = "<="
	GreaterCondition        = ">"
	LessCondition           = "<"
	BetweenCondition        = "BETWEEN"
	InCondition             = "IN"
	ExistsCondition         = "EXISTS"
	ContainsCondition       = "CONTAINS"
	BeginsWithCondition     = "BEGINS_WITH"
)

// Criteria Movies application DSL query lang
type Criteria struct {
	Limit      int    `json:"limit"`
	NextPage   string `json:"next_page"`
	ActiveOnly bool   `json:"active_only"`
	Query      Query  `json:"query"`
}

// NewCriteria creates and ensures criteria
func NewCriteria(l int, nextPage string) *Criteria {
	if l <= 0 || l > 100 {
		l = 100
	}
	return &Criteria{
		Limit:      l,
		NextPage:   nextPage,
		ActiveOnly: true,
		Query: Query{
			Filters: map[string]Filter{},
		},
	}
}

// Query DTO used to fetch specific data
type Query struct {
	Condition string            `json:"condition"`
	Filters   map[string]Filter `json:"filters"`
}

// FilterExists verifies if the given filter key exists in the current query map
func (q Query) FilterExists(f string) bool {
	return q.Filters[f] != Filter{}
}

// IsAnd returns true if the current global condition is AND
func (q Query) IsAnd() bool {
	return q.Condition == AndCondition
}

// IsOr returns true if the current global condition is OR
func (q Query) IsOr() bool {
	return q.Condition == OrCondition
}

// Filter DTO which contains a condition
type Filter struct {
	Condition string      `json:"condition"`
	Negate    bool        `json:"negate"`
	Value     interface{} `json:"value"`
	AltValue  interface{} `json:"alt_value"`
}
