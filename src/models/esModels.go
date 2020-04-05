package models

type ESResponse struct {
	OuterHits OuterHits `json:"hits"`
}

type OuterHits struct {
	Total     Total      `json:"total"`
	InnerHits []InnerHit `json:"hits"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type InnerHit struct {
	ID     string `json:"_id"`
	Source User   `json:"_source"`
}
