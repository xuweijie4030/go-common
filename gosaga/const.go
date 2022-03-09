package gosaga

const (
	SerialExecType     = 0
	ConcurrentExecType = 1
)

const (
	RpcTransaction = 1
	ApiTransaction = 2
)

const (
	BackwardCompensateType = 1
	ForwardCompensateType  = 2
)

const (
	TransactionSuccess = 1
	TransactionFailed  = 0
)

const (
	CompensateSuccess = 1
	CompensateFailed  = 0
)

const (
	DefaultTimeoutToFail             = 3
	DefaultRetryInterval             = 1
	DefaultMaxRetryTimes             = 3
	DefaultTransactionCompensateType = BackwardCompensateType
	DefaultGlobalCompensateType      = 0
)

const (
	SagaCreatedEvent = iota
	TransactionBeginEvent
	TransactionEndEvent
	TransactionAbortedEvent
	TransactionCompensateEvent
	SagaEndEvent
)
