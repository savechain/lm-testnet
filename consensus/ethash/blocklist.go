// Copyright (c) 2019 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package ethash

import (

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)
const (
	// AddressLength length of address in bytes.
	AddressLength = common.AddressLength
)
// Address address of account.
//type Address common.Address


var blocklist = make(map[common.Address]bool)

func init() {
	var list = params.Mainnetblocks
	for _, str := range list {
		blocklist[common.HexToAddress(str)] = true
	}
}
// IsOriginBlocked returns true if the address is in the blocklist
func IsOriginBlocked(origin common.Address) bool {
	return blocklist[origin]
}

