package hedera

/*-
 *
 * Hedera Go SDK
 *
 * Copyright (C) 2020 - 2023 Hedera Hashgraph, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import (
	"time"

	"github.com/CaioLuColaco/hedera-protobufs-local/services"
	protobuf "google.golang.org/protobuf/proto"
)

type StakingInfoHedera struct {
	DeclineStakingReward bool
	StakePeriodStart     *time.Time
	PendingReward        int64
	PendingHbarReward    Hbar
	StakedToMe           Hbar
	StakedAccountID      *AccountID
	StakedNodeID         *int64
}

func _StakingInfoHederaFromProtobuf(pb *services.StakingInfoHedera) StakingInfoHedera {
	var start time.Time
	if pb.StakePeriodStart != nil {
		start = _TimeFromProtobuf(pb.StakePeriodStart)
	}

	body := StakingInfoHedera{
		DeclineStakingReward: pb.DeclineReward,
		StakePeriodStart:     &start,
		PendingReward:        pb.PendingReward,
		PendingHbarReward:    HbarFromTinybar(pb.PendingReward),
		StakedToMe:           HbarFromTinybar(pb.StakedToMe),
	}

	switch temp := pb.StakedId.(type) {
	case *services.StakingInfoHedera_StakedAccountId:
		body.StakedAccountID = _AccountIDFromProtobuf(temp.StakedAccountId)
	case *services.StakingInfoHedera_StakedNodeId:
		body.StakedNodeID = &temp.StakedNodeId
	}

	return body
}

func (StakingInfoHedera *StakingInfoHedera) _ToProtobuf() *services.StakingInfoHedera { // nolint
	var pendingReward int64

	if StakingInfoHedera.PendingReward > 0 {
		pendingReward = StakingInfoHedera.PendingReward
	} else {
		pendingReward = StakingInfoHedera.PendingHbarReward.AsTinybar()
	}

	body := services.StakingInfoHedera{
		DeclineReward: StakingInfoHedera.DeclineStakingReward,
		PendingReward: pendingReward,
		StakedToMe:    StakingInfoHedera.StakedToMe.AsTinybar(),
	}

	if StakingInfoHedera.StakePeriodStart != nil {
		body.StakePeriodStart = _TimeToProtobuf(*StakingInfoHedera.StakePeriodStart)
	}

	if StakingInfoHedera.StakedAccountID != nil {
		body.StakedId = &services.StakingInfoHedera_StakedAccountId{StakedAccountId: StakingInfoHedera.StakedAccountID._ToProtobuf()}
	} else if StakingInfoHedera.StakedNodeID != nil {
		body.StakedId = &services.StakingInfoHedera_StakedNodeId{StakedNodeId: *StakingInfoHedera.StakedNodeID}
	}

	return &body
}

// ToBytes returns the byte representation of the StakingInfoHedera
func (StakingInfoHedera *StakingInfoHedera) ToBytes() []byte {
	data, err := protobuf.Marshal(StakingInfoHedera._ToProtobuf())
	if err != nil {
		return make([]byte, 0)
	}

	return data
}

// StakingInfoHederaFromBytes returns a StakingInfoHedera object from a raw byte array
func StakingInfoHederaFromBytes(data []byte) (StakingInfoHedera, error) {
	if data == nil {
		return StakingInfoHedera{}, errByteArrayNull
	}
	pb := services.StakingInfoHedera{}
	err := protobuf.Unmarshal(data, &pb)
	if err != nil {
		return StakingInfoHedera{}, err
	}

	info := _StakingInfoHederaFromProtobuf(&pb)

	return info, nil
}
