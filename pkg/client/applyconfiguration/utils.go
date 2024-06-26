// This file is part of MinIO Operator
// Copyright (c) 2024 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/minio/operator/pkg/apis/job.min.io/v1alpha1"
	v2 "github.com/minio/operator/pkg/apis/minio.min.io/v2"
	stsminiov1alpha1 "github.com/minio/operator/pkg/apis/sts.min.io/v1alpha1"
	v1beta1 "github.com/minio/operator/pkg/apis/sts.min.io/v1beta1"
	jobminiov1alpha1 "github.com/minio/operator/pkg/client/applyconfiguration/job.min.io/v1alpha1"
	miniominiov2 "github.com/minio/operator/pkg/client/applyconfiguration/minio.min.io/v2"
	applyconfigurationstsminiov1alpha1 "github.com/minio/operator/pkg/client/applyconfiguration/sts.min.io/v1alpha1"
	stsminiov1beta1 "github.com/minio/operator/pkg/client/applyconfiguration/sts.min.io/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=job.min.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("CommandSpec"):
		return &jobminiov1alpha1.CommandSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CommandStatus"):
		return &jobminiov1alpha1.CommandStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MinIOJob"):
		return &jobminiov1alpha1.MinIOJobApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MinIOJobSpec"):
		return &jobminiov1alpha1.MinIOJobSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MinIOJobStatus"):
		return &jobminiov1alpha1.MinIOJobStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TenantRef"):
		return &jobminiov1alpha1.TenantRefApplyConfiguration{}

		// Group=minio.min.io, Version=v2
	case v2.SchemeGroupVersion.WithKind("Bucket"):
		return &miniominiov2.BucketApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("CertificateConfig"):
		return &miniominiov2.CertificateConfigApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("CertificateStatus"):
		return &miniominiov2.CertificateStatusApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("CustomCertificateConfig"):
		return &miniominiov2.CustomCertificateConfigApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("CustomCertificates"):
		return &miniominiov2.CustomCertificatesApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("ExposeServices"):
		return &miniominiov2.ExposeServicesApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("Features"):
		return &miniominiov2.FeaturesApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("KESConfig"):
		return &miniominiov2.KESConfigApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("LocalCertificateReference"):
		return &miniominiov2.LocalCertificateReferenceApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("Logging"):
		return &miniominiov2.LoggingApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("Pool"):
		return &miniominiov2.PoolApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("PoolStatus"):
		return &miniominiov2.PoolStatusApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("ServiceMetadata"):
		return &miniominiov2.ServiceMetadataApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("SideCars"):
		return &miniominiov2.SideCarsApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("Tenant"):
		return &miniominiov2.TenantApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("TenantDomains"):
		return &miniominiov2.TenantDomainsApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("TenantScheduler"):
		return &miniominiov2.TenantSchedulerApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("TenantSpec"):
		return &miniominiov2.TenantSpecApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("TenantStatus"):
		return &miniominiov2.TenantStatusApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("TenantUsage"):
		return &miniominiov2.TenantUsageApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("TierUsage"):
		return &miniominiov2.TierUsageApplyConfiguration{}

		// Group=sts.min.io, Version=v1alpha1
	case stsminiov1alpha1.SchemeGroupVersion.WithKind("Application"):
		return &applyconfigurationstsminiov1alpha1.ApplicationApplyConfiguration{}
	case stsminiov1alpha1.SchemeGroupVersion.WithKind("PolicyBinding"):
		return &applyconfigurationstsminiov1alpha1.PolicyBindingApplyConfiguration{}
	case stsminiov1alpha1.SchemeGroupVersion.WithKind("PolicyBindingSpec"):
		return &applyconfigurationstsminiov1alpha1.PolicyBindingSpecApplyConfiguration{}
	case stsminiov1alpha1.SchemeGroupVersion.WithKind("PolicyBindingStatus"):
		return &applyconfigurationstsminiov1alpha1.PolicyBindingStatusApplyConfiguration{}
	case stsminiov1alpha1.SchemeGroupVersion.WithKind("PolicyBindingUsage"):
		return &applyconfigurationstsminiov1alpha1.PolicyBindingUsageApplyConfiguration{}

		// Group=sts.min.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Application"):
		return &stsminiov1beta1.ApplicationApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PolicyBinding"):
		return &stsminiov1beta1.PolicyBindingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PolicyBindingSpec"):
		return &stsminiov1beta1.PolicyBindingSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PolicyBindingStatus"):
		return &stsminiov1beta1.PolicyBindingStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PolicyBindingUsage"):
		return &stsminiov1beta1.PolicyBindingUsageApplyConfiguration{}

	}
	return nil
}
