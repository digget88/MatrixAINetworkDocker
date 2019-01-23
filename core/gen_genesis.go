// Copyright (c) 2018-2019 The MATRIX Authors
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php
// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package core

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/matrix/go-matrix/base58"
	"github.com/matrix/go-matrix/common"
	"github.com/matrix/go-matrix/common/hexutil"
	"github.com/matrix/go-matrix/common/math"
	"github.com/matrix/go-matrix/params"
)

var _ = (*genesisSpecMarshaling)(nil)

type GenesisAddress common.Address

// UnmarshalJSON parses a hash in hex syntax.
func (a *GenesisAddress) UnmarshalJSON(input []byte) error {
	*a = GenesisAddress(base58.Base58DecodeToAddress(string(input[1 : len(input)-1])))
	return nil
}

//func (a *GenesisAddress) MarshalJSON() ([]byte, error) {
//	buff := base58.Base58EncodeToString("MAN", common.Address(*a))
//	return []byte(buff),nil
//}
// MarshalText returns the hex representation of a.
func (a GenesisAddress) MarshalText() ([]byte, error) {
	buff := base58.Base58EncodeToString("MAN", common.Address(a))
	return []byte(buff), nil
}

// UnmarshalText parses a hash in hex syntax.
func (a *GenesisAddress) UnmarshalText(input []byte) error {
	err := hexutil.UnmarshalFixedText("GenesisAddress", input, a[:])
	return err
}

func (g Genesis) MarshalJSON() ([]byte, error) {
	type Genesis struct {
		Config            *params.ChainConfig               `json:"config,omitempty"`
		Nonce             math.HexOrDecimal64               `json:"nonce"`
		Timestamp         math.HexOrDecimal64               `json:"timestamp"`
		ExtraData         hexutil.Bytes                     `json:"extraData"`
		Version           string                            `json:"version"`
		VersionSignatures []common.Signature                `json:"versionSignatures"`
		VrfValue          hexutil.Bytes                     `json:"vrfvalue"`
		Leader            GenesisAddress                    `json:"leader"`
		NextElect         []GenesisElect                    `json:"nextElect"        gencodec:"required"`
		NetTopology       GenesisNetTopology                `json:"nettopology"        gencodec:"required"`
		Signatures        []common.Signature                `json:"signatures" gencodec:"required"`
		GasLimit          math.HexOrDecimal64               `json:"gasLimit"   gencodec:"required"`
		Difficulty        *math.HexOrDecimal256             `json:"difficulty" gencodec:"required"`
		Mixhash           common.Hash                       `json:"mixHash"`
		Coinbase          GenesisAddress                    `json:"coinbase"`
		Alloc             map[GenesisAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		MState            *GenesisMState                    `json:"mstate,omitempty"`
		Number            math.HexOrDecimal64               `json:"number"`
		GasUsed           math.HexOrDecimal64               `json:"gasUsed"`
		ParentHash        common.Hash                       `json:"parentHash"`
		Root              common.Hash                       `json:"stateRoot,omitempty"`
		TxHash            common.Hash                       `json:"transactionsRoot,omitempty"`
	}
	var enc Genesis
	enc.Config = g.Config
	enc.Nonce = math.HexOrDecimal64(g.Nonce)
	enc.Timestamp = math.HexOrDecimal64(g.Timestamp)
	enc.ExtraData = g.ExtraData
	enc.Version = g.Version
	enc.VersionSignatures = g.VersionSignatures
	enc.VrfValue = g.VrfValue
	enc.Leader = GenesisAddress(g.Leader)
	enc.NextElect = make([]GenesisElect, len(g.NextElect))
	for i, elec := range g.NextElect {
		enc.NextElect[i] = TransToGenesisElect(elec)
	}
	enc.NetTopology = TransToGenesisNetTopology(g.NetTopology)
	enc.Signatures = g.Signatures
	enc.GasLimit = math.HexOrDecimal64(g.GasLimit)
	enc.Difficulty = (*math.HexOrDecimal256)(g.Difficulty)
	enc.Mixhash = g.Mixhash
	enc.Coinbase = GenesisAddress(g.Coinbase)
	if g.Alloc != nil {
		enc.Alloc = make(map[GenesisAddress]GenesisAccount, len(g.Alloc))
		for k, v := range g.Alloc {
			enc.Alloc[GenesisAddress(k)] = v
		}
	}
	enc.MState = g.MState
	enc.Number = math.HexOrDecimal64(g.Number)
	enc.GasUsed = math.HexOrDecimal64(g.GasUsed)
	enc.ParentHash = g.ParentHash
	enc.Root = g.Root
	enc.TxHash = g.TxHash
	return json.Marshal(&enc)
}

type GenesisElect struct {
	Account GenesisAddress
	Stock   uint16
	Type    common.ElectRoleType
	VIP     common.VIPRoleType
}

func TransToGenesisElect(elect common.Elect) GenesisElect {
	return GenesisElect{
		GenesisAddress(elect.Account),
		elect.Stock,
		elect.Type,
		elect.VIP,
	}
}
func TransToCommonElect(elect GenesisElect) common.Elect {
	return common.Elect{
		common.Address(elect.Account),
		elect.Stock,
		elect.Type,
		elect.VIP,
	}
}

type GenesisNetTopologyData struct {
	Account  GenesisAddress
	Position uint16
}
type GenesisNetTopology struct {
	Type            uint8
	NetTopologyData []GenesisNetTopologyData
}

func TransToGenesisNetTopology(topology common.NetTopology) GenesisNetTopology {
	gtopology := GenesisNetTopology{
		topology.Type,
		make([]GenesisNetTopologyData, len(topology.NetTopologyData)),
	}
	for i, item := range topology.NetTopologyData {
		gtopology.NetTopologyData[i].Account = GenesisAddress(item.Account)
		gtopology.NetTopologyData[i].Position = item.Position
	}
	return gtopology
}
func TransToCommonNetTopology(gtopology GenesisNetTopology) common.NetTopology {
	topology := common.NetTopology{
		gtopology.Type,
		make([]common.NetTopologyData, len(gtopology.NetTopologyData)),
	}
	for i, item := range gtopology.NetTopologyData {
		topology.NetTopologyData[i].Account = common.Address(item.Account)
		topology.NetTopologyData[i].Position = item.Position
	}
	return topology
}

func (g *Genesis) UnmarshalJSON(input []byte) error {
	type Genesis struct {
		Config            *params.ChainConfig               `json:"config,omitempty"`
		Nonce             *math.HexOrDecimal64              `json:"nonce"`
		Timestamp         *math.HexOrDecimal64              `json:"timestamp"`
		ExtraData         *hexutil.Bytes                    `json:"extraData"`
		Version           *string                           `json:"version"`
		VersionSignatures *[]common.Signature               `json:"versionSignatures"`
		VrfValue          *hexutil.Bytes                    `json:"vrfvalue"`
		Leader            *GenesisAddress                   `json:"leader"`
		NextElect         *[]GenesisElect                   `json:"nextElect" gencodec:"required"`
		NetTopology       *GenesisNetTopology               `json:"nettopology"        gencodec:"required"`
		Signatures        *[]common.Signature               `json:"signatures" gencodec:"required"`
		GasLimit          *math.HexOrDecimal64              `json:"gasLimit"   gencodec:"required"`
		Difficulty        *math.HexOrDecimal256             `json:"difficulty" gencodec:"required"`
		Mixhash           *common.Hash                      `json:"mixHash"`
		Coinbase          *GenesisAddress                   `json:"coinbase"`
		Alloc             map[GenesisAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		MState            *GenesisMState                    `json:"mstate"`
		Number            *math.HexOrDecimal64              `json:"number"`
		GasUsed           *math.HexOrDecimal64              `json:"gasUsed"`
		ParentHash        *common.Hash                      `json:"parentHash"`
		Root              *common.Hash                      `json:"stateRoot,omitempty"`
		TxHash            *common.Hash                      `json:"transactionsRoot,omitempty"`
	}
	var dec Genesis
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Config != nil {
		g.Config = dec.Config
	}
	if dec.Nonce != nil {
		g.Nonce = uint64(*dec.Nonce)
	}
	if dec.Timestamp != nil {
		g.Timestamp = uint64(*dec.Timestamp)
	}
	if dec.ExtraData != nil {
		g.ExtraData = *dec.ExtraData
	}
	if dec.VersionSignatures != nil {
		g.VersionSignatures = *dec.VersionSignatures
	}
	if dec.Version != nil {
		g.Version = *dec.Version
	}
	if g.VrfValue != nil {
		g.VrfValue = *dec.VrfValue
	}

	if dec.Leader != nil {
		g.Leader = common.Address(*dec.Leader)
	}
	if dec.NextElect != nil {
		g.NextElect = make([]common.Elect, len(*dec.NextElect))
		for i, item := range *dec.NextElect {
			g.NextElect[i] = TransToCommonElect(item)
		}
	}
	if dec.NetTopology != nil {
		g.NetTopology = TransToCommonNetTopology(*dec.NetTopology)
	}
	if dec.Signatures != nil {
		g.Signatures = *dec.Signatures
	}
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Genesis")
	}
	g.GasLimit = uint64(*dec.GasLimit)
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Genesis")
	}
	g.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Mixhash != nil {
		g.Mixhash = *dec.Mixhash
	}
	if dec.Coinbase != nil {
		g.Coinbase = common.Address(*dec.Coinbase)
	}
	if dec.Alloc == nil {
		return errors.New("missing required field 'alloc' for Genesis")
	}
	g.Alloc = make(GenesisAlloc, len(dec.Alloc))
	if dec.MState != nil {
		g.MState = dec.MState
	}
	for k, v := range dec.Alloc {
		g.Alloc[common.Address(k)] = v
	}
	if dec.Number != nil {
		g.Number = uint64(*dec.Number)
	}
	if dec.GasUsed != nil {
		g.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.ParentHash != nil {
		g.ParentHash = *dec.ParentHash
	}
	if dec.Root != nil {
		g.Root = *dec.Root
	}
	if dec.TxHash != nil {
		g.TxHash = *dec.TxHash
	}
	return nil
}
