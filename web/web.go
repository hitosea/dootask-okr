package web

import "embed"

//go:embed dist/assets/*
var Assets embed.FS

//go:embed src/assets/*
var SrcAssets embed.FS

//go:embed dist/index.html
var IndexByte []byte

//go:embed favicon.ico
var Favicon embed.FS
