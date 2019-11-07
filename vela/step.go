// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// StepService handles retriving steps for builds
// from the server methods of the Vela API.
type StepService service

// Get returns the provided step.
func (s *StepService) Get(org, repo string, buildNum, target int) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, buildNum, target)

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetAll returns a list of all steps.
func (s *StepService) GetAll(org, repo string, buildNum int) (*[]library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps", org, repo, buildNum)

	// slice library Step type we want to return
	v := new([]library.Step)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Add constructs a step with the provided details.
func (s *StepService) Add(org, repo string, buildNum int, target *library.Step) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps", org, repo, buildNum)

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := s.client.Call("POST", u, target, v)
	return v, resp, err
}

// Update modifies a step with the provided details.
func (s *StepService) Update(org, repo string, buildNum int, target *library.Step) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, buildNum, *target.Number)

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Remove deletes the provided step.
func (s *StepService) Remove(org, repo string, buildNum, target int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, buildNum, target)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}