package proto

import "time"

type SagaInfo struct {
	ExecType             int           `json:"exec_type"`
	GlobalCompensateType int           `json:"global_compensate_type"` // 定义Saga中所有事务的补偿类型，优先级高于事务内部的补偿类型，默认为0
	Transaction          []Transaction `json:"transaction"`
}

type (
	SubmitSagaRequest = SagaInfo

	SubmitSagaResponse struct {
		TraceId string
	}
)

type (
	RecoverSagaRequest struct {
		SagaId []int
	}
	RecoverSagaResponse struct {
	}
)

type TransactionBaseOption struct {
	TimeoutToFail  time.Duration `json:"timeout_to_fail"` // second, default 3 second
	RetryInterval  time.Duration `json:"retry_interval"`  // second, default 1 second
	CompensateType int           `json:"compensate_type"` // 1-backward recovery 2-forward recovery
	MaxRetryTimes  int           `json:"max_retry_times"` // max retry times, default 3 times, auto rollback after reaching the number of MaxRetryTimes
}

type Transaction struct {
	CallType   int                    `json:"call_type"`   // 1-rpc 2-api
	BasePath   string                 `json:"base_path"`   // use to rpc
	ServerName string                 `json:"server_name"` // use to prc
	Action     string                 `json:"action"`      // node transaction call addr
	Compensate string                 `json:"compensate"`  // node compensate call addr
	Headers    map[string]string      `json:"headers"`     // use to api
	Data       map[string]interface{} `json:"data"`
	TransactionBaseOption
}

type (
	TransactionRequest struct {
		TraceId string                 `json:"trace_id"`
		SpanId  string                 `json:"span_id"`
		Params  map[string]interface{} `json:"params"`
		// If it is a serial transaction, the parameter is the parameter when the transaction is committed.
		// Otherwise, the return value of each node will be appended in turn according to the node order
		PreviousResult []map[string]interface{} `json:"previous_result"`
	}
	TransactionResponse struct {
		Status int                    `json:"status"` // 0-failed 1-success
		Msg    string                 `json:"msg"`
		Data   map[string]interface{} `json:"data"`
	}
)

type (
	CompensateRequest struct {
		TraceId string `json:"trace_id"`
		SpanId  string `json:"span_id"`
		// compensate request params eq action request params
		Params map[string]interface{} `json:"params"`
	}
	CompensateResponse struct {
		Status int                    `json:"status"` // 0-failed 1-success
		Msg    string                 `json:"msg"`
		Data   map[string]interface{} `json:"data"`
	}
)

type (
	GetRunningSagaIdRequest struct {
	}
	GetRunningSagaIdResponse struct {
		SagaId []int
	}
)

