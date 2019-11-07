// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestVela_NewClient(t *testing.T) {
	// setup types
	addr := "http://localhost:8080"

	url, err := url.Parse(addr)
	if err != nil {
		t.Errorf("Unable to parse url: %v", err)
	}

	want := &Client{
		client:    http.DefaultClient,
		baseURL:   url,
		UserAgent: userAgent,
	}
	want.Authentication = &AuthenticationService{client: want}
	want.Authorization = &AuthorizationService{client: want}
	want.Log = &LogService{client: want}
	want.Build = &BuildService{client: want}
	want.Repo = &RepoService{client: want}
	want.Secret = &SecretService{client: want}
	want.Step = &StepService{client: want}
	want.Svc = &SvcService{client: want}

	// run test
	got, err := NewClient(addr, nil)
	if err != nil {
		t.Errorf("NewClient returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewClient is %v, want %v", got, want)
	}
}

func TestVela_NewClient_EmptyUrl(t *testing.T) {
	// run test
	got, err := NewClient("", nil)
	if err == nil {
		t.Errorf("NewClient should have returned err")
	}

	if got != nil {
		t.Errorf("NewClient is %v, want nil", got)
	}
}

func TestVela_NewClient_BadUrl(t *testing.T) {
	// run test
	got, err := NewClient("!@#$%^&*()", nil)
	if err == nil {
		t.Errorf("NewClient should have returned err")
	}

	if got != nil {
		t.Errorf("NewClient is %v, want nil", got)
	}
}

func TestVela_buildURLForRequest_NoSlash(t *testing.T) {
	// setup types
	want := "http://localhost:8080/test"
	c, err := NewClient("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("test")
	if err != nil {
		t.Errorf("buildURLForRequest returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildURLForRequest is %v, want %v", got, want)
	}
}

func TestVela_buildURLForRequest_PrefixSlash(t *testing.T) {
	// setup types
	want := "http://localhost:8080/test"
	c, err := NewClient("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("/test")
	if err != nil {
		t.Errorf("buildURLForRequest returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildURLForRequest is %v, want %v", got, want)
	}
}

func TestVela_buildURLForRequest_SuffixSlash(t *testing.T) {
	// setup types
	want := "http://localhost:8080/test/"
	c, err := NewClient("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("test/")
	if err != nil {
		t.Errorf("buildURLForRequest returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildURLForRequest is %v, want %v", got, want)
	}
}

func TestVela_buildURLForRequest_BadUrl(t *testing.T) {
	// setup types
	c, err := NewClient("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("!@#$%^&*()")
	if err == nil {
		t.Errorf("buildURLForRequest should have returned err")
	}

	if len(got) > 0 {
		t.Errorf("buildURLForRequest is %v, want \"\"", got)
	}
}

func TestVela_addAuthentication(t *testing.T) {
	// setup types
	want := "Bearer foobar"
	c, err := NewClient("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	r, err := http.NewRequest("GET", "http://localhost:8080/health", nil)
	if err != nil {
		t.Errorf("Unable to create new request: %v", err)
	}

	// run test
	c.Authentication.SetTokenAuth("foobar")
	c.addAuthentication(r)

	got := r.Header.Get("Authorization")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addAuthentication is %v, want %v", got, want)
	}
}

func TestVela_NewRequest(t *testing.T) {
	// setup types
	want, err := http.NewRequest("GET", "http://localhost:8080/health", nil)
	if err != nil {
		t.Errorf("Unable to create new request: %v", err)
	}
	want.Header.Add("Content-Type", "application/json")

	c, err := NewClient("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Errorf("NewRequest returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewRequest is %v, want %v", got, want)
	}
}

func TestVela_NewRequest_BadUrl(t *testing.T) {
	// setup types
	c, err := NewClient("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.NewRequest("GET", "!@#$%^&*()", nil)
	if err == nil {
		t.Errorf("NewRequest should have returned err")
	}

	if got != nil {
		t.Errorf("NewRequest is %v, want nil", got)
	}
}

type options struct {
	ShowAll bool `url:"all"`
	Page    int  `url:"page"`
}

func TestVela_addOptions(t *testing.T) {
	// setup types
	want := "http://localhost:8080?all=true&page=1"
	options := options{ShowAll: true, Page: 1}

	// run test
	got, err := addOptions("http://localhost:8080", options)
	if err != nil {
		t.Errorf("addOptions returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}

func TestVela_addOptions_BadOptions(t *testing.T) {
	// setup types
	want := "http://localhost:8080"

	// run test
	got, err := addOptions("http://localhost:8080", 87)
	if err == nil {
		t.Errorf("addOptions should have returned err")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}

func TestVela_addOptions_BadUrl(t *testing.T) {
	// setup types
	want := "!@#$%^&*()"
	options := options{ShowAll: true, Page: 1}

	// run test
	got, err := addOptions("!@#$%^&*()", options)
	if err == nil {
		t.Errorf("addOptions should have returned err")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}

func TestVela_addOptions_NilOptions(t *testing.T) {
	// setup types
	want := "http://localhost:8080"

	// run test
	got, err := addOptions("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("addOptions returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}