package repositories

import "context"

type TransactionId int

type TransactionHandler func(iTransactionId TransactionId) error

type TransactionLevel int
const (
	ReadUncommitted TransactionLevel = 0
	ReadCommitted = 1
	RepeatableRead = 2
	Serializable = 3
)

type RepositoryFactoryI interface {
	/// start a transaction and automatically commit or abort based
	/// on whether the handler returns error and closes the transaction
	/// after the handler exits. The handler is given a 
	/// transaction ID which can be used to create different 
	/// repositories. Repositories instances don't need to be
	/// the same for the same transaction id.
	/// return the error returns by transaction handler and the result of the 
	/// begin / commit / abort
	StartTransactionAuto(
		iContext context.Context, 
		iIsolationLevel TransactionLevel,
		iHandler TransactionHandler,
	) (error, error)
	/// return common.NotFound if iId has not been created
	GetTaskRepository(iId TransactionId) (TaskRepositoryI, error)
	/// return common.NotFound if iId has not been created
	GetUserRepository(iId TransactionId) (UserRepositoryI, error)
}