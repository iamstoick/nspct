package main

import (
	"reflect"
	"testing"
)

func Test_randSeq(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randSeq(tt.args.n); got != tt.want {
				t.Errorf("randSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getURLHeaders(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getURLHeaders(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getURLHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printHelp(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printHelp()
		})
	}
}

func Test_getURLHeader(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getURLHeader(tt.args.url); got != tt.want {
				t.Errorf("getURLHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dnsQuery(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dnsQuery(tt.args.url)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
