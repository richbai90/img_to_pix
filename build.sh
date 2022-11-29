#! /usr/bin/env bash
ROOT="$(readlink -f "$(dirname "$0")")"
mkdir -p "$ROOT/bin/linux"
mkdir -p "$ROOT/bin/macos_x86"
mkdir -p "$ROOT/bin/windows"
mkdir -p "$ROOT/bin/macos_arm"
GOOS=linux GOARCH=amd64 go build -o "$ROOT/bin/linux/img_to_pix_linux"
GOOS=windows GOARCH=amd64 go build -o "$ROOT/bin/windows/img_to_pix.exe"
GOOS=darwin GOARCH=amd64 go build -o "$ROOT/bin/macos_x86/img_to_pix_macos_x86"
GOOS=darwin GOARCH=arm64 go build -o "$ROOT/bin/macos_arm/img_to_pix_macos_arm"

