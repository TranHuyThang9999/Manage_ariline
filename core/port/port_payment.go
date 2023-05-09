package port

import "context"

type RepositoryPayment interface {
	Payed(ctx context.Context)
	GetBill(ctx context.Context)
	GetAllByForm(ctx context.Context)
}
