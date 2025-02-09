// Copyright (c) 2022 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GlobalThreatFeedLister helps list GlobalThreatFeeds.
// All objects returned here must be treated as read-only.
type GlobalThreatFeedLister interface {
	// List lists all GlobalThreatFeeds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.GlobalThreatFeed, err error)
	// Get retrieves the GlobalThreatFeed from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.GlobalThreatFeed, error)
	GlobalThreatFeedListerExpansion
}

// globalThreatFeedLister implements the GlobalThreatFeedLister interface.
type globalThreatFeedLister struct {
	indexer cache.Indexer
}

// NewGlobalThreatFeedLister returns a new GlobalThreatFeedLister.
func NewGlobalThreatFeedLister(indexer cache.Indexer) GlobalThreatFeedLister {
	return &globalThreatFeedLister{indexer: indexer}
}

// List lists all GlobalThreatFeeds in the indexer.
func (s *globalThreatFeedLister) List(selector labels.Selector) (ret []*v3.GlobalThreatFeed, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.GlobalThreatFeed))
	})
	return ret, err
}

// Get retrieves the GlobalThreatFeed from the index for a given name.
func (s *globalThreatFeedLister) Get(name string) (*v3.GlobalThreatFeed, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("globalthreatfeed"), name)
	}
	return obj.(*v3.GlobalThreatFeed), nil
}
