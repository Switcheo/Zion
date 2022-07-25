/*
 * Copyright (C) 2021 The Zion Authors
 * This file is part of The Zion library.
 *
 * The Zion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Zion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The Zion.  If not, see <http://www.gnu.org/licenses/>.
 */
package boot

import (
	"github.com/Switcheo/Zion/contracts/native/cross_chain_manager"
	"github.com/Switcheo/Zion/contracts/native/economic"
	"github.com/Switcheo/Zion/contracts/native/governance/neo3_state_manager"
	"github.com/Switcheo/Zion/contracts/native/governance/node_manager"
	"github.com/Switcheo/Zion/contracts/native/governance/relayer_manager"
	"github.com/Switcheo/Zion/contracts/native/governance/side_chain_manager"
	"github.com/Switcheo/Zion/contracts/native/governance/signature_manager"
	"github.com/Switcheo/Zion/contracts/native/header_sync"
	"github.com/Switcheo/Zion/contracts/native/utils"
	"github.com/Switcheo/Zion/log"
)

func InitNativeContracts() {
	header_sync.InitHeaderSync()
	cross_chain_manager.InitCrossChainManager()
	neo3_state_manager.InitNeo3StateManager()
	node_manager.InitNodeManager()
	relayer_manager.InitRelayerManager()
	side_chain_manager.InitSideChainManager()

	signature_manager.InitSignatureManager()

	economic.InitEconomic()

	log.Info("Initialize main chain native contracts",
		"header sync", utils.HeaderSyncContractAddress.Hex(),
		"cross chain manager", utils.CrossChainManagerContractAddress.Hex(),
		"neo3 state manager", utils.Neo3StateManagerContractAddress.Hex(),
		"node manager", utils.NodeManagerContractAddress.Hex(),
		"relayer manager", utils.RelayerManagerContractAddress.Hex(),
		"side chain manager", utils.SideChainManagerContractAddress.Hex(),
		"signature manager", utils.SignatureManagerContractAddress.Hex(),
		"economic", utils.EconomicContractAddress.Hex())
}
