/*
Copyright The Godel Scheduler Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubewharf/godel-scheduler-api/pkg/apis/scheduling/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReservationLister helps list Reservations.
// All objects returned here must be treated as read-only.
type ReservationLister interface {
	// List lists all Reservations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Reservation, err error)
	// Reservations returns an object that can list and get Reservations.
	Reservations(namespace string) ReservationNamespaceLister
	ReservationListerExpansion
}

// reservationLister implements the ReservationLister interface.
type reservationLister struct {
	indexer cache.Indexer
}

// NewReservationLister returns a new ReservationLister.
func NewReservationLister(indexer cache.Indexer) ReservationLister {
	return &reservationLister{indexer: indexer}
}

// List lists all Reservations in the indexer.
func (s *reservationLister) List(selector labels.Selector) (ret []*v1alpha1.Reservation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Reservation))
	})
	return ret, err
}

// Reservations returns an object that can list and get Reservations.
func (s *reservationLister) Reservations(namespace string) ReservationNamespaceLister {
	return reservationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReservationNamespaceLister helps list and get Reservations.
// All objects returned here must be treated as read-only.
type ReservationNamespaceLister interface {
	// List lists all Reservations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Reservation, err error)
	// Get retrieves the Reservation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Reservation, error)
	ReservationNamespaceListerExpansion
}

// reservationNamespaceLister implements the ReservationNamespaceLister
// interface.
type reservationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Reservations in the indexer for a given namespace.
func (s reservationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Reservation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Reservation))
	})
	return ret, err
}

// Get retrieves the Reservation from the indexer for a given namespace and name.
func (s reservationNamespaceLister) Get(name string) (*v1alpha1.Reservation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("reservation"), name)
	}
	return obj.(*v1alpha1.Reservation), nil
}
