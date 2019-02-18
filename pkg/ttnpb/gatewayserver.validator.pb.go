// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/gatewayserver.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *GatewayUp) Validate() error {
	for _, item := range this.UplinkMessages {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UplinkMessages", err)
			}
		}
	}
	if this.GatewayStatus != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GatewayStatus); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GatewayStatus", err)
		}
	}
	if this.TxAcknowledgment != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TxAcknowledgment); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TxAcknowledgment", err)
		}
	}
	return nil
}
func (this *GatewayDown) Validate() error {
	if this.DownlinkMessage != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DownlinkMessage); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DownlinkMessage", err)
		}
	}
	return nil
}
func (this *ScheduleDownlinkResponse) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.Delay)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("Delay", err)
	}
	return nil
}
