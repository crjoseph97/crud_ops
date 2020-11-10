package backend

//Error messagess
const (
	ErrSQLConnection        string = "backend: unexpected error while closing the sql connection"
	ErrQueryExecutionFailed string = "backend: unexpected error during query execution"
	ErrScanFailed           string = "backend: unexpected error while scanning through the obtained results"
	ErrIterationFailed      string = "backend: unexpected error while iterating the results set"
)
