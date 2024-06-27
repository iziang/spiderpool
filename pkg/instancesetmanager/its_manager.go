// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package instancesetmanager

import (
	"context"
	"fmt"

	apitypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kbv1alpha1 "github.com/spidernet-io/spiderpool/kbapi/workloads/v1alpha1"

	"github.com/spidernet-io/spiderpool/pkg/constant"
)

type InstanceSetManager interface {
	GetInstanceSetByName(ctx context.Context, namespace, name string, cached bool) (*kbv1alpha1.InstanceSet, error)
	ListInstanceSets(ctx context.Context, cached bool, opts ...client.ListOption) (*kbv1alpha1.InstanceSetList, error)
	IsValidInstanceSetPod(ctx context.Context, namespace, podName, podControllerType string) (bool, error)
}

type instanceSetManager struct {
	client    client.Client
	apiReader client.Reader
}

func NewInstanceSetManager(client client.Client, apiReader client.Reader) (InstanceSetManager, error) {
	if client == nil {
		return nil, fmt.Errorf("k8s client %w", constant.ErrMissingRequiredParam)
	}
	if apiReader == nil {
		return nil, fmt.Errorf("api reader %w", constant.ErrMissingRequiredParam)
	}

	return &instanceSetManager{
		client:    client,
		apiReader: apiReader,
	}, nil
}

func (sm *instanceSetManager) GetInstanceSetByName(ctx context.Context, namespace, name string, cached bool) (*kbv1alpha1.InstanceSet, error) {
	reader := sm.apiReader
	if cached == constant.UseCache {
		reader = sm.client
	}

	var sts kbv1alpha1.InstanceSet
	if err := reader.Get(ctx, apitypes.NamespacedName{Namespace: namespace, Name: name}, &sts); err != nil {
		return nil, err
	}

	return &sts, nil
}

func (sm *instanceSetManager) ListInstanceSets(ctx context.Context, cached bool, opts ...client.ListOption) (*kbv1alpha1.InstanceSetList, error) {
	reader := sm.apiReader
	if cached == constant.UseCache {
		reader = sm.client
	}

	var itsList kbv1alpha1.InstanceSetList
	if err := reader.List(ctx, &itsList, opts...); err != nil {
		return nil, err
	}

	return &itsList, nil
}

// IsValidInstanceSetPod only serves for InstanceSet pod, it will check the pod whether need to be cleaned up with the given params podNS, podName.
// Once the pod's controller InstanceSet was deleted, the pod's corresponding IPPool IP and Endpoint need to be cleaned up.
// Or the pod's controller InstanceSet decreased its replicas and the pod's index is out of replicas, it needs to be cleaned up too.
func (sm *instanceSetManager) IsValidInstanceSetPod(ctx context.Context, namespace, podName, podControllerType string) (bool, error) {
	if podControllerType != constant.KindInstanceSet {
		return false, fmt.Errorf("pod '%s/%s' is controlled by '%s' instead of InstanceSet", namespace, podName, podControllerType)
	}

	stsName, _, found := getInstanceSetNameAndOrdinal(podName)
	if !found {
		return false, nil
	}

	sts, err := sm.GetInstanceSetByName(ctx, namespace, stsName, constant.IgnoreCache)
	if err != nil {
		return false, client.IgnoreNotFound(err)
	}

	return sts.DeletionTimestamp == nil, nil
}
