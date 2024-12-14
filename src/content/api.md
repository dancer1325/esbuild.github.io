* API
  * ways to access to it
    * CL
      * 👀flags forms 👀
        * form `--foo` -- for enabling -- boolean flags
          * _Example:_ [`--minify`](#minify)
        * form `--foo=bar` -- for -- flags / have 1! value & are specified 1!
          * _Example:_ [`--platform=`](#platform)
        * form `--foo:bar`-- for -- flags / have multiple values & can be re-specified multiple times
          * _Example:_ [`--external:`](#external)
      * general == NOT esbuild-specific
        * 👀-> your current shell -- interprets the -- command's arguments | BEFORE the command running sees them👀
            * == specific behavior -- depends on the -- used shell
            * _Example:_ `echo` command just writes out / it reads in
                * `echo "foo"` can print `foo` instead of `"foo"`
                * `echo *.json` can print `package.json` instead of `*.json`
      * recommendations
        * use esbuild's JavaScript or Go APIs
    * JS
      * see
        * [JS-specific details](#js-details)
        * [browser](#browser)
        * [TypeScript type definitions](https://github.com/evanw/esbuild/blob/main/lib/shared/types.ts)
    * Go
      * see 
        * [`pkg/api`](https://pkg.go.dev/github.com/evanw/esbuild/pkg/api)
        * [`pkg/cli`](https://pkg.go.dev/github.com/evanw/esbuild/pkg/cli) 
  * concepts & parameters
    * == | ALL allowed ways to access == SAME documentation

# Overview

* 👀MOST commonly-used esbuild APIs 👀
  * [build](#build)
  * [transform](#transform)

## Build

* PRIMARY interface to esbuild
* arguments
  * \>=1 [entry point](#entry-points) files
  * VARIOUS options
* how does it work?
  * 💡arguments -- writes the -- results | file system 💡

* _Example:_ enables [bundling](#bundle) + [output directory](#outdir)
  * | CLI
    ```
    esbuild app.ts --bundle --outdir=dist
    ```  
  * | JS
    ```
    node module.mjs
    ```
  * | GO  
    * Problems: How to run?
      * Attempt1: `go main.go`

* build context
  * allows
    * using deeply the build API
  * == explicit object | JS and Go
  * == implicit | CLI
  * ALL builds | SAME context -> SAME build options -> subsequent builds are done incrementally
    * == they reuse some work -- from -- previous builds / improve performance
  * use cases
    * | development
      * Reason: 🧠 esbuild can rebuild your app | background & while you work 🧠

* incremental build APIs
  * [**Watch mode**](#watch)
    * TODO: tells esbuild to watch the file system and
                    automatically rebuild for you whenever you edit and save a file that
                    could invalidate the build. Here's an example:

                - example:
                    noCheck: true

                    cli:
                      - $: |
                          esbuild app.ts --bundle --outdir=dist --watch
                      - expect: |
                          [watch] build finished, watching for changes...

                    mjs: |
                      let ctx = await esbuild.context({
                        entryPoints: ['app.ts'],
                        bundle: true,
                        outdir: 'dist',
                      })

                      await ctx.watch()

                    go: |
                      ctx, err := api.Context(api.BuildOptions{
                        EntryPoints: []string{"app.ts"},
                        Bundle:      true,
                        Outdir:      "dist",
                      })

                      err2 := ctx.Watch(api.WatchOptions{})

  * [**Serve mode**](#serve)
    * starts a local development server that serves the
                        results of the latest build. Incoming requests automatically start new
                        builds so your web app is always up to date when you reload the page in
                        the browser. Here's an example:

                    - example:
                        noCheck: true

                        cli:
                          - $: |
                              esbuild app.ts --bundle --outdir=dist --serve
                          - expect: |2

             >                   Local:   http://127.0.0.1:8000/
             >                   Network: http://192.168.0.1:8000/

                              127.0.0.1:61302 - "GET /" 200 [1ms]

                        mjs: |
                          let ctx = await esbuild.context({
                            entryPoints: ['app.ts'],
                            bundle: true,
                            outdir: 'dist',
                          })

                          let { host, port } = await ctx.serve()

                        go: |
                          ctx, err := api.Context(api.BuildOptions{
                            EntryPoints: []string{"app.ts"},
                            Bundle:      true,
                            Outdir:      "dist",
                          })

                          server, err2 := ctx.Serve(api.ServeOptions{})

                    - ul:
                      - >
                        [**Rebuild mode**](#rebuild) lets you manually invoke a build. This is useful
                        when integrating esbuild with other tools (e.g. using a custom file watcher
                        or development server instead of esbuild's built-in ones). Here's an example:

                    - example:
                        noCheck: true

                        cli: |
                          # The CLI does not have an API for "rebuild"

                        mjs: |
                          let ctx = await esbuild.context({
                            entryPoints: ['app.ts'],
                            bundle: true,
                            outdir: 'dist',
                          })

                          for (let i = 0; i < 5; i++) {
                            let result = await ctx.rebuild()
                          }

                        go: |
                          ctx, err := api.Context(api.BuildOptions{
                            EntryPoints: []string{"app.ts"},
                            Bundle:      true,
                            Outdir:      "dist",
                          })

                          for i := 0; i < 5; i++ {
                            result := ctx.Rebuild()
                          }

                    - p: >
                        These three incremental build APIs can be combined. To enable [live reloading](#live-reload)
                        (automatically reloading the page when you edit and save a file) you'll need
                        to enable [watch](#watch) and [serve](#serve) together on the same context.

                    - p: >
                        When you are done with a context object, you can call `dispose()` on the
                        context to wait for existing builds to finish, stop watch and/or serve mode,
                        and free up resources.

                    - p: >
                        The build and context APIs both take the following options:

                    - available-options:
                      - Alias
                      - Allow overwrite
                      - Analyze
                      - Asset names
                      - Banner
                      - Bundle
                      - Cancel
                      - Charset
                      - Chunk names
                      - Color
                      - Conditions
                      - Define
                      - Drop
                      - Drop labels
                      - Entry names
                      - Entry points
                      - External
                      - Footer
                      - Format
                      - Format messages
                      - Global name
                      - Ignore annotations
                      - Inject
                      - JSX
                      - JSX dev
                      - JSX factory
                      - JSX fragment
                      - JSX import source
                      - JSX side effects
                      - Keep names
                      - Legal comments
                      - Line limit
                      - Live reload
                      - Loader
                      - Log level
                      - Log limit
                      - Log override
                      - Main fields
                      - Mangle props
                      - Metafile
                      - Minify
                      - Node paths
                      - Out extension
                      - Outbase
                      - Outdir
                      - Outfile
                      - Packages
                      - Platform
                      - Preserve symlinks
                      - Public path
                      - Pure
                      - Rebuild
                      - Resolve extensions
                      - Serve
                      - Source root
                      - Sourcefile
                      - Sourcemap
                      - Sources content
                      - Splitting
                      - Stdin
                      - Supported
                      - Target
                      - Tree shaking
                      - Tsconfig
                      - Tsconfig raw
                      - Watch
                      - Working directory
                      - Write

## Transform

  - p: >
      This is a limited special-case of [build](#build) that transforms a string
      of code representing an in-memory file in an isolated environment that's
      completely disconnected from any other files. Common uses include minifying
      code and transforming TypeScript into JavaScript. Here's an example:

  - example:
      cli:
        - $: |
            echo 'let x: number = 1' | esbuild --loader=ts
        - expect: |
            let x = 1;

      mjs: |
        import * as esbuild from 'esbuild'

        let ts = 'let x: number = 1'
        let result = await esbuild.transform(ts, {
          loader: 'ts',
        })
        console.log(result)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          ts := "let x: number = 1"
          result := api.Transform(ts, api.TransformOptions{
            Loader: api.LoaderTS,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      Taking a string instead of a file as input is more ergonomic for certain
      use cases. File system isolation has certain advantages (e.g. works in the
      browser, not affected by nearby `package.json` files) and certain
      disadvantages (e.g. can't be used with [bundling](#bundle) or
      [plugins](/plugins/)). If your use case doesn't fit the transform API then
      you should use the more general [build](#build) API instead.

  - p: >
      The transform API takes the following options:

  - available-options:
    - Banner
    - Charset
    - Color
    - Define
    - Drop
    - Drop labels
    - Footer
    - Format
    - Format messages
    - Global name
    - Ignore annotations
    - JSX
    - JSX dev
    - JSX factory
    - JSX fragment
    - JSX import source
    - JSX side effects
    - Keep names
    - Legal comments
    - Line limit
    - Loader
    - Log level
    - Log limit
    - Log override
    - Mangle props
    - Minify
    - Platform
    - Pure
    - Source root
    - Sourcefile
    - Sourcemap
    - Sources content
    - Supported
    - Target
    - Tree shaking
    - Tsconfig raw

## js-details: JS-specific details

  - p: >
      The JS API for esbuild comes in both asynchronous and synchronous flavors.
      The [asynchronous API](#js-async) is recommended because it works in all
      environments and it's faster and more powerful. The [synchronous API](#js-sync)
      only works in node and can only do certain things, but it's sometimes
      necessary in certain node-specific situations. In detail:

  - h4#js-async: Async API

  - p: >
      Asynchronous API calls return their results using a promise. Note that
      you'll likely have to use the `.mjs` file extension in node due to the
      use of the `import` and top-level `await` keywords:

  - pre.js: |
      import * as esbuild from 'esbuild'

      let result1 = await esbuild.transform(code, options)
      let result2 = await esbuild.build(options)

  - p: >
      Pros:

  - ul:
    - >
      You can use [plugins](/plugins/) with the asynchronous API
    - >
      The current thread is not blocked so you can perform other work in the meantime
    - >
      You can run many simultaneous esbuild API calls concurrently which are
      then spread across all available CPUs for maximum performance

  - p: >
      Cons:

  - ul:
    - >
      Using promises can result in messier code, especially in CommonJS where
      [top-level await](https://v8.dev/features/top-level-await) is not available
    - >
      Doesn't work in situations that must be synchronous such as within
      <a href="https://nodejs.org/api/modules.html#requireextensions"><code>require<wbr>.extensions</code></a>

  - h4#js-sync: Sync API

  - p: >
      Synchronous API calls return their results inline:

  - pre.js: |
      let esbuild = require('esbuild')

      let result1 = esbuild.transformSync(code, options)
      let result2 = esbuild.buildSync(options)

  - p: >
      Pros:

  - ul:
    - >
      Avoiding promises can result in cleaner code, especially when
      [top-level await](https://v8.dev/features/top-level-await) is not available
    - >
      Works in situations that must be synchronous such as within
      <a href="https://nodejs.org/api/modules.html#requireextensions"><code>require<wbr>.extensions</code></a>

  - p: >
      Cons:

  - ul:
    - >
      You can't use [plugins](/plugins/) with the synchronous API since plugins are asynchronous
    - >
      It blocks the current thread so you can't perform other work in the meantime
    - >
      Using the synchronous API prevents esbuild from parallelizing esbuild API calls

## browser: In the browser

  - p: >
      The esbuild API can also run in the browser using WebAssembly in a Web
      Worker. To take advantage of this you will need to install the
      `esbuild-wasm` package instead of the `esbuild` package:

  - pre: |
      npm install esbuild-wasm

  - p: >
      The API for the browser is similar to the API for node except that you
      need to call `initialize()` first, and you need to pass the URL of
      the WebAssembly binary. The synchronous versions of the API are also
      not available. Assuming you are using a bundler, that would look
      something like this:

  - pre.js: |
      import * as esbuild from 'esbuild-wasm'

      await esbuild.initialize({
        wasmURL: './node_modules/esbuild-wasm/esbuild.wasm',
      })

      let result1 = await esbuild.transform(code, options)
      let result2 = esbuild.build(options)

  - p: >
      If you're already running this code from a worker and don't want
      `initialize` to create another worker, you can pass <code>worker: <wbr>false</code>
      to it. Then it will create a WebAssembly module in the same thread
      as the thread that calls `initialize`.

  - p: >
      You can also use esbuild's API as a script tag in a HTML file without
      needing to use a bundler by loading the `lib/browser.min.js` file with
      a `<script>` tag. In this case the API creates a global called `esbuild`
      that holds the API object:

  - pre.html: |
      <script src="./node_modules/esbuild-wasm/lib/browser.min.js"></script>
      <script>
        esbuild.initialize({
          wasmURL: './node_modules/esbuild-wasm/esbuild.wasm',
        }).then(() => {
          ...
        })
      </script>

  - p: >
      If you want to use this API with ECMAScript modules, you should import
      the `esm/browser.min.js` file instead:

  - pre.html: |
      <script type="module">
        import * as esbuild from './node_modules/esbuild-wasm/esm/browser.min.js'

        await esbuild.initialize({
          wasmURL: './node_modules/esbuild-wasm/esbuild.wasm',
        })

        ...
      </script>

  # General options

  ## Bundle

  - p: >
      To bundle a file means to inline any imported dependencies into the
      file itself. This process is recursive so dependencies of dependencies
      (and so on) will also be inlined. By default esbuild will _not_ bundle
      the input files. Bundling must be explicitly enabled like this:

  - example:
      in:
        in.js: '1 + 2'

      cli: |
          esbuild in.js --bundle

      mjs: |
        import * as esbuild from 'esbuild'

        console.log(await esbuild.build({
          entryPoints: ['in.js'],
          bundle: true,
          outfile: 'out.js',
        }))

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"in.js"},
            Bundle:      true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Refer to the [getting started guide](/getting-started/#your-first-bundle)
      for an example of bundling with real-world code.

  - p: >
      Note that bundling is different than file concatenation. Passing
      esbuild multiple input files with bundling enabled will create multiple
      separate bundles instead of joining the input files together. To join
      a set of files together with esbuild, import them all into a single
      entry point file and bundle just that one file with esbuild.

  - h4: Non-analyzable imports

  - p: >
      Import paths are currently only bundled if they are a string literal or
      a [glob pattern](#glob). Other forms of import paths are not bundled,
      and are instead preserved verbatim in the generated output. This is
      because bundling is a compile-time operation and esbuild doesn't
      support all forms of run-time path resolution. Here are some examples:

  - pre.js: |
      // Analyzable imports (will be bundled by esbuild)
      import 'pkg';
      import('pkg');
      require('pkg');
      import(`./locale-${foo}.json`);
      require(`./locale-${foo}.json`);

      // Non-analyzable imports (will not be bundled by esbuild)
      import(`pkg/${foo}`);
      require(`pkg/${foo}`);
      ['pkg'].map(require);

  - p: >
      The way to work around non-analyzable imports is to mark the package
      containing this problematic code as [external](#external) so that
      it's not included in the bundle. You will then need to ensure that a
      copy of the external package is available to your bundled code at run-time.

  - p: >
      Some bundlers such as [Webpack](https://webpack.js.org/) try to support
      all forms of run-time path resolution by including all potentially-reachable
      files in the bundle and then emulating a file system at run-time. However,
      run-time file system emulation is out of scope and will not be implemented
      in esbuild. If you really need to bundle code that does this, you will
      likely need to use another bundler instead of esbuild.

  - h4#glob: Glob-style imports

  - p: >
      Import paths that are evaluated at run-time can now be bundled in certain
      limited situations. The import path expression must be a form of string
      concatenation and must start with either `./` or `../`. Each non-string
      expression in the string concatenation chain becomes a wildcard in a
      [glob](https://en.wikipedia.org/wiki/Glob_(programming)) pattern. Some
      examples:

  - pre.js: |
      // These two forms are equivalent
      const json1 = require('./data/' + kind + '.json')
      const json2 = require(`./data/${kind}.json`)

  - p: >
      When you do this, esbuild will search the file system for all files that
      match the pattern and include all of them in the bundle along with a map
      that maps the matching import path to the bundled module. The import
      expression will be replaced with a lookup into that map. An error will
      be thrown at run-time if the import path is not present in the map. The
      generated code will look something like this (unimportant parts were
      omitted for brevity):

  - pre.js: |
      // data/bar.json
      var require_bar = ...;

      // data/foo.json
      var require_foo = ...;

      // require("./data/**/*.json") in example.js
      var globRequire_data_json = __glob({
        "./data/bar.json": () => require_bar(),
        "./data/foo.json": () => require_foo()
      });

      // example.js
      var json1 = globRequire_data_json("./data/" + kind + ".json");
      var json2 = globRequire_data_json(`./data/${kind}.json`);

  - p: >
      This feature works with `require(...)` and `import(...)` because these
      can all accept run-time expressions. It does not work with `import` and
      `export` statements because these cannot accept run-time expressions. If
      you want to prevent esbuild from trying to bundle these imports, you
      should move the string concatenation expression outside of the
      `require(...)` or `import(...)`. For example:

  - pre.js: |
      // This will be bundled
      const json1 = require('./data/' + kind + '.json')

      // This will not be bundled
      const path = './data/' + kind + '.json'
      const json2 = require(path)

  - p: >
      Note that using this feature means esbuild will potentially do a lot of
      file system I/O to find all possible files that might match the pattern.
      This is by design, and is not a bug. If this is a concern, there are two
      ways to reduce the amount of file system I/O that esbuild does:

  - ol:
    - >
      <p>
      The simplest approach is to put all files that you want to import for a
      given run-time import expression in a subdirectory and then include the
      subdirectory in the pattern. This limits esbuild to searching inside that
      subdirectory since esbuild doesn't consider `..` path elements during
      pattern-matching.
      </p>

    - >
      <p>
      Another approach is to prevent esbuild from searching into any subdirectory
      at all. The pattern matching algorithm that esbuild uses only allows
      a wildcard to match something containing a `/` path separator if that
      wildcard has a `/` before it in the pattern. So for example <code>'./data/' + <wbr>x + <wbr>'.json'</code>
      will match `x` with anything in any subdirectory while <code>'./data-' + <wbr>x + <wbr>'.json'</code>
      will only match `x` with anything in the top-level directory (but not in
      any subdirectory).
      </p>

  ## Cancel

  - p: >
      If you are using [rebuild](#rebuild) to manually invoke incremental
      builds, you may want to use this cancel API to end the current build
      early so that you can start a new one. You can do that like this:

  - example:
      noCheck: true

      cli: |
        # The CLI does not have an API for "cancel"

      mjs: |
        import * as esbuild from 'esbuild'
        import process from 'node:process'

        let ctx = await esbuild.context({
          entryPoints: ['app.ts'],
          bundle: true,
          outdir: 'www',
          logLevel: 'info',
        })

        // Whenever we get some data over stdin
        process.stdin.on('data', async () => {
          try {
            // Cancel the already-running build
            await ctx.cancel()

            // Then start a new build
            console.log('build:', await ctx.rebuild())
          } catch (err) {
            console.error(err)
          }
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          ctx, err := api.Context(api.BuildOptions{
            EntryPoints: []string{"app.ts"},
            Bundle:      true,
            Outdir:      "www",
            LogLevel:    api.LogLevelInfo,
          })
          if err != nil {
            os.Exit(1)
          }

          // Whenever we get some data over stdin
          buf := make([]byte, 100)
          for {
            if n, err := os.Stdin.Read(buf); err != nil || n == 0 {
              break
            }
            go func() {
              // Cancel the already-running build
              ctx.Cancel()

              // Then start a new build
              result := ctx.Rebuild()
              fmt.Fprintf(os.Stderr, "build: %v\n", result)
            }()
          }
        }

  - p: >
      Make sure to wait until the cancel operation is done before starting a
      new build (i.e. `await` the returned promise when using JavaScript),
      otherwise the next [rebuild](#rebuild) will give you the just-canceled
      build that still hasn't ended yet. Note that plugin [on-end callbacks](/plugins/#on-end)
      will still be run regardless of whether or not the build was canceled.

  ## Live reload

  - p: >
      Live reload is an approach to development where you have your browser
      open and visible at the same time as your code editor. When you edit
      and save your source code, the browser automatically reloads and the
      reloaded version of the app contains your changes. This means you can
      iterate faster because you don't have to manually switch to your browser,
      reload, and then switch back to your code editor after every change.
      It's very helpful when changing CSS, for example.

  - p: >
      There is no esbuild API for live reloading directly. Instead, you can
      construct live reloading by combining [watch mode](#watch) (to
      automatically start a build when you edit and save a file) and
      [serve mode](#serve) (to serve the latest build, but block until it's
      done) plus a small bit of client-side JavaScript code that you add to
      your app only during development.

  - p: >
      The first step is to enable [watch](#watch) and [serve](#serve) together:

  - example:
      noCheck: true

      cli: |
        esbuild app.ts --bundle --outdir=www --watch --servedir=www

      mjs: |
        import * as esbuild from 'esbuild'

        let ctx = await esbuild.context({
          entryPoints: ['app.ts'],
          bundle: true,
          outdir: 'www',
        })

        await ctx.watch()

        let { host, port } = await ctx.serve({
          servedir: 'www',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          ctx, err := api.Context(api.BuildOptions{
            EntryPoints: []string{"app.ts"},
            Bundle:      true,
            Outdir:      "www",
          })
          if err != nil {
            os.Exit(1)
          }

          err2 := ctx.Watch(api.WatchOptions{})
          if err2 != nil {
            os.Exit(1)
          }

          result, err3 := ctx.Serve(api.ServeOptions{
            Servedir: "www",
          })
          if err3 != nil {
            os.Exit(1)
          }
        }

  - p: >
      The second step is to add some code to your JavaScript that subscribes to
      the `/esbuild` [server-sent event](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events)
      source. When you get the `change` event, you can reload the page to get
      the latest version of the app. You can do this in a single line of code:

  - pre.js: |
      new EventSource('/esbuild').addEventListener('change', () => location.reload())

  - p: >
      That's it! If you load your app in the browser, the page should now
      automatically reload when you edit and save a file (assuming there are
      no build errors).

  - p: >
      This should only be included during development, and should not be included
      in production. One way to remove this code in production is to guard it with
      an if statement such as `if (!window.IS_PRODUCTION)` and then use [define](#define)
      to set `window.IS_PRODUCTION` to `true` in production.

  - h4#live-reload-caveats: Live reload caveats

  - p: >
      Implementing live reloading like this has a few known caveats:

  - ul:
    - >
      <p>
      These events only trigger when esbuild's output changes. They do not trigger
      when files unrelated to the build being watched are changed. If your HTML
      file references other files that esbuild doesn't know about and those files
      are changed, you can either manually reload the page or you can implement
      your own live reloading infrastructure instead of using esbuild's built-in
      behavior.
      </p>

    - >
      <p>
      The `EventSource` API is supposed to automatically reconnect for you. However,
      there's [a bug in Firefox](https://bugzilla.mozilla.org/show_bug.cgi?id=1809332)
      that breaks this if the server is ever temporarily unreachable. Workarounds
      are to use any other browser, to manually reload the page if this happens, or
      to write more complicated code that manually closes and re-creates the
      `EventSource` object if there is a connection error.
      </p>

    - >
      <p>
      Browser vendors have decided to not implement HTTP/2 without TLS. This
      means that when using the `http://` protocol, each `/esbuild` event source
      will take up one of your precious 6 simultaneous per-domain HTTP/1.1
      connections. So if you open more than six HTTP tabs that use this
      live-reloading technique, you will be unable to use live reloading in
      some of those tabs (and other things will likely also break). The
      workaround is to [enable the `https://` protocol](#https).
      </p>


  - h4#hot-reloading-css: Hot-reloading for CSS

  - p: >
      The `change` event also contains additional information to enable more
      advanced use cases. It currently contains the `added`, `removed`, and
      `updated` arrays with the paths of the files that have changed since the
      previous build, which can be described by the following TypeScript
      interface:

  - pre.js: |
      interface ChangeEvent {
        added: string[]
        removed: string[]
        updated: string[]
      }

  - p: >
      The code sample below enables "hot reloading" for CSS, which is when the
      CSS is automatically updated in place without reloading the page. If an
      event arrives that isn't CSS-related, then the whole page will be
      reloaded as a fallback:

  - pre.js: |
      new EventSource('/esbuild').addEventListener('change', e => {
        const { added, removed, updated } = JSON.parse(e.data)

        if (!added.length && !removed.length && updated.length === 1) {
          for (const link of document.getElementsByTagName("link")) {
            const url = new URL(link.href)

            if (url.host === location.host && url.pathname === updated[0]) {
              const next = link.cloneNode()
              next.href = updated[0] + '?' + Math.random().toString(36).slice(2)
              next.onload = () => link.remove()
              link.parentNode.insertBefore(next, link.nextSibling)
              return
            }
          }
        }

        location.reload()
      })

  - h4#hot-reloading-js: Hot-reloading for JavaScript

  - p: >
      Hot-reloading for JavaScript is not currently implemented by esbuild.
      It's possible to transparently implement hot-reloading for CSS because
      CSS is stateless, but JavaScript is stateful so you cannot transparently
      implement hot-reloading for JavaScript like you can for CSS.

  - p: >
      Some other development servers implement hot-reloading for JavaScript
      anyway, but it requires additional APIs, sometimes requires framework-specific
      hacks, and sometimes introduces transient state-related bugs during an
      editing session. Doing this is outside of esbuild's scope. You are
      welcome to use other tools instead of esbuild if hot-reloading for
      JavaScript is one of your requirements.

  - p: >
      However, with esbuild's live-reloading you can persist your app's current
      JavaScript state in [`sessionStorage`](https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage)
      to more easily restore your app's JavaScript state after a page reload.
      If your app loads quickly (which it already should for your users' sake),
      live-reloading with JavaScript can be almost as fast as hot-reloading
      with JavaScript would be.

  ## Platform

  - p: >
      By default, esbuild's bundler is configured to generate code intended for
      the browser. If your bundled code is intended to run in node instead, you
      should set the platform to `node`:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --bundle --platform=node

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          platform: 'node',
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Platform:    api.PlatformNode,
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      When the platform is set to `browser` (the default value):

  - ul:
    - >
      <p>
      When [bundling](#bundle) is enabled the default output [format](#format)
      is set to `iife`, which wraps the generated JavaScript code in an
      immediately-invoked function expression to prevent variables from leaking
      into the global scope.
      </p>
    - >
      <p>
      If a package specifies a map for the
      [`browser`](https://gist.github.com/defunctzombie/4339901/49493836fb873ddaa4b8a7aa0ef2352119f69211)
      field in its `package.json` file, esbuild will use that map to replace
      specific files or modules with their browser-friendly versions. For
      example, a package might contain a substitution of [`path`](https://nodejs.org/api/path.html)
      with [`path-browserify`](https://www.npmjs.com/package/path-browserify).
      </p>
    - >
      <p>
      The [main fields](#main-fields) setting is set to <code>browser,<wbr>module,<wbr>main</code>
      but with some additional special behavior: if a package provides `module`
      and `main` entry points but not a `browser` entry point then `main` is
      used instead of `module` if that package is ever imported using `require()`.
      This behavior improves compatibility with CommonJS modules that export a
      function by assigning it to `module.exports`. If you want to disable this
      additional special behavior, you can explicitly set the [main fields](#main-fields)
      setting to <code>browser,<wbr>module,<wbr>main</code>.
      </p>
    - >
      <p>
      The [conditions](#conditions) setting automatically includes the `browser`
      condition. This changes how the `exports` field in `package.json` files
      is interpreted to prefer browser-specific code.
      </p>
    - >
      <p>
      If no custom [conditions](#conditions) are configured, the Webpack-specific
      `module` condition is also included. The `module` condition is used by
      package authors to provide a tree-shakable ESM alternative to a CommonJS
      file without creating a [dual package hazard](https://nodejs.org/api/packages.html#dual-package-hazard).
      You can prevent the `module` condition from being included by explicitly
      configuring some custom conditions (even an empty list).
      </p>
    - >
      <p>
      When using the [build](#build) API, all <code>process.<wbr>env.<wbr>NODE_ENV</code>
      expressions are automatically [defined](#define) to `"production"` if all
      [minification](#minify) options are enabled and `"development"` otherwise.
      This only happens if `process`, `process.env`, and `process.env.NODE_ENV`
      are not already defined. This substitution is necessary to avoid React-based
      code crashing instantly (since `process` is a node API, not a web API).
      </p>
    - >
      <p>
      The character sequence `</script>` will be escaped in JavaScript code and
      the character sequence `</style>` will be escaped in CSS code. This is
      done in case you inline esbuild's output directly into an HTML file. This
      can be disabled with esbuild's [supported](#supported) feature by setting
      `inline-script` (for JavaScript) and/or `inline-style` (for CSS) to `false`.
      </p>

  - p: >
      When the platform is set to `node`:

  - ul:
    - >
      <p>
      When [bundling](#bundle) is enabled the default output [format](#format)
      is set to `cjs`, which stands for  CommonJS (the module format used by node).
      ES6-style exports using `export` statements will be converted into getters
      on the CommonJS `exports` object.
      </p>
    - >
      <p>
      All [built-in node modules](https://nodejs.org/docs/latest/api/) such
      as `fs` are automatically marked as [external](#external) so they don't
      cause errors when the bundler tries to bundle them.
      </p>
    - >
      <p>
      The [main fields](#main-fields) setting is set to <code>main,<wbr>module</code>. This
      means tree shaking will likely not happen for packages that provide both
      `module` and `main` since tree shaking works with ECMAScript modules but
      not with CommonJS modules.
      </p>
      <p>
      Unfortunately some packages incorrectly treat `module` as meaning "browser
      code" instead of "ECMAScript module code" so this default behavior is
      required for compatibility. You can manually configure the [main fields](#main-fields)
      setting to <code>module,<wbr>main</code> if you want to enable tree shaking and know it
      is safe to do so.
      </p>
    - >
      <p>
      The [conditions](#conditions) setting automatically includes the `node`
      condition. This changes how the `exports` field in `package.json` files
      is interpreted to prefer node-specific code.
      </p>
    - >
      <p>
      If no custom [conditions](#conditions) are configured, the Webpack-specific
      `module` condition is also included. The `module` condition is used by
      package authors to provide a tree-shakable ESM alternative to a CommonJS
      file without creating a [dual package hazard](https://nodejs.org/api/packages.html#dual-package-hazard).
      You can prevent the `module` condition from being included by explicitly
      configuring some custom conditions (even an empty list).
      </p>
    - >
      <p>
      When the [format](#format) is set to `cjs` but the entry point is ESM,
      esbuild will add special annotations for any named exports to enable
      importing those named exports using ESM syntax from the resulting
      CommonJS file. Node's documentation has more information about
      [node's detection of CommonJS named exports](https://nodejs.org/api/esm.html#commonjs-namespaces).
      </p>
    - >
      <p>
      The [`binary`](/content-types/#binary) loader will make use of node's
      built-in [`Buffer.from`](https://nodejs.org/api/buffer.html#static-method-bufferfromstring-encoding)
      API to decode the base64 data embedded in the bundle into a `Uint8Array`.
      This is faster than what esbuild can do otherwise since it's implemented
      by node in native code.
      </p>

  - p: >
      When the platform is set to `neutral`:

  - ul:
    - >
      <p>
      When [bundling](#bundle) is enabled the default output [format](#format)
      is set to `esm`, which uses the `export` syntax introduced with ECMAScript
      2015 (i.e. ES6). You can change the output format if this default is not
      appropriate.
      </p>
    - >
      <p>
      The [main fields](#main-fields) setting is empty by default. If you want
      to use npm-style packages, you will likely have to configure this to be
      something else such as `main` for the standard main field used by node.
      </p>
    - >
      <p>
      The [conditions](#conditions) setting does not automatically include any
      platform-specific values.
      </p>

  - p: >
      See also [bundling for the browser](/getting-started/#bundling-for-the-browser)
      and [bundling for node](/getting-started/#bundling-for-node).

  ## Rebuild

  - p: >
      You may want to use this API if your use case involves calling esbuild's
      [build](#build) API repeatedly with the same options. For example,
      this is useful if you are implementing your own file watcher service.
      Rebuilding is more efficient than building again because some of the data
      from the previous build is cached and can be reused if the original files
      haven't changed since the previous build. There are currently two forms
      of caching used by the rebuild API:

  - ul:
    - >
      <p>
      Files are stored in memory and are not re-read from the file system if
      the file metadata hasn't changed since the last build. This optimization
      only applies to file system paths. It does not apply to virtual modules
      created by [plugins](/plugins/).
      </p>

    - >
      <p>
      Parsed [ASTs](https://en.wikipedia.org/wiki/Abstract_syntax_tree) are
      stored in memory and re-parsing the AST is avoided if the file contents
      haven't changed since the last build. This optimization applies to
      virtual modules created by plugins in addition to file system modules,
      as long as the virtual module path remains the same.
      </p>

  - p: >
      Here's how to do a rebuild:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        # The CLI does not have an API for "rebuild"

      mjs: |
        import * as esbuild from 'esbuild'

        let ctx = await esbuild.context({
          entryPoints: ['app.js'],
          bundle: true,
          outfile: 'out.js',
        })

        // Call "rebuild" as many times as you want
        for (let i = 0; i < 5; i++) {
          let result = await ctx.rebuild()
        }

        // Call "dispose" when you're done to free up resources
        ctx.dispose()

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          ctx, err := api.Context(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Outfile:     "out.js",
          })
          if err != nil {
            os.Exit(1)
          }

          // Call "Rebuild" as many times as you want
          for i := 0; i < 5; i++ {
            result := ctx.Rebuild()
            if len(result.Errors) > 0 {
              os.Exit(1)
            }
          }

          // Call "Dispose" when you're done to free up resources
          ctx.Dispose()
        }

  ## Serve

  - info: >
      If you want your app to automatically reload as you edit, you should
      read about [live reloading](#live-reload). It combines serve mode with
      [watch mode](#watch) to listen for changes to the file system.

  - p: >
      Serve mode starts a web server that serves your code to your browser on
      your device. Here's an example that bundles `src/app.ts` into `www/js/app.js`
      and then also serves the `www` directory over `http://localhost:8000/`:

  - example:
      noCheck: true

      cli: |
        esbuild src/app.ts --outdir=www/js --bundle --servedir=www

      mjs: |
        import * as esbuild from 'esbuild'

        let ctx = await esbuild.context({
          entryPoints: ['src/app.ts'],
          outdir: 'www/js',
          bundle: true,
        })

        let { host, port } = await ctx.serve({
          servedir: 'www',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          ctx, err := api.Context(api.BuildOptions{
            EntryPoints: []string{"src/app.ts"},
            Outdir:     "www/js",
            Bundle:      true,
          })
          if err != nil {
            os.Exit(1)
          }

          server, err2 := ctx.Serve(api.ServeOptions{
            Servedir: "www",
          })
          if err2 != nil {
            os.Exit(1)
          }

          // Returning from main() exits immediately in Go.
          // Block forever so we keep serving and don't exit.
          <-make(chan struct{})
        }

  - p: >
      If you create the file `www/index.html` with the following contents, the
      code contained in `src/app.ts` will load when you navigate to `http://localhost:8000/`:

  - pre.html: |
      <script src="js/app.js"></script>

  - p: >
      One benefit of using esbuild's built-in web server instead of another web
      server is that whenever you reload, the files that esbuild serves are
      always up to date. That's not necessarily the case with other development
      setups. One common setup is to run a local file watcher that rebuilds
      output files whenever their input files change, and then separately to run
      a local file server to serve those output files. But that means reloading
      after an edit may reload the old output files if the rebuild hasn't finished
      yet. With esbuild's web server, each incoming request starts a rebuild if
      one is not already in progress, and then waits for the current rebuild to
      complete before serving the file. This means esbuild never serves stale
      build results.

  - p: >
      Note that this web server is intended to only be used in development. _Do
      not use this in production._

  - h4#serve-arguments: Arguments

  - p: >
      The arguments to the serve API are as follows:

  - example:
      noCheck: true

      cli: |
        # Enable serve mode
        --serve

        # Set the port
        --serve=9000

        # Set the host and port (IPv4)
        --serve=127.0.0.1:9000

        # Set the host and port (IPv6)
        --serve=[::1]:9000

        # Set the directory to serve
        --servedir=www

        # Enable HTTPS
        --keyfile=your.key --certfile=your.cert

        # Specify a fallback HTML file
        --serve-fallback=some-file.html

      js: |
        interface ServeOptions {
          port?: number
          host?: string
          servedir?: string
          keyfile?: string
          certfile?: string
          fallback?: string
          onRequest?: (args: ServeOnRequestArgs) => void
        }

        interface ServeOnRequestArgs {
          remoteAddress: string
          method: string
          path: string
          status: number
          timeInMS: number
        }

      go: |
        type ServeOptions struct {
          Port      uint16
          Host      string
          Servedir  string
          Keyfile   string
          Certfile  string
          Fallback  string
          OnRequest func(ServeOnRequestArgs)
        }

        type ServeOnRequestArgs struct {
          RemoteAddress string
          Method        string
          Path          string
          Status        int
          TimeInMS      int
        }

  - ul:
    - >
      `host`
      <p>
      By default, esbuild makes the web server available on all IPv4 network
      interfaces. This corresponds to a host address of `0.0.0.0`. If you would
      like to configure a different host (for example, to only serve on the
      `127.0.0.1` loopback interface without exposing anything to the network),
      you can specify the host using this argument.
      </p>
      <p>
      If you need to use IPv6 instead of IPv4, you just need to specify an IPv6
      host address. The equivalent to the `127.0.0.1` loopback interface in IPv6
      is `::1` and the equivalent to the `0.0.0.0` universal interface in IPv6 is
      `::`.
      </p>

    - >
      `port`
      <p>
      The HTTP port can optionally be configured here. If omitted, it will
      default to an open port with a preference for ports in the range 8000 to
      8009.
      </p>

    - >
      `servedir`
      <p>
      This is a directory of extra content for esbuild's HTTP server to serve
      instead of a 404 when incoming requests don't match any of the generated
      output file paths. This lets you use esbuild as a general-purpose local
      web server.
      </p>
      <p>
      For example, you might want to create an `index.html` file and then set
      `servedir` to `"."` to serve the current directory (which includes the
      `index.html` file). If you don't set `servedir` then esbuild will only
      serve the build results, but not any other files.
      </p>

    - >
      `keyfile` and `certfile`
      <p>
      If you pass a private key and certificate to esbuild using `keyfile` and
      `certfile`, then esbuild's web server will use the `https://` protocol
      instead of the `http://` protocol. See [enabling HTTPS](#https) for more
      information.
      </p>

    - >
      `fallback`
      <p>
      This is a HTML file for esbuild's HTTP server to serve instead of a 404
      when incoming requests don't match any of the generated output file
      paths. You can use this for a custom "not found" page. You can also use
      this as the entry point of a [single-page application](https://en.wikipedia.org/wiki/Single-page_application)
      that mutates the current URL and therefore needs to be served from many
      different URLs simultaneously.
      </p>

    - >
      `onRequest`
      <p>
      This is called once for each incoming request with some information about
      the request. This callback is used by the CLI to print out a log message
      for each request. The time field is the time to generate the data for the
      request, but it does not include the time to stream the request to the
      client.
      </p>
      <p>
      Note that this is called after the request has completed. It's not possible
      to use this callback to modify the request in any way. If you want to do
      this, you should [put a proxy in front of esbuild](#serve-proxy) instead.
      </p>

  - h4#serve-return-values: Return values

  - example:
      noCheck: true

      cli: |
        # The CLI will print the host and port like this:

         > Local: http://127.0.0.1:8000/

      js: |
        interface ServeResult {
          host: string
          port: number
        }

      go: |
        type ServeResult struct {
          Host string
          Port uint16
        }

  - ul:
    - >
      `host`
      <p>
      This is the host that ended up being used by the web server. It will be
      `0.0.0.0` (i.e. serving on all available network interfaces) unless a
      custom host was configured. If you are using the CLI and the host is
      `0.0.0.0`, all available network interfaces will be printed as hosts
      instead.
      </p>

    - >
      `port`
      <p>
      This is the port that ended up being used by the web server. You'll want to
      use this if you don't specify a port since esbuild will end up picking
      an arbitrary open port, and you need to know which port it picked to be
      able to connect to it.
      </p>

  - h4#https: Enabling HTTPS

  - p: >
      By default, esbuild's web server uses the `http://` protocol. However,
      certain modern web features are unavailable to HTTP websites. If you want
      to use these features, then you'll need to tell esbuild to use the
      `https://` protocol instead.

  - p: >
      To enable HTTPS with esbuild:

  - ol:
    - >
      <p>
      Generate a self-signed certificate. There are many ways to do this. Here's
      one way, assuming you have the `openssl` command installed:
      </p>
      <pre>
      openssl req -x509 -newkey rsa:4096 -keyout your.key -out your.cert -days 9999 -nodes -subj /CN=127.0.0.1
      </pre>

    - >
      <p>
      Pass `your.key` and `your.cert` to esbuild using the `keyfile` and `certfile`
      [serve arguments](#serve-arguments).
      </p>

    - >
      <p>
      Click past the scary warning in your browser when you load your page
      (self-signed certificates aren't secure, but that doesn't matter since
      we're just doing local development).
      </p>

  - p: >
      If you have more complex needs than this, you can still [put a proxy in front of esbuild](#serve-proxy)
      and use that for HTTPS instead. Note that if you see the message
      <code>Client <wbr>sent <wbr>an <wbr>HTTP <wbr>request <wbr>to <wbr>an <wbr>HTTPS <wbr>server</code>
      when you load your page, then you are using the incorrect protocol.
      Replace `http://` with `https://` in your browser's URL bar.

  - p: >
      Keep in mind that esbuild's HTTPS support has nothing to do with security.
      The only reason to enable HTTPS in esbuild is because browsers have made
      it impossible to do local development with certain modern web features
      without jumping through these extra hoops. *Please do not use esbuild's
      development server for anything that needs to be secure.* It's only
      intended for local development and no considerations have been made for
      production environments whatsoever.

  - h4#serve-proxy: Customizing server behavior

  - p: >
      It's not possible to hook into esbuild's local server to customize the
      behavior of the server itself. Instead, behavior should be customized
      by putting a proxy in front of esbuild.

  - p: >
      Here's a simple example of a proxy server to get you started, using node's
      built-in [`http`](https://nodejs.org/api/http.html) module. It adds a
      custom 404 page instead of esbuild's default 404 page:

  - pre.js: |
      import * as esbuild from 'esbuild'
      import http from 'node:http'

      // Start esbuild's server on a random local port
      let ctx = await esbuild.context({
        // ... your build options go here ...
      })

      // The return value tells us where esbuild's local server is
      let { host, port } = await ctx.serve({ servedir: '.' })

      // Then start a proxy server on port 3000
      http.createServer((req, res) => {
        const options = {
          hostname: host,
          port: port,
          path: req.url,
          method: req.method,
          headers: req.headers,
        }

        // Forward each incoming request to esbuild
        const proxyReq = http.request(options, proxyRes => {
          // If esbuild returns "not found", send a custom 404 page
          if (proxyRes.statusCode === 404) {
            res.writeHead(404, { 'Content-Type': 'text/html' })
            res.end('<h1>A custom 404 page</h1>')
            return
          }

          // Otherwise, forward the response from esbuild to the client
          res.writeHead(proxyRes.statusCode, proxyRes.headers)
          proxyRes.pipe(res, { end: true })
        })

        // Forward the body of the request to esbuild
        req.pipe(proxyReq, { end: true })
      }).listen(3000)

  - p: >
      This code starts esbuild's server on random local port and then starts a
      proxy server on port 3000. During development you would load [http://localhost:3000](http://localhost:3000)
      in your browser, which talks to the proxy. This example demonstrates
      modifying a response after esbuild has handled the request, but you can
      also modify or replace the request before esbuild has handled it.

  - p: >
      You can do many things with a proxy like this including:

  - ul:
    - Injecting your own 404 page (the example above)
    - Customizing the mapping of routes to files on the file system
    - Redirecting some routes to an API server instead of to esbuild

  - p: >
      You can also use a real proxy such as [nginx](https://nginx.org/en/docs/beginners_guide.html#proxy)
      if you have more advanced needs.

  ## Tsconfig

  - p: >
      Normally the [build](#build) API automatically discovers `tsconfig.json`
      files and reads their contents during a build. However, you can also
      configure a custom `tsconfig.json` file to use instead. This can be
      useful if you need to do multiple builds of the same code with different
      settings:

  - example:
      in:
        app.ts: '1 + 2'
        custom-tsconfig.json: '{}'

      cli: |
          esbuild app.ts --bundle --tsconfig=custom-tsconfig.json

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.ts'],
          bundle: true,
          tsconfig: 'custom-tsconfig.json',
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.ts"},
            Bundle:      true,
            Tsconfig:    "custom-tsconfig.json",
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## Tsconfig raw

  - p: >
      This option can be used to pass your `tsconfig.json` file to the
      [transform](#transform) API, which doesn't access the file system.
      It can also be used to pass the contents of your `tsconfig.json`
      file to the [build](#build) API inline without writing it to a
      file. Using it looks like this:

  - example:
      cli: |
          echo 'class Foo { foo }' | esbuild --loader=ts --tsconfig-raw='{"compilerOptions":{"useDefineForClassFields":false}}'

      mjs: |
        import * as esbuild from 'esbuild'

        let ts = 'class Foo { foo }'
        let result = await esbuild.transform(ts, {
          loader: 'ts',
          tsconfigRaw: `{
            "compilerOptions": {
              "useDefineForClassFields": false,
            },
          }`,
        })
        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          ts := "class Foo { foo }"

          result := api.Transform(ts, api.TransformOptions{
            Loader: api.LoaderTS,
            TsconfigRaw: `{
              "compilerOptions": {
                "useDefineForClassFields": false,
              },
            }`,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  ## Watch

  - p: >
      Enabling watch mode tells esbuild to listen for changes on the file system
      and to automatically rebuild whenever a file changes that could invalidate
      the build. Using it looks like this:

  - example:
      noCheck: true

      cli:
        - $: |
            esbuild app.js --outfile=out.js --bundle --watch
        - expect: |
            [watch] build finished, watching for changes...

      mjs: |
        import * as esbuild from 'esbuild'

        let ctx = await esbuild.context({
          entryPoints: ['app.js'],
          outfile: 'out.js',
          bundle: true,
        })

        await ctx.watch()
        console.log('watching...')

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          ctx, err := api.Context(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Outfile:     "out.js",
            Bundle:      true,
            Write:       true,
          })
          if err != nil {
            os.Exit(1)
          }

          err2 := ctx.Watch(api.WatchOptions{})
          if err2 != nil {
            os.Exit(1)
          }
          fmt.Printf("watching...\n")

          // Returning from main() exits immediately in Go.
          // Block forever so we keep watching and don't exit.
          <-make(chan struct{})
        }

  - p: >
      If you want to stop watch mode at some point in the future, you can call
      `dispose` on the context object to terminate the file watcher:

  - example:
      noCheck: true

      cli: |
        # Use Ctrl+C to stop the CLI in watch mode

      mjs: |
        import * as esbuild from 'esbuild'

        let ctx = await esbuild.context({
          entryPoints: ['app.js'],
          outfile: 'out.js',
          bundle: true,
        })

        await ctx.watch()
        console.log('watching...')

        await new Promise(r => setTimeout(r, 10 * 1000))
        await ctx.dispose()
        console.log('stopped watching')

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"
        import "os"
        import "time"

        func main() {
          ctx, err := api.Context(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Outfile:     "out.js",
            Bundle:      true,
            Write:       true,
          })
          if err != nil {
            os.Exit(1)
          }

          err2 := ctx.Watch(api.WatchOptions{})
          if err2 != nil {
            os.Exit(1)
          }
          fmt.Printf("watching...\n")

          time.Sleep(10 * time.Second)
          ctx.Dispose()
          fmt.Printf("stopped watching\n")
        }

  - p: >
      Watch mode in esbuild is implemented using polling instead of OS-specific
      file system APIs for portability. The polling system is designed to use
      relatively little CPU vs. a more traditional polling system that scans
      the whole directory tree at once. The file system is still scanned
      regularly but each scan only checks a random subset of your files, which
      means a change to a file will be picked up soon after the change is made
      but not necessarily instantly.

  - p: >
      With the current heuristics, large projects should be completely scanned
      around every 2 seconds so in the worst case it could take up to 2 seconds
      for a change to be noticed. However, after a change has been noticed the
      change's path goes on a short list of recently changed paths which are
      checked on every scan, so further changes to recently changed files
      should be noticed almost instantly.

  - p: >
      Note that it is still possible to implement watch mode yourself using
      esbuild's [rebuild](#rebuild) API and a file watcher library of your
      choice if you don't want to use a polling-based approach.

  - p: >
      If you are using the CLI, keep in mind that watch mode will be terminated
      when esbuild's stdin is closed. This prevents esbuild from accidentally
      outliving the parent process and unexpectedly continuing to consume
      resources on the system. If you have a use case that requires esbuild to
      continue to watch forever even when the parent process has finished, you
      may use <code>--watch=<wbr>forever</code> instead of `--watch`.

# Input

## Entry points

* == `[arrayOfFiles]` / 
  * each one == input to the bundling algorithm
  * `[arrayOfFilePaths]`
    * SIMPLEST way
* Reason of the "entry points" naming: 🧠EACH one == initial script / evaluate which AFTERWARD loads ALL other code's aspects / it represents 🧠
* if you want to use -> use `import` 
  * ALTERNATIVE to `<script>`

* use cases
  * | SIMPLE apps, ONLY need 1! entry point
  * MULTIPLE entry points, |  
    * MULTIPLE logically-independent groups of code / main thread + worker thread, or
    * app / separate relatively unrelated areas (_Example:_ landing page, editor page, and a settings page)

* Separate entry points
  * allows
    * separation of concerns
    * reduce the amount of unnecessary code / browser needs to download
  * if applicable -> enable [code splitting](#splitting)
    * Reason: 🧠reduce further download sizes | browse to a second page /
      * entry point -- shares some ALREADY-downloaded code with a -- first page ALREADY visited 🧠

* ways to specify
  * via CLI
    ```
    esbuild src/content/examples/input/entryPoints/home.ts src/content/examples/input/entryPoints/settings.ts --bundle --outdir=out
    ```
  * | JS
    ```
    node src/content/examples/input/entryPoints/module.mjs
    ```
  * | go  
    * Problems: How to run?
      * Attempt1: `go src/content/examples/input/entryPoints/main.go`

* TODO:

- p: >
    This will generate two output files, `out/home.js` and `out/settings.js`
    corresponding to the two entry points `home.ts` and `settings.ts`.

- p: >
    For further control over how the paths of the output files are derived
    from the corresponding input entry points, you should look into these
    options:

- ul:
  - '[Entry names](#entry-names)'
  - '[Out extension](#out-extension)'
  - '[Outbase](#outbase)'
  - '[Outdir](#outdir)'
  - '[Outfile](#outfile)'

- p: >
    In addition, you can also specify a fully custom output path for each
    individual entry point using an alternative entry point syntax:

- example:
    in:
      home.ts: '1 + 2'
      settings.ts: '1 + 2'

    cli: |
      esbuild out1=home.ts out2=settings.ts --bundle --outdir=out

    mjs: |
      import * as esbuild from 'esbuild'

      await esbuild.build({
        entryPoints: [
          { out: 'out1', in: 'home.ts'},
          { out: 'out2', in: 'settings.ts'},
        ],
        bundle: true,
        write: true,
        outdir: 'out',
      })

    go: |
      package main

      import "github.com/evanw/esbuild/pkg/api"
      import "os"

      func main() {
        result := api.Build(api.BuildOptions{
          EntryPointsAdvanced: []api.EntryPoint{{
            OutputPath: "out1",
            InputPath:  "home.ts",
          }, {
            OutputPath: "out2",
            InputPath:  "settings.ts",
          }},
          Bundle: true,
          Write:  true,
          Outdir: "out",
        })

        if len(result.Errors) > 0 {
          os.Exit(1)
        }
      }

- p: >
    This will generate two output files, `out/out1.js` and `out/out2.js`
    corresponding to the two entry points `home.ts` and `settings.ts`.

### Glob-style entry points

- p: >
    If an entry point contains the `*` character, then it's considered to be
    a [glob](https://en.wikipedia.org/wiki/Glob_(programming)) pattern. This
    means esbuild will use that entry point as a pattern to search for files
    on the file system and will then replace that entry point with any
    matching files that were found. So for example, an entry point of `*.js`
    will cause esbuild to consider all files in the current directory that
    end in `.js` to be entry points.

- p: >
    The glob matcher that esbuild implements is intentionally simple, and does
    not support more advanced features found in certain other glob libraries.
    Only two kinds of wildcards are supported:

- ul:
  - >
    `*`
    <p>
    This wildcard matches any number of characters (including none) except
    that it does not match a slash character (i.e. a `/`), which means it
    does not cause esbuild to traverse into subdirectories. For example,
    `*.js` will match `foo.js` but not `bar/foo.js`.
    </p>

  - >
    `/**/`
    <p>
    This wildcard matches zero or more path segments, which means it can be
    used to tell esbuild to match against a whole directory tree. For example,
    `./**/*.js` will match `./foo.js` and `./bar/foo.js` and `./a/b/c/foo.js`.
    </p>

- p: >
    If you are using esbuild via the CLI, keep in mind that if you do not
    quote arguments that contain shell metacharacters before you pass them
    to esbuild, your shell will likely expand them before esbuild sees them.
    So if you run `esbuild "*.js"` (with quotes) then esbuild will see an
    entry point of `*.js` and the glob-style entry point rules described
    above will apply. But if you run `esbuild *.js` (without quotes)
    then esbuild will see whatever your current shell decided to expand
    `*.js` into (which may include seeing nothing at all if your shell
    expanded it into nothing). Using esbuild's built-in glob pattern support
    can be a convenient way to ensure cross-platform consistency by avoiding
    shell-specific behavior, but it requires you to quote your arguments
    correctly so that your shell doesn't interpret them.

  ## Loader

  - p: >
      This option changes how a given input file is interpreted. For example,
      the [`js`](/content-types/#javascript) loader interprets the file as
      JavaScript and the [`css`](/content-types/#css) loader interprets the
      file as CSS. See the [content types](/content-types/) page for a
      complete list of all built-in loaders.

  - p: >
      Configuring a loader for a given file type lets you load that file type
      with an `import` statement or a `require` call. For example, configuring
      the `.png` file extension to use the [data URL](/content-types/#data-url)
      loader means importing a `.png` file gives you a data URL containing the
      contents of that image:

  - pre.js: |
      import url from './example.png'
      let image = new Image
      image.src = url
      document.body.appendChild(image)

      import svg from './example.svg'
      let doc = new DOMParser().parseFromString(svg, 'application/xml')
      let node = document.importNode(doc.documentElement, true)
      document.body.appendChild(node)

  - p: >
      The above code can be bundled using the [build](#build) API call like
      this:

  - example:
      in:
        app.js: |
          import url from './example.png'
          let image = new Image
          image.src = url
          document.body.appendChild(image)

          import svg from './example.svg'
          let doc = new DOMParser().parseFromString(svg, 'application/xml')
          let node = document.importNode(doc.documentElement, true)
          document.body.appendChild(node)

        example.png: |
          this is some data

        example.svg: |
          this is some data

      cli: |
        esbuild app.js --bundle --loader:.png=dataurl --loader:.svg=text

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          loader: {
            '.png': 'dataurl',
            '.svg': 'text',
          },
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Loader: map[string]api.Loader{
              ".png": api.LoaderDataURL,
              ".svg": api.LoaderText,
            },
            Write: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      This option is specified differently if you are using the build API with
      input from [stdin](#stdin), since stdin does not have a file extension.
      Configuring a loader for stdin with the build API looks like this:

  - example:
      in:
        pkg.js: |
          module.exports = 123

      cli: |
        echo 'import pkg = require("./pkg")' | esbuild --loader=ts --bundle

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          stdin: {
            contents: 'import pkg = require("./pkg")',
            loader: 'ts',
            resolveDir: '.',
          },
          bundle: true,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            Stdin: &api.StdinOptions{
              Contents:   "import pkg = require('./pkg')",
              Loader:     api.LoaderTS,
              ResolveDir: ".",
            },
            Bundle: true,
          })
          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      The [transform](#transform) API call just takes a single loader since
      it doesn't involve interacting with the file system, and therefore doesn't
      deal with file extensions. Configuring a loader (in this case the
      [`ts`](/content-types/#typescript) loader) for the transform API looks
      like this:

  - example:
      cli:
        - $: |
            echo 'let x: number = 1' | esbuild --loader=ts
        - expect: |
            let x = 1;

      mjs: |
        import * as esbuild from 'esbuild'

        let ts = 'let x: number = 1'
        let result = await esbuild.transform(ts, {
          loader: 'ts',
        })
        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          ts := "let x: number = 1"
          result := api.Transform(ts, api.TransformOptions{
            Loader: api.LoaderTS,
          })
          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  ## Stdin

  - p: >
      Normally the build API call takes one or more file names as input.
      However, this option can be used to run a build without a module existing
      on the file system at all. It's called "stdin" because it corresponds to
      piping a file to stdin on the command line.

  - p: >
      In addition to specifying the contents of the stdin file, you can
      optionally also specify the resolve directory (used to determine where
      relative imports are located), the [sourcefile](#sourcefile) (the file
      name to use in error messages and source maps), and the [loader](#loader)
      (which determines how the file contents are interpreted). The CLI doesn't
      have a way to specify the resolve directory. Instead, it's automatically
      set to the current working directory.

  - p: >
      Here's how to use this feature:

  - example:
      in:
        another-file.js: 'export let foo = 123'

      cli: |
        echo 'export * from "./another-file"' | esbuild --bundle --sourcefile=imaginary-file.js --loader=ts --format=cjs

      mjs: |
        import * as esbuild from 'esbuild'

        let result = await esbuild.build({
          stdin: {
            contents: `export * from "./another-file"`,

            // These are all optional:
            resolveDir: './src',
            sourcefile: 'imaginary-file.js',
            loader: 'ts',
          },
          format: 'cjs',
          write: false,
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            Stdin: &api.StdinOptions{
              Contents: "export * from './another-file'",

              // These are all optional:
              ResolveDir: "./src",
              Sourcefile: "imaginary-file.js",
              Loader:     api.LoaderTS,
            },
            Format: api.FormatCommonJS,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  # Output contents

  ## Banner

  - p: >
      Use this to insert an arbitrary string at the beginning of generated
      JavaScript and CSS files. This is commonly used to insert comments:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --banner:js=//comment --banner:css=/*comment*/

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          banner: {
            js: '//comment',
            css: '/*comment*/',
          },
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Banner: map[string]string{
              "js":  "//comment",
              "css": "/*comment*/",
            },
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      This is similar to [footer](#footer) which inserts at the end instead
      of the beginning.

  - p: >
      Note that if you are inserting non-comment code into a CSS file, be aware
      that CSS ignores all `@import` rules that come after a non-`@import` rule
      (other than a `@charset` rule), so using a banner to inject CSS rules may
      accidentally disable imports of external stylesheets.

  ## Charset

  - p: >
      By default esbuild's output is ASCII-only. Any non-ASCII characters are
      escaped using backslash escape sequences. One reason is because non-ASCII
      characters are misinterpreted by the browser by default, which causes
      confusion. You have to explicitly add <code>&lt;meta <wbr>charset=<wbr>"utf-8"&gt;</code> to your
      HTML or serve it with the correct <code>Content-<wbr>Type</code> header for the browser
      to not mangle your code. Another reason is that non-ASCII characters can
      significantly [slow down the browser's parser](https://v8.dev/blog/scanner).
      However, using escape sequences makes the generated output slightly bigger,
      and also makes it harder to read.

  - p: >
      If you would like for esbuild to print the original characters without
      using escape sequences and you have ensured that the browser will
      interpret your code as UTF-8, you can disable character escaping by
      setting the charset:

  - example:
      cli:
        - $: |
            echo 'let π = Math.PI' | esbuild
        - expect: |
            let \u03C0 = Math.PI;
        - $: |
            echo 'let π = Math.PI' | esbuild --charset=utf8
        - expect: |
            let π = Math.PI;

      mjs:
        - $: |
            import * as esbuild from 'esbuild'
        - $: |
            let js = 'let π = Math.PI'
        - $: |
            (await esbuild.transform(js)).code
        - expect: |
            'let \\u03C0 = Math.PI;\n'
        - $: |
            (await esbuild.transform(js, {
              charset: 'utf8',
            })).code
        - expect: |
            'let π = Math.PI;\n'

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "let π = Math.PI"

          result1 := api.Transform(js, api.TransformOptions{})

          if len(result1.Errors) == 0 {
            fmt.Printf("%s", result1.Code)
          }

          result2 := api.Transform(js, api.TransformOptions{
            Charset: api.CharsetUTF8,
          })

          if len(result2.Errors) == 0 {
            fmt.Printf("%s", result2.Code)
          }
        }

  - p: >
      Some caveats:

  - ul:
    - >
      <p>
      This does not yet escape non-ASCII characters embedded in regular
      expressions. This is because esbuild does not currently parse the contents
      of regular expressions at all. The flag was added despite this limitation
      because it's still useful for code that doesn't contain cases like this.
      </p>
    - >
      <p>
      This flag does not apply to comments. I believe preserving non-ASCII data
      in comments should be fine because even if the encoding is wrong, the run
      time environment should completely ignore the contents of all comments.
      For example, the [V8 blog post](https://v8.dev/blog/scanner) mentions an
      optimization that avoids decoding comment contents completely. And all
      comments other than license-related comments are stripped out by esbuild
      anyway.
      </p>
    - >
      <p>
      This option simultaneously applies to all output file types (JavaScript,
      CSS, and JSON). So if you configure your web server to send the correct
      <code>Content-<wbr>Type</code> header and want to use the UTF-8 charset,
      make sure your web server is configured to treat both `.js` and `.css`
      files as UTF-8.
      </p>

  ## Footer

  - p: >
      Use this to insert an arbitrary string at the end of generated JavaScript
      and CSS files. This is commonly used to insert comments:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --footer:js=//comment --footer:css=/*comment*/

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          footer: {
            js: '//comment',
            css: '/*comment*/',
          },
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Footer: map[string]string{
              "js":  "//comment",
              "css": "/*comment*/",
            },
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      This is similar to [banner](#banner) which inserts at the beginning
      instead of the end.

  ## Format

  - p: >
      This sets the output format for the generated JavaScript files. There are
      currently three possible values that can be configured: `iife`, `cjs`,
      and `esm`. When no output format is specified, esbuild picks an output
      format for you if [bundling](#bundle) is enabled (as described below),
      or doesn't do any format conversion if [bundling](#bundle) is disabled.

  - h4#format-iife: IIFE

  - p: >
      The `iife` format stands for "immediately-invoked function expression" and
      is intended to be run in the browser. Wrapping your code in a function
      expression ensures that any variables in your code don't accidentally
      conflict with variables in the global scope. If your entry point has
      exports that you want to expose as a global in the browser, you can
      configure that global's name using the [global name](#global-name)
      setting. The `iife` format will automatically be enabled when no output
      format is specified, [bundling](#bundle) is enabled, and [platform](#platform)
      is set to `browser` (which it is by default). Specifying the `iife` format
      looks like this:

  - example:
      cli:
        - $: |
            echo 'alert("test")' | esbuild --format=iife
        - expect: |
            (() => {
              alert("test");
            })();

      mjs: |
        import * as esbuild from 'esbuild'

        let js = 'alert("test")'
        let result = await esbuild.transform(js, {
          format: 'iife',
        })
        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "alert(\"test\")"

          result := api.Transform(js, api.TransformOptions{
            Format: api.FormatIIFE,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - h4#format-commonjs: CommonJS

  - p: >
      The `cjs` format stands for "CommonJS" and is intended to be run in node.
      It assumes the environment contains `exports`, `require`, and `module`.
      Entry points with exports in ECMAScript module syntax will be converted
      to a module with a getter on `exports` for each export name. The `cjs`
      format will automatically be enabled when no output format is specified,
      [bundling](#bundle) is enabled, and [platform](#platform) is set to `node`.
      Specifying the `cjs` format looks like this:

  - example:
      cli:
        - $: |
            echo 'export default "test"' | esbuild --format=cjs
        - expect: |
            ...
            var stdin_exports = {};
            __export(stdin_exports, {
              default: () => stdin_default
            });
            module.exports = __toCommonJS(stdin_exports);
            var stdin_default = "test";

      mjs: |
        import * as esbuild from 'esbuild'

        let js = 'export default "test"'
        let result = await esbuild.transform(js, {
          format: 'cjs',
        })
        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "export default 'test'"

          result := api.Transform(js, api.TransformOptions{
            Format: api.FormatCommonJS,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - h4#format-esm: ESM

  - p: >
      The `esm` format stands for "ECMAScript module". It assumes the environment
      supports `import` and `export` syntax. Entry points with exports in CommonJS
      module syntax will be converted to a single `default` export of the value
      of `module.exports`. The `esm` format will automatically be enabled when no
      output format is specified, [bundling](#bundle) is enabled, and [platform](#platform)
      is set to `neutral`. Specifying the `esm` format looks like this:

  - example:
      cli:
        - $: |
            echo 'module.exports = "test"' | esbuild --format=esm
        - expect: |
            ...
            var require_stdin = __commonJS({
              "<stdin>"(exports, module) {
                module.exports = "test";
              }
            });
            export default require_stdin();

      mjs: |
        import * as esbuild from 'esbuild'

        let js = 'module.exports = "test"'
        let result = await esbuild.transform(js, {
          format: 'esm',
        })
        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "module.exports = 'test'"

          result := api.Transform(js, api.TransformOptions{
            Format: api.FormatESModule,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      The `esm` format can be used either in the browser or in node, but you
      have to explicitly load it as a module. This happens automatically if you
      `import` it from another module. Otherwise:

  - ul:
    - >
      In the browser, you can load a module using <code>&lt;script <wbr>src="<wbr>file.js" <wbr>type="<wbr>module"&gt;<wbr>&lt;/script&gt;</code>.
      Do not forget <code>type="<wbr>module"</code> as this will break your
      code in subtle and confusing ways (omitting <code>type="<wbr>module"</code>
      means that all top-level variables will end up in the global scope, which
      will then collide with top-level variables that have the same name in other
      JavaScript files).
      <br>&nbsp;

    - >
      In node, you can load a module using <code>node <wbr>file.mjs</code>.
      Note that node requires the `.mjs` extension unless you have configured
      <code>"type": <wbr>"module"</code> in your `package.json` file.
      You can use the [out extension](#out-extension) setting in esbuild to
      customize the output extension for the files esbuild generates. You can
      read more about using ECMAScript modules in node [here](https://nodejs.org/api/esm.html#enabling).

  ## Global name

  - p: >
      This option only matters when the [format](#format) setting is `iife`
      (which stands for immediately-invoked function expression). It sets the
      name of the global variable which is used to store the exports from the
      entry point:

  - example:
      cli: |
        echo 'module.exports = "test"' | esbuild --format=iife --global-name=xyz

      mjs: |
        import * as esbuild from 'esbuild'

        let js = 'module.exports = "test"'
        let result = await esbuild.transform(js, {
          format: 'iife',
          globalName: 'xyz',
        })
        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "module.exports = 'test'"

          result := api.Transform(js, api.TransformOptions{
            Format:     api.FormatIIFE,
            GlobalName: "xyz",
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      Specifying the global name with the `iife` format will generate code that
      looks something like this:

  - pre.js: |
      var xyz = (() => {
        ...
        var require_stdin = __commonJS((exports, module) => {
          module.exports = "test";
        });
        return require_stdin();
      })();

  - p: >
      The global name can also be a compound property expression, in which case
      esbuild will generate a global variable with that property. Existing
      global variables that conflict will not be overwritten. This can be used
      to implement "namespacing" where multiple independent scripts add their
      exports onto the same global object. For example:

  - example:
      cli: |
        echo 'module.exports = "test"' | esbuild --format=iife --global-name='example.versions["1.0"]'

      mjs: |
        import * as esbuild from 'esbuild'

        let js = 'module.exports = "test"'
        let result = await esbuild.transform(js, {
          format: 'iife',
          globalName: 'example.versions["1.0"]',
        })
        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "module.exports = 'test'"

          result := api.Transform(js, api.TransformOptions{
            Format:     api.FormatIIFE,
            GlobalName: `example.versions["1.0"]`,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      The compound global name used above generates code that looks like this:

  - pre.js: |
      var example = example || {};
      example.versions = example.versions || {};
      example.versions["1.0"] = (() => {
        ...
        var require_stdin = __commonJS((exports, module) => {
          module.exports = "test";
        });
        return require_stdin();
      })();

  ## Legal comments

  - p: >
      A "legal comment" is considered to be any statement-level comment in JS
      or rule-level  comment in CSS that contains `@license` or `@preserve` or
      that starts with `//!` or `/*!`. These comments are preserved in output
      files by default since that follows the intent of the original authors
      of the code. However, this behavior can be configured by using one of
      the following options:

  - ul:
    - >
      <p>`none`<br>Do not preserve any legal comments.</p>

    - >
      <p>`inline`<br>Preserve all legal comments.</p>

    - >
      <p>`eof`<br>Move all legal comments to the end of the file.</p>

    - >
      <p>`linked`<br>Move all legal comments to a `.LEGAL.txt` file and link to them with a comment.</p>

    - >
      <p>`external`<br>Move all legal comments to a `.LEGAL.txt` file but to not link to them.</p>

  - p: >
      The default behavior is `eof` when [bundling](#bundle) is enabled and `inline` otherwise.
      Setting the legal comment mode looks like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --legal-comments=eof

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          legalComments: 'eof',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:   []string{"app.js"},
            LegalComments: api.LegalCommentsEndOfFile,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Note that "statement-level" for JS and "rule-level" for CSS means the
      comment must appear in a context where multiple statements or rules
      are allowed such as in the top-level scope or in a statement or rule
      block. So comments inside expressions or at the declaration level are
      not considered legal comments.

  ## Line limit

  - p: >
      This setting is a way to prevent esbuild from generating output files
      with really long lines, which can help editing performance in
      poorly-implemented text editors. Set this to a positive integer to tell
      esbuild to end a given line soon after it passes that number of bytes.
      For example, this wraps long lines soon after they pass ~80 characters:

  - example:
      in:
        app.ts: '1 + 2'

      cli: |
        esbuild app.ts --line-limit=80

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.ts'],
          lineLimit: 80,
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.ts"},
            LineLimit:   80,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Lines are truncated after they pass the limit instead of before because
      it's simpler to check when the limit is passed then to predict when the
      limit is about to be passed, and because it's faster to avoid backing up
      and rewriting things when generating an output file. So the limit is only
      approximate.

  - p: >
      This setting applies to both JavaScript and CSS, and works even when
      minification is disabled. Note that turning this setting on will make
      your files bigger, as the extra newlines take up additional space in
      the file (even after gzip compression).

  ## Splitting

  - warning: >
      Code splitting is still a work in progress. It currently only works with
      the `esm` output [format](#format). There is also a known
      [ordering issue](https://github.com/evanw/esbuild/issues/399) with
      `import` statements across code splitting chunks. You can follow
      [the tracking issue](https://github.com/evanw/esbuild/issues/16) for
      updates about this feature.

  - p: >
      This enables "code splitting" which serves two purposes:

  - ul:
      - >
        <p>
        Code shared between multiple entry points is split off into a separate
        shared file that both entry points import. That way if the user first
        browses to one page and then to another page, they don't have to
        download all of the JavaScript for the second page from scratch if the
        shared part has already been downloaded and cached by their browser.
        </p>

      - >
        <p>
        Code referenced through an asynchronous `import()` expression will be
        split off into a separate file and only loaded when that expression is
        evaluated. This allows you to improve the initial download time of your
        app by only downloading the code you need at startup, and then lazily
        downloading additional code if needed later.
        </p>
        <p>
        Without code splitting enabled, an `import()` expression becomes
        <code>Promise<wbr>.resolve()<wbr>.then(() =&gt; <wbr>require())</code>
        instead. This still preserves the asynchronous semantics of the
        expression but it means the imported code is included in the same
        bundle instead of being split off into a separate file.
        </p>

  - p: >
      When you enable code splitting you must also configure the output
      directory using the [outdir](#outdir) setting:

  - example:
      in:
        home.ts: '1 + 2'
        about.ts: '1 + 2'

      cli: |
        esbuild home.ts about.ts --bundle --splitting --outdir=out --format=esm

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['home.ts', 'about.ts'],
          bundle: true,
          splitting: true,
          outdir: 'out',
          format: 'esm',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"home.ts", "about.ts"},
            Bundle:      true,
            Splitting:   true,
            Outdir:      "out",
            Format:      api.FormatESModule,
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  # Output location

  ## Allow overwrite

  - p: >
      Enabling this setting allows output files to overwrite input files. It's
      not  enabled by default because doing so means overwriting your source
      code, which can lead to data loss if your code is not checked in. But
      supporting this makes certain workflows easier by avoiding the need for
      a temporary directory. So you can enable this when you want to deliberately
      overwrite your source code:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --outdir=. --allow-overwrite

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          outdir: '.',
          allowOverwrite: true,
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:    []string{"app.js"},
            Outdir:         ".",
            AllowOverwrite: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## Asset names

  - p: >
      This option controls the file names of the additional output files generated
      when the [loader](#loader) is set to [`file`](/content-types/#external-file).
      It configures the output paths using a template with placeholders that will
      be substituted with values specific to the file when the output path is
      generated. For example, specifying an asset name template of
      <code>assets/<wbr>[name]-<wbr>[hash]</code> puts all assets into a
      subdirectory called `assets` inside of the output directory and includes
      the content hash of the asset in the file name. Doing that looks like this:

  - example:
      in:
        app.js: 'import "./file.png"'
        file.png: 'a png file'

      cli: |
        esbuild app.js --asset-names=assets/[name]-[hash] --loader:.png=file --bundle --outdir=out

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          assetNames: 'assets/[name]-[hash]',
          loader: { '.png': 'file' },
          bundle: true,
          outdir: 'out',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            AssetNames:  "assets/[name]-[hash]",
            Loader: map[string]api.Loader{
              ".png": api.LoaderFile,
            },
            Bundle: true,
            Outdir: "out",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      There are four placeholders that can be used in asset path templates:

  - ul:
    - >
      `[dir]`
      <p>
      This is the relative path from the directory containing the asset file to
      the [outbase](#outbase) directory. Its purpose is to help asset output
      paths look more aesthetically pleasing by mirroring the input directory
      structure inside of the output directory.
      </p>

    - >
      `[name]`
      <p>
      This is the original file name of the asset without the extension. For
      example, if the asset was originally named `image.png` then `[name]` will
      be substituted with `image` in the template. It is not necessary to use
      this placeholder; it only exists to provide human-friendly asset names to
      make debugging easier.
      </p>

    - >
      `[hash]`
      <p>
      This is the content hash of the asset, which is useful to avoid name
      collisions. For example, your code may import <code>components/<wbr>button/<wbr>icon.png</code>
      and <code>components/<wbr>select/<wbr>icon.png</code> in which case
      you'll need the hash to distinguish between the two assets that are both
      named `icon`.
      </p>

    - >
      `[ext]`
      <p>
      This is the file extension of the asset (i.e. everything after the end of
      the last `.` character). It can be used to put different types of assets
      into different directories. For example, <code>--asset-names=<wbr>assets/<wbr>[ext]/<wbr>[name]-[hash]</code>
      might write out an asset named `image.png` as <code>assets/<wbr>png/<wbr>image-CQFGD2NG.png</code>.
      </p>

  - p: >
      Asset path templates do not need to include a file extension. The original
      file extension of the asset will be automatically added to the end of the
      output path after template substitution.

  - p: >
      This option is similar to the [chunk names](#chunk-names) and
      [entry names](#entry-names) options.

  ## Chunk names

  - p: >
      This option controls the file names of the chunks of shared code that are
      automatically generated when [code splitting](#splitting) is enabled.
      It configures the output paths using a template with placeholders that will
      be substituted with values specific to the chunk when the output path is
      generated. For example, specifying a chunk name template of
      <code>chunks/<wbr>[name]-<wbr>[hash]</code> puts all generated chunks into
      a subdirectory called `chunks` inside of the output directory and includes
      the content hash of the chunk in the file name. Doing that looks like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --chunk-names=chunks/[name]-[hash] --bundle --outdir=out --splitting --format=esm

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          chunkNames: 'chunks/[name]-[hash]',
          bundle: true,
          outdir: 'out',
          splitting: true,
          format: 'esm',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            ChunkNames:  "chunks/[name]-[hash]",
            Bundle:      true,
            Outdir:      "out",
            Splitting:   true,
            Format:      api.FormatESModule,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      There are three placeholders that can be used in chunk path templates:

  - ul:
    - >
      `[name]`
      <p>
      This will currently always be the text `chunk`, although this placeholder
      may take on additional values in future releases.
      </p>

    - >
      `[hash]`
      <p>
      This is the content hash of the chunk. Including this is necessary to
      distinguish different chunks from each other in the case where multiple
      chunks of shared code are generated.
      </p>

    - >
      `[ext]`
      <p>
      This is the file extension of the chunk (i.e. everything after the end of
      the last `.` character). It can be used to put different types of chunks
      into different directories. For example, <code>--chunk-names=<wbr>chunks/<wbr>[ext]/<wbr>[name]-[hash]</code>
      might write out a chunk as <code>chunks/<wbr>css/<wbr>chunk-DEFJT7KY.css</code>.
      </p>

  - p: >
      Chunk path templates do not need to include a file extension. The
      configured [out extension](#out-extension) for the appropriate content
      type will be automatically added to the end of the output path after
      template substitution.

  - p: >
      Note that this option only controls the names for automatically-generated
      chunks of shared code. It does _not_ control the names for output files
      related to entry points. The names of these are currently determined from
      the path of the original entry point file relative to the [outbase](#outbase)
      directory, and this behavior cannot be changed. An additional API option
      will be added in the future to let you change the file names of entry
      point output files.

  - p: >
      This option is similar to the [asset names](#asset-names) and
      [entry names](#entry-names) options.

  ## Entry names

  - p: >
      This option controls the file names of the output files corresponding to each
      input entry point file. It configures the output paths using a template with
      placeholders that will be substituted with values specific to the file when
      the output path is generated. For example, specifying an entry name template
      of <code>[dir]/<wbr>[name]-<wbr>[hash]</code> includes a hash of the output
      file in the file name and puts the files into the output directory,
      potentially under a subdirectory (see the details about `[dir]` below).
      Doing that looks like this:

  - example:
      in:
        src/main-app/app.js: '1 + 2'

      cli: |
        esbuild src/main-app/app.js --entry-names=[dir]/[name]-[hash] --outbase=src --bundle --outdir=out

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['src/main-app/app.js'],
          entryNames: '[dir]/[name]-[hash]',
          outbase: 'src',
          bundle: true,
          outdir: 'out',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"src/main-app/app.js"},
            EntryNames:  "[dir]/[name]-[hash]",
            Outbase:     "src",
            Bundle:      true,
            Outdir:      "out",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      There are four placeholders that can be used in entry path templates:

  - ul:
    - >
      `[dir]`
      <p>
      This is the relative path from the directory containing the input entry
      point file to the [outbase](#outbase) directory. Its purpose is to help
      you avoid collisions between identically-named entry points in different
      subdirectories.
      </p>
      <p>
      For example, if there are two entry points
      <code>src/<wbr>pages/<wbr>home/<wbr>index.ts</code> and
      <code>src/<wbr>pages/<wbr>about/<wbr>index.ts</code>, the outbase
      directory is `src`, and the entry names template is `[dir]/[name]`,
      the output directory will contain <code>pages/<wbr>home/<wbr>index.js</code>
      and <code>pages/<wbr>about/<wbr>index.js</code>. If the entry names
      template had been just `[name]` instead, bundling would have failed
      because there would have been two output files with the same output path
      `index.js` inside the output directory.
      </p>

    - >
      `[name]`
      <p>
      This is the original file name of the entry point without the extension.
      For example, if the input entry point file is named `app.js` then `[name]`
      will be substituted with `app` in the template.
      </p>

    - >
      `[hash]`
      <p>
      This is the content hash of the output file, which can be used to take
      optimal advantage of browser caching. Adding `[hash]` to your entry
      point names means esbuild will calculate a hash that relates to all
      content in the corresponding output file (and any output file it imports
      if [code splitting](#splitting) is active). The hash is designed to change
      if and only if any of the input files relevant to that output file are
      changed.
      </p>
      <p>
      After that, you can have your web server tell browsers that to cache these
      files forever (in practice you can say they expire a very long time from
      now such as in a year). You can then use the information in the
      [metafile](#metafile) to determine which output file path corresponds to
      which input entry point so you know what path to include in your `<script>`
      tag.
      </p>

    - >
      `[ext]`
      <p>
      This is the file extension that the entry point file will be written out to
      (i.e. the [out extension](#out-extension) setting, not the original file
      extension). It can be used to put different types of entry points into different
      directories. For example, <code>--entry-names=<wbr>entries/<wbr>[ext]/<wbr>[name]</code>
      might write the output file for `app.ts` to <code>entries/<wbr>js/<wbr>app.js</code>.
      </p>

  - p: >
      Entry path templates do not need to include a file extension. The appropriate
      [out extension](#out-extension) based on the file type will be automatically
      added to the end of the output path after template substitution.

  - p: >
      This option is similar to the [asset names](#asset-names) and
      [chunk names](#chunk-names) options.

  ## Out extension

  - p: >
      This option lets you customize the file extension of the files that
      esbuild generates to something other than `.js` or `.css`. In particular,
      the `.mjs` and `.cjs` file extensions have special meaning in node (they
      indicate a file in ESM and CommonJS format, respectively). This option is
      useful if you are using esbuild to generate multiple files and you have
      to use the [outdir](#outdir) option instead of the [outfile](#outfile)
      option. You can use it like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --bundle --outdir=dist --out-extension:.js=.mjs

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          outdir: 'dist',
          outExtension: { '.js': '.mjs' },
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Outdir:      "dist",
            OutExtension: map[string]string{
              ".js": ".mjs",
            },
            Write: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## Outbase

  - p: >
      If your build contains multiple entry points in separate directories, the
      directory structure will be replicated into the [output directory](#outdir)
      relative to the outbase directory. For example, if there are two entry
      points <code>src/<wbr>pages/<wbr>home/<wbr>index.ts</code> and
      <code>src/<wbr>pages/<wbr>about/<wbr>index.ts</code> and the outbase directory is
      `src`, the output directory will contain <code>pages/<wbr>home/<wbr>index.js</code>
      and <code>pages/<wbr>about/<wbr>index.js</code>. Here's how to use it:

  - example:
      in:
        src/pages/home/index.ts: '1 + 2'
        src/pages/about/index.ts: '3 + 4'

      cli: |
        esbuild src/pages/home/index.ts src/pages/about/index.ts --bundle --outdir=out --outbase=src

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: [
            'src/pages/home/index.ts',
            'src/pages/about/index.ts',
          ],
          bundle: true,
          outdir: 'out',
          outbase: 'src',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{
              "src/pages/home/index.ts",
              "src/pages/about/index.ts",
            },
            Bundle:  true,
            Outdir:  "out",
            Outbase: "src",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      If the outbase directory isn't specified, it defaults to the
      [lowest common ancestor](https://en.wikipedia.org/wiki/Lowest_common_ancestor)
      directory among all input entry point paths. This is <code>src/<wbr>pages</code>
      in the example above, which means by default the output directory will
      contain <code>home/<wbr>index.js</code> and <code>about/<wbr>index.js</code>
      instead.

  ## Outdir

  - p: >
      This option sets the output directory for the build operation. For
      example, this command will generate a directory called `out`:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --bundle --outdir=out

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          outdir: 'out',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Outdir:      "out",
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      The output directory will be generated if it does not already exist, but
      it will not be cleared if it already contains some files. Any generated
      files will silently overwrite existing files with the same name. You
      should clear the output directory yourself before running esbuild if you
      want the output directory to only contain files from the current run of
      esbuild.

  - p: >
      If your build contains multiple entry points in separate directories, the
      directory structure will be replicated into the output directory starting
      from the [lowest common ancestor](https://en.wikipedia.org/wiki/Lowest_common_ancestor)
      directory among all input entry point paths. For example, if there are
      two entry points <code>src/<wbr>home/<wbr>index.ts</code> and
      <code>src/<wbr>about/<wbr>index.ts</code>, the output directory will
      contain <code>home/<wbr>index.js</code> and <code>about/<wbr>index.js</code>.
      If you want to customize this behavior, you should change the
      [outbase directory](#outbase).

  ## Outfile

  - p: >
      This option sets the output file name for the build operation. This is
      only applicable if there is a single entry point. If there are multiple
      entry points, you must use the [outdir](#outdir) option instead to
      specify an output directory. Using outfile looks like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --bundle --outfile=out.js

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Outfile:     "out.js",
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## Public path

  - p: >
      This is useful in combination with the [external file](/content-types/#external-file)
      loader. By default that loader exports the name of the imported file as a
      string using the `default` export. The public path option lets you
      prepend a base path to the exported string of each file loaded by this
      loader:

  - example:
      in:
        app.js: |
          import url from './example.png'
          let image = new Image
          image.src = url
          document.body.appendChild(image)
        example.png: |
          this is some data

      cli: |
        esbuild app.js --bundle --loader:.png=file --public-path=https://www.example.com/v1 --outdir=out

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.png': 'file' },
          publicPath: 'https://www.example.com/v1',
          outdir: 'out',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Loader: map[string]api.Loader{
              ".png": api.LoaderFile,
            },
            Outdir:     "out",
            PublicPath: "https://www.example.com/v1",
            Write:      true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## Write

  - p: >
      The build API call can either write to the file system directly or return
      the files that would have been written as in-memory buffers. By default
      the CLI and JavaScript APIs write to the file system and the Go API
      doesn't. To use the in-memory buffers:

  - example:
      in:
        app.js: '1 + 2'

      mjs: |
        import * as esbuild from 'esbuild'

        let result = await esbuild.build({
          entryPoints: ['app.js'],
          sourcemap: 'external',
          write: false,
          outdir: 'out',
        })

        for (let out of result.outputFiles) {
          console.log(out.path, out.contents, out.hash, out.text)
        }

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Sourcemap:   api.SourceMapExternal,
            Write:       false,
            Outdir:      "out",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }

          for _, out := range result.OutputFiles {
            fmt.Printf("%v %v %s\n", out.Path, out.Contents, out.Hash)
          }
        }

  - p: >
      The `hash` property is a hash of the `contents` field and has been provided
      for convenience. The hash algorithm (currently [XXH64](https://xxhash.com/))
      is implementation-dependent and may be changed at any time in between esbuild
      versions.

  # Path resolution

  ## Alias

  - p: >
      This feature lets you substitute one package for another when bundling.
      The example below substitutes the package `oldpkg` with the package
      `newpkg`:

  - example:
      in:
        app.js: 'import "oldpkg/foo"'
        node_modules/newpkg/foo.js: 'works'

      cli: |
        esbuild app.js --bundle --alias:oldpkg=newpkg

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          write: true,
          alias: {
            'oldpkg': 'newpkg',
          },
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Write:       true,
            Alias: map[string]string{
              "oldpkg": "newpkg",
            },
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      These new substitutions happen first before all of esbuild's other path
      resolution logic. One use case for this feature is replacing a node-only
      package with a browser-friendly package in third-party code that you
      don't control.

  - p: >
      Note that when an import path is substituted using an alias, the resulting
      import path is resolved in the working directory instead of in the directory
      containing the source file with the import path. If needed, the working
      directory that esbuild uses can be set with the [working directory](#working-directory)
      feature.

  ## Conditions

  - p: >
      This feature controls how the `exports` field in `package.json` is
      interpreted. Custom conditions can be added using the conditions setting.
      You can specify as many of these as you want and the meaning of these is
      entirely up to package authors. Node has currently only endorsed the
      `development` and `production` custom conditions for recommended use.
      Here is an example of adding the custom conditions `custom1` and `custom2`:

  - example:
      in:
        src/app.js: 'import "pkg"'
        src/node_modules/pkg/package.json: '{ "exports": { "custom1": "./foo.js" } }'
        src/node_modules/pkg/foo.js: 'console.log(123)'

      cli: |
        esbuild src/app.js --bundle --conditions=custom1,custom2

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['src/app.js'],
          bundle: true,
          conditions: ['custom1', 'custom2'],
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"src/app.js"},
            Bundle:      true,
            Conditions:  []string{"custom1", "custom2"},
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ### How conditions work

  - p: >
      Conditions allow you to redirect the same import path to different file
      locations in different situations. The redirect map containing the
      conditions and paths is stored in the `exports` field in the package's
      `package.json` file. For example, this would remap `require('pkg/foo')`
      to `pkg/required.cjs` and `import 'pkg/foo'` to `pkg/imported.mjs`
      using the `import` and `require` conditions:

  - pre.json: |
      {
        "name": "pkg",
        "exports": {
          "./foo": {
            "import": "./imported.mjs",
            "require": "./required.cjs",
            "default": "./fallback.js"
          }
        }
      }

  - p: >
      Conditions are checked in the order that they appear within the JSON file.
      So the example above behaves sort of like this:

  - pre.js: |
      if (importPath === './foo') {
        if (conditions.has('import')) return './imported.mjs'
        if (conditions.has('require')) return './required.cjs'
        return './fallback.js'
      }

  - p: >
      By default there are five conditions with special behavior that are built
      in to esbuild, and cannot be disabled:

  - ul:
    - >
      `default`
      <p>
      This condition is always active. It is intended to come last and lets you
      provide a fallback for when no other condition applies. This condition is
      also active when you run your code natively in node.
      </p>
    - >
      `import`
      <p>
      This condition is only active when the import path is from an ESM `import`
      statement or `import()` expression. It can be used to provide ESM-specific
      code. This condition is also active when you run your code natively in node
      (but only in an ESM context).
      </p>
    - >
      `require`
      <p>
      This condition is only active when the import path is from a CommonJS
      `require()` call. It can be used to provide CommonJS-specific code. This
      condition is also active when you run your code natively in node (but only
      in a CommonJS context).
      </p>
    - >
      `browser`
      <p>
      This condition is only active when esbuild's [platform](#platform) setting
      is set to `browser`. It can be used to provide browser-specific code. This
      condition is not active when you run your code natively in node.
      </p>
    - >
      `node`
      <p>
      This condition is only active when esbuild's [platform](#platform) setting
      is set to `node`. It can be used to provide node-specific code. This
      condition is also active when you run your code natively in node.
      </p>

  - p: >
      The following condition is also automatically included when the [platform](#platform)
      is set to either `browser` or `node` and no custom conditions are
      configured. If there are any custom conditions configured (even an empty
      list) then this condition will no longer be automatically included:

  - ul:
    - >
      `module`
      <p>
      This condition can be used to tell esbuild to pick the ESM variant for a
      given import path to provide better tree-shaking when bundling. This
      condition is not active when you run your code natively in node. It is
      specific to bundlers, and originated from Webpack.
      </p>

  - p: >
      Note that when you use the `require` and `import` conditions, _your
      package may end up in the bundle multiple times!_ This is a subtle issue
      that can cause bugs due to duplicate copies of your code's state in
      addition to bloating the resulting bundle. This is commonly known as the
      [dual package hazard](https://nodejs.org/docs/latest/api/packages.html#packages_dual_package_hazard).

  - p: >
      One way of avoiding the dual package hazard that works both for bundlers
      and when running natively in node is to put all of your code in the
      `require` condition as CommonJS and have the `import` condition just be a
      light ESM wrapper that calls `require` on your package and re-exports the
      package using ESM syntax. This approach doesn't provide good tree-shaking,
      however, as esbuild doesn't tree-shake CommonJS modules.

  - p: >
      Another way of avoiding a dual package hazard is to use the bundler-specific
      `module` condition to direct bundlers to always load the ESM version of your
      package while letting node always fall back to the CommonJS version of
      your package. Both `import` and `module` are intended to be used with ESM
      but unlike `import`, the `module` condition is always active even if the
      import path was loaded using a `require` call. This works well with bundlers
      because bundlers support loading ESM using `require`, but it's not something
      that can work with node because node deliberately doesn't implement loading
      ESM using `require`.

  ## External

  - p: >
      You can mark a file or a package as external to exclude it from your
      build. Instead of being bundled, the import will be preserved (using
      `require` for the `iife` and `cjs` formats and using `import` for the
      `esm` format) and will be evaluated at run time instead.

  - p: >
      This has several uses. First of all, it can be used to trim unnecessary
      code from your bundle for a code path that you know will never be
      executed. For example, a package may contain code that only runs in node
      but you will only be using that package in the browser. It can also be
      used to import code in node at run time from a package that cannot be
      bundled. For example, the `fsevents` package contains a native extension,
      which esbuild doesn't support. Marking something as external looks like
      this:

  - example:
      cli:
        - $: |
            echo 'require("fsevents")' > app.js
        - $: |
            esbuild app.js --bundle --external:fsevents --platform=node
        - expect: |
            // app.js
            require("fsevents");

      mjs: |
        import * as esbuild from 'esbuild'
        import fs from 'node:fs'

        fs.writeFileSync('app.js', 'require("fsevents")')

        await esbuild.build({
          entryPoints: ['app.js'],
          outfile: 'out.js',
          bundle: true,
          platform: 'node',
          external: ['fsevents'],
        })

      go: |
        package main

        import "io/ioutil"
        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          ioutil.WriteFile("app.js", []byte("require(\"fsevents\")"), 0644)

          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Outfile:     "out.js",
            Bundle:      true,
            Write:       true,
            Platform:    api.PlatformNode,
            External:    []string{"fsevents"},
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      You can also use the `*` wildcard character in an external path to mark
      all files matching that pattern as external. For example, you can use
      `*.png` to remove all `.png` files or `/images/*` to remove all paths
      starting with `/images/`:

  - example:
      in:
        app.js: 'import "/images/*"; import "*.png"'

      cli: |
        esbuild app.js --bundle "--external:*.png" "--external:/images/*"

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          outfile: 'out.js',
          bundle: true,
          external: ['*.png', '/images/*'],
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Outfile:     "out.js",
            Bundle:      true,
            Write:       true,
            External:    []string{"*.png", "/images/*"},
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      External paths are applied both before and after path resolution, which
      lets you match against both the import path in the source code and the
      absolute file system path. The path is considered to be external if the
      external path matches in either case. The specific behavior is as follows:

  - ul:
    - >
      <p>
      Before path resolution begins, import paths are checked against all
      external paths. In addition, if the external path looks like a package
      path (i.e. doesn't start with `/` or `./` or `../`), import paths are
      checked to see if they have that package path as a path prefix.
      </p>
      <p>
      This means that <code>--external:<wbr>@foo/<wbr>bar</code> implicitly
      also means <code>--external:<wbr>@foo/<wbr>bar/\*</code> which matches
      the import path <code>@foo/<wbr>bar/<wbr>baz</code>. So it marks all
      paths inside the `@foo/bar` package as external too.
      </p>

    - >
      <p>
      After path resolution ends, the resolved absolute paths are checked
      against all external paths that don't look like a package path (i.e.
      those that start with `/` or `./` or `../`). But before checking, the
      external path is joined with the current working directory and then
      normalized, becoming an absolute path (even if it contains a `*`
      wildcard character).
      </p>
      <p>
      This means that you can mark everything in the directory `dir` as
      external using <code>--external:<wbr>./dir/\*</code>. Note that the
      leading `./` is important. Using <code>--external:<wbr>dir/\*</code>
      instead is treated as a package path and is not checked for after
      path resolution ends.
      </p>

  ## Main fields

  - p: >
      When you import a package in node, the `main` field in that package's
      `package.json` file determines which file is imported (along with
      [a lot of other rules](https://nodejs.org/api/modules.html#all-together)).
      Major JavaScript bundlers including esbuild let you specify additional
      `package.json` fields to try when resolving a package. There are at least
      three such fields commonly in use:

  - ul:
    - >
      `main`
      <p>
      This is [the standard field](https://docs.npmjs.com/files/package.json#main)
      for all packages that are meant to be used with node. The name `main` is
      hard-coded in to node's module resolution logic itself. Because it's
      intended for use with node, it's reasonable to expect that the file path
      in this field is a CommonJS-style module.
      </p>

    - >
      `module`
      <p>
      This field came from [a proposal](https://github.com/dherman/defense-of-dot-js/blob/f31319be735b21739756b87d551f6711bd7aa283/proposal.md)
      for how to integrate ECMAScript modules into node. Because of this, it's
      reasonable to expect that the file path in this field is an
      ECMAScript-style module. This proposal wasn't adopted by node (node uses
      <code>"type": <wbr>"module"</code> instead) but it was adopted by major
      bundlers because ECMAScript-style modules lead to better [tree shaking](#tree-shaking),
      or dead code removal.
      </p>
      <p>
      For package authors: Some packages incorrectly use the `module` field for
      browser-specific code, leaving node-specific code for the `main` field.
      This is probably because node ignores the `module` field and people
      typically only use bundlers for browser-specific code. However, bundling
      node-specific code is valuable too (e.g. it decreases download and boot
      time) and packages that put browser-specific code in `module` prevent
      bundlers from being able to do tree shaking effectively. If you are
      trying to publish browser-specific code in a package, use the `browser`
      field instead.
      </p>

    - >
      `browser`
      <p>
      This field came from [a proposal](https://gist.github.com/defunctzombie/4339901/49493836fb873ddaa4b8a7aa0ef2352119f69211)
      that allows bundlers to replace node-specific files or modules with their
      browser-friendly versions. It lets you specify an alternate
      browser-specific entry point. Note that it is possible for a package to
      use both the `browser` and `module` field together (see the note below).
      </p>

  - p: >
      The default main fields depend on the current [platform](#platform)
      setting. These defaults should be the most widely compatible with the
      existing package ecosystem. But you can customize them like this if you
      want to:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --bundle --main-fields=module,main

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          mainFields: ['module', 'main'],
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            MainFields:  []string{"module", "main"},
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - h4#main-fields-for-package-authors: For package authors

  - p: >
      If you want to author a package that uses the `browser` field in
      combination with the `module` field, then you'll probably want to fill
      out **_all four entries_** in the full CommonJS-vs-ESM and
      browser-vs-node compatibility matrix. For that you'll need to use the
      expanded form of the `browser` field that is a map instead of just a
      string:

  - pre.json: >
      {
        "main": "./node-cjs.js",
        "module": "./node-esm.js",
        "browser": {
          "./node-cjs.js": "./browser-cjs.js",
          "./node-esm.js": "./browser-esm.js"
        }
      }

  - p: >
      The `main` field is expected to be CommonJS while the `module` field is
      expected to be ESM. The decision about which module format to use is
      independent from the decision about whether to use a browser-specific
      or node-specific variant. If you omit one of these four entries, then you
      risk the wrong variant being chosen. For example, if you omit the entry
      for the CommonJS browser build, then the CommonJS node build could be
      chosen instead.

  - p: >
      Note that using `main`, `module`, and `browser` is the old way of doing
      this. There is also a newer way to do this that you may prefer to use
      instead: the [`exports` field](#how-conditions-work) in `package.json`.
      It provides a different set of trade-offs. For example, it gives you more
      precise control over imports for all sub-paths in your package (while
      `main` fields only give you control over the entry point), but it may
      cause your package to be imported multiple times depending on how you
      configure it.

  ## Node paths

  - p: >
      Node's module resolution algorithm supports an environment variable called
      [`NODE_PATH`](https://nodejs.org/api/modules.html#modules_loading_from_the_global_folders)
      that contains a list of global directories to use when resolving import
      paths. These paths are searched for packages in addition to the
      `node_modules` directories in all parent directories. You can pass this
      list of directories to esbuild using an environment variable with the CLI
      and using an array with the JS and Go APIs:

  - example:
      in:
        app.js: 'import {x} from "test"'
        someDir/test.js: 'export let x'

      cli: |
        NODE_PATH=someDir esbuild app.js --bundle --outfile=out.js

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          nodePaths: ['someDir'],
          entryPoints: ['app.js'],
          bundle: true,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            NodePaths:   []string{"someDir"},
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Outfile:     "out.js",
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      If you are using the CLI and want to pass multiple directories using
      `NODE_PATH`, you will have to separate them with `:` on Unix and `;` on
      Windows. This is the same format that Node itself uses.

  ## Packages

  - p: >
      Use this setting to control whether all of your package's dependencies
      are excluded from the bundle or not. This is useful when
      [bundling for node](/getting-started/#bundling-for-node)
      because many npm packages use node-specific features that esbuild doesn't
      support while bundling (such as `__dirname`, `import.meta.url`, `fs.readFileSync`,
      and `*.node` native binary modules). There are two possible values:

  - ul:
    - >
      `bundle`
      <p>
      This is the default value. It means that package imports are allowed to
      be bundled. Note that this value doesn't mean all packages will be
      bundled, just that they are allowed to be. You can still exclude
      individual packages from the bundle using [external](#external).
      </p>

    - >
      `external`
      <p>
      This means that all package imports considered external to the bundle,
      and are not bundled. _Note that your dependencies must still be present
      on the file system when your bundle is run._ It has the same effect as
      manually passing each dependency to [external](#external) but is more
      concise. If you want to customize which of your dependencies are external
      and which ones aren't, then you should set this to `bundle` instead and
      then use [external](#external) for individual dependencies.
      </p>
      <p>
      This setting considers all import paths that "look like" package imports
      in the original source code to be package imports. Specifically import paths
      that don't start with a path segment of `/` or `.` or `..` are considered
      to be package imports. The only two exceptions to this rule are
      [subpath imports](https://nodejs.org/api/packages.html#subpath-imports)
      (which start with a `#` character) and TypeScript path remappings via
      [`paths`](https://www.typescriptlang.org/tsconfig/#paths) and/or
      [`baseUrl`](https://www.typescriptlang.org/tsconfig/#baseUrl) in
      `tsconfig.json` (which are applied first).
      </p>

  - p: >
      Using it looks like this:

  - example:
      in:
        app.js: 'import "pkg"'

      cli: |
        esbuild app.js --bundle --packages=external

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          packages: 'external',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Bundle:      true,
            Packages:    api.PackagesExternal,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Note that this setting only has an effect when [bundling](#bundle) is
      enabled. Also note that marking an import path as external happens after
      the import path is rewritten by any configured [aliases](#alias), so the
      alias feature still has an effect when this setting is used.

  ## Preserve symlinks

  - p: >
      This setting mirrors the [`--preserve-symlinks`](https://nodejs.org/api/cli.html#cli_preserve_symlinks)
      setting in node. If you use that setting (or the similar [`resolve.symlinks`](https://webpack.js.org/configuration/resolve/#resolvesymlinks)
      setting in Webpack), you will likely need to enable this setting in
      esbuild too. It can be enabled like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --bundle --preserve-symlinks --outfile=out.js

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          preserveSymlinks: true,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:      []string{"app.js"},
            Bundle:           true,
            PreserveSymlinks: true,
            Outfile:          "out.js",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Enabling this setting causes esbuild to determine file identity by the
      original file path (i.e. the path without following symlinks) instead of
      the real file path (i.e. the path after following symlinks). This can be
      beneficial with certain directory structures. Keep in mind that this means
      a file may be given multiple identities if there are multiple symlinks
      pointing to it, which can result in it appearing multiple times in
      generated output files.

  - p: >
      _Note: The term "symlink" means [symbolic link](https://en.wikipedia.org/wiki/Symbolic_link)
      and refers to a file system feature where a path can redirect to another
      path._

  ## Resolve extensions

  - p: >
      The [resolution algorithm used by node](https://nodejs.org/api/modules.html#modules_file_modules)
      supports implicit file extensions. You can <code>require(<wbr>'./file')</code> and it
      will check for `./file`, `./file.js`, `./file.json`, and `./file.node` in
      that order. Modern bundlers including esbuild extend this concept to other
      file types as well. The full order of implicit file extensions in esbuild
      can be customized using the resolve extensions setting, which defaults to
      <code>.tsx,<wbr>.ts,<wbr>.jsx,<wbr>.js,<wbr>.css,<wbr>.json</code>:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
          esbuild app.js --bundle --resolve-extensions=.ts,.js

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          resolveExtensions: ['.ts', '.js'],
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:       []string{"app.js"},
            Bundle:            true,
            ResolveExtensions: []string{".ts", ".js"},
            Write:             true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Note that esbuild deliberately does not include the new `.mjs` and `.cjs`
      extensions in this list. Node's resolution algorithm doesn't treat these
      as implicit file extensions, so esbuild doesn't either. If you want to
      import files with these extensions you should either explicitly add the
      extensions in your import paths or change this setting to include the
      additional extensions that you want to be implicit.

  ## Working directory

  - p: >
      This API option lets you specify the working directory to use for the
      build. It normally defaults to the current [working directory](https://en.wikipedia.org/wiki/Working_directory)
      of the process you are using to call esbuild's API. The working directory
      is used by esbuild for a few different things including resolving relative
      paths given as API options to absolute paths and pretty-printing absolute
      paths as relative paths in log messages. Here is how to customize esbuild's
      working directory:

  - example:
      in:
        /var/tmp/custom/working/directory/file.js: 'export let foo = 123'

      cli: |
        cd "/var/tmp/custom/working/directory"

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['file.js'],
          absWorkingDir: '/var/tmp/custom/working/directory',
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:   []string{"file.js"},
            AbsWorkingDir: "/var/tmp/custom/working/directory",
            Outfile:       "out.js",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Note: If you are using [Yarn Plug'n'Play](https://yarnpkg.com/features/pnp/),
      keep in mind that this working directory is used to search for Yarn's
      manifest file. If you are running esbuild from an unrelated directory,
      you will have to set this working directory to the directory containing
      the manifest file (or one of its child directories) for the manifest
      file to be found by esbuild.

  # Transformation

  ## JSX

  - p: >
      This option tells esbuild what to do about JSX syntax. Here are the available options:

  - ul:
    - >
      `transform`
      <p>
      This tells esbuild to transform JSX to JS using a general-purpose
      transform that's shared between many libraries that use JSX syntax.
      Each JSX element is turned into a call to the [JSX factory](#jsx-factory)
      function with the element's component (or with the [JSX fragment](#jsx-fragment)
      for fragments) as the first argument. The second argument is an array of
      props (or `null` if there are no props). Any child elements present
      become additional arguments after the second argument.
      </p>
      <p>
      If you want to configure this setting on a per-file basis, you can do
      that by using a <code>// @jsxRuntime <wbr>classic</code> comment. This is
      a convention from [Babel's JSX plugin](https://babeljs.io/docs/en/babel-preset-react/)
      that esbuild follows.
      </p>

    - >
      `preserve`
      <p>
      This preserves the JSX syntax in the output instead of transforming it
      into function calls. JSX elements are treated as first-class syntax and
      are still affected by other settings such as [minification](#minify) and
      [property mangling](#mangle-props).
      </p>
      <p>
      Note that this means the output files are no longer valid JavaScript code.
      This feature is intended to be used when you want to transform the JSX
      syntax in esbuild's output files by another tool after bundling.
      </p>

    - >
      `automatic`
      <p>
      This transform was [introduced in React 17+](https://reactjs.org/blog/2020/09/22/introducing-the-new-jsx-transform.html)
      and is very specific to React. It automatically generates `import`
      statements from the [JSX import source](#jsx-import-source) and introduces
      many special cases regarding how the syntax is handled. The details are too
      complicated to describe here. For more information, please read
      [React's documentation about their new JSX transform](https://github.com/reactjs/rfcs/blob/createlement-rfc/text/0000-create-element-changes.md).
      If you want to enable the development mode version of this transform, you
      need to additionally enable the [JSX dev](#jsx-dev) setting.
      </p>
      <p>
      If you want to configure this setting on a per-file basis, you can do
      that by using a <code>// @jsxRuntime <wbr>automatic</code> comment. This is
      a convention from [Babel's JSX plugin](https://babeljs.io/docs/en/babel-preset-react/)
      that esbuild follows.
      </p>

  - p: >
      Here's an example of setting the JSX transform to `preserve`:

  - example:
      cli:
        - $: |
            echo '<div/>' | esbuild --jsx=preserve --loader=jsx
        - expect: |
            <div />;

      mjs: |
        import * as esbuild from 'esbuild'

        let result = await esbuild.transform('<div/>', {
          jsx: 'preserve',
          loader: 'jsx',
        })

        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          result := api.Transform("<div/>", api.TransformOptions{
            JSX:    api.JSXPreserve,
            Loader: api.LoaderJSX,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  ## JSX dev

  - p: >
      If the [JSX](#jsx) transform has been set to `automatic`, then enabling
      this setting causes esbuild to automatically inject the file name and
      source location into each JSX element. Your JSX library can then use
      this information to help with debugging. If the JSX transform has been
      set to something other than `automatic`, then this setting does nothing.
      Here's an example of enabling this setting:

  - example:
      in:
        app.jsx: '<a/>'

      cli:
        - $: |
            echo '<a/>' | esbuild --loader=jsx --jsx=automatic
        - expect: |
            import { jsx } from "react/jsx-runtime";
            /* @__PURE__ */ jsx("a", {});

        - $: |
            echo '<a/>' | esbuild --loader=jsx --jsx=automatic --jsx-dev
        - expect: |
            import { jsxDEV } from "react/jsx-dev-runtime";
            /* @__PURE__ */ jsxDEV("a", {}, void 0, false, {
              fileName: "<stdin>",
              lineNumber: 1,
              columnNumber: 1
            }, this);

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.jsx'],
          jsxDev: true,
          jsx: 'automatic',
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.jsx"},
            JSXDev:      true,
            JSX:         api.JSXAutomatic,
            Outfile:     "out.js",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## JSX factory

  - p: >
      This sets the function that is called for each JSX element. Normally a
      JSX expression such as this:

  - pre.xml: |
      <div>Example text</div>

  - p: >
      is compiled into a function call to `React.createElement` like this:

  - pre.js: |
      React.createElement("div", null, "Example text");

  - p: >
      You can call something other than `React.createElement` by changing the
      JSX factory. For example, to call the function `h` instead (which is
      used by other libraries such as [Preact](https://preactjs.com/)):

  - example:
      cli:
        - $: |
            echo '<div/>' | esbuild --jsx-factory=h --loader=jsx
        - expect: |
            /* @__PURE__ */ h("div", null);

      mjs: |
        import * as esbuild from 'esbuild'

        let result = await esbuild.transform('<div/>', {
          jsxFactory: 'h',
          loader: 'jsx',
        })

        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          result := api.Transform("<div/>", api.TransformOptions{
            JSXFactory: "h",
            Loader:     api.LoaderJSX,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      Alternatively, if you are using TypeScript, you can just configure JSX
      for TypeScript by adding this to your `tsconfig.json` file and esbuild
      should pick it up automatically without needing to be configured:

  - pre.json: |
      {
        "compilerOptions": {
          "jsxFactory": "h"
        }
      }

  - p: >
      If you want to configure this on a per-file basis, you can do that by
      using a <code>// @jsx <wbr>h</code> comment. Note that this setting
      does not apply when the [JSX](#jsx) transform has been set to `automatic`.

  ## JSX fragment

  - p: >
      This sets the function that is called for each JSX fragment. Normally a
      JSX fragment expression such as this:

  - pre.xml: |
      <>Stuff</>

  - p: >
      is compiled into a use of the `React.Fragment` component like this:

  - pre.js: |
      React.createElement(React.Fragment, null, "Stuff");

  - p: >
      You can use a component other than `React.Fragment` by changing the
      JSX fragment. For example, to use the component `Fragment` instead
      (which is used by other libraries such as [Preact](https://preactjs.com/)):

  - example:
      cli:
        - $: |
            echo '<>x</>' | esbuild --jsx-fragment=Fragment --loader=jsx
        - expect: |
            /* @__PURE__ */ React.createElement(Fragment, null, "x");

      mjs: |
        import * as esbuild from 'esbuild'

        let result = await esbuild.transform('<>x</>', {
          jsxFragment: 'Fragment',
          loader: 'jsx',
        })

        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          result := api.Transform("<>x</>", api.TransformOptions{
            JSXFragment: "Fragment",
            Loader:      api.LoaderJSX,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      Alternatively, if you are using TypeScript, you can just configure JSX
      for TypeScript by adding this to your `tsconfig.json` file and esbuild
      should pick it up automatically without needing to be configured:

  - pre.json: |
      {
        "compilerOptions": {
          "jsxFragmentFactory": "Fragment"
        }
      }

  - p: >
      If you want to configure this on a per-file basis, you can do that by
      using a <code>// @jsxFrag <wbr>Fragment</code> comment. Note that this
      setting does not apply when the [JSX](#jsx) transform has been set to
      `automatic`.

  ## JSX import source

  - p: >
      If the [JSX](#jsx) transform has been set to `automatic`, then setting
      this lets you change which library esbuild uses to automatically import
      its JSX helper functions from. Note that this only works with the JSX
      transform that's [specific to React 17+](https://reactjs.org/blog/2020/09/22/introducing-the-new-jsx-transform.html).
      If you set the JSX import source to `your-pkg`, then that package must
      expose at least the following exports:

  - pre.js: |
      import { createElement } from "your-pkg"
      import { Fragment, jsx, jsxs } from "your-pkg/jsx-runtime"
      import { Fragment, jsxDEV } from "your-pkg/jsx-dev-runtime"

  - p: >
      The `/jsx-runtime` and `/jsx-dev-runtime` subpaths are hard-coded by
      design and cannot be changed. The `jsx` and `jsxs` imports are used when
      [JSX dev mode](#jsx-dev) is off and the `jsxDEV` import is used when
      JSX dev mode is on. The meaning of these is described in
      [React's documentation about their new JSX transform](https://github.com/reactjs/rfcs/blob/createlement-rfc/text/0000-create-element-changes.md).
      The `createElement` import is used regardless of the JSX dev mode when an
      element has a prop spread followed by a `key` prop, which looks like this:

  - pre.jsx: |
      return <div {...props} key={key} />

  - p: >
      Here's an example of setting the JSX import source to [`preact`](https://preactjs.com/):

  - example:
      in:
        app.jsx: '<a/>'

      cli: |
        esbuild app.jsx --jsx-import-source=preact --jsx=automatic

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.jsx'],
          jsxImportSource: 'preact',
          jsx: 'automatic',
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:     []string{"app.jsx"},
            JSXImportSource: "preact",
            JSX:             api.JSXAutomatic,
            Outfile:         "out.js",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Alternatively, if you are using TypeScript, you can just configure the
      JSX import source for TypeScript by adding this to your `tsconfig.json`
      file and esbuild should pick it up automatically without needing to be
      configured:

  - pre.json: |
      {
        "compilerOptions": {
          "jsx": "react-jsx",
          "jsxImportSource": "preact"
        }
      }

  - p: >
      And if you want to control this setting on the per-file basis, you can do
      that with a <code>// @jsxImportSource <wbr>your-pkg</code> comment in each
      file. You may also need to add a <code>// @jsxRuntime <wbr>automatic</code>
      comment as well if the [JSX](#jsx) transform has not already been set by
      other means, or if you want that to be set on a per-file basis as well.

  ## JSX side effects

  - p: >
      By default esbuild assumes that JSX expressions are side-effect free,
      which means they are annoated with [`/* @__PURE__ */` comments](#pure) and
      are removed during bundling when they are unused. This follows the common
      use of JSX for virtual DOM and applies to the vast majority of JSX libraries.
      However, some people have written JSX libraries that don't have this property
      (specifically JSX expressions can have arbitrary side effects and can't be
      removed when unused). If you are using such a library, you can use this
      setting to tell esbuild that JSX expressions have side effects:

  - example:
      in:
        app.jsx: '<a/>'

      cli: |
        esbuild app.jsx --jsx-side-effects

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.jsx'],
          outfile: 'out.js',
          jsxSideEffects: true,
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:    []string{"app.jsx"},
            Outfile:        "out.js",
            JSXSideEffects: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## Supported

  - p: >
      This setting lets you customize esbuild's set of unsupported syntax
      features at the individual syntax feature level. For example, you can use
      this to tell esbuild that [BigInts](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/BigInt)
      are not supported so that esbuild generates an error when you try to use
      one. Usually this is configured for you when you use the [`target`](#target)
      setting, which you should typically be using instead of this setting.
      If the target is specified in addition to this setting, this setting will
      override whatever is specified by the target.

  - p: >
      Here are some examples of why you might want to use this setting instead
      of or in addition to setting the target:

  - ul:
    - >
      <p>
      JavaScript runtimes often do a quick implementation of newer syntax
      features that is slower than the equivalent older JavaScript, and you
      can get a speedup by telling esbuild to pretend this syntax feature isn't
      supported. For example, [V8](https://v8.dev/) has a [long-standing performance bug regarding object spread](https://bugs.chromium.org/p/v8/issues/detail?id=11536)
      that can be avoided by manually copying properties instead of using
      object spread syntax.
      </p>

    - >
      <p>
      There are many other JavaScript implementations in addition to the ones
      that esbuild's `target` setting recognizes, and they may not support certain
      features. If you are targeting such an implementation, you can use this setting
      to configure esbuild with a custom syntax feature compatibility set without
      needing to change esbuild itself. For example, [TypeScript's](https://www.typescriptlang.org/) JavaScript
      parser may not support [arbitrary module namespace identifier names](https://github.com/microsoft/TypeScript/issues/40594)
      so you may want to turn those off when targeting TypeScript's JavaScript
      parser.
      </p>

    - >
      <p>
      You may be processing esbuild's output with another tool, and you may want
      esbuild to transform certain features and the other tool to transform
      certain other features. For example, if you are using esbuild to transform
      files individually to ES5 but you are then feeding the output into [Webpack](https://webpack.js.org/)
      for bundling, you may want to preserve `import()` expressions even though
      they are a syntax error in ES5.
      </p>

  - p: >
      If you want esbuild to consider a certain syntax feature to be unsupported,
      you can specify that like this:

  - example:
      in:
        app.js: ''

      cli: |
        esbuild app.js --supported:bigint=false

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          supported: {
            'bigint': false,
          },
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Supported: map[string]bool{
              "bigint": false,
            },
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Syntax features are specified using esbuild-specific feature names. The
      full set of feature names is as follows:

      <p>**JavaScript:**</p>
      <ul>
      <li>`arbitrary-module-namespace-names`</li>
      <li>`array-spread`</li>
      <li>`arrow`</li>
      <li>`async-await`</li>
      <li>`async-generator`</li>
      <li>`bigint`</li>
      <li>`class`</li>
      <li>`class-field`</li>
      <li>`class-private-accessor`</li>
      <li>`class-private-brand-check`</li>
      <li>`class-private-field`</li>
      <li>`class-private-method`</li>
      <li>`class-private-static-accessor`</li>
      <li>`class-private-static-field`</li>
      <li>`class-private-static-method`</li>
      <li>`class-static-blocks`</li>
      <li>`class-static-field`</li>
      <li>`const-and-let`</li>
      <li>`decorators`</li>
      <li>`default-argument`</li>
      <li>`destructuring`</li>
      <li>`dynamic-import`</li>
      <li>`exponent-operator`</li>
      <li>`export-star-as`</li>
      <li>`for-await`</li>
      <li>`for-of`</li>
      <li>`function-name-configurable`</li>
      <li>`function-or-class-property-access`</li>
      <li>`generator`</li>
      <li>`hashbang`</li>
      <li>`import-assertions`</li>
      <li>`import-attributes`</li>
      <li>`import-meta`</li>
      <li>`inline-script`</li>
      <li>`logical-assignment`</li>
      <li>`nested-rest-binding`</li>
      <li>`new-target`</li>
      <li>`node-colon-prefix-import`</li>
      <li>`node-colon-prefix-require`</li>
      <li>`nullish-coalescing`</li>
      <li>`object-accessors`</li>
      <li>`object-extensions`</li>
      <li>`object-rest-spread`</li>
      <li>`optional-catch-binding`</li>
      <li>`optional-chain`</li>
      <li>`regexp-dot-all-flag`</li>
      <li>`regexp-lookbehind-assertions`</li>
      <li>`regexp-match-indices`</li>
      <li>`regexp-named-capture-groups`</li>
      <li>`regexp-set-notation`</li>
      <li>`regexp-sticky-and-unicode-flags`</li>
      <li>`regexp-unicode-property-escapes`</li>
      <li>`rest-argument`</li>
      <li>`template-literal`</li>
      <li>`top-level-await`</li>
      <li>`typeof-exotic-object-is-object`</li>
      <li>`unicode-escapes`</li>
      <li>`using`</li>
      </ul>

      <p>**CSS:**</p>
      <ul>
      <li>`color-functions`</li>
      <li>`gradient-double-position`</li>
      <li>`gradient-interpolation`</li>
      <li>`gradient-midpoints`</li>
      <li>`hwb`</li>
      <li>`hex-rgba`</li>
      <li>`inline-style`</li>
      <li>`inset-property`</li>
      <li>`is-pseudo-class`</li>
      <li>`modern-rgb-hsl`</li>
      <li>`nesting`</li>
      <li>`rebecca-purple`</li>
      </ul>

  ## Target

  - p: >
      This sets the target environment for the generated JavaScript and/or CSS
      code. It tells esbuild to transform JavaScript syntax that is too new for
      these environments into older JavaScript syntax that will work in these
      environments. For example, the `??` operator was introduced in Chrome 80
      so esbuild will convert it into an equivalent (but more verbose)
      conditional expression when targeting Chrome 79 or earlier.

  - p: >
      Note that this is only concerned with syntax features, not APIs. It does
      *not* automatically add [polyfills](https://developer.mozilla.org/en-US/docs/Glossary/Polyfill)
      for new APIs that are not used by these environments. You will have to
      explicitly import polyfills for the APIs you need (e.g. by importing
      [`core-js`](https://www.npmjs.com/package/core-js)). Automatic polyfill
      injection is outside of esbuild's scope.

  - p: >
      Each target environment is an environment name followed by a version
      number. The following environment names are currently supported:

  - ul:
    - '`chrome`'
    - '`deno`'
    - '`edge`'
    - '`firefox`'
    - '`hermes`'
    - '`ie`'
    - '`ios`'
    - '`node`'
    - '`opera`'
    - '`rhino`'
    - '`safari`'

  - p: >
      In addition, you can also specify JavaScript language versions such as
      `es2020`. The default target is `esnext` which means that by default,
      esbuild will assume all of the latest JavaScript and CSS features are
      supported. Here is an example that configures multiple target environments.
      You don't need to specify all of them; you can just specify the subset
      of target environments that your project cares about. You can also be more
      precise about version numbers if you'd like (e.g. `node12.19.0` instead of
      just `node12`):

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --target=es2020,chrome58,edge16,firefox57,node12,safari11

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          target: [
            'es2020',
            'chrome58',
            'edge16',
            'firefox57',
            'node12',
            'safari11',
          ],
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Target:      api.ES2020,
            Engines: []api.Engine{
              {Name: api.EngineChrome, Version: "58"},
              {Name: api.EngineEdge, Version: "16"},
              {Name: api.EngineFirefox, Version: "57"},
              {Name: api.EngineNode, Version: "12"},
              {Name: api.EngineSafari, Version: "11"},
            },
            Write: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      You can refer to the [JavaScript loader](/content-types/#javascript) for
      the details about which syntax features were introduced with which language
      versions. Keep in mind that while JavaScript language versions such as
      `es2020` are identified by year, that is the year the specification is
      approved. It has nothing to do with the year all major browsers implement
      that specification which often happens earlier or later than that year.

  - p: >
      If you use a syntax feature that esbuild doesn't yet have support for
      transforming to your current language target, esbuild will generate an
      error where the unsupported syntax is used. This is often the case when
      targeting the `es5` language version, for example, since esbuild only
      supports transforming most newer JavaScript syntax features to `es6`.

  - p: >
      If you need to customize the set of supported syntax features at the
      individual feature level in addition to or instead of what `target`
      provides, you can do that with the [`supported`](#supported) setting.

  # Optimization

  ## Define

  - p: >
      This feature provides a way to replace global identifiers with
      constant expressions. It can be a way to change the behavior some code
      between builds without changing the code itself:

  - example:
      cli:
        - $: |
            echo 'hooks = DEBUG && require("hooks")' | esbuild --define:DEBUG=true
        - expect: |
            hooks = require("hooks");

        - $: |
            echo 'hooks = DEBUG && require("hooks")' | esbuild --define:DEBUG=false
        - expect: |
            hooks = false;

      mjs:
        - $: import * as esbuild from 'esbuild'
        - $: let js = 'hooks = DEBUG && require("hooks")'

        - $: |
            (await esbuild.transform(js, {
              define: { DEBUG: 'true' },
            })).code
        - expect: |
            'hooks = require("hooks");\n'

        - $: |
            (await esbuild.transform(js, {
              define: { DEBUG: 'false' },
            })).code
        - expect: |
            'hooks = false;\n'

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "hooks = DEBUG && require('hooks')"

          result1 := api.Transform(js, api.TransformOptions{
            Define: map[string]string{"DEBUG": "true"},
          })

          if len(result1.Errors) == 0 {
            fmt.Printf("%s", result1.Code)
          }

          result2 := api.Transform(js, api.TransformOptions{
            Define: map[string]string{"DEBUG": "false"},
          })

          if len(result2.Errors) == 0 {
            fmt.Printf("%s", result2.Code)
          }
        }

  - p: >
      Each `define` entry maps an identifier to a string of code containing
      an expression. The expression in the string must either be a JSON object
      (null, boolean, number, string, array, or object) or a single identifier.
      Replacement expressions other than arrays and objects are substituted
      inline, which means that they can participate in constant folding. Array
      and object replacement expressions are stored in a variable and then
      referenced using an identifier instead of being substituted inline, which
      avoids substituting repeated copies of the value but means that the values
      don't participate in constant folding.

  - p: >
      If you want to replace something with a string literal, keep in mind that
      the replacement value passed to esbuild must itself contain quotes because
      each `define` entry maps to a string containing code. Omitting the quotes
      means the replacement value is an identifier instead. This is demonstrated
      in the example below:

  - example:
      cli:
        - $: |
            echo 'id, str' | esbuild --define:id=text --define:str=\"text\"
        - expect: |
            text, "text";

      mjs:
        - $: import * as esbuild from 'esbuild'
        - $: |
            (await esbuild.transform('id, str', {
              define: { id: 'text', str: '"text"' },
            })).code
        - expect: |
            'text, "text";\n'

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          result := api.Transform("id, text", api.TransformOptions{
            Define: map[string]string{
              "id":  "text",
              "str": "\"text\"",
            },
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      If you're using the CLI, keep in mind that different shells have different
      rules for how to escape double-quote characters (which are necessary when
      the replacement value is a string). Use a `\"` backslash escape because it
      works in both bash and Windows command prompt. Other methods of escaping
      double quotes that work in bash such as surrounding them with single
      quotes will not work on Windows, since Windows command prompt does not
      remove the single quotes. This is relevant when using the CLI from a npm
      script in your `package.json` file, which people will expect to work on
      all platforms:

  - pre.json: |
      {
        "scripts": {
          "build": "esbuild --define:process.env.NODE_ENV=\\\"production\\\" app.js"
        }
      }

  - p: >
      If you still run into cross-platform quote escaping issues with different
      shells, you will probably want to switch to using the [JavaScript API](/api/)
      instead. There you can use regular JavaScript syntax to eliminate
      cross-platform differences.

  - p: >
      If you're looking for a more advanced form of the define feature that can
      replace an expression with something other than a constant (e.g. replacing
      a global variable with a shim), you may be able to use the similar
      [inject](#inject) feature to do that.

  ## Drop

  - p: >
      This tells esbuild to edit your source code before building to drop
      certain constructs. There are currently two possible things that can be
      dropped:

  - ul:
    - >
      `debugger`
      <p>
      Passing this flag causes all [`debugger` statements](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/debugger)
      to be removed from the output. This is similar to the `drop_debugger: true`
      flag available in the popular [UglifyJS](https://github.com/mishoo/UglifyJS)
      and [Terser](https://github.com/terser/terser) JavaScript minifiers.
      </p>
      <p>
      JavaScript's `debugger` statements cause the active debugger to treat the
      statement as an automatically-configured breakpoint. Code containing this
      statement will automatically be paused when the debugger is open. If no
      debugger is open, the statement does nothing. Dropping these statements
      from your code just prevents the debugger from automatically stopping
      when your code runs.
      </p>
      <p>
      You can drop `debugger` statements like this:
      </p>

  - example:
      in:
        app.js: 'debugger'

      cli: |
        esbuild app.js --drop:debugger

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          drop: ['debugger'],
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Drop:        api.DropDebugger,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - ul:
    - >
      `console`
      <p>
      Passing this flag causes all [`console` API calls](https://developer.mozilla.org/en-US/docs/Web/API/console#methods)
      to be removed from the output. This is similar to the `drop_console: true`
      flag available in the popular [UglifyJS](https://github.com/mishoo/UglifyJS)
      and [Terser](https://github.com/terser/terser) JavaScript minifiers.
      </p>
      <p>
      WARNING: Using this flag can introduce bugs into your code! This flag
      removes the entire call expression including all call arguments. If any
      of those arguments had important side effects, using this flag will
      change the behavior of your code. Be very careful when using this flag.
      </p>
      <p>
      If you want to remove console API calls without removing the arguments
      with side effects (so you do not introduce bugs), you should mark the
      relevant API calls as [pure](#pure) instead. For example, you can mark
      `console.log` as pure using <code>--pure:<wbr>console.log</code>. This
      will cause these API calls to be removed safely when minification is
      enabled.
      </p>
      <p>
      You can drop `console` API calls like this:
      </p>

  - example:
      in:
        app.js: 'console.log()'

      cli: |
        esbuild app.js --drop:console

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          drop: ['console'],
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Drop:        api.DropConsole,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ## Drop labels

  - p: >
      This tells esbuild to edit your source code before building to drop
      [labeled statements](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/label)
      with specific label names. For example, consider the following code:

  - pre.js: |
      function example() {
        DEV: doAnExpensiveCheck()
        return normalCodePath()
      }

  - p: >
      If you use this option to drop all labels named `DEV`, then esbuild will
      give you this:

  - pre.js: |
      function example() {
        return normalCodePath();
      }

  - p: >
      You can configure this feature like this (which will drop both the `DEV`
      and `TEST` labels):

  - example:
      in:
        app.js: |
          DEV: dev()
          TEST: test()
          PROD: prod()

      cli: |
        esbuild app.js --drop-labels=DEV,TEST

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          dropLabels: ['DEV', 'TEST'],
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            DropLabels:  []string{"DEV", "TEST"},
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Note that this is not the only way to conditionally remove code. Another
      more common way is to use the [define](#define) feature to replace
      specific global variables with a boolean value. For example, consider the
      following code:

  - pre.js: |
      function example() {
        DEV && doAnExpensiveCheck()
        return normalCodePath()
      }

  - p: >
      If you define `DEV` to `false`, then esbuild will give you this:

  - pre.js: |
      function example() {
        return normalCodePath();
      }

  - p: >
      This is pretty much the same thing as using a label. However, an advantage
      of using a label instead of a global variable to conditionally remove code
      is that you don't have to worry about the global variable not being defined
      because someone forgot to configure esbuild to replace it with something.
      Some drawbacks of using the label approach are that it makes conditionally
      removing code when the label is *not* dropped slightly harder to read,
      and it doesn't work for code embedded within nested expressions. Which
      approach to use for a given project comes down to personal preference.

  ## Ignore annotations

  - p: >
      Since JavaScript is a dynamic language, identifying unused code is
      sometimes very difficult for a compiler, so the community has developed
      certain annotations to help tell compilers what code should be considered
      side-effect free and available for removal. Currently there are two forms
      of side-effect annotations that esbuild supports:

  - ul:
    - >
      <p>
      Inline `/* @__PURE__ */` comments before function calls tell esbuild that
      the function call can be removed if the resulting value isn't used. See
      the [pure](#pure) API option for more information.
      </p>

    - >
      <p>
      The `sideEffects` field in `package.json` can be used to tell esbuild
      which files in your package can be removed if all imports from that
      file end up being unused. This is a convention from Webpack and many
      libraries published to npm already have this field in their package
      definition. You can learn more about this field in
      [Webpack's documentation](https://webpack.js.org/guides/tree-shaking/)
      for this field.
      </p>

  - p: >
      These annotations can be problematic because the compiler depends
      completely on developers for accuracy, and developers occasionally
      publish packages with incorrect annotations. The `sideEffects` field is
      particularly error-prone for developers because by default it causes all
      files in your package to be considered dead code if no imports are used.
      If you add a new file containing side effects and forget to update that
      field, your package will likely break when people try to bundle it.

  - p: >
      This is why esbuild includes a way to ignore side-effect annotations.
      You should only enable this if you encounter a problem where the bundle
      is broken because necessary code was unexpectedly removed from the bundle:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
          esbuild app.js --bundle --ignore-annotations

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          bundle: true,
          ignoreAnnotations: true,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:       []string{"app.js"},
            Bundle:            true,
            IgnoreAnnotations: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Enabling this means esbuild will no longer respect `/* @__PURE__ */`
      comments or the `sideEffects` field. It will still do automatic
      [tree shaking](#tree-shaking) of unused imports, however, since that doesn't rely on
      annotations from developers. Ideally this flag is only a temporary
      workaround. You should report these issues to the maintainer of the
      package to get them fixed since they indicate a problem with the package
      and they will likely trip up other people too.

  ## Inject

  - p: >
      This option allows you to automatically replace a global variable with an
      import from another file. This can be a useful tool for adapting code that
      you don't control to a new environment. For example, assume you have a
      file called `process-cwd-shim.js` that exports a shim using the export
      name `process.cwd`:

  - pre.js: |
      // process-cwd-shim.js
      let processCwdShim = () => ''
      export { processCwdShim as 'process.cwd' }

  - pre.js: |
      // entry.js
      console.log(process.cwd())

  - p: >
      This is intended to replace uses of node's `process.cwd()` function to
      prevent packages that call it from crashing when run in the browser.
      You can use the inject feature to replace all references to the global
      property `process.cwd` with an import from that file:

  - example:
      in:
        process-cwd-shim.js: |
          let processCwdShim = () => ''
          export { processCwdShim as 'process.cwd' }
        entry.js: 'console.log(process.cwd())'

      cli: |
        esbuild entry.js --inject:./process-cwd-shim.js --outfile=out.js

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['entry.js'],
          inject: ['./process-cwd-shim.js'],
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"entry.js"},
            Inject:      []string{"./process-cwd-shim.js"},
            Outfile:     "out.js",
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      That results in something like this:

  - pre.js: |
      // out.js
      var processCwdShim = () => "";
      console.log(processCwdShim());

  - p: >
      You can think of the inject feature as similar to the [define](#define)
      feature, except it replaces an expression with an import to a file instead
      of with a constant, and the expression to replace is specified using an
      export name in a file instead of using an inline string in esbuild's API.

  ### Auto-import for [JSX](/content-types/#jsx)

  - p: >
      React (the library for which JSX syntax was originally created) has a
      mode they call `automatic` where you don't have to `import` anything to
      use JSX syntax. Instead, the JSX-to-JS transformer will automatically
      import the correct JSX factory function for you. You can enable
      `automatic` JSX mode with esbuild's [`jsx`](#jsx) setting. If you want
      auto-import for JSX and you are using a sufficiently new version of
      React, then you should be using the `automatic` JSX mode.

  - p: >
      However, setting `jsx` to `automatic` unfortunately also means you are
      using a highly React-specific JSX transform instead of the default
      general-purpose JSX transform. This means writing a JSX factory function
      is more complicated, and it also means that the `automatic` mode doesn't
      work with libraries that expect to be used with the standard JSX transform
      (including older versions of React).

  - p: >
      You can use esbuild's inject feature to automatically import the [factory](#jsx-factory)
      and [fragment](#jsx-fragment) for JSX expressions when the JSX transform
      is not set to `automatic`. Here's an example file that can be injected
      to do this:

  - pre.js: |
      const { createElement, Fragment } = require('react')
      export {
        createElement as 'React.createElement',
        Fragment as 'React.Fragment',
      }

  - p: >
      This code uses the React library as an example, but you can use this
      approach with any other JSX library as well with appropriate changes.

  ### Injecting files without imports

  - p: >
      You can also use this feature with files that have no exports. In that
      case the injected file just comes first before the rest of the output
      as if every input file contained <code>import <wbr>"./file.js"</code>.
      Because of the way ECMAScript modules work, this injection is still
      "hygienic" in that symbols with the same name in different files are
      renamed so they don't collide with each other.

  ### Conditionally injecting a file

  - p: >
      If you want to _conditionally_ import a file only if the export is
      actually used, you should mark the injected file as not having side
      effects by putting it in a package and adding <code>"sideEffects": <wbr>false</code>
      in that package's `package.json` file. This setting is a
      [convention from Webpack](https://webpack.js.org/guides/tree-shaking/#mark-the-file-as-side-effect-free)
      that esbuild respects for any imported file, not just files used with inject.

  ## Keep names

  - p: >
      In JavaScript the `name` property on functions and classes defaults to a
      nearby identifier in the source code. These syntax forms all set the `name`
      property of the function to `"fn"`:

  - pre.js: |
      function fn() {}
      let fn = function() {};
      fn = function() {};
      let [fn = function() {}] = [];
      let {fn = function() {}} = {};
      [fn = function() {}] = [];
      ({fn = function() {}} = {});

  - p: >
      However, [minification](#minify) renames symbols to reduce code size and
      [bundling](#bundle) sometimes need to rename symbols to avoid collisions.
      That changes value of the `name` property for many of these cases. This
      is usually fine because the `name` property is normally only used for
      debugging. However, some frameworks rely on the `name` property for
      registration and binding purposes. If this is the case, you can enable
      this option to preserve the original `name` values even in minified code:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --minify --keep-names

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          minify: true,
          keepNames: true,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:       []string{"app.js"},
            MinifyWhitespace:  true,
            MinifyIdentifiers: true,
            MinifySyntax:      true,
            KeepNames:         true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Note that this feature is unavailable if the [target](#target) has been
      set to an old environment that doesn't allow esbuild to mutate the `name`
      property on functions and classes. This is the case for environments that
      don't support ES6.

  ## Mangle props

  - warning: >
      **Using this feature can break your code in subtle ways.** Do not use this
      feature unless you know what you are doing, and you know exactly how it
      will affect both your code and all of your dependencies.

  - p: >
      This setting lets you pass a regular expression to esbuild to tell
      esbuild to automatically rename all properties that match this regular
      expression. It's useful when you want to minify certain property names
      in your code either to make the generated code smaller or to somewhat
      obfuscate your code's intent.

  - p: >
      Here's an example that uses the regular expression `_$` to mangle all
      properties ending in an underscore, such as `foo_`. This mangles
      <code>print({ <wbr>foo_: 0 <wbr>}.foo_)</code>
      into <code>print({ <wbr>a: 0 <wbr>}.a)</code>:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --mangle-props=_$

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          mangleProps: /_$/,
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            MangleProps: "_$",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Only mangling properties that end in an underscore is a reasonable
      heuristic because normal JS code doesn't typically contain identifiers
      like that. Browser APIs also don't use this naming convention so this
      also avoids conflicts with browser APIs. If you want to avoid mangling
      names such as [`__defineGetter__`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/__defineGetter__)
      you could consider using a more complex regular expression such as
      `[^_]_$` (i.e. must end in a non-underscore followed by an underscore).

  - p: >
      This is a separate setting instead of being part of the [minify](#minify)
      setting because it's an unsafe transformation that does not work on
      arbitrary JavaScript code. It only works if the provided regular
      expression matches all of the properties that you want mangled and does
      not match any of the properties that you don't want mangled. It also only
      works if you do not under any circumstances reference a mangled property
      indirectly. For example, it means you can't use `obj[prop]` to reference a
      property where `prop` is a string containing the property name. Specifically
      the following syntax constructs are the only ones eligible for property mangling:

  - table: |
      | Syntax                           | Example                                   |
      |----------------------------------|-------------------------------------------|
      | Dot property accesses            | `x.foo_`                                  |
      | Dot optional chains              | `x?.foo_`                                 |
      | Object properties                | `x = { foo_: y }`                         |
      | Object methods                   | `x = { foo_() {} }`                       |
      | Class fields                     | `class x { foo_ = y }`                    |
      | Class methods                    | `class x { foo_() {} }`                   |
      | Object destructuring bindings    | `let { foo_: x } = y`                     |
      | Object destructuring assignments | `({ foo_: x } = y)`                       |
      | JSX element member expression    | `<X.foo_></X.foo_>`                       |
      | JSX attribute names              | `<X foo_={y} />`                          |
      | TypeScript namespace exports     | `namespace x { export let foo_ = y }`     |
      | TypeScript parameter properties  | `class x { constructor(public foo_) {} }` |

  - p: >
      When using this feature, keep in mind that property names are only
      consistently mangled within a single esbuild API call but not across
      esbuild API calls. Each esbuild API call does an independent property
      mangling operation so output files generated by two different API calls
      may mangle the same property to two different names, which could cause
      the resulting code to behave incorrectly.

  - h4#mangle-quoted: >
      Quoted properties

  - p: >
      By default, esbuild doesn't modify the contents of string literals. This means
      you can avoid property mangling for an individual property by quoting it
      as a string. However, you must consistently use quotes or no quotes for
      a given property everywhere for this to work. For example,
      <code>print({ <wbr>foo_: 0 <wbr>}.foo_)</code> will be mangled into
      <code>print({ <wbr>a: 0 <wbr>}.a)</code> while
      <code>print({ <wbr>'foo_': 0 <wbr>}['foo_'])</code> will not be mangled.

  - p: >
      If you would like for esbuild to also mangle the contents of string literals,
      you can explicitly enable that behavior like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --mangle-props=_$ --mangle-quoted

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          mangleProps: /_$/,
          mangleQuoted: true,
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:  []string{"app.js"},
            MangleProps:  "_$",
            MangleQuoted: api.MangleQuotedTrue,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Enabling this makes the following syntax constructs also eligible for property
      mangling:

  - table: |
      | Syntax                                  | Example                   |
      |-----------------------------------------|---------------------------|
      | Quoted property accesses                | `x['foo_']`               |
      | Quoted optional chains                  | `x?.['foo_']`             |
      | Quoted object properties                | `x = { 'foo_': y }`       |
      | Quoted object methods                   | `x = { 'foo_'() {} }`     |
      | Quoted class fields                     | `class x { 'foo_' = y }`  |
      | Quoted class methods                    | `class x { 'foo_'() {} }` |
      | Quoted object destructuring bindings    | `let { 'foo_': x } = y`   |
      | Quoted object destructuring assignments | `({ 'foo_': x } = y)`     |
      | String literals to the left of `in`     | `'foo_' in x`             |

  - h4#mangle-key: >
      Mangling other strings

  - p: >
      Mangling [quoted properties](#mangle-quoted) still only mangles strings
      in property name position. Sometimes you may also need to mangle property
      names in strings at arbitrary other locations in your code. To do that,
      you can prefix the string with a `/* @__KEY__ */` comment to tell esbuild
      that the contents of a string should be treated as a property name that
      can be mangled. For example:

  - pre.js: |
      let obj = {}
      Object.defineProperty(
        obj,
        /* @__KEY__ */ 'foo_',
        { get: () => 123 },
      )
      console.log(obj.foo_)

  - p: >
      This will cause the contents of the string `'foo_'` to be mangled as a
      property name (assuming [property mangling](#mangle-props) is enabled
      and `foo_` is eligible for renaming). The `/* @__KEY__ */` comment is a
      convention from [Terser](https://github.com/terser/terser), a popular
      JavaScript minifier with a similar property mangling feature.

  - h4#reserve-props: >
      Preventing renaming

  - p: >
      If you would like to exclude certain properties from mangling, you can
      reserve them with an additional setting. For example, this uses the
      regular expression `^__.*__$` to reserve all properties that start and
      end with two underscores, such as `__foo__`:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --mangle-props=_$ "--reserve-props=^__.*__$"

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          mangleProps: /_$/,
          reserveProps: /^__.*__$/,
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints:  []string{"app.js"},
            MangleProps:  "_$",
            ReserveProps: "^__.*__$",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - h4#mangle-cache: >
      Persisting renaming decisions

  - p: >
      Advanced usage of the property mangling feature involves storing the
      mapping from original name to mangled name in a persistent cache. When
      enabled, all mangled property renamings are recorded in the cache during
      the initial build. Subsequent builds reuse the renamings stored in the
      cache and add additional renamings for any newly-added properties. This
      has a few consequences:

  - ul:
    - >
      <p>
      You can customize what mangled properties are renamed to by editing the
      cache before passing it to esbuild.
      </p>

    - >
      <p>
      The cache serves as a list of all properties that were mangled. You can
      easily scan it to see if there are any unexpected property renamings.
      </p>

    - >
      <p>
      You can disable mangling for individual properties by setting the renamed
      value to `false` instead of to a string. This is similar to the [reserve props](#reserve-props)
      setting but on a per-property basis.
      </p>

    - >
      <p>
      You can ensure consistent renaming between builds (e.g. a main-thread
      file and a web worker, or a library and a plugin). Without this feature,
      each build would do an independent renaming operation and the mangled
      property names likely wouldn't be consistent.
      </p>

  - p: >
      For example, consider the following input file:

  - pre.js: |
      console.log({
        someProp_: 1,
        customRenaming_: 2,
        disabledRenaming_: 3
      });

  - p: >
      If we want `customRenaming_` to be renamed to `cR_` and we don't want
      `disabledRenaming_` to be renamed at all, we can pass the following
      mangle cache JSON to esbuild:

  - pre.json: |
      {
        "customRenaming_": "cR_",
        "disabledRenaming_": false
      }

  - p: >
      The mangle cache JSON can be passed to esbuild like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --mangle-props=_$ --mangle-cache=cache.json

      mjs: |
        import * as esbuild from 'esbuild'

        let result = await esbuild.build({
          entryPoints: ['app.js'],
          mangleProps: /_$/,
          mangleCache: {
            customRenaming_: "cR_",
            disabledRenaming_: false
          },
        })

        console.log('updated mangle cache:', result.mangleCache)

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            MangleProps: "_$",
            MangleCache: map[string]interface{}{
              "customRenaming_":   "cR_",
              "disabledRenaming_": false,
            },
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }

          fmt.Println("updated mangle cache:", result.MangleCache)
        }

  - p: >
      When property naming is enabled, that will result in the following
      output file:

  - pre.js: |
      console.log({
        a: 1,
        cR_: 2,
        disabledRenaming_: 3
      });

  - p: >
      And the following updated mangle cache:

  - pre.json: |
      {
        "customRenaming_": "cR_",
        "disabledRenaming_": false,
        "someProp_": "a"
      }

  ## Minify

  - p: >
      When enabled, the generated code will be minified instead of
      pretty-printed. Minified code is generally equivalent to non-minified
      code but is smaller, which means it downloads faster but is harder to
      debug. Usually you minify code in production but not in development.

  - p: >
      Enabling minification in esbuild looks like this:

  - example:
      cli:
        - $: |
            echo 'fn = obj => { return obj.x }' | esbuild --minify
        - expect: |
            fn=n=>n.x;

      mjs:
        - $: import * as esbuild from 'esbuild'
        - $: |
            var js = 'fn = obj => { return obj.x }'
        - $: |
            (await esbuild.transform(js, {
              minify: true,
            })).code
        - expect: |
            'fn=n=>n.x;\n'

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "fn = obj => { return obj.x }"

          result := api.Transform(js, api.TransformOptions{
            MinifyWhitespace:  true,
            MinifyIdentifiers: true,
            MinifySyntax:      true,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      This option does three separate things in combination: it removes
      whitespace, it rewrites your syntax to be more compact, and it renames
      local variables to be shorter. Usually you want to do all of these
      things, but these options can also be enabled individually if necessary:

  - example:
      cli:
        - $: |
            echo 'fn = obj => { return obj.x }' | esbuild --minify-whitespace
        - expect: |
            fn=obj=>{return obj.x};

        - $: |
            echo 'fn = obj => { return obj.x }' | esbuild --minify-identifiers
        - expect: |
            fn = (n) => {
              return n.x;
            };

        - $: |
            echo 'fn = obj => { return obj.x }' | esbuild --minify-syntax
        - expect: |
            fn = (obj) => obj.x;

      mjs:
        - $: import * as esbuild from 'esbuild'

        - $: |
            var js = 'fn = obj => { return obj.x }'

        - $: |
            (await esbuild.transform(js, {
              minifyWhitespace: true,
            })).code
        - expect: |
            'fn=obj=>{return obj.x};\n'

        - $: |
            (await esbuild.transform(js, {
              minifyIdentifiers: true,
            })).code
        - expect: |
            'fn = (n) => {\n  return n.x;\n};\n'

        - $: |
            (await esbuild.transform(js, {
              minifySyntax: true,
            })).code
        - expect: |
            'fn = (obj) => obj.x;\n'

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          css := "div { color: yellow }"

          result1 := api.Transform(css, api.TransformOptions{
            Loader:           api.LoaderCSS,
            MinifyWhitespace: true,
          })

          if len(result1.Errors) == 0 {
            fmt.Printf("%s", result1.Code)
          }

          result2 := api.Transform(css, api.TransformOptions{
            Loader:            api.LoaderCSS,
            MinifyIdentifiers: true,
          })

          if len(result2.Errors) == 0 {
            fmt.Printf("%s", result2.Code)
          }

          result3 := api.Transform(css, api.TransformOptions{
            Loader:       api.LoaderCSS,
            MinifySyntax: true,
          })

          if len(result3.Errors) == 0 {
            fmt.Printf("%s", result3.Code)
          }
        }

  - p: >
      These same concepts also apply to CSS, not just to JavaScript:

  - example:
      cli:
        - $: |
            echo 'div { color: yellow }' | esbuild --loader=css --minify
        - expect: |
            div{color:#ff0}

      mjs:
        - $: import * as esbuild from 'esbuild'
        - $: |
            var css = 'div { color: yellow }'
        - $: |
            (await esbuild.transform(css, {
              loader: 'css',
              minify: true,
            })).code
        - expect: |
            'div{color:#ff0}\n'

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          css := "div { color: yellow }"

          result := api.Transform(css, api.TransformOptions{
            Loader:            api.LoaderCSS,
            MinifyWhitespace:  true,
            MinifyIdentifiers: true,
            MinifySyntax:      true,
          })

          if len(result.Errors) == 0 {
            fmt.Printf("%s", result.Code)
          }
        }

  - p: >
      The JavaScript minification algorithm in esbuild usually generates output
      that is very close to the minified output size of industry-standard
      JavaScript minification tools. [This benchmark](https://github.com/privatenumber/minification-benchmarks#readme)
      has an example comparison of output sizes between different minifiers.
      While esbuild is not the optimal JavaScript minifier in all cases (and
      doesn't try to be), it strives to generate minified output within a few
      percent of the size of dedicated minification tools for most code, and
      of course to do so much faster than other tools.

  - h4#minify-considerations: Considerations

  - p: >
      Here are some things to keep in mind when using esbuild as a minifier:

  - ul:
    - >
      <p>
      You should probably also set the [target](#target) option when
      minification is enabled. By default esbuild takes advantage of modern
      JavaScript features to make your code smaller. For example,
      <code>a ===<wbr> undefined<wbr> || a ===<wbr> null<wbr> ? 1 : a</code>
      could be minified to <code>a ?? 1</code>. If you do not want esbuild
      to take advantage of modern JavaScript features when minifying, you
      should use an older language target such as <code>--target=es6</code>.
      </p>

    - >
      <p>
      The character escape sequence `\n` will be replaced with a newline
      character in JavaScript template literals. String literals will also be
      converted into template literals if the [target](#target) supports them
      and if doing so would result in smaller output. **This is not a bug.**
      Minification means you are asking for smaller output, and the escape
      sequence `\n` takes two bytes while the newline character takes one byte.
      You can read more about this in the [FAQ entry on this topic](/faq/#minified-newlines).
      </p>

    - >
      <p>
      By default esbuild won't minify the names of top-level declarations. This
      is because esbuild doesn't know what you will be doing with the output.
      You might be injecting the minified code into the middle of some other
      code, in which case minifying top-level declaration names would be unsafe.
      Setting an output [format](#format) (or enabling [bundling](#bundle),
      which picks an output format for you if you haven't set one) tells
      esbuild that the output will be run within its own scope, which means
      it's then safe to minify top-level declaration names.
      </p>

    - >
      <p>
      Minification is not safe for 100% of all JavaScript code. This is true
      for esbuild as well as for other popular JavaScript minifiers such as
      [terser](https://github.com/terser/terser). In particular, esbuild is
      not designed to preserve the value of calling `.toString()` on a
      function. The reason for this is because if all code inside all
      functions had to be preserved verbatim, minification would hardly do
      anything at all and would be virtually useless. However, this means that
      JavaScript code relying on the return value of `.toString()` will likely
      break when minified. For example, some patterns in the [AngularJS](https://angularjs.org/)
      framework break when code is minified because AngularJS uses `.toString()`
      to read the argument names of functions. A workaround is to use [explicit
      annotations instead](https://docs.angularjs.org/api/auto/service/$injector#injection-function-annotation).
      </p>

    - >
      <p>
      By default esbuild does not preserve the value of `.name` on function
      and class objects. This is because most code doesn't rely on this property
      and using shorter names is an important size optimization. However, some
      code does rely on the `.name` property for registration and binding
      purposes. If you need to rely on this you should enable the
      [keep names](#keep-names) option.
      </p>

    - >
      <p>
      The minifier assumes that built-in JavaScript features behave the way they
      are expected to behave. These assumptions help esbuild generate more
      compact code. If you want a JavaScript minifier that doesn't make any
      assumptions about the behavior of built-in JavaScript features, then
      esbuild may not be the right JavaScript minifier for you. Here are some
      examples of these kinds of assumptions (note that this is not an exhaustive
      list):
      </p>
      <ul>
        <li><p>
        It's expected that `Array.prototype.join` behaves [as specified](https://tc39.es/ecma262/#sec-array.prototype.join).
        This means it's safe for esbuild's minifier to transform <code>x = [<wbr>1, <wbr>2, <wbr>3] +<wbr> ''</code>
        into `x="1,2,3";`.
        </p></li>
        <li><p>
        Accessing the `log` property on the `console` global is expected to not
        have any side effects. This means it's safe for esbuild's minifier to
        transform <code>var a, b =<wbr> a ?<wbr> console.<wbr>log(<wbr>x) :<wbr> console.<wbr>log(<wbr>y);</code>
        into <code>var a,b=<wbr>console.<wbr>log(<wbr>a?x:y);</code> (i.e. esbuild
        is assuming evaluating `console.log` can't change the value of `a`).
        </p></li>
      </ul>

    - |
      <p>
      Use of certain JavaScript features can disable many of esbuild's
      optimizations including minification. Specifically, using direct `eval`
      and/or the `with` statement prevent esbuild from renaming identifiers
      to smaller names since these features cause identifier binding to happen
      at run time instead of compile time. This is almost always unintentional,
      and only happens because people are unaware of what direct `eval` is and
      why it's bad.
      </p>
      <p>
      If you are thinking about writing some code like this:
      </p>
      <pre>
      // Direct eval (will disable minification for the whole file)
      let result = eval(something)
      </pre>
      <p>
      You should probably write your code like this instead so your code can be minified:
      </p>
      <pre>
      // Indirect eval (has no effect on the surrounding code)
      let result = (0, eval)(something)
      </pre>
      <p>
      There is more information about the consequences of direct `eval` and the
      available alternatives [here](/content-types/#direct-eval).
      </p>

    - >
      <p>
      The minification algorithm in esbuild does not yet do advanced code
      optimizations. In particular, the following code optimizations are
      possible for JavaScript code but are not done by esbuild (not an
      exhaustive list):
      </p>
      <ul>
        <li>Dead-code elimination within function bodies</li>
        <li>Function inlining</li>
        <li>Cross-statement constant propagation</li>
        <li>Object shape modeling</li>
        <li>Allocation sinking</li>
        <li>Method devirtualization</li>
        <li>Symbolic execution</li>
        <li>JSX expression hoisting</li>
        <li>TypeScript enum detection and inlining</li>
      </ul>
      <p>
      If your code makes use of patterns that require some of these forms of
      code optimization to be compact, or if you are searching for the optimal
      JavaScript minification algorithm for your use case, you should consider
      using other tools. Some examples of tools that implement some of these
      advanced code optimizations include [Terser](https://github.com/terser/terser#readme)
      and [Google Closure Compiler](https://github.com/google/closure-compiler#readme).
      </p>

  ## Pure

  - p: >
      There is a convention used by various JavaScript tools where a special
      comment containing either `/* @__PURE__ */` or `/* #__PURE__ */` before
      a new or call expression means that that expression can be removed if the
      resulting value is unused. It looks like this:

  - pre.js: |
      let button = /* @__PURE__ */ React.createElement(Button, null);

  - p: >
      This information is used by bundlers such as esbuild during tree shaking
      (a.k.a. dead code removal) to perform fine-grained removal of unused
      imports across module boundaries in situations where the bundler is not
      able to prove by itself that the removal is safe due to the dynamic
      nature of JavaScript code.

  - p: >
      Note that while the comment says "pure", it confusingly does _not_
      indicate that the function being called is pure. For example, it does not
      indicate that it is ok to cache repeated calls to that function. The name
      is essentially just an abstract shorthand for "ok to be removed if unused".

  - p: >
      Some expressions such as JSX and certain built-in globals are automatically
      annotated as `/* @__PURE__ */` in esbuild. You can also configure additional
      globals to be marked `/* @__PURE__ */` as well. For example, you can mark
      the global <code>document.<wbr>createElement</code> function as such to have
      it be automatically removed from your bundle when the bundle is minified as
      long as the result isn't used.

  - p: >
      It's worth mentioning that the effect of the annotation only extends to
      the call itself, not to the arguments. Arguments with side effects are
      still kept even when minification is enabled:

  - example:
      cli:
        - $: |
            echo 'document.createElement(elemName())' | esbuild --pure:document.createElement
        - expect: |
            /* @__PURE__ */ document.createElement(elemName());
        - $: |
            echo 'document.createElement(elemName())' | esbuild --pure:document.createElement --minify
        - expect: |
            elemName();

      mjs:
        - $: import * as esbuild from 'esbuild'
        - $: |
            let js = 'document.createElement(elemName())'
        - $: |
            (await esbuild.transform(js, {
              pure: ['document.createElement'],
            })).code
        - expect: |
            '/* @__PURE__ */ document.createElement(elemName());\n'
        - $: |
            (await esbuild.transform(js, {
              pure: ['document.createElement'],
              minify: true,
            })).code
        - expect: |
            'elemName();\n'

      go: |
        package main

        import "fmt"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js := "document.createElement(elemName())"

          result1 := api.Transform(js, api.TransformOptions{
            Pure: []string{"document.createElement"},
          })

          if len(result1.Errors) == 0 {
            fmt.Printf("%s", result1.Code)
          }

          result2 := api.Transform(js, api.TransformOptions{
            Pure:         []string{"document.createElement"},
            MinifySyntax: true,
          })

          if len(result2.Errors) == 0 {
            fmt.Printf("%s", result2.Code)
          }
        }

  - p: >
      Note that if you are trying to remove all calls to `console` API methods
      such as `console.log` and also want to remove the evaluation of arguments
      with side effects, there is a special case available for this: you can use
      the [drop feature](#drop) instead of marking `console` API calls as pure.
      However, this mechanism is specific to the `console` API and doesn't work
      with other call expressions.

  ## Tree shaking

  - p: >
      Tree shaking is the term the JavaScript community uses for dead code
      elimination, a common compiler optimization that automatically removes
      unreachable code. Within esbuild, this term specifically refers to
      declaration-level dead code removal.

  - p: >
      Tree shaking is easiest to explain with an example. Consider the following
      file. There is one used function and one unused function:

  - pre.js: |
      // input.js
      function one() {
        console.log('one')
      }
      function two() {
        console.log('two')
      }
      one()

  - p: >
      If you bundle this file with <code>esbuild <wbr>--bundle <wbr>input.js <wbr>--outfile=<wbr>output.js</code>,
      the unused function will automatically be discarded leaving you with the
      following output:

  - pre.js: |
      // input.js
      function one() {
        console.log("one");
      }
      one();

  - p: >
      This even works if we split our functions off into a separate library
      file and import them using an `import` statement:

  - pre.js: |
      // lib.js
      export function one() {
        console.log('one')
      }
      export function two() {
        console.log('two')
      }

  - pre.js: |
      // input.js
      import * as lib from './lib.js'
      lib.one()

  - p: >
      If you bundle this file with <code>esbuild <wbr>--bundle <wbr>input.js <wbr>--outfile=<wbr>output.js</code>,
      the unused function and unused import will still be automatically
      discarded leaving you with the following output:

  - pre.js: |
      // lib.js
      function one() {
        console.log("one");
      }

      // input.js
      one();

  - p: >
      This way esbuild will only bundle the parts of your packages that you
      actually use, which can sometimes be a substantial size savings. Note
      that esbuild's tree shaking implementation relies on the use of ECMAScript
      module `import` and `export` statements. It does not work with CommonJS
      modules. Many packages on npm include both formats and esbuild tries to
      pick the format that works with tree shaking by default. You can
      customize which format esbuild picks using the [main fields](#main-fields)
      and/or [conditions](#conditions) options depending on the package.

  - p: >
      By default, tree shaking is only enabled either when [bundling](#bundle)
      is enabled or when the output [format](#format) is set to `iife`, otherwise
      tree shaking is disabled. You can force-enable tree shaking by setting it
      to `true`:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
          esbuild app.js --tree-shaking=true

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          treeShaking: true,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            TreeShaking: api.TreeShakingTrue,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      You can also force-disable tree shaking by setting it to `false`:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
          esbuild app.js --tree-shaking=false

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          treeShaking: false,
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            TreeShaking: api.TreeShakingFalse,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  ### Tree shaking and side effects

  - p: >
      The side effect detection used for tree shaking is conservative, meaning
      that esbuild only considers code removable as dead code if it can be sure
      that there are no hidden side effects. For example, primitive literals
      such as `12.34` and `"abcd"` are side-effect free and can be removed while
      expressions such as `"ab" + cd` and `foo.bar` are not side-effect free
      (joining strings invokes `toString()` which can have side effects, and
      member access can invoke a getter which can also have side effects). Even
      referencing a global identifier is considered to be a side effect because
      it will throw a `ReferenceError` if there is no global with that name.
      Here's an example:

  - pre.js: |
      // These are considered side-effect free
      let a = 12.34;
      let b = "abcd";
      let c = { a: a };

      // These are not considered side-effect free
      // since they could cause some code to run
      let x = "ab" + cd;
      let y = foo.bar;
      let z = { [x]: x };

  - p: >
      Sometimes it's desirable to allow some code to be tree shaken even if that
      code can't be automatically determined to have no side effects. This can
      be done with a [pure annotation comment](#pure) which tells esbuild to
      trust the author of the code that there are no side effects within the
      annotated code. The annotation comment is `/* @__PURE__ */` and can only
      precede a new or call expression. You can annotate an immediately-invoked
      function expression and put arbitrary side effects inside the function
      body:

  - pre.js: |
      // This is considered side-effect free due to
      // the annotation, and will be removed if unused
      let gammaTable = /* @__PURE__ */ (() => {
        // Side-effect detection is skipped in here
        let table = new Uint8Array(256);
        for (let i = 0; i < 256; i++)
          table[i] = Math.pow(i / 255, 2.2) * 255;
        return table;
      })();

  - p: >
      While the fact that `/* @__PURE__ */` only works on call expressions
      can sometimes make code more verbose, a big benefit of this syntax is
      that it's portable across many other tools in the JavaScript ecosystem
      including the popular [UglifyJS](https://github.com/mishoo/uglifyjs) and
      [Terser](https://github.com/terser/terser) JavaScript minifiers (which are
      used by other major tools including [Webpack](https://github.com/webpack/webpack)
      and [Parcel](https://github.com/parcel-bundler/parcel)).

  - p: >
      Note that the annotations cause esbuild to assume that the annotated
      code is side-effect free. If the annotations are wrong and the code
      actually does have important side effects, these annotations can result
      in broken code. If you are bundling third-party code with annotations
      that have been authored incorrectly, you may need to enable
      [ignoring annotations](#ignore-annotations) to make sure the bundled
      code is correct.

# Source maps

## Source root

  - p: >
      This feature is only relevant when [source maps](#sourcemap) are enabled.
      It lets you set the value of the `sourceRoot` field in the source map,
      which specifies the path that all other paths in the source map are
      relative to. If this field is not present, all paths in the source map
      are interpreted as being relative to the directory containing the source
      map instead.

  - p: >
      You can configure `sourceRoot` like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        esbuild app.js --sourcemap --source-root=https://raw.githubusercontent.com/some/repo/v1.2.3/

      mjs: |
        import * as esbuild from 'esbuild'

        await esbuild.build({
          entryPoints: ['app.js'],
          sourcemap: true,
          sourceRoot: 'https://raw.githubusercontent.com/some/repo/v1.2.3/',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.js"},
            Sourcemap:   api.SourceMapInline,
            SourceRoot:  "https://raw.githubusercontent.com/some/repo/v1.2.3/",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

## Sourcefile

  - p: >
      This option sets the file name when using an input which has no file
      name. This happens when using the transform API and when using the build
      API with stdin. The configured file name is reflected in error messages
      and in source maps. If it's not configured, the file name defaults to
      `<stdin>`. It can be configured like this:

  - example:
      in:
        app.js: '1 + 2'

      cli: |
        cat app.js | esbuild --sourcefile=example.js --sourcemap

      mjs: |
        import * as esbuild from 'esbuild'
        import fs from 'node:fs'

        let js = fs.readFileSync('app.js', 'utf8')
        let result = await esbuild.transform(js, {
          sourcefile: 'example.js',
          sourcemap: 'inline',
        })

        console.log(result.code)

      go: |
        package main

        import "fmt"
        import "io/ioutil"
        import "github.com/evanw/esbuild/pkg/api"

        func main() {
          js, err := ioutil.ReadFile("app.js")
          if err != nil {
            panic(err)
          }

          result := api.Transform(string(js),
            api.TransformOptions{
              Sourcefile: "example.js",
              Sourcemap:  api.SourceMapInline,
            })

          if len(result.Errors) == 0 {
            fmt.Printf("%s %s", result.Code)
          }
        }

## Sourcemap

* allows
  * making easier to debug your code
    * Reason: 🧠encode the information necessary to translate from a line/column offset | generated output file back -- to a -- line/column offset | corresponding original input file 🧠
* use cases
  * your generated code is sufficiently different vs your original code 
    * _Example:_ your original code
      * is TypeScript or
      * you enabled [minification](#minify)
  * look at INDIVIDUAL files | your browser's developer tools
    * vs 1 big bundled file

* supported for
  * JS
  * CSS

* different modes for source map generation
  * linked
    * TODO:
              <p>
              This mode means the source map is generated into a separate `.js.map`
              output file alongside the `.js` output file, and the `.js` output file contains
              a special `//# sourceMappingURL=` comment that points to the `.js.map` output file.
              That way the browser knows where to find the source map for a given file
              when you open the debugger. Use `linked` source map mode like this:
              </p>
              </li></ol>

          - example:
              in:
                app.ts: 'let x: number = 1'

              cli: |
                esbuild app.ts --sourcemap --outfile=out.js

              mjs: |
                import * as esbuild from 'esbuild'

                await esbuild.build({
                  entryPoints: ['app.ts'],
                  sourcemap: true,
                  outfile: 'out.js',
                })

              go: |
                package main

                import "github.com/evanw/esbuild/pkg/api"
                import "os"

                func main() {
                  result := api.Build(api.BuildOptions{
                    EntryPoints: []string{"app.ts"},
                    Sourcemap:   api.SourceMapLinked,
                    Outfile:     "out.js",
                    Write:       true,
                  })

                  if len(result.Errors) > 0 {
                    os.Exit(1)
                  }
                }

  * external
            <p>
            This mode means the source map is generated into a separate `.js.map`
            output file alongside the `.js` output file, but unlike `linked` mode the `.js`
            output file does not contain a `//# sourceMappingURL=` comment. Use `external`
            source map mode like this:
            </p>
            </li></ol>

        - example:
            in:
              app.ts: 'let x: number = 1'

            cli: |
              esbuild app.ts --sourcemap=external --outfile=out.js

            mjs: |
              import * as esbuild from 'esbuild'

              await esbuild.build({
                entryPoints: ['app.ts'],
                sourcemap: 'external',
                outfile: 'out.js',
              })

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints: []string{"app.ts"},
                  Sourcemap:   api.SourceMapExternal,
                  Outfile:     "out.js",
                  Write:       true,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }
              }

  * inline
            <p>
            This mode means the source map is appended to the end of the `.js` output
            file as a base64 payload inside a `//# sourceMappingURL=` comment. No
            additional `.js.map` output file is generated. Keep in mind that source
            maps are usually very big because they contain all of your original source
            code, so you usually do not want to ship code containing `inline` source
            maps. To remove the source code from the source map (keeping only the file
            names and the line/column mappings), use the [sources content](#sources-content) option.
            Use `inline` source map mode like this:
            </p>
            </li></ol>

        - example:
            in:
              app.ts: 'let x: number = 1'

            cli: |
              esbuild app.ts --sourcemap=inline --outfile=out.js

            mjs: |
              import * as esbuild from 'esbuild'

              await esbuild.build({
                entryPoints: ['app.ts'],
                sourcemap: 'inline',
                outfile: 'out.js',
              })

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints: []string{"app.ts"},
                  Sourcemap:   api.SourceMapInline,
                  Outfile:     "out.js",
                  Write:       true,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }
              }

  * both
            <p>
            This mode is a combination of `inline` and `external`. The source map is
            appended inline to the end of the `.js` output file, and another copy of
            the same source map is written to a separate `.js.map` output file
            alongside the `.js` output file. Use `both` source map mode like this:
            </p>
            </li></ol>

        - example:
            in:
              app.ts: 'let x: number = 1'

            cli: |
              esbuild app.ts --sourcemap=both --outfile=out.js

            mjs: |
              import * as esbuild from 'esbuild'

              await esbuild.build({
                entryPoints: ['app.ts'],
                sourcemap: 'both',
                outfile: 'out.js',
              })

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints: []string{"app.ts"},
                  Sourcemap:   api.SourceMapInlineAndExternal,
                  Outfile:     "out.js",
                  Write:       true,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }
              }

        - p: >
            The [build](#build) API supports all four source map modes listed above,
            but the [transform](#transform) API does not support the `linked` mode.
            This is because the output returned from the transform API does not have an
            associated filename. If you want the output of the transform API to have a
            source map comment, you can append one yourself. In addition, the CLI form
            of the transform API only supports the `inline` mode because the output is
            written to stdout so generating multiple output files is not possible.

        - p: >
            If you want to "peek under the hood" to see what a source map does (or to
            debug problems with your source map), you can upload the relevant output
            file and the associated source map here:
            [Source Map Visualization](https://evanw.github.io/source-map-visualization/).

### Using source maps

        - p: >
            In the browser, source maps should be automatically picked up by the
            browser's developer tools as long as the source map setting is enabled.
            Note that the browser only uses the source maps to alter the display of
            stack traces when they are logged to the console. The stack traces
            themselves are not modified so inspecting <code>error.<wbr>stack</code>
            in your code will still give the unmapped stack trace containing compiled
            code. Here's how to enable this setting in your browser's developer tools:

        - ul:
            - 'Chrome: ⚙ → Enable JavaScript source maps'
            - 'Safari: ⚙ → Sources → Enable source maps'
            - 'Firefox: ··· → Enable Source Maps'

        - p: >
            In node, source maps are supported natively starting with [version v12.12.0](https://nodejs.org/en/blog/release/v12.12.0/).
            This feature is disabled by default but can be enabled with a flag. Unlike
            in the browser, the actual stack traces are also modified in node so
            inspecting <code>error.<wbr>stack</code> in your code will give the mapped
            stack trace containing your original source code. Here's how to enable this
            setting in node (the <code>--enable-<wbr>source-<wbr>maps</code> flag must
            come before the script file name):

        - pre.sh: |
            node --enable-source-maps app.js

        ## Sources content

        - p: >
            [Source maps](#sourcemap) are generated using [version 3](https://sourcemaps.info/spec.html)
            of the source map format, which is by far the most widely-supported
            variant. Each source map will look something like this:

        - pre.json: |
            {
              "version": 3,
              "sources": ["bar.js", "foo.js"],
              "sourcesContent": ["bar()", "foo()\nimport './bar'"],
              "mappings": ";AAAA;;;ACAA;",
              "names": []
            }

        - p: >
            The `sourcesContent` field is an optional field that contains all of the
            original source code. This is helpful for debugging because it means the
            original source code will be available in the debugger.

        - p: >
            However, it's not needed in some scenarios. For example, if you are just
            using source maps in production to generate stack traces that contain the
            original file name, you don't need the original source code because there
            is no debugger involved. In that case it can be desirable to omit the
            `sourcesContent` field to make the source map smaller:

        - example:
            in:
              app.js: '1 + 2'

            cli: |
              esbuild --bundle app.js --sourcemap --sources-content=false

            mjs: |
              import * as esbuild from 'esbuild'

              await esbuild.build({
                bundle: true,
                entryPoints: ['app.js'],
                sourcemap: true,
                sourcesContent: false,
                outfile: 'out.js',
              })

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  Bundle:         true,
                  EntryPoints:    []string{"app.js"},
                  Sourcemap:      api.SourceMapInline,
                  SourcesContent: api.SourcesContentExclude,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }
              }

        # Build metadata

        ## Analyze

        - info: >
            If you're looking for an interactive visualization, try esbuild's
            [Bundle Size Analyzer](/analyze/) instead. You can upload your esbuild
            [metafile](#metafile) to see a bundle size breakdown.

        - p: >
            Using the analyze feature generates an easy-to-read report about the contents of your bundle:

        - example:
            install:
              react: '17.0.2'
              react-dom: '17.0.2'

            in:
              example.jsx: |
                import * as React from 'react'
                import * as Server from 'react-dom/server'

                let Greet = () => <h1>Hello, world!</h1>
                console.log(Server.renderToString(<Greet />))

            cli:
              - $: |
                  esbuild --bundle example.jsx --outfile=out.js --minify --analyze

              - expect: |2

                    out.js                                                                    27.6kb  100.0%
                     ├ node_modules/react-dom/cjs/react-dom-server.browser.production.min.js  19.2kb   69.8%
                     ├ node_modules/react/cjs/react.production.min.js                          5.9kb   21.4%
                     ├ node_modules/object-assign/index.js                                     962b     3.4%
                     ├ example.jsx                                                             137b     0.5%
                     ├ node_modules/react-dom/server.browser.js                                 50b     0.2%
                     └ node_modules/react/index.js                                              50b     0.2%

                  ...

            mjs: |
              import * as esbuild from 'esbuild'

              let result = await esbuild.build({
                entryPoints: ['example.jsx'],
                outfile: 'out.js',
                minify: true,
                metafile: true,
              })

              console.log(await esbuild.analyzeMetafile(result.metafile))

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "fmt"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints:       []string{"example.jsx"},
                  Outfile:           "out.js",
                  MinifyWhitespace:  true,
                  MinifyIdentifiers: true,
                  MinifySyntax:      true,
                  Metafile:          true,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }

                fmt.Printf("%s", api.AnalyzeMetafile(result.Metafile, api.AnalyzeMetafileOptions{}))
              }

        - p: >
            The information shows which input files ended up in each output file as
            well as the percentage of the output file they ended up taking up. If you
            would like additional information, you can enable the "verbose" mode.
            This currently shows the import path from the entry point to each input
            file which tells you why a given input file is being included in the bundle:

        - example:
            install:
              react: '17.0.2'
              react-dom: '17.0.2'

            in:
              example.jsx: |
                import * as React from 'react'
                import * as Server from 'react-dom/server'

                let Greet = () => <h1>Hello, world!</h1>
                console.log(Server.renderToString(<Greet />))

            cli:
              - $: |
                  esbuild --bundle example.jsx --outfile=out.js --minify --analyze=verbose

              - expect: |2

                    out.js ─────────────────────────────────────────────────────────────────── 27.6kb ─ 100.0%
                     ├ node_modules/react-dom/cjs/react-dom-server.browser.production.min.js ─ 19.2kb ── 69.8%
                     │  └ node_modules/react-dom/server.browser.js
                     │     └ example.jsx
                     ├ node_modules/react/cjs/react.production.min.js ───────────────────────── 5.9kb ── 21.4%
                     │  └ node_modules/react/index.js
                     │     └ example.jsx
                     ├ node_modules/object-assign/index.js ──────────────────────────────────── 962b ──── 3.4%
                     │  └ node_modules/react-dom/cjs/react-dom-server.browser.production.min.js
                     │     └ node_modules/react-dom/server.browser.js
                     │        └ example.jsx
                     ├ example.jsx ──────────────────────────────────────────────────────────── 137b ──── 0.5%
                     ├ node_modules/react-dom/server.browser.js ──────────────────────────────── 50b ──── 0.2%
                     │  └ example.jsx
                     └ node_modules/react/index.js ───────────────────────────────────────────── 50b ──── 0.2%
                        └ example.jsx

                  ...

            mjs: |
              import * as esbuild from 'esbuild'

              let result = await esbuild.build({
                entryPoints: ['example.jsx'],
                outfile: 'out.js',
                minify: true,
                metafile: true,
              })

              console.log(await esbuild.analyzeMetafile(result.metafile, {
                verbose: true,
              }))

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "fmt"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints:       []string{"example.jsx"},
                  Outfile:           "out.js",
                  MinifyWhitespace:  true,
                  MinifyIdentifiers: true,
                  MinifySyntax:      true,
                  Metafile:          true,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }

                fmt.Printf("%s", api.AnalyzeMetafile(result.Metafile, api.AnalyzeMetafileOptions{
                  Verbose: true,
                }))
              }

        - p: >
            This analysis is just a visualization of the information that can be found
            in the [metafile](#metafile). If this analysis doesn't exactly suit your
            needs, you are welcome to build your own visualization using the information
            in the metafile.

        - p: >
            Note that this formatted analysis summary is intended for humans, not
            machines. The specific formatting may change over time which will likely
            break any tools that try to parse it. You should not write a tool to parse
            this data. You should be using the information in the [JSON metadata file](#metafile)
            instead. Everything in this visualization is derived from the JSON metadata
            so you are not losing out on any information by not parsing esbuild's
            formatted analysis summary.

        ## Metafile

        - p: >
            This option tells esbuild to produce some metadata about the build in
            JSON format. The following example puts the metadata in a file called
            `meta.json`:

        - example:
            in:
              app.js: '1 + 2'

            cli: |
              esbuild app.js --bundle --metafile=meta.json --outfile=out.js

            mjs: |
              import * as esbuild from 'esbuild'
              import fs from 'node:fs'

              let result = await esbuild.build({
                entryPoints: ['app.js'],
                bundle: true,
                metafile: true,
                outfile: 'out.js',
              })

              fs.writeFileSync('meta.json', JSON.stringify(result.metafile))

            go: |
              package main

              import "io/ioutil"
              import "github.com/evanw/esbuild/pkg/api"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints: []string{"app.js"},
                  Bundle:      true,
                  Metafile:    true,
                  Outfile:     "out.js",
                  Write:       true,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }

                ioutil.WriteFile("meta.json", []byte(result.Metafile), 0644)
              }

        - p: >
            This data can then be analyzed by other tools. For an interactive
            visualization, you can use esbuild's own [Bundle Size Analyzer](/analyze/).
            For a quick textual analysis, you can use esbuild's build-in [analyze](#analyze)
            feature. Or you can write your own analysis which uses this information.

        - p: >
            The metadata JSON format looks like this (described using a TypeScript
            interface):

        - pre.ts: |
            interface Metafile {
              inputs: {
                [path: string]: {
                  bytes: number
                  imports: {
                    path: string
                    kind: string
                    external?: boolean
                    original?: string
                    with?: Record<string, string>
                  }[]
                  format?: string
                  with?: Record<string, string>
                }
              }
              outputs: {
                [path: string]: {
                  bytes: number
                  inputs: {
                    [path: string]: {
                      bytesInOutput: number
                    }
                  }
                  imports: {
                    path: string
                    kind: string
                    external?: boolean
                  }[]
                  exports: string[]
                  entryPoint?: string
                  cssBundle?: string
                }
              }
            }

        # Logging

        ## Color

        - p: >
            This option enables or disables colors in the error and warning messages
            that esbuild writes to stderr file descriptor in the terminal. By
            default, color is automatically enabled if stderr is a TTY session and
            automatically disabled otherwise. Colored output in esbuild looks like this:

        - pre.raw: >
            {{ FORMAT_MESSAGES('import log from "logger"\nlog(typeof x == "null")', { sourcefile: 'example.js', bundle: true }) }}

        - p: >
            Colored output can be force-enabled by setting color to `true`. This is
            useful if you are piping esbuild's stderr output into a TTY yourself:

        - example:
            cli: |
              echo 'typeof x == "null"' | esbuild --color=true 2> stderr.txt

            mjs: |
              import * as esbuild from 'esbuild'

              let js = 'typeof x == "null"'
              await esbuild.transform(js, {
                color: true,
              })

            go: |
              package main

              import "fmt"
              import "github.com/evanw/esbuild/pkg/api"

              func main() {
                js := "typeof x == 'null'"

                result := api.Transform(js, api.TransformOptions{
                  Color: api.ColorAlways,
                })

                if len(result.Errors) == 0 {
                  fmt.Printf("%s", result.Code)
                }
              }

        - p: >
            Colored output can also be set to `false` to disable colors.

        ## Format messages

        - p: >
            This API call can be used to format the log errors and warnings returned
            by the [build](#build) API and [transform](#transform) APIs as a
            string using the same formatting that esbuild itself uses. This is useful
            if you want to customize the way esbuild's logging works, such as processing
            the log messages before they are printed or printing them to somewhere other
            than to the console. Here's an example:

        - example:
            in:
              app.js: '1 + 2'

            mjs: |
              import * as esbuild from 'esbuild'

              let formatted = await esbuild.formatMessages([
                {
                  text: 'This is an error',
                  location: {
                    file: 'app.js',
                    line: 10,
                    column: 4,
                    length: 3,
                    lineText: 'let foo = bar',
                  },
                },
              ], {
                kind: 'error',
                color: false,
                terminalWidth: 100,
              })

              console.log(formatted.join('\n'))

            go: |
              package main

              import "fmt"
              import "github.com/evanw/esbuild/pkg/api"
              import "strings"

              func main() {
                formatted := api.FormatMessages([]api.Message{
                  {
                    Text: "This is an error",
                    Location: &api.Location{
                      File:     "app.js",
                      Line:     10,
                      Column:   4,
                      Length:   3,
                      LineText: "let foo = bar",
                    },
                  },
                }, api.FormatMessagesOptions{
                  Kind:          api.ErrorMessage,
                  Color:         false,
                  TerminalWidth: 100,
                })

                fmt.Printf("%s", strings.Join(formatted, "\n"))
              }

        ### Options

        - p: >
            The following options can be provided to control the formatting:

        - example:
            noCheck: true

            js: |
              interface FormatMessagesOptions {
                kind: 'error' | 'warning';
                color?: boolean;
                terminalWidth?: number;
              }

            go: |
              type FormatMessagesOptions struct {
                Kind          MessageKind
                Color         bool
                TerminalWidth int
              }

        - ul:
          - >
            `kind`
            <p>Controls whether these log messages are printed as errors or warnings.</p>
          - >
            `color`
            <p>If this is `true`, Unix-style terminal escape codes are included for
            colored output.</p>
          - >
            `terminalWidth`
            <p>Provide a positive value to wrap long lines so that they don't overflow
            past the provided column width. Provide `0` to disable word wrapping.</p>

        ## Log level

        - p: >
            The log level can be changed to prevent esbuild from printing warning
            and/or error messages to the terminal. The six log levels are:

        - ul:
          - >
            <p>`silent`<br>Do not show any log output. This is the default log level
            when using the JS [transform](#transform) API.</p>
          - >
            <p>`error`<br>Only show errors.</p>
          - >
            <p>`warning`<br>Only show warnings and errors. This is the default log
            level when using the JS [build](#build) API.</p>
          - >
            <p>`info`<br>Show warnings, errors, and an output file summary. This is
            the default log level when using the CLI.</p>
          - >
            <p>`debug`<br>Log everything from `info` and some additional messages
            that may help you debug a broken bundle. This log level has a performance
            impact and some of the messages may be false positives, so this information
            is not shown by default.</p>
          - >
            <p>`verbose`<br>This generates a torrent of log messages and was added to
            debug issues with file system drivers. It's not intended for general use.</p>

        - p: >
            The log level can be set like this:

        - example:
            cli: |
              echo 'typeof x == "null"' | esbuild --log-level=error

            mjs: |
              import * as esbuild from 'esbuild'

              let js = 'typeof x == "null"'
              await esbuild.transform(js, {
                logLevel: 'error',
              })

            go: |
              package main

              import "fmt"
              import "github.com/evanw/esbuild/pkg/api"

              func main() {
                js := "typeof x == 'null'"

                result := api.Transform(js, api.TransformOptions{
                  LogLevel: api.LogLevelError,
                })

                if len(result.Errors) == 0 {
                  fmt.Printf("%s", result.Code)
                }
              }

        ## Log limit

        - p: >
            By default, esbuild stops reporting log messages after 10 messages have
            been reported. This avoids the accidental generation of an overwhelming number
            of log messages, which can easily lock up slower terminal emulators such
            as Windows command prompt. It also avoids accidentally using up the
            whole scroll buffer for terminal emulators with limited scroll buffers.

        - p: >
            The log limit can be changed to another value, and can also be disabled
            completely by setting it to zero. This will show all log messages:

        - example:
            in:
              app.js: '1 + 2'

            cli: |
              esbuild app.js --log-limit=0

            mjs: |
              import * as esbuild from 'esbuild'

              await esbuild.build({
                entryPoints: ['app.js'],
                logLimit: 0,
                outfile: 'out.js',
              })

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints: []string{"app.js"},
                  LogLimit:    0,
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }
              }

        ## Log override

        - p: >
            This feature lets you change the log level of individual types of log
            messages. You can use it to silence a particular type of warning, to
            enable additional warnings that aren't enabled by default, or even to
            turn warnings into errors.

        - p: >
            For example, when targeting older browsers, esbuild automatically
            transforms regular expression literals which use features that are
            too new for those browsers into <code>new <wbr>RegExp()</code> calls
            to allow the generated code to run without being considered a syntax
            error by the browser. However, these calls will still throw at runtime
            if you don't add a polyfill for `RegExp` because that regular
            expression syntax is still unsupported. If you want esbuild to generate
            a warning when you use newer unsupported regular expression syntax, you
            can do that like this:

        - example:
            in:
              app.js: '/./d'

            cli: |
              esbuild app.js --log-override:unsupported-regexp=warning --target=chrome50

            mjs: |
              import * as esbuild from 'esbuild'

              await esbuild.build({
                entryPoints: ['app.js'],
                logOverride: {
                  'unsupported-regexp': 'warning',
                },
                target: 'chrome50',
              })

            go: |
              package main

              import "github.com/evanw/esbuild/pkg/api"
              import "os"

              func main() {
                result := api.Build(api.BuildOptions{
                  EntryPoints: []string{"app.js"},
                  LogOverride: map[string]api.LogLevel{
                    "unsupported-regexp": api.LogLevelWarning,
                  },
                  Engines: []api.Engine{
                    {Name: api.EngineChrome, Version: "50"},
                  },
                })

                if len(result.Errors) > 0 {
                  os.Exit(1)
                }
              }

        - p: >
            The log level for each message type can be overridden to any value
            supported by the [log level](#log-level) setting. All
            currently-available message types are listed below (click on each one
            for an example log message):

        - p: |
            <ul>
              <li>
                **JS:**
                <ul>
                  <li><details><summary>`assert-to-with`</summary><pre>{{ FORMAT_MESSAGES('import data from "./data.json" assert { type: "json" }', { sourcefile: 'example.js', supported: { 'import-assertions': false } }) }}</pre></details></li>
                  <li><details><summary>`assert-type-json`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'import * as data from "./data.json" assert { type: "json" }; console.log(data.value)', 'data.json': '{}' }, { bundle: true }) }}</pre></details></li>
                  <li><details><summary>`assign-to-constant`</summary><pre>{{ FORMAT_MESSAGES('const foo = 1; foo = 2', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`assign-to-define`</summary><pre>{{ FORMAT_MESSAGES('DEFINE = false', { sourcefile: 'example.js', define: { DEFINE: 'true' } }) }}</pre></details></li>
                  <li><details><summary>`assign-to-import`</summary><pre>{{ FORMAT_MESSAGES('import foo from "foo"; foo = null', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`call-import-namespace`</summary><pre>{{ FORMAT_MESSAGES('import * as foo from "foo"; foo()', { sourcefile: 'example.js', format: 'esm' }) }}</pre></details></li>
                  <li><details><summary>`class-name-will-throw`</summary><pre>{{ FORMAT_MESSAGES('class Foo { static key = "foo"; static [Foo.key] = 123 }', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`commonjs-variable-in-esm`</summary><pre>{{ FORMAT_MESSAGES('exports.foo = 1; export let bar = 2', { sourcefile: 'example.js', format: 'esm' }) }}</pre></details></li>
                  <li><details><summary>`delete-super-property`</summary><pre>{{ FORMAT_MESSAGES('class Foo extends Object { foo() { delete super.foo } }', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`direct-eval`</summary><pre>{{ FORMAT_MESSAGES('let apparentlyUnused; eval("actuallyUse(apparentlyUnused)")', { sourcefile: 'example.js', logOverride: { 'direct-eval': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`duplicate-case`</summary><pre>{{ FORMAT_MESSAGES('switch (foo) { case 1: return 1; case 1: return 2 }', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`duplicate-class-member`</summary><pre>{{ FORMAT_MESSAGES('class Foo { x = 1; x = 2 }', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`duplicate-object-key`</summary><pre>{{ FORMAT_MESSAGES('foo = { bar: 1, bar: 2 }', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`empty-import-meta`</summary><pre>{{ FORMAT_MESSAGES('foo = import.meta', { sourcefile: 'example.js', target: 'chrome50' }) }}</pre></details></li>
                  <li><details><summary>`equals-nan`</summary><pre>{{ FORMAT_MESSAGES('foo = foo.filter(x => x !== NaN)', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`equals-negative-zero`</summary><pre>{{ FORMAT_MESSAGES('foo = foo.filter(x => x !== -0)', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`equals-new-object`</summary><pre>{{ FORMAT_MESSAGES('foo = foo.filter(x => x !== [])', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`html-comment-in-js`</summary><pre>{{ FORMAT_MESSAGES('<!-- comment -->', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`impossible-typeof`</summary><pre>{{ FORMAT_MESSAGES('foo = foo.map(x => typeof x !== "null")', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`indirect-require`</summary><pre>{{ FORMAT_MESSAGES('let r = require, fs = r("fs")', { sourcefile: 'example.js', bundle: true, logOverride: { 'indirect-require': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`private-name-will-throw`</summary><pre>{{ FORMAT_MESSAGES('class Foo { get #foo() {} bar() { this.#foo++ } }', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`semicolon-after-return`</summary><pre>{{ FORMAT_MESSAGES('return\nx', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`suspicious-boolean-not`</summary><pre>{{ FORMAT_MESSAGES('if (!foo in bar) {\n}', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`suspicious-define`</summary><pre>{{ FORMAT_MESSAGES('', { sourcefile: 'example.js', define: { 'process.env.NODE_ENV': 'production' } }) }}</pre></details></li>
                  <li><details><summary>`suspicious-logical-operator`</summary><pre>{{ FORMAT_MESSAGES('const isInRange = x => 0 && x <= 1', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`suspicious-nullish-coalescing`</summary><pre>{{ FORMAT_MESSAGES('return name === user.name ?? ""', { sourcefile: 'example.js' }) }}</pre></details></li>
                  <li><details><summary>`this-is-undefined-in-esm`</summary><pre>{{ FORMAT_MESSAGES('this.foo = 1; export let bar = 2', { sourcefile: 'example.js', bundle: true, logOverride: { 'this-is-undefined-in-esm': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`unsupported-dynamic-import`</summary><pre>{{ FORMAT_MESSAGES('import(foo)', { sourcefile: 'example.js', bundle: true, logOverride: { 'unsupported-dynamic-import': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`unsupported-jsx-comment`</summary><pre>{{ FORMAT_MESSAGES('// @jsx 123', { sourcefile: 'example.jsx', loader: 'jsx' }) }}</pre></details></li>
                  <li><details><summary>`unsupported-regexp`</summary><pre>{{ FORMAT_MESSAGES('/./d', { sourcefile: 'example.js', target: 'chrome50', logOverride: { 'unsupported-regexp': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`unsupported-require-call`</summary><pre>{{ FORMAT_MESSAGES('require(foo)', { sourcefile: 'example.js', bundle: true, logOverride: { 'unsupported-require-call': 'warning' } }) }}</pre></details></li>
                </ul>
                <br>
              </li>

              <li>
                **CSS:**
                <ul>
                  <li><details><summary>`css-syntax-error`</summary><pre>{{ FORMAT_MESSAGES('div[] {\n}', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`invalid-@charset`</summary><pre>{{ FORMAT_MESSAGES('div { color: red } @charset "UTF-8";', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`invalid-@import`</summary><pre>{{ FORMAT_MESSAGES('div { color: red } @import "foo.css";', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`invalid-@layer`</summary><pre>{{ FORMAT_MESSAGES('@layer initial {\n}', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`invalid-calc`</summary><pre>{{ FORMAT_MESSAGES('div { z-index: calc(-(1+2)); }', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`js-comment-in-css`</summary><pre>{{ FORMAT_MESSAGES('// comment', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`undefined-composes-from`</summary><pre>{{ FORMAT_MESSAGES({ 'example.module.css': '.foo { composes: bar from "lib.module.css"; zoom: 1; }', 'lib.module.css': '.bar { zoom: 2 }' }, { bundle: true }) }}</pre></details></li>
                  <li><details><summary>`unsupported-@charset`</summary><pre>{{ FORMAT_MESSAGES('@charset "ASCII";', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`unsupported-@namespace`</summary><pre>{{ FORMAT_MESSAGES('@namespace "ns";', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`unsupported-css-property`</summary><pre>{{ FORMAT_MESSAGES('div { widht: 1px }', { sourcefile: 'example.css', loader: 'css' }) }}</pre></details></li>
                  <li><details><summary>`unsupported-css-nesting`</summary><pre>{{ FORMAT_MESSAGES('a b {\n.foo & {\n}\n}', { sourcefile: 'example.css', loader: 'css', target: 'chrome50' }) }}</pre></details></li>
                </ul>
                <br>
              </li>

              <li>
                **Bundler:**
                <ul>
                  <li><details><summary>`ambiguous-reexport`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'export * from "./a"; export * from "./b"', 'a.js': 'export let foo = 1', 'b.js': 'export let foo = 2' }, { bundle: true, logOverride: { 'ambiguous-reexport': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`different-path-case`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'import "./foo.js"\nimport "./Foo.js"', 'foo.js': '' }, { bundle: true }) }}</pre></details></li>
                  <li><details><summary>`empty-glob`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'function getIcon(name) {\n  return import("./icon-" + name + ".json")\n}' }, { bundle: true }) }}</pre></details></li>
                  <li><details><summary>`ignored-bare-import`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'import "foo"', 'node_modules/foo/index.js': 'foo', 'node_modules/foo/package.json': '{\n  "sideEffects": false\n}' }, { bundle: true }) }}</pre></details></li>
                  <li><details><summary>`ignored-dynamic-import`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'import("foo").catch(e => {\n})' }, { bundle: true, logOverride: { 'ignored-dynamic-import': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`import-is-undefined`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'import { foo } from "./foo"', 'foo.js': 'let foo = 1' }, { bundle: true, logOverride: { 'import-is-undefined': 'warning' } }) }}</pre></details></li>
                  <li><details><summary>`require-resolve-not-external`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': 'let foo = require.resolve("foo")' }, { bundle: true, format: 'cjs' }) }}</pre></details></li>
                </ul>
                <br>
              </li>

              <li>
                **Source maps:**
                <ul>
                  <li><details><summary>`invalid-source-mappings`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': '//# sourceMappingURL=example.js.map', 'example.js.map': '{ "version": 3, "sources": ["example.js"],\n  "mappings": "aAAFA,UAAU;;"\n}' }, { bundle: true, sourcemap: true }) }}</pre></details></li>
                  <li><details><summary>`sections-in-source-map`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': '//# sourceMappingURL=example.js.map', 'example.js.map': '{\n  "sections": []\n}' }, { bundle: true, sourcemap: true }) }}</pre></details></li>
                  <li><details><summary>`missing-source-map`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': '//# sourceMappingURL=.' }, { bundle: true, sourcemap: true }) }}</pre></details></li>
                  <li><details><summary>`unsupported-source-map-comment`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': '//# sourceMappingURL=data:application/json,"%"' }, { bundle: true, sourcemap: true }) }}</pre></details></li>
                </ul>
                <br>
              </li>

              <li>
                **Resolver:**
                <ul>
                  <li><details><summary>`package.json`</summary><pre>{{ FORMAT_MESSAGES({ 'example.js': '', 'package.json': '{ "type": "esm" }' }, { bundle: true }) }}</pre></details></li>
                  <li><details><summary>`tsconfig.json`</summary><pre>{{ FORMAT_MESSAGES({ 'example.ts': '', 'tsconfig.json': '{ "compilerOptions": { "target": "ES4" } }' }, { bundle: true }) }}</pre></details></li>
                </ul>
                <br>
              </li>
            </ul>

        - p: >
            These message types should be reasonably stable but new ones may be added
            and old ones may occasionally be removed in the future. If a message type
            is removed, any overrides for that message type will just be silently ignored.