package action

type ActionEntity struct {
	ID          string `json:"id"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Label       string `json:"label"`
	Description string `json:"description"`
}
