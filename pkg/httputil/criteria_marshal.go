package httputil

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/maestre3d/academy-go-q12021/internal/repository"
)

// UnmarshalCriteriaJSON parses the given request body into a repository.Criteria struct
func UnmarshalCriteriaJSON(r *http.Request) (repository.Criteria, error) {
	c := repository.Criteria{}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return repository.Criteria{}, err
	}
	return c, nil
}

// UnmarshalCriteria parses the given request queries into a repository.Criteria struct
func UnmarshalCriteria(r *http.Request) repository.Criteria {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	criteria := *repository.NewCriteria(limit, r.URL.Query().Get("next_page"))
	criteria.ActiveOnly, _ = strconv.ParseBool(r.URL.Query().Get("active_only"))
	criteria.Query.Filters["items_per_worker"] = repository.Filter{
		Condition: repository.EqualsCondition,
		Value:     r.URL.Query().Get("items_per_worker"),
	}
	criteria.Query.Filters["type"] = repository.Filter{
		Condition: repository.EqualsCondition,
		Value:     r.URL.Query().Get("type"),
	}
	return criteria
}
