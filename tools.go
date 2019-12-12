package main

import (
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/go/ast/astutil"
	_ "golang.org/x/tools/go/buildutil"
	_ "golang.org/x/tools/go/loader"
	_ "golang.org/x/tools/imports"
	_ "golang.org/x/tools/internal/fastwalk"
	_ "golang.org/x/tools/internal/gopathwalk"
)
