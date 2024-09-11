// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	database "github.com/jellysmack-tech/provider-mysql/internal/controller/mysql/database"
	grant "github.com/jellysmack-tech/provider-mysql/internal/controller/mysql/grant"
	role "github.com/jellysmack-tech/provider-mysql/internal/controller/mysql/role"
	user "github.com/jellysmack-tech/provider-mysql/internal/controller/mysql/user"
	providerconfig "github.com/jellysmack-tech/provider-mysql/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		grant.Setup,
		role.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
