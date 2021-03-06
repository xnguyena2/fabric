/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	channelconfig "github.com/hyperledger/fabric/common/config/channel"
	genesisconfig "github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	cb "github.com/hyperledger/fabric/protos/common"
)

func newChainRequest(consensusType, creationPolicy, newChannelId string) *cb.Envelope {

	//JCS: added application organizations to avoid the config transaction from being rejected
	conf := genesisconfig.Load(genesisconfig.SampleSingleMSPChannelProfile)
	orgs := conf.Application.Organizations

	var orgNames []string
	for _, org := range orgs {
		orgNames = append(orgNames, org.Name)
	}

	env, err := channelconfig.MakeChainCreationTransaction(newChannelId, genesisconfig.SampleConsortiumName, signer, orgNames...) // JCS: added `orgsNames` at the end
	if err != nil {
		panic(err)
	}
	return env
}
