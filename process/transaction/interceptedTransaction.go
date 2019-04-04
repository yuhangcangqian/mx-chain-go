package transaction

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go-sandbox/crypto"
	"github.com/ElrondNetwork/elrond-go-sandbox/data/state"
	"github.com/ElrondNetwork/elrond-go-sandbox/data/transaction"
	"github.com/ElrondNetwork/elrond-go-sandbox/process"
	"github.com/ElrondNetwork/elrond-go-sandbox/sharding"
)

// InterceptedTransaction holds and manages a transaction based struct with extended functionality
type InterceptedTransaction struct {
	*transaction.Transaction

	txBuffWithoutSig         []byte
	hash                     []byte
	rcvShard                 uint32
	sndShard                 uint32
	isAddressedToOtherShards bool
	addrConv                 state.AddressConverter
	singleSigner             crypto.SingleSigner
	keyGen                   crypto.KeyGenerator
}

// NewInterceptedTransaction returns a new instance of InterceptedTransaction
func NewInterceptedTransaction(signer crypto.SingleSigner) *InterceptedTransaction {
	return &InterceptedTransaction{
		Transaction:  &transaction.Transaction{},
		singleSigner: signer,
	}
}

// IntegrityAndValidity returns a non nil error if transaction failed some checking tests
func (inTx *InterceptedTransaction) IntegrityAndValidity(coordinator sharding.Coordinator) error {
	if coordinator == nil {
		return process.ErrNilShardCoordinator
	}

	err := inTx.Integrity(coordinator)
	if err != nil {
		return err
	}

	if inTx.addrConv == nil {
		return process.ErrNilAddressConverter
	}

	sndAddr, err := inTx.addrConv.CreateAddressFromPublicKeyBytes(inTx.SndAddr)
	if err != nil {
		return process.ErrInvalidSndAddr
	}

	rcvAddr, err := inTx.addrConv.CreateAddressFromPublicKeyBytes(inTx.RcvAddr)
	if err != nil {
		return process.ErrInvalidRcvAddr
	}

	inTx.rcvShard = coordinator.ComputeId(rcvAddr)
	inTx.sndShard = coordinator.ComputeId(sndAddr)

	inTx.isAddressedToOtherShards = inTx.rcvShard != coordinator.SelfId() && inTx.sndShard != coordinator.SelfId()

	return nil
}

// Integrity checks for not nil fields and negative value
func (inTx *InterceptedTransaction) Integrity(coordinator sharding.Coordinator) error {
	if inTx.Transaction == nil {
		return process.ErrNilTransaction
	}

	if inTx.Signature == nil {
		return process.ErrNilSignature
	}

	if inTx.RcvAddr == nil {
		return process.ErrNilRcvAddr
	}

	if inTx.SndAddr == nil {
		return process.ErrNilSndAddr
	}

	if inTx.Transaction.Value == nil {
		return process.ErrNilValue
	}

	if inTx.Transaction.Value.Cmp(big.NewInt(0)) < 0 {
		return process.ErrNegativeValue
	}

	return nil
}

// VerifySig checks if the tx is correctly signed
func (inTx *InterceptedTransaction) VerifySig() error {
	if inTx.Transaction == nil {
		return process.ErrNilTransaction
	}

	if inTx.keyGen == nil {
		return process.ErrNilKeyGen
	}

	if inTx.singleSigner == nil {
		return process.ErrNilSingleSigner
	}

	senderPubKey, err := inTx.keyGen.PublicKeyFromByteArray(inTx.SndAddr)
	if err != nil {
		return err
	}

	err = inTx.singleSigner.Verify(senderPubKey, inTx.txBuffWithoutSig, inTx.Signature)

	if err != nil {
		return err
	}

	return nil
}

// RcvShard returns the receiver shard
func (inTx *InterceptedTransaction) RcvShard() uint32 {
	return inTx.rcvShard
}

// SndShard returns the sender shard
func (inTx *InterceptedTransaction) SndShard() uint32 {
	return inTx.sndShard
}

// IsAddressedToOtherShards returns true if this transaction is not meant to be processed by the node from this shard
func (inTx *InterceptedTransaction) IsAddressedToOtherShards() bool {
	return inTx.isAddressedToOtherShards
}

// SetAddressConverter sets the AddressConverter implementation used in address processing
func (inTx *InterceptedTransaction) SetAddressConverter(converter state.AddressConverter) {
	inTx.addrConv = converter
}

// AddressConverter returns the AddressConverter implementation used in address processing
func (inTx *InterceptedTransaction) AddressConverter() state.AddressConverter {
	return inTx.addrConv
}

// GetTransaction returns the transaction pointer that actually holds the data
func (inTx *InterceptedTransaction) GetTransaction() *transaction.Transaction {
	return inTx.Transaction
}

// SetHash sets the hash of this transaction. The hash will also be the ID of this object
func (inTx *InterceptedTransaction) SetHash(hash []byte) {
	inTx.hash = hash
}

// Hash gets the hash of this transaction
func (inTx *InterceptedTransaction) Hash() []byte {
	return inTx.hash
}

// SetTxBuffWithoutSig sets the byte slice buffer of this transaction having nil in Signature field.
func (inTx *InterceptedTransaction) SetTxBuffWithoutSig(txBuffWithoutSig []byte) {
	inTx.txBuffWithoutSig = txBuffWithoutSig
}

// TxBuffWithoutSig gets the byte slice buffer of this transaction having nil in Signature field
func (inTx *InterceptedTransaction) TxBuffWithoutSig() []byte {
	return inTx.txBuffWithoutSig
}

// SingleSignKeyGen returns the key generator that is used to create a new public key verifier that will be used
// for validating transaction's signature
func (inTx *InterceptedTransaction) SingleSignKeyGen() crypto.KeyGenerator {
	return inTx.keyGen
}

// SetSingleSignKeyGen sets the key generator that is used to create a new public key verifier that will be used
// for validating transaction's signature
func (inTx *InterceptedTransaction) SetSingleSignKeyGen(generator crypto.KeyGenerator) {
	inTx.keyGen = generator
}