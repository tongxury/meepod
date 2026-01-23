package verification_comp

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"strconv"
)

func Assemble(proofStore IProofStore, poster IPoster) IAuthVerifyComp {
	return &instance{
		proofStore: proofStore,
		poster:     poster,
	}
}

type instance struct {
	proofStore IProofStore
	poster     IPoster
}

func (i *instance) SendProof(ctx context.Context, dest string) (string, error) {

	should, err := i.proofStore.ShouldSend(ctx, dest)
	if err != nil {
		return "", err
	}

	if !should {
		return "", fmt.Errorf("too frequently")
	}

	proof := strconv.Itoa(helper.GetRandom(100000, 999999))

	if err := i.poster.Post(ctx, dest, proof); err != nil {
		return "", err
	}

	if err := i.proofStore.Save(ctx, dest, proof); err != nil {
		return "", err
	}

	return proof, nil
}

func (i *instance) VerifyProof(ctx context.Context, dest, sourceProof string) (string, error) {

	proof, err := i.proofStore.Find(ctx, dest)
	if err != nil {
		return "", err
	}

	if proof == "" {
		return ERR_NOT_FOUND, xerror.Wrapf("invalid proof: not found")
	}

	if proof != sourceProof {
		return ERR_NOT_MATCHING, xerror.Wrapf("invalid code: %s", sourceProof)
	}

	return "", nil
}
