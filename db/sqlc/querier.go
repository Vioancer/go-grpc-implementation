// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package generated_db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	// sqlc.arg(amount) is a placeholder for the actual value, sqlc when generating the go
	//code will name the function AddAccountBalance and the parameters required to call it
	// will be named AddAccountBalanceParams and the struct will have a field named Amount
	// and ID instead of Balance and ID
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAllAccounts(ctx context.Context, arg ListAllAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	ListUserAccounts(ctx context.Context, arg ListUserAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	// Note you cannot mix the positional parameters($1, $2, $3) with
	// named parameters(@set_full_name, @full_name, @set_email
	// Also sqlc.arg(set_hashed_password) is similar to @set_hashed_password
	// Using COALESCE function of postgresql to update only the fields that
	// are provided by the user.
	// COALESCE function returns the first non-null value from the list of arguments
	// provided to it.
	// sqlc.narg is the same as sqlc. arg , but always marks the parameter as nullable.
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
