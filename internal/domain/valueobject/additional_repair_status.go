package valueobject

type AdditionalRepairStatus string

const (
	StatusARAberta              AdditionalRepairStatus = "ABERTA"
	StatusARAguardandoAprovacao AdditionalRepairStatus = "AGUARDANDO_APROVACAO"
	StatusAAprovada             AdditionalRepairStatus = "APROVADA"
	StatusARRejeitada           AdditionalRepairStatus = "REJEITADA"
	StatusARCancelada           AdditionalRepairStatus = "CANCELADA"
)

func ParseAdditionalRepairStatus(status string) AdditionalRepairStatus {
	switch status {
	case "ABERTA":
		return StatusARAberta
	case "AGUARDANDO_APROVACAO":
		return StatusARAguardandoAprovacao
	case "APROVADA":
		return StatusAAprovada
	case "REJEITADA":
		return StatusARRejeitada
	case "CANCELADA":
		return StatusARCancelada
	default:
		return AdditionalRepairStatus(status)
	}
}

func (s AdditionalRepairStatus) String() string {
	return string(s)
}

func (s AdditionalRepairStatus) IsAberta() bool {
	return s == StatusARAberta
}

func (s AdditionalRepairStatus) IsAguardandoAprovacao() bool {
	return s == StatusARAguardandoAprovacao
}

func (s AdditionalRepairStatus) IsAprovada() bool {
	return s == StatusAAprovada
}

func (s AdditionalRepairStatus) IsRejeitada() bool {
	return s == StatusARRejeitada
}

func (s AdditionalRepairStatus) IsCancelada() bool {
	return s == StatusARCancelada
}
