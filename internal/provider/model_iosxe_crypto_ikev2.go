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
	"regexp"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CryptoIKEv2 struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	DeleteMode       types.String `tfsdk:"delete_mode"`
	NatKeepalive     types.Int64  `tfsdk:"nat_keepalive"`
	Dpd              types.Int64  `tfsdk:"dpd"`
	DpdRetryInterval types.Int64  `tfsdk:"dpd_retry_interval"`
	DpdQuery         types.String `tfsdk:"dpd_query"`
}

type CryptoIKEv2Data struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	NatKeepalive     types.Int64  `tfsdk:"nat_keepalive"`
	Dpd              types.Int64  `tfsdk:"dpd"`
	DpdRetryInterval types.Int64  `tfsdk:"dpd_retry_interval"`
	DpdQuery         types.String `tfsdk:"dpd_query"`
}

func (data CryptoIKEv2) getPath() string {
	return "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2"
}

func (data CryptoIKEv2Data) getPath() string {
	return "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2"
}

// if last path element has a key -> remove it
func (data CryptoIKEv2) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CryptoIKEv2) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.NatKeepalive.IsNull() && !data.NatKeepalive.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"nat.keepalive", strconv.FormatInt(data.NatKeepalive.ValueInt64(), 10))
	}
	if !data.Dpd.IsNull() && !data.Dpd.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dpd-container.dpd", strconv.FormatInt(data.Dpd.ValueInt64(), 10))
	}
	if !data.DpdRetryInterval.IsNull() && !data.DpdRetryInterval.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dpd-container.retry-interval", strconv.FormatInt(data.DpdRetryInterval.ValueInt64(), 10))
	}
	if !data.DpdQuery.IsNull() && !data.DpdQuery.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dpd-container.dpd-query", data.DpdQuery.ValueString())
	}
	return body
}

func (data *CryptoIKEv2) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "nat.keepalive"); value.Exists() && !data.NatKeepalive.IsNull() {
		data.NatKeepalive = types.Int64Value(value.Int())
	} else {
		data.NatKeepalive = types.Int64Null()
	}
	if value := res.Get(prefix + "dpd-container.dpd"); value.Exists() && !data.Dpd.IsNull() {
		data.Dpd = types.Int64Value(value.Int())
	} else {
		data.Dpd = types.Int64Null()
	}
	if value := res.Get(prefix + "dpd-container.retry-interval"); value.Exists() && !data.DpdRetryInterval.IsNull() {
		data.DpdRetryInterval = types.Int64Value(value.Int())
	} else {
		data.DpdRetryInterval = types.Int64Null()
	}
	if value := res.Get(prefix + "dpd-container.dpd-query"); value.Exists() && !data.DpdQuery.IsNull() {
		data.DpdQuery = types.StringValue(value.String())
	} else {
		data.DpdQuery = types.StringNull()
	}
}

func (data *CryptoIKEv2Data) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "nat.keepalive"); value.Exists() {
		data.NatKeepalive = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dpd-container.dpd"); value.Exists() {
		data.Dpd = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dpd-container.retry-interval"); value.Exists() {
		data.DpdRetryInterval = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "dpd-container.dpd-query"); value.Exists() {
		data.DpdQuery = types.StringValue(value.String())
	}
}

func (data *CryptoIKEv2) getDeletedListItems(ctx context.Context, state CryptoIKEv2) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *CryptoIKEv2) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *CryptoIKEv2) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.NatKeepalive.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/nat/keepalive", data.getPath()))
	}
	if !data.Dpd.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dpd-container", data.getPath()))
	}
	if !data.DpdRetryInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dpd-container", data.getPath()))
	}
	if !data.DpdQuery.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dpd-container", data.getPath()))
	}
	return deletePaths
}
