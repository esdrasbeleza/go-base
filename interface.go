package main

import "context"

//go:generate go tool mockgen -source=$GOFILE -destination=interface_mock.go -package=$GOPACKAGE
type MyInterface interface {
	SomeFunction(ctx context.Context, arg1, arg2 string) (string, error)
}
