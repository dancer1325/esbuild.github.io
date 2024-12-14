import * as esbuild from 'esbuild'

await esbuild.build({
    entryPoints: ['src/content/examples/input/entryPoints/home.ts', 'src/content/examples/input/entryPoints/settings.ts'],
    bundle: true,
    write: true,
    outdir: 'out',
})