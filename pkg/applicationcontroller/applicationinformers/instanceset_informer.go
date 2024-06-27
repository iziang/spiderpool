// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"context"

	"k8s.io/client-go/tools/cache"

	"github.com/spidernet-io/spiderpool/pkg/logutils"
)

func (c *Controller) AddInstanceSetHandler(informer cache.SharedIndexInformer) error {
	controllersLogger.Info("Setting up InstanceSet handlers")

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onInstanceSetAdd,
		UpdateFunc: c.onInstanceSetUpdate,
		DeleteFunc: c.onInstanceSetDelete,
	})
	if nil != err {
		return err
	}

	return nil
}

func (c *Controller) onInstanceSetAdd(obj interface{}) {
	err := c.reconcileFunc(logutils.IntoContext(context.TODO(), controllersLogger), nil, obj)
	if nil != err {
		controllersLogger.Sugar().Errorf("onInstanceSetAdd: %v", err)
	}
}

func (c *Controller) onInstanceSetUpdate(oldObj interface{}, newObj interface{}) {
	err := c.reconcileFunc(logutils.IntoContext(context.TODO(), controllersLogger), oldObj, newObj)
	if nil != err {
		controllersLogger.Sugar().Errorf("onInstanceSetUpdate: %v", err)
	}
}

func (c *Controller) onInstanceSetDelete(obj interface{}) {
	err := c.cleanupFunc(logutils.IntoContext(context.TODO(), controllersLogger), obj)
	if nil != err {
		controllersLogger.Sugar().Errorf("onInstanceSetDelete: %v", err)
	}
}
