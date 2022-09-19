package client_go

import (
	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/proto"
)

const (
	HTTP Protocol = "http"
	GRPC Protocol = "grpc"
)

type (
	Protocol string
)

type (
	TranType consts.TransactionType
	Action   consts.BranchAction
	Level    consts.Level
)

const (
	TransactionUnknown TranType = "unknown"
	TCC                TranType = "tcc"
	SAGA               TranType = "saga"

	// Try Try、Confirm and Cancel branch type for TCC

	ActionUnknown        = "unknown"
	Try           Action = "try"
	Confirm       Action = "confirm"
	Cancel        Action = "cancel"

	// Normal and Compensation branch type for SAGA
	Normal       Action = "normal"
	Compensation Action = "compensation"
)

func ConvertBranchActionToGrpc(action Action) proto.Action {
	switch action {
	case Try:
		return proto.Action_TRY
	case Confirm:
		return proto.Action_CONFIRM
	case Cancel:
		return proto.Action_CANCEL
	case Normal:
		return proto.Action_NORMAL
	case Compensation:
		return proto.Action_COMPENSATION
	default:
	}
	return proto.Action_UN_KNOW_TRANSACTION_TYPE
}

func ConvertTranTypeToGrpc(tranType TranType) proto.TranType {
	switch tranType {
	case TCC:
		return proto.TranType_TCC
	case SAGA:
		return proto.TranType_SAGE
	default:
	}
	return proto.TranType_UN_KNOW
}
