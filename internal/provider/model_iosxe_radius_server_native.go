// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RadiusServerNative struct {
	Device                types.String `tfsdk:"device"`
	Id                    types.String `tfsdk:"id"`
	DeleteMode            types.String `tfsdk:"delete_mode"`
	Name                  types.String `tfsdk:"name"`
	RadiusHostAddressIpv4 types.String `tfsdk:"radius_host_address_ipv4"`
	AddressAuthPort       types.Int64  `tfsdk:"address_auth_port"`
	Timeout               types.Int64  `tfsdk:"timeout"`
	Retransmit            types.Int64  `tfsdk:"retransmit"`
	KeyKey                types.String `tfsdk:"key_key"`
}

type RadiusServerNativeData struct {
	Device                types.String `tfsdk:"device"`
	Id                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	RadiusHostAddressIpv4 types.String `tfsdk:"radius_host_address_ipv4"`
	AddressAuthPort       types.Int64  `tfsdk:"address_auth_port"`
	Timeout               types.Int64  `tfsdk:"timeout"`
	Retransmit            types.Int64  `tfsdk:"retransmit"`
	KeyKey                types.String `tfsdk:"key_key"`
}

func (data RadiusServerNative) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/radius/Cisco-IOS-XE-aaa:server=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data RadiusServerNativeData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/radius/Cisco-IOS-XE-aaa:server=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data RadiusServerNative) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data RadiusServerNative) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Name.ValueString())
	}
	if !data.RadiusHostAddressIpv4.IsNull() && !data.RadiusHostAddressIpv4.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address.ipv4", data.RadiusHostAddressIpv4.ValueString())
	}
	if !data.AddressAuthPort.IsNull() && !data.AddressAuthPort.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address.auth-port", strconv.FormatInt(data.AddressAuthPort.ValueInt64(), 10))
	}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timeout", strconv.FormatInt(data.Timeout.ValueInt64(), 10))
	}
	if !data.Retransmit.IsNull() && !data.Retransmit.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"retransmit", strconv.FormatInt(data.Retransmit.ValueInt64(), 10))
	}
	if !data.KeyKey.IsNull() && !data.KeyKey.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"key.key", data.KeyKey.ValueString())
	}
	return body
}

func (data *RadiusServerNative) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "address.ipv4"); value.Exists() && !data.RadiusHostAddressIpv4.IsNull() {
		data.RadiusHostAddressIpv4 = types.StringValue(value.String())
	} else {
		data.RadiusHostAddressIpv4 = types.StringNull()
	}
	if value := res.Get(prefix + "address.auth-port"); value.Exists() && !data.AddressAuthPort.IsNull() {
		data.AddressAuthPort = types.Int64Value(value.Int())
	} else {
		data.AddressAuthPort = types.Int64Null()
	}
	if value := res.Get(prefix + "timeout"); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(prefix + "retransmit"); value.Exists() && !data.Retransmit.IsNull() {
		data.Retransmit = types.Int64Value(value.Int())
	} else {
		data.Retransmit = types.Int64Null()
	}
	if value := res.Get(prefix + "key.key"); value.Exists() && !data.KeyKey.IsNull() {
		data.KeyKey = types.StringValue(value.String())
	} else {
		data.KeyKey = types.StringNull()
	}
}

func (data *RadiusServerNativeData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "address.ipv4"); value.Exists() {
		data.RadiusHostAddressIpv4 = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "address.auth-port"); value.Exists() {
		data.AddressAuthPort = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "timeout"); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "retransmit"); value.Exists() {
		data.Retransmit = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "key.key"); value.Exists() {
		data.KeyKey = types.StringValue(value.String())
	}
}

func (data *RadiusServerNative) getDeletedListItems(ctx context.Context, state RadiusServerNative) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *RadiusServerNative) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *RadiusServerNative) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.RadiusHostAddressIpv4.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/ipv4", data.getPath()))
	}
	if !data.AddressAuthPort.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/auth-port", data.getPath()))
	}
	if !data.Timeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timeout", data.getPath()))
	}
	if !data.Retransmit.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/retransmit", data.getPath()))
	}
	if !data.KeyKey.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/key/key", data.getPath()))
	}
	return deletePaths
}
