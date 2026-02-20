package valueobject

import "strings"

type ServiceOrderStatus string

const (
	StatusRecebida            ServiceOrderStatus = "RECEBIDA"
	StatusEmDiagnostico       ServiceOrderStatus = "EM DIAGNÓSTICO"
	StatusAguardandoAprovacao ServiceOrderStatus = "AGUARDANDO APROVAÇÃO"
	StatusAprovada            ServiceOrderStatus = "APROVADA"
	StatusRejeitada           ServiceOrderStatus = "REJEITADA"
	StatusEmExecucao          ServiceOrderStatus = "EM EXECUÇÃO"
	StatusFinalizada          ServiceOrderStatus = "FINALIZADA"
	StatusEntregue            ServiceOrderStatus = "ENTREGUE"
	StatusCancelada           ServiceOrderStatus = "CANCELADA"
)

var (
	normalizedRecebida            = normalizeStatus(string(StatusRecebida))
	normalizedEmDiagnostico       = normalizeStatus(string(StatusEmDiagnostico))
	normalizedAguardandoAprovacao = normalizeStatus(string(StatusAguardandoAprovacao))
	normalizedAprovada            = normalizeStatus(string(StatusAprovada))
	normalizedRejeitada           = normalizeStatus(string(StatusRejeitada))
	normalizedEmExecucao          = normalizeStatus(string(StatusEmExecucao))
	normalizedFinalizada          = normalizeStatus(string(StatusFinalizada))
	normalizedEntregue            = normalizeStatus(string(StatusEntregue))
	normalizedCancelada           = normalizeStatus(string(StatusCancelada))
)

func ParseServiceOrderStatus(status string) ServiceOrderStatus {
	switch normalizeStatus(status) {
	case normalizedRecebida:
		return StatusRecebida
	case normalizedEmDiagnostico:
		return StatusEmDiagnostico
	case normalizedAguardandoAprovacao:
		return StatusAguardandoAprovacao
	case normalizedAprovada:
		return StatusAprovada
	case normalizedRejeitada:
		return StatusRejeitada
	case normalizedEmExecucao:
		return StatusEmExecucao
	case normalizedFinalizada:
		return StatusFinalizada
	case normalizedEntregue:
		return StatusEntregue
	case normalizedCancelada:
		return StatusCancelada
	default:
		return ServiceOrderStatus(status)
	}
}

func (s ServiceOrderStatus) IsValid() bool {
	switch normalizeStatus(string(s)) {
	case normalizedRecebida,
		normalizedEmDiagnostico,
		normalizedAguardandoAprovacao,
		normalizedAprovada,
		normalizedRejeitada,
		normalizedEmExecucao,
		normalizedFinalizada,
		normalizedEntregue,
		normalizedCancelada:
		return true
	default:
		return false
	}
}

func (s ServiceOrderStatus) IsSame(c ServiceOrderStatus) bool {
	return normalizeStatus(string(s)) == normalizeStatus(string(c))
}

func (s ServiceOrderStatus) IsRecebida() bool {
	return normalizeStatus(string(s)) == normalizedRecebida
}

func (s ServiceOrderStatus) IsEmDiagnostico() bool {
	return normalizeStatus(string(s)) == normalizedEmDiagnostico
}

func (s ServiceOrderStatus) IsAguardandoAprovacao() bool {
	return normalizeStatus(string(s)) == normalizedAguardandoAprovacao
}

func (s ServiceOrderStatus) IsAprovada() bool {
	return normalizeStatus(string(s)) == normalizedAprovada
}

func (s ServiceOrderStatus) IsRejeitada() bool {
	return normalizeStatus(string(s)) == normalizedRejeitada
}

func (s ServiceOrderStatus) IsEmExecucao() bool {
	return normalizeStatus(string(s)) == normalizedEmExecucao
}

func (s ServiceOrderStatus) IsFinalizada() bool {
	return normalizeStatus(string(s)) == normalizedFinalizada
}

func (s ServiceOrderStatus) IsEntregue() bool {
	return normalizeStatus(string(s)) == normalizedEntregue
}

func (s ServiceOrderStatus) IsCancelada() bool {
	return normalizeStatus(string(s)) == normalizedCancelada
}

func (s ServiceOrderStatus) String() string {
	return string(s)
}

func normalizeStatus(status string) string {
	status = strings.TrimSpace(status)
	status = strings.ToUpper(status)
	status = strings.Join(strings.Fields(status), " ")
	return status
}
