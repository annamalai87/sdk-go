// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/database"
	"github.com/go-vela/types/library"
)

// BuildService handles retriving builds from
// the server methods of the Vela API.
type BuildService service

// Get returns the provided build.
func (s *BuildService) Get(org, repo string, target int) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, target)

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetAll returns a list of all builds.
func (s *BuildService) GetAll(org, repo string) (*[]library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds", org, repo)

	// slice library Build type we want to return
	v := new([]library.Build)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetLogs returns the provided build logs.
func (s *BuildService) GetLogs(org, repo string, target int) (*[]database.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/logs", org, repo, target)

	// slice database Log type we want to return
	v := new([]database.Log)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Add constructs a build with the provided details.
func (s *BuildService) Add(org, repo string, target *library.Build) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds", org, repo)

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := s.client.Call("POST", u, target, v)
	return v, resp, err
}

// Update modifies a build with the provided details.
func (s *BuildService) Update(org, repo string, target *library.Build) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, *target.Number)

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Remove deletes the provided build.
func (s *BuildService) Remove(org, repo string, target int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, target)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// Restart takes the build provided and restarts it
func (s *BuildService) Restart(org, repo string, target int) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, target)

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := s.client.Call("POST", u, nil, v)
	return v, resp, err
}