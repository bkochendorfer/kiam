// Copyright 2017 uSwitch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kiam

import (
	"context"
	"testing"

	"github.com/uswitch/kiam/pkg/aws/sts"
)

func TestAddsPrefix(t *testing.T) {
	resolver := sts.DefaultResolver("arn:aws:iam::account-id:role/")
	role, _ := resolver.Resolve(context.Background(), "myrole")

	if role != "arn:aws:iam::account-id:role/myrole" {
		t.Error("unexpected role, was:", role)
	}
}

func TestUsesAbsoluteARN(t *testing.T) {
	resolver := sts.DefaultResolver("arn:aws:iam::account-id:role/")
	role, _ := resolver.Resolve(context.Background(), "arn:aws:iam::some-other-account:role/another-role")

	if role != "arn:aws:iam::some-other-account:role/another-role" {
		t.Error("unexpected role, was:", role)
	}
}
