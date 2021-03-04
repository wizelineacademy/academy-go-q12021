package cards

// AuthorResponse struct defines response fields
type CardResponse struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  Type      string `json:"type"`
  Level     int    `json:"level"`
  Race      string `json:"race"`
  Attribute string `json:"attribute"`
  ATK       int    `json:"atk"`
  DEF       int    `json:"def"`
}

// ListResponse struct defines authors list response structure
type ListResponse struct {
  Data []CardResponse `json:"data"`
}
