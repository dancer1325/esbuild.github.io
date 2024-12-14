package main

import "github.com/evanw/esbuild/pkg/api"
import "os"

func main() {
result := api.Build(api.BuildOptions{
EntryPoints: []string{"app.ts"},
Bundle:      true,
Outdir:      "dist",
})
if len(result.Errors) != 0 {
os.Exit(1)
}
}