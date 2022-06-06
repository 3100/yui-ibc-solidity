// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcidentifier

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcidentifierABI is the input ABI used to generate the binding from.
const IbcidentifierABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"consensusCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"consensusStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"packetAcknowledgementCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"portCapabilityPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCapabilityPath\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Ibcidentifier is an auto generated Go binding around an Ethereum contract.
type Ibcidentifier struct {
	IbcidentifierCaller     // Read-only binding to the contract
	IbcidentifierTransactor // Write-only binding to the contract
	IbcidentifierFilterer   // Log filterer for contract events
}

// IbcidentifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcidentifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcidentifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcidentifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcidentifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcidentifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcidentifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcidentifierSession struct {
	Contract     *Ibcidentifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcidentifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcidentifierCallerSession struct {
	Contract *IbcidentifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IbcidentifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcidentifierTransactorSession struct {
	Contract     *IbcidentifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IbcidentifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcidentifierRaw struct {
	Contract *Ibcidentifier // Generic contract binding to access the raw methods on
}

// IbcidentifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcidentifierCallerRaw struct {
	Contract *IbcidentifierCaller // Generic read-only contract binding to access the raw methods on
}

// IbcidentifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcidentifierTransactorRaw struct {
	Contract *IbcidentifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcidentifier creates a new instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifier(address common.Address, backend bind.ContractBackend) (*Ibcidentifier, error) {
	contract, err := bindIbcidentifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcidentifier{IbcidentifierCaller: IbcidentifierCaller{contract: contract}, IbcidentifierTransactor: IbcidentifierTransactor{contract: contract}, IbcidentifierFilterer: IbcidentifierFilterer{contract: contract}}, nil
}

// NewIbcidentifierCaller creates a new read-only instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifierCaller(address common.Address, caller bind.ContractCaller) (*IbcidentifierCaller, error) {
	contract, err := bindIbcidentifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcidentifierCaller{contract: contract}, nil
}

// NewIbcidentifierTransactor creates a new write-only instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcidentifierTransactor, error) {
	contract, err := bindIbcidentifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcidentifierTransactor{contract: contract}, nil
}

// NewIbcidentifierFilterer creates a new log filterer instance of Ibcidentifier, bound to a specific deployed contract.
func NewIbcidentifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcidentifierFilterer, error) {
	contract, err := bindIbcidentifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcidentifierFilterer{contract: contract}, nil
}

// bindIbcidentifier binds a generic wrapper to an already deployed contract.
func bindIbcidentifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcidentifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcidentifier *IbcidentifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ibcidentifier.Contract.IbcidentifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcidentifier *IbcidentifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.IbcidentifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcidentifier *IbcidentifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.IbcidentifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcidentifier *IbcidentifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ibcidentifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcidentifier *IbcidentifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcidentifier *IbcidentifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.contract.Transact(opts, method, params...)
}

// ChannelCapabilityPath is a paid mutator transaction binding the contract method 0x3bc3339f.
//
// Solidity: function channelCapabilityPath(string portId, string channelId) returns(bytes)
func (_Ibcidentifier *IbcidentifierTransactor) ChannelCapabilityPath(opts *bind.TransactOpts, portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "channelCapabilityPath", portId, channelId)
}

// ChannelCapabilityPath is a paid mutator transaction binding the contract method 0x3bc3339f.
//
// Solidity: function channelCapabilityPath(string portId, string channelId) returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) ChannelCapabilityPath(portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ChannelCapabilityPath(&_Ibcidentifier.TransactOpts, portId, channelId)
}

// ChannelCapabilityPath is a paid mutator transaction binding the contract method 0x3bc3339f.
//
// Solidity: function channelCapabilityPath(string portId, string channelId) returns(bytes)
func (_Ibcidentifier *IbcidentifierTransactorSession) ChannelCapabilityPath(portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ChannelCapabilityPath(&_Ibcidentifier.TransactOpts, portId, channelId)
}

// ChannelCommitmentKey is a paid mutator transaction binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ChannelCommitmentKey(opts *bind.TransactOpts, portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "channelCommitmentKey", portId, channelId)
}

// ChannelCommitmentKey is a paid mutator transaction binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ChannelCommitmentKey(portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentKey(&_Ibcidentifier.TransactOpts, portId, channelId)
}

// ChannelCommitmentKey is a paid mutator transaction binding the contract method 0xc6c9159c.
//
// Solidity: function channelCommitmentKey(string portId, string channelId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ChannelCommitmentKey(portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentKey(&_Ibcidentifier.TransactOpts, portId, channelId)
}

// ChannelCommitmentSlot is a paid mutator transaction binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ChannelCommitmentSlot(opts *bind.TransactOpts, portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "channelCommitmentSlot", portId, channelId)
}

// ChannelCommitmentSlot is a paid mutator transaction binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ChannelCommitmentSlot(portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentSlot(&_Ibcidentifier.TransactOpts, portId, channelId)
}

// ChannelCommitmentSlot is a paid mutator transaction binding the contract method 0x3560a02d.
//
// Solidity: function channelCommitmentSlot(string portId, string channelId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ChannelCommitmentSlot(portId string, channelId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ChannelCommitmentSlot(&_Ibcidentifier.TransactOpts, portId, channelId)
}

// ClientCommitmentKey is a paid mutator transaction binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ClientCommitmentKey(opts *bind.TransactOpts, clientId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "clientCommitmentKey", clientId)
}

// ClientCommitmentKey is a paid mutator transaction binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ClientCommitmentKey(clientId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ClientCommitmentKey(&_Ibcidentifier.TransactOpts, clientId)
}

// ClientCommitmentKey is a paid mutator transaction binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ClientCommitmentKey(clientId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ClientCommitmentKey(&_Ibcidentifier.TransactOpts, clientId)
}

// ClientStateCommitmentSlot is a paid mutator transaction binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ClientStateCommitmentSlot(opts *bind.TransactOpts, clientId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "clientStateCommitmentSlot", clientId)
}

// ClientStateCommitmentSlot is a paid mutator transaction binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ClientStateCommitmentSlot(clientId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ClientStateCommitmentSlot(&_Ibcidentifier.TransactOpts, clientId)
}

// ClientStateCommitmentSlot is a paid mutator transaction binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ClientStateCommitmentSlot(clientId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ClientStateCommitmentSlot(&_Ibcidentifier.TransactOpts, clientId)
}

// ConnectionCommitmentKey is a paid mutator transaction binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ConnectionCommitmentKey(opts *bind.TransactOpts, connectionId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "connectionCommitmentKey", connectionId)
}

// ConnectionCommitmentKey is a paid mutator transaction binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConnectionCommitmentKey(connectionId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentKey(&_Ibcidentifier.TransactOpts, connectionId)
}

// ConnectionCommitmentKey is a paid mutator transaction binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ConnectionCommitmentKey(connectionId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentKey(&_Ibcidentifier.TransactOpts, connectionId)
}

// ConnectionCommitmentSlot is a paid mutator transaction binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ConnectionCommitmentSlot(opts *bind.TransactOpts, connectionId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "connectionCommitmentSlot", connectionId)
}

// ConnectionCommitmentSlot is a paid mutator transaction binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConnectionCommitmentSlot(connectionId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentSlot(&_Ibcidentifier.TransactOpts, connectionId)
}

// ConnectionCommitmentSlot is a paid mutator transaction binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ConnectionCommitmentSlot(connectionId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConnectionCommitmentSlot(&_Ibcidentifier.TransactOpts, connectionId)
}

// ConsensusCommitmentKey is a paid mutator transaction binding the contract method 0xbff19ae3.
//
// Solidity: function consensusCommitmentKey(string clientId, HeightData height) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ConsensusCommitmentKey(opts *bind.TransactOpts, clientId string, height HeightData) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "consensusCommitmentKey", clientId, height)
}

// ConsensusCommitmentKey is a paid mutator transaction binding the contract method 0xbff19ae3.
//
// Solidity: function consensusCommitmentKey(string clientId, HeightData height) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConsensusCommitmentKey(clientId string, height HeightData) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConsensusCommitmentKey(&_Ibcidentifier.TransactOpts, clientId, height)
}

// ConsensusCommitmentKey is a paid mutator transaction binding the contract method 0xbff19ae3.
//
// Solidity: function consensusCommitmentKey(string clientId, HeightData height) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ConsensusCommitmentKey(clientId string, height HeightData) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConsensusCommitmentKey(&_Ibcidentifier.TransactOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a paid mutator transaction binding the contract method 0x956f5239.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, HeightData height) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) ConsensusStateCommitmentSlot(opts *bind.TransactOpts, clientId string, height HeightData) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "consensusStateCommitmentSlot", clientId, height)
}

// ConsensusStateCommitmentSlot is a paid mutator transaction binding the contract method 0x956f5239.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, HeightData height) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) ConsensusStateCommitmentSlot(clientId string, height HeightData) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConsensusStateCommitmentSlot(&_Ibcidentifier.TransactOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a paid mutator transaction binding the contract method 0x956f5239.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, HeightData height) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) ConsensusStateCommitmentSlot(clientId string, height HeightData) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.ConsensusStateCommitmentSlot(&_Ibcidentifier.TransactOpts, clientId, height)
}

// PacketAcknowledgementCommitmentKey is a paid mutator transaction binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) PacketAcknowledgementCommitmentKey(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "packetAcknowledgementCommitmentKey", portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentKey is a paid mutator transaction binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentKey(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentKey is a paid mutator transaction binding the contract method 0xe334f11b.
//
// Solidity: function packetAcknowledgementCommitmentKey(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) PacketAcknowledgementCommitmentKey(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentKey(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentSlot is a paid mutator transaction binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) PacketAcknowledgementCommitmentSlot(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "packetAcknowledgementCommitmentSlot", portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentSlot is a paid mutator transaction binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketAcknowledgementCommitmentSlot(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentSlot(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PacketAcknowledgementCommitmentSlot is a paid mutator transaction binding the contract method 0xc50839ec.
//
// Solidity: function packetAcknowledgementCommitmentSlot(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) PacketAcknowledgementCommitmentSlot(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketAcknowledgementCommitmentSlot(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a paid mutator transaction binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) PacketCommitmentKey(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "packetCommitmentKey", portId, channelId, sequence)
}

// PacketCommitmentKey is a paid mutator transaction binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketCommitmentKey(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PacketCommitmentKey is a paid mutator transaction binding the contract method 0xdae1b0f8.
//
// Solidity: function packetCommitmentKey(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) PacketCommitmentKey(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketCommitmentKey(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PacketCommitmentSlot is a paid mutator transaction binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactor) PacketCommitmentSlot(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "packetCommitmentSlot", portId, channelId, sequence)
}

// PacketCommitmentSlot is a paid mutator transaction binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierSession) PacketCommitmentSlot(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketCommitmentSlot(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PacketCommitmentSlot is a paid mutator transaction binding the contract method 0xb0db45ba.
//
// Solidity: function packetCommitmentSlot(string portId, string channelId, uint64 sequence) returns(bytes32)
func (_Ibcidentifier *IbcidentifierTransactorSession) PacketCommitmentSlot(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PacketCommitmentSlot(&_Ibcidentifier.TransactOpts, portId, channelId, sequence)
}

// PortCapabilityPath is a paid mutator transaction binding the contract method 0x2570dae0.
//
// Solidity: function portCapabilityPath(string portId) returns(bytes)
func (_Ibcidentifier *IbcidentifierTransactor) PortCapabilityPath(opts *bind.TransactOpts, portId string) (*types.Transaction, error) {
	return _Ibcidentifier.contract.Transact(opts, "portCapabilityPath", portId)
}

// PortCapabilityPath is a paid mutator transaction binding the contract method 0x2570dae0.
//
// Solidity: function portCapabilityPath(string portId) returns(bytes)
func (_Ibcidentifier *IbcidentifierSession) PortCapabilityPath(portId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PortCapabilityPath(&_Ibcidentifier.TransactOpts, portId)
}

// PortCapabilityPath is a paid mutator transaction binding the contract method 0x2570dae0.
//
// Solidity: function portCapabilityPath(string portId) returns(bytes)
func (_Ibcidentifier *IbcidentifierTransactorSession) PortCapabilityPath(portId string) (*types.Transaction, error) {
	return _Ibcidentifier.Contract.PortCapabilityPath(&_Ibcidentifier.TransactOpts, portId)
}
