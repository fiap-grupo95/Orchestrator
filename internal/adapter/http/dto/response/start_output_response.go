package response

type StartOutput struct {
	OSID     string `json:"os_id"`
	BudgetID string `json:"budget_id"`
	Status   string `json:"status"` // COMPLETED / COMPENSATED / FAILED
}
