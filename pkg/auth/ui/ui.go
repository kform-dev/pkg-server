/*
Copyright 2024 Nokia.

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

package ui

import (
	"context"

	"github.com/kform-dev/pkg-server/pkg/auth"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/endpoints/request"
)

type ApiserverUserInfoProvider struct{}

var _ auth.UserInfoProvider = &ApiserverUserInfoProvider{}

func (p *ApiserverUserInfoProvider) GetUserInfo(ctx context.Context) *auth.UserInfo {
	userinfo, ok := request.UserFrom(ctx)
	if !ok {
		return nil
	}

	name := userinfo.GetName()
	if name == "" {
		return nil
	}

	for _, group := range userinfo.GetGroups() {
		if group == user.AllAuthenticated {
			return &auth.UserInfo{
				Name:  name, // k8s authentication only provides single name; use it for both values for now.
				Email: name,
			}
		}
	}

	return nil
}
