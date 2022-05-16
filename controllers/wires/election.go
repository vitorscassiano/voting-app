package wires

type PostElectionIn struct {
	Description string   `json:"description" binding:"required"`
	Candidates  []string `json:"candidates" binding:"required"`
}
