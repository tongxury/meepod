package verification_comp

import "context"

type IAuthVerifyComp interface {
	SendProof(ctx context.Context, dest string) (string, error)
	VerifyProof(ctx context.Context, dest, sourceProof string) (string, error)
}

type IProofStore interface {
	ShouldSend(ctx context.Context, dest string) (bool, error)
	Save(ctx context.Context, dest, proof string) error
	Find(ctx context.Context, dest string) (string, error)
}

type IPoster interface {
	GetSubject(ctx context.Context) string
	Post(ctx context.Context, dest, proof string) error
}
