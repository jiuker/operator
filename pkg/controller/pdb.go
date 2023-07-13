// Copyright (C) 2023, MinIO, Inc.
//
// This code is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License, version 3,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License, version 3,
// along with this program.  If not, see <http://www.gnu.org/licenses/>

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	v2 "github.com/minio/operator/pkg/apis/minio.min.io/v2"
	v1 "k8s.io/api/policy/v1"
	"k8s.io/api/policy/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/klog/v2"
)

// DeletePDB - delete PDB for tenant
func (c *Controller) DeletePDB(ctx context.Context, t *v2.Tenant) (err error) {
	available := c.PDBAvailable()
	if !available.Available() {
		return nil
	}
	if available.V1Available() {
		err := c.kubeClientSet.PolicyV1().PodDisruptionBudgets(t.Namespace).DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{
			LabelSelector: metav1.SetAsLabelSelector(labels.Set{
				v2.TenantLabel: t.Name,
			}).String(),
		})
		if err != nil {
			return err
		}
	}
	if available.V1BetaAvailable() {
		err := c.kubeClientSet.PolicyV1beta1().PodDisruptionBudgets(t.Namespace).DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{
			LabelSelector: metav1.SetAsLabelSelector(labels.Set{
				v2.TenantLabel: t.Name,
			}).String(),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateOrUpdatePDB - hold PDB as expected
func (c *Controller) CreateOrUpdatePDB(ctx context.Context, t *v2.Tenant) (err error) {
	available := c.PDBAvailable()
	if !available.Available() {
		return nil
	}
	for _, pool := range t.Spec.Pools {
		if strings.TrimSpace(pool.Name) == "" {
			continue
		}
		if available.V1Available() {
			var pdb *v1.PodDisruptionBudget
			var isCreate = false
			pdb, err = c.kubeClientSet.PolicyV1().PodDisruptionBudgets(t.Namespace).Get(ctx, pool.Name, metav1.GetOptions{})
			if err != nil {
				if k8serrors.IsNotFound(err) {
					pdb = &v1.PodDisruptionBudget{}
					isCreate = true
				} else {
					return err
				}
			}
			if !isCreate {
				// exist and as expected
				if pdb.Spec.MaxUnavailable != nil && pdb.Spec.MaxUnavailable.IntValue() == (int(pool.Servers/2)) {
					return nil
				}
			}
			// set filed we expected
			pdb.Name = pool.Name
			pdb.Namespace = t.Namespace
			maxUnavailable := intstr.FromInt(int(pool.Servers / 2))
			pdb.Spec.MaxUnavailable = &maxUnavailable
			pdb.Labels = map[string]string{
				v2.TenantLabel: t.Name,
				v2.PoolLabel:   pool.Name,
			}
			pdb.Spec.Selector = metav1.SetAsLabelSelector(labels.Set{
				v2.PoolLabel:   pool.Name,
				v2.TenantLabel: t.Name,
			})
			if isCreate {
				_, err = c.kubeClientSet.PolicyV1().PodDisruptionBudgets(t.Namespace).Create(ctx, pdb, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			} else {
				patchData := map[string]interface{}{
					"spec": map[string]interface{}{
						"maxUnavailable": pdb.Spec.MaxUnavailable,
					},
				}
				pData, err := json.Marshal(patchData)
				if err != nil {
					return err
				}
				_, err = c.kubeClientSet.PolicyV1().PodDisruptionBudgets(t.Namespace).Patch(ctx, t.Name, types.MergePatchType, pData, metav1.PatchOptions{})
				if err != nil {
					return nil
				}
			}
		}
		if available.V1BetaAvailable() {
			var pdb *v1beta1.PodDisruptionBudget
			var isCreate = false
			pdb, err = c.kubeClientSet.PolicyV1beta1().PodDisruptionBudgets(t.Namespace).Get(ctx, pool.Name, metav1.GetOptions{})
			if err != nil {
				if k8serrors.IsNotFound(err) {
					pdb = &v1beta1.PodDisruptionBudget{}
					isCreate = true
				} else {
					return err
				}
			}
			if !isCreate {
				// exist and as expected
				if pdb.Spec.MaxUnavailable != nil && pdb.Spec.MaxUnavailable.IntValue() == (int(pool.Servers/2)) {
					return nil
				}
			}
			// set filed we expected
			pdb.Name = pool.Name
			pdb.Namespace = t.Namespace
			maxUnavailable := intstr.FromInt(int(pool.Servers / 2))
			pdb.Spec.MaxUnavailable = &maxUnavailable
			pdb.Labels = map[string]string{
				v2.TenantLabel: t.Name,
				v2.PoolLabel:   pool.Name,
			}
			pdb.Spec.Selector = metav1.SetAsLabelSelector(labels.Set{
				v2.PoolLabel:   pool.Name,
				v2.TenantLabel: t.Name,
			})
			if isCreate {
				_, err = c.kubeClientSet.PolicyV1beta1().PodDisruptionBudgets(t.Namespace).Create(ctx, pdb, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			} else {
				patchData := map[string]interface{}{
					"spec": map[string]interface{}{
						"maxUnavailable": pdb.Spec.MaxUnavailable,
					},
				}
				pData, err := json.Marshal(patchData)
				if err != nil {
					return err
				}
				_, err = c.kubeClientSet.PolicyV1beta1().PodDisruptionBudgets(t.Namespace).Patch(ctx, t.Name, types.MergePatchType, pData, metav1.PatchOptions{})
				if err != nil {
					return nil
				}
			}
		}

	}
	if len(t.Spec.Pools) == 0 {
		return fmt.Errorf("%s empty pools", t.Name)
	}
	return nil
}

type PDBAvailable struct {
	v1     bool
	v1beta bool
}

// V1Available - show if it support PDB v1
func (p *PDBAvailable) V1Available() bool {
	return p.v1
}

// V1BetaAvailable - show if it support PDB v1beta
func (p *PDBAvailable) V1BetaAvailable() bool {
	return p.v1beta
}

// Available - show if it support PDB
func (p *PDBAvailable) Available() bool {
	return p.v1 || p.v1beta
}

var globalPDBAvailable = PDBAvailable{}
var globalPDBAvailableOnce = sync.Once{}

func (c *Controller) PDBAvailable() PDBAvailable {
	globalPDBAvailableOnce.Do(func() {
		defer func() {
			if globalPDBAvailable.v1 {
				klog.Infof("PodDisruptionBudget: v1")
			} else if globalPDBAvailable.v1beta {
				klog.Infof("PodDisruptionBudget: v1beta")
			} else {
				klog.Infof("PodDisruptionBudget: unsupport")
			}
		}()
		resouces, _ := c.kubeClientSet.Discovery().ServerPreferredResources()
		for _, r := range resouces {
			if r.GroupVersion == "policy/v1" {
				for _, api := range r.APIResources {
					if api.Kind == "PodDisruptionBudget" {
						globalPDBAvailable.v1 = true
						return
					}
				}
			}
			if r.GroupVersion == "policy/v1beta" {
				for _, api := range r.APIResources {
					if api.Kind == "PodDisruptionBudget" {
						globalPDBAvailable.v1beta = true
						return
					}
				}
			}
		}
	})
	return globalPDBAvailable
}
