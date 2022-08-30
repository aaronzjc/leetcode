@echo off
if exist coverage.out (
    del coverage.out
)

set GO111MODULE=on
go fmt ./...
if "%1" == "test" (
    go test -cover -coverprofile=coverage.out ./...   
) else if "%1" == "test_profile" (
    go test -cover -coverprofile=coverage.out ./...
    go tool cover -func coverage.out
) else if "%1" == "test_profile_html" (
    go test -cover -coverprofile=coverage.out ./...
    go tool cover -html coverage.out
)