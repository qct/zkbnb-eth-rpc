package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

/*
	LoadZkBNBInstance: load zkbnb instance if it is already deployed
*/
func LoadZkBNBInstance(cli *rpc.ProviderClient, addr string) (instance *ZkBNB, err error) {
	instance, err = NewZkBNB(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
	CommitBlocks: commit blocks
*/
func CommitBlocks(
	cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	CommitBlocks: commit blocks
*/
func CommitBlocksWithNonce(
	cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithNonce(authCli, gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocks(
	cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocksWithNonce(authCli *rpc.AuthClient, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithNonce(authCli, gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func RevertBlocks(
	cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	revertBlocks []StorageStoredBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.RevertBlocks(transactOpts, revertBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func RevertBlocksWithNonce(authCli *rpc.AuthClient, instance *ZkBNB,
	revertBlocks []StorageStoredBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOptsWithNonce(authCli, gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	tx, err := instance.RevertBlocks(transactOpts, revertBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func PerformDesert(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	exitData ExodusVerifierExitData, assetMerkleProof [15]*big.Int, accountMerkleProof [31]*big.Int, nftMerkleProof [39]*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.PerformDesert(transactOpts, exitData, assetMerkleProof, accountMerkleProof, nftMerkleProof)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func PerformDesertNft(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB,
	exitNftData ExodusVerifierExitNftData, accountRoot *big.Int, nftMerkleProof [39]*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.PerformDesertNft(transactOpts, exitNftData, accountRoot, nftMerkleProof)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func WithdrawPendingBalance(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB, owner common.Address, token common.Address, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.WithdrawPendingBalance(transactOpts, owner, token, amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func CancelOutstandingDepositsForExodusMode(cli *rpc.ProviderClient, authCli *rpc.AuthClient, instance *ZkBNB, priorityRequestId uint64, depositsPubData [][]byte, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CancelOutstandingDepositsForExodusMode(transactOpts, priorityRequestId, depositsPubData)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
