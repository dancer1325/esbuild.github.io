package main

import "github.com/evanw/esbuild/pkg/api"
import "os"

func main() {
result := api.Build(api.BuildOptions{
  EntryPoints: []string{"home.ts", "settings.ts"},
  Bundle:      true,
  Write:       true,
  Outdir:      "out",
})

if len(result.Errors) > 0 {
  os.Exit(1)
}
}