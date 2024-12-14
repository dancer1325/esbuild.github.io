# Install esbuild
* ways to install it
  * `npm install --save-exact --save-dev esbuild`
    * download & install the `esbuild` locally
      * check the installation | `node_modules/`
      * run it, checking the version
        * | unix, `./node_modules/.bin/esbuild --version`
        * | windows, `.\node_modules\.bin\esbuild --version`
  * [other ways](#other-ways-to-install)

# Your first bundle
* bundle a code
  * == your code + required libraries / bundled together
    * JSX ALSO accepted |
      * ".jsx" / WITHOUT configuration
      * ".js" / requires `--loader:.js=jsx` -- see [API documentation](api.md)
  * / COMPLETELY self-contained
    * Reason: 🧠 NO longer -- depends on -- your `node_modules` directory 🧠
  * uses
    * run it

* _Example:_ 
  * `npm install react react-dom`
    * install the `react` and `react-dom` packages
  * create a file / named `app.jsx`
  * bundle the file -- via -- esbuild
    * | unix, `./node_modules/.bin/esbuild src/content/examples/app.jsx --bundle --outfile=out.js`
    * | windows, `.\node_modules\.bin\esbuild src\content\examples\app.jsx --bundle --outfile=out.js`
  * `node out.js`
    * run it

# Build scripts

* uses
  * run repeatedly & automatically, your build command

* ways to create
  * add a build script | your `package.json`
    ```
    {
      "scripts": {
        "build": "esbuild app.jsx --bundle --outfile=out.js"
        // `esbuild` command directly WITHOUT a relative path
        // Reason: 🧠it works because everything | `scripts`, run with the `esbuild` command | path 🧠
      }
    }
    ```
    * invoke the script
      * `npm run build`
  * write | JS -- via -- esbuild's JavaScript API
    * allows
      * passing MANY options to esbuild
    * save as `.mjs`
      * Reason: 🧠it uses the `import` keyword 🧠
      ```.mjs
      import * as esbuild from 'esbuild'

      await esbuild.build({
        entryPoints: ['app.jsx'],
        bundle: true,
        outfile: 'out.js',
      })
      ```

* `build`
  * runs the esbuild executable | child process
  * returns a promise / resolves | complete the build
  * recommended use cases
    * build scripts
      * Reason: 🧠[plugins](plugins.md) ONLY work | asynchronous API 🧠
  * see [API documentation](api.md#build)
* `buildSync` API
  * == sync build

# Bundling for the browser

* bundler outputs
  * 👀by default, for the browser 👀
    * -> NO additional configuration
  * for
    * development builds
      * typically to enable
        * [source maps](api.md#sourcemap) `--sourcemap`
    * production builds
      * enable [minification](/api/#minify) `--minify`
      * configure the [target](/api/#target) environment | supported browsers
        * JS syntax / too new -- will be transformed into -- older JS syntax 

* _Example:_
  * via CLI 
    ```
    esbuild app.jsx --bundle --minify --sourcemap --target=chrome58,firefox57,safari11,edge16`
    ```
  * | JS file
    ```.mjs
    import * as esbuild from 'esbuild'
  
    await esbuild.build({
      entryPoints: ['app.jsx'],
      bundle: true,
      minify: true,
      sourcemap: true,
      target: ['chrome58', 'firefox57', 'safari11', 'edge16'],
      outfile: 'out.js',
    })
    ```
  * | go
    ```
    import "github.com/evanw/esbuild/pkg/api"
    import "os"
  
    func main() {
    result := api.Build(api.BuildOptions{
      EntryPoints:       []string{"app.jsx"},
      Bundle:            true,
      MinifyWhitespace:  true,
      MinifyIdentifiers: true,
      MinifySyntax:      true,
      Engines: []api.Engine{
        {api.EngineChrome, "58"},
        {api.EngineFirefox, "57"},
        {api.EngineSafari, "11"},
        {api.EngineEdge, "16"},
      },
      Write: true,
    })
  
    if len(result.Errors) > 0 {
      os.Exit(1)
    }
    }
    ```

* esbuild's OTHER configurations
  * conditionally work & successfully bundle
* undefined globals -- can be replaced with --
  * [define](api.md#define) | simple cases
  * [inject](api.md#inject) | MORE complex cases

# Bundling for node

* TODO:
          - p: >
              Even though a bundler is not necessary when using node, sometimes it
              can still be beneficial to process your code with esbuild before running
              it in node. Bundling can automatically strip TypeScript types, convert
              ECMAScript module syntax to CommonJS, and transform newer JavaScript
              syntax into older syntax for a specific version of node. And it may be
              beneficial to bundle your package before publishing it so that it's
              a smaller download and so it spends less time reading from the file
              system when being loaded.

          - p: >
              If you are bundling code that will be run in node, you should configure
              the [platform](/api/#platform) setting by passing <code>--platform=<wbr>node</code>
              to esbuild. This simultaneously changes a few different settings to
              node-friendly default values. For example, all packages that are
              built-in to node such as `fs` are automatically marked as external so
              esbuild doesn't try to bundle them. This setting also disables the
              interpretation of the browser field in `package.json`.

          - p: >
              If your code uses newer JavaScript syntax that doesn't work in your
              version of node, you will want to configure the [target](/api/#target)
              version of node:

          - example:
              in:
                app.js: '1 + 2'

              cli: |
                esbuild app.js --bundle --platform=node --target=node10.4

              mjs: |
                import * as esbuild from 'esbuild'

                await esbuild.build({
                  entryPoints: ['app.js'],
                  bundle: true,
                  platform: 'node',
                  target: ['node10.4'],
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
                    Engines: []api.Engine{
                      {api.EngineNode, "10.4"},
                    },
                    Write: true,
                  })

                  if len(result.Errors) > 0 {
                    os.Exit(1)
                  }
                }

          - p: >
              You also may not want to bundle your dependencies with esbuild. There
              are many node-specific features that esbuild doesn't support while
              bundling such as `__dirname`, `import.meta.url`, `fs.readFileSync`,
              and `*.node` native binary modules. You can exclude all of your
              dependencies from the bundle by setting [packages](/api/#packages)
              to external:

          - example:
              in:
                app.jsx: '<div/>'

              cli: |
                esbuild app.jsx --bundle --platform=node --packages=external

              js: |
                require('esbuild').buildSync({
                  entryPoints: ['app.jsx'],
                  bundle: true,
                  platform: 'node',
                  packages: 'external',
                  outfile: 'out.js',
                })

              go: |
                package main

                import "github.com/evanw/esbuild/pkg/api"
                import "os"

                func main() {
                  result := api.Build(api.BuildOptions{
                    EntryPoints: []string{"app.jsx"},
                    Bundle:      true,
                    Platform:    api.PlatformNode,
                    Packages:    api.PackagesExternal,
                    Write:       true,
                  })

                  if len(result.Errors) > 0 {
                    os.Exit(1)
                  }
                }

          - p: >
              If you do this, your dependencies must still be present on the file
              system at run-time since they are no longer included in the bundle.

# Simultaneous platforms

          - p: >
              You cannot install esbuild on one OS, copy the `node_modules` directory
              to another OS without reinstalling, and then run esbuild on that other OS.
              This won't work because esbuild is written with native code and needs to
              install a platform-specific binary executable. Normally this isn't an
              issue because you typically check your `package.json` file into version
              control, not your `node_modules` directory, and then everyone runs
              `npm install` on their local machine after cloning the repository.

          - p: >
              However, people sometimes get into this situation by installing esbuild
              on Windows or macOS and copying their `node_modules` directory into a
              [Docker](https://www.docker.com/) image that runs Linux, or by copying
              their `node_modules` directory between Windows and [WSL](https://docs.microsoft.com/en-us/windows/wsl/)
              environments. The way to get this to work depends on your package manager:

          - ul:
            - >
                <p>
                **npm/pnpm:**
                If you are installing with npm or pnpm, you can try not copying the
                `node_modules` directory when you copy the files over, and running
                `npm ci` or `npm install` on the destination platform after the copy.
                Or you could consider using [Yarn](https://yarnpkg.com/) instead which
                has built-in support for installing a package on multiple platforms
                simultaneously.
                </p>

            - >
                <p>
                **Yarn:**
                If you are installing with Yarn, you can try listing both this platform and the
                other platform in your `.yarnrc.yml` file using
                [the `supportedArchitectures` feature](https://yarnpkg.com/configuration/yarnrc/#supportedArchitectures).
                Keep in mind that this means multiple copies of esbuild will be present
                on the file system.
                </p>

          - p: >
              You can also get into this situation on a macOS computer with an ARM
              processor if you install esbuild using the ARM version of npm but then
              try to run esbuild with the x86-64 version of node running inside of
              [Rosetta](https://en.wikipedia.org/wiki/Rosetta_(software)). In that
              case, an easy fix is to run your code using the ARM version of node
              instead, which can be downloaded here: [https://nodejs.org/en/download/](https://nodejs.org/en/download/).

          - p: >
              Another alternative is to [use the `esbuild-wasm` package instead](#wasm),
              which works the same way on all platforms. But it comes with a heavy
              performance cost and can sometimes be 10x slower than the `esbuild` package,
              so you may also not want to do that.

# yarn-pnp: Using Yarn Plug'n'Play

          - p: >
              Yarn's [Plug'n'Play](https://yarnpkg.com/features/pnp/) package
              installation strategy is supported natively by esbuild. To use it, make
              sure you are running esbuild such that the [current working directory](/api/#working-directory)
              contains Yarn's generated package manifest JavaScript file (either
              `.pnp.cjs` or `.pnp.js`). If a Yarn Plug'n'Play package manifest is
              detected, esbuild will automatically resolve package imports to paths
              inside the `.zip` files in Yarn's package cache, and will automatically
              extract these files on the fly during bundling.

          - p: >
              Because esbuild is written in Go, support for Yarn Plug'n'Play has been
              completely re-implemented in Go instead of relying on Yarn's JavaScript
              API. This allows Yarn Plug'n'Play package resolution to integrate well
              with esbuild's fully parallelized bundling pipeline for maximum speed.
              Note that Yarn's command-line interface adds a lot of unavoidable
              performance overhead to every command. For maximum esbuild performance,
              you may want to consider running esbuild without using Yarn's CLI (i.e.
              not using `yarn esbuild`). This can result in esbuild running 10x faster.

# Other ways to install

          - p: >
              The recommended way to install esbuild is to [install the native executable using npm](#install-esbuild).
              But you can also install esbuild in these ways:

## Download a build

          - p: >
              If you have a Unix system, you can use the following command to download
              the `esbuild` binary executable for your current platform (it will be
              downloaded to the current working directory):

          - pre: |
              curl -fsSL https://esbuild.github.io/dl/vCURRENT_ESBUILD_VERSION | sh

          - p: >
              You can also use `latest` instead of the version number to download the
              most recent version of esbuild:

          - pre: |
              curl -fsSL https://esbuild.github.io/dl/latest | sh

          - p: >
              If you don't want to evaluate a shell script from the internet to
              download esbuild, you can also manually download the package from
              npm yourself instead (which is all the above shell script is doing).
              Although the precompiled native executables are hosted using npm, you don't
              actually need npm installed to download them. The npm package registry is
              a normal HTTP server and packages are normal gzipped tar files.

          - p: >
              Here is an example of downloading a binary executable directly:

          - example:
              noCheck: true

              cli:
                - $: |
                    curl -O https://registry.npmjs.org/@esbuild/darwin-x64/-/darwin-x64-CURRENT_ESBUILD_VERSION.tgz
                - $: |
                    tar xzf ./darwin-x64-CURRENT_ESBUILD_VERSION.tgz
                - $: |
                    ./package/bin/esbuild
                - expect: |
                    Usage:
                      esbuild [options] [entry points]

                    ...

          - p: >
              The native executable in the `@esbuild/darwin-x64` package is for the macOS
              operating system and the 64-bit Intel architecture. As of writing, this is the
              full list of native executable packages for the platforms esbuild supports:

          - table: |
              | Package name                                                                       | OS                   | Architecture           | Download                                                                                                                  |
              |------------------------------------------------------------------------------------|----------------------|------------------------|---------------------------------------------------------------------------------------------------------------------------|
              | [`@esbuild/aix-ppc64`](https://www.npmjs.org/package/@esbuild/aix-ppc64)           | `aix`                | `ppc64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/aix-ppc64/-/android-arm-CURRENT_ESBUILD_VERSION.tgz"></a>         |
              | [`@esbuild/android-arm`](https://www.npmjs.org/package/@esbuild/android-arm)       | `android`            | `arm`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/android-arm/-/android-arm-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/android-arm64`](https://www.npmjs.org/package/@esbuild/android-arm64)   | `android`            | `arm64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/android-arm64/-/android-arm64-CURRENT_ESBUILD_VERSION.tgz"></a>   |
              | [`@esbuild/android-x64`](https://www.npmjs.org/package/@esbuild/android-x64)       | `android`            | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/android-x64/-/android-x64-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/darwin-arm64`](https://www.npmjs.org/package/@esbuild/darwin-arm64)     | `darwin`             | `arm64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/darwin-arm64/-/darwin-arm64-CURRENT_ESBUILD_VERSION.tgz"></a>     |
              | [`@esbuild/darwin-x64`](https://www.npmjs.org/package/@esbuild/darwin-x64)         | `darwin`             | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/darwin-x64/-/darwin-x64-CURRENT_ESBUILD_VERSION.tgz"></a>         |
              | [`@esbuild/freebsd-arm64`](https://www.npmjs.org/package/@esbuild/freebsd-arm64)   | `freebsd`            | `arm64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/freebsd-arm64/-/freebsd-arm64-CURRENT_ESBUILD_VERSION.tgz"></a>   |
              | [`@esbuild/freebsd-x64`](https://www.npmjs.org/package/@esbuild/freebsd-x64)       | `freebsd`            | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/freebsd-x64/-/freebsd-x64-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/linux-arm`](https://www.npmjs.org/package/@esbuild/linux-arm)           | `linux`              | `arm`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-arm/-/linux-arm-CURRENT_ESBUILD_VERSION.tgz"></a>           |
              | [`@esbuild/linux-arm64`](https://www.npmjs.org/package/@esbuild/linux-arm64)       | `linux`              | `arm64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-arm64/-/linux-arm64-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/linux-ia32`](https://www.npmjs.org/package/@esbuild/linux-ia32)         | `linux`              | `ia32`                 | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-ia32/-/linux-ia32-CURRENT_ESBUILD_VERSION.tgz"></a>         |
              | [`@esbuild/linux-loong64`](https://www.npmjs.org/package/@esbuild/linux-loong64)   | `linux`              | `loong64`<sup>2</sup>  | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-loong64/-/linux-loong64-CURRENT_ESBUILD_VERSION.tgz"></a>   |
              | [`@esbuild/linux-mips64el`](https://www.npmjs.org/package/@esbuild/linux-mips64el) | `linux`              | `mips64el`<sup>2</sup> | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-mips64el/-/linux-mips64el-CURRENT_ESBUILD_VERSION.tgz"></a> |
              | [`@esbuild/linux-ppc64`](https://www.npmjs.org/package/@esbuild/linux-ppc64)       | `linux`              | `ppc64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-ppc64/-/linux-ppc64-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/linux-riscv64`](https://www.npmjs.org/package/@esbuild/linux-riscv64)   | `linux`              | `riscv64`<sup>2</sup>  | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-riscv64/-/linux-riscv64-CURRENT_ESBUILD_VERSION.tgz"></a>   |
              | [`@esbuild/linux-s390x`](https://www.npmjs.org/package/@esbuild/linux-s390x)       | `linux`              | `s390x`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-s390x/-/linux-s390x-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/linux-x64`](https://www.npmjs.org/package/@esbuild/linux-x64)           | `linux`              | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/linux-x64/-/linux-x64-CURRENT_ESBUILD_VERSION.tgz"></a>           |
              | [`@esbuild/netbsd-x64`](https://www.npmjs.org/package/@esbuild/netbsd-x64)         | `netbsd`<sup>1</sup> | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/netbsd-x64/-/netbsd-x64-CURRENT_ESBUILD_VERSION.tgz"></a>         |
              | [`@esbuild/openbsd-arm64`](https://www.npmjs.org/package/@esbuild/openbsd-arm64)   | `openbsd`            | `arm64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/openbsd-arm64/-/openbsd-arm64-CURRENT_ESBUILD_VERSION.tgz"></a>   |
              | [`@esbuild/openbsd-x64`](https://www.npmjs.org/package/@esbuild/openbsd-x64)       | `openbsd`            | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/openbsd-x64/-/openbsd-x64-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/sunos-x64`](https://www.npmjs.org/package/@esbuild/sunos-x64)           | `sunos`              | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/sunos-x64/-/sunos-x64-CURRENT_ESBUILD_VERSION.tgz"></a>           |
              | [`@esbuild/win32-arm64`](https://www.npmjs.org/package/@esbuild/win32-arm64)       | `win32`              | `arm64`                | <a class="dl" href="https://registry.npmjs.org/@esbuild/win32-arm64/-/win32-arm64-CURRENT_ESBUILD_VERSION.tgz"></a>       |
              | [`@esbuild/win32-ia32`](https://www.npmjs.org/package/@esbuild/win32-ia32)         | `win32`              | `ia32`                 | <a class="dl" href="https://registry.npmjs.org/@esbuild/win32-ia32/-/win32-ia32-CURRENT_ESBUILD_VERSION.tgz"></a>         |
              | [`@esbuild/win32-x64`](https://www.npmjs.org/package/@esbuild/win32-x64)           | `win32`              | `x64`                  | <a class="dl" href="https://registry.npmjs.org/@esbuild/win32-x64/-/win32-x64-CURRENT_ESBUILD_VERSION.tgz"></a>           |

          - p: >
              **Why this is not recommended:**
              This approach only works on Unix systems that can run shell scripts, so
              it will require [WSL](https://learn.microsoft.com/en-us/windows/wsl/) on
              Windows. An additional drawback is that you cannot use [plugins](/plugins/)
              with the native version of esbuild.

          - p: >
              If you choose to write your own code to download esbuild directly from
              npm, then you are relying on internal implementation details of esbuild's
              native executable installer. These details may change at some point, in
              which case this approach will no longer work for new esbuild versions. This
              is only a minor drawback though since the approach should still work
              forever for existing esbuild versions (packages published to npm are
              immutable).

          - p: >
              <footer>
              <sup>1</sup> This operating system is not on [node's list of supported platforms](https://nodejs.org/api/process.html#process_process_platform)
              <br>
              <sup>2</sup> This architecture is not on [node's list of supported architectures](https://nodejs.org/api/process.html#processarch)
              </footer>

## #wasm: Install the WASM version

          - p: >
              In addition to the `esbuild` npm package, there is also an `esbuild-wasm`
              package that functions similarly but that uses WebAssembly instead of
              native code. Installing it will also install an executable called `esbuild`:

          - pre: |
              npm install --save-exact esbuild-wasm

          - p: >
              **Why this is not recommended:**
              The WebAssembly version is much, much slower than the native version. In
              many cases it is an order of magnitude (i.e. 10x) slower. This is for
              various reasons including a) node re-compiles the WebAssembly code from
              scratch on every run, b) Go's WebAssembly compilation approach is
              single-threaded, and c) node has WebAssembly bugs that can delay the
              exiting of the process by many seconds. The WebAssembly version also
              excludes some features such as the local file server. You should only
              use the WebAssembly package like this if there is no other option, such
              as when you want to use esbuild on an unsupported platform. The WebAssembly
              package is primarily intended to only be used [in the browser](/api/#browser).

## deno: Deno instead of node

          - p: >
              There is also basic support for the [Deno](https://deno.land) JavaScript
              environment if you'd like to use esbuild with that instead. The package
              is hosted at [https://deno.land/x/esbuild](https://deno.land/x/esbuild)
              and uses the native esbuild executable. The executable will be downloaded
              and cached from npm at run-time so your computer will need network access
              to registry.npmjs.org to make use of this package. Using the package
              looks like this:

          - pre.js: |
              import * as esbuild from 'https://deno.land/x/esbuild@vCURRENT_ESBUILD_VERSION/mod.js'
              let ts = 'let test: boolean = true'
              let result = await esbuild.transform(ts, { loader: 'ts' })
              console.log('result:', result)
              await esbuild.stop()

          - p: >
              It has basically the same API as esbuild's npm package with one addition:
              you need to call `stop()` when you're done because unlike node, Deno doesn't
              provide the necessary APIs to allow Deno to exit while esbuild's internal
              child process is still running.

          - p: >
              If you would like to use esbuild's WebAssembly implementation instead of
              esbuild's native implementation with Deno, you can do that by importing
              `wasm.js` instead of `mod.js` like this:

          - pre.js: |
              import * as esbuild from 'https://deno.land/x/esbuild@vCURRENT_ESBUILD_VERSION/wasm.js'
              let ts = 'let test: boolean = true'
              let result = await esbuild.transform(ts, { loader: 'ts' })
              console.log('result:', result)
              await esbuild.stop()

          - p: >
              Using WebAssembly instead of native means you do not need to specify Deno's
              `--allow-run` permission, and WebAssembly the only option in situations where
              the file system is unavailable such as with [Deno Deploy](https://deno.com/deploy).
              However, keep in mind that the WebAssembly version of esbuild is a lot
              slower than the native version. Another thing to know about WebAssembly
              is that Deno currently has a bug where process termination is unnecessarily
              delayed until all loaded WebAssembly modules are fully optimized, which
              can take many seconds. You may want to manually call `Deno.exit(0)` after
              your code is done if you are writing a short-lived script that uses
              esbuild's WebAssembly implementation so that your code exits in a
              reasonable timeframe.

          - p: >
              **Why this is not recommended:**
              Deno is newer than node, less widely used, and supports fewer platforms
              than node, so node is recommended as the primary way to run esbuild.
              Deno also uses the internet as a package system instead of existing
              JavaScript package ecosystems, and esbuild is designed around and
              optimized for npm-style package management. You should still be able to
              use esbuild with Deno, but you will need a plugin if you would like to
              be able to bundle HTTP URLs.

## Build from source

          - p: >
              To build esbuild from source:

          - ol:
            - |
              Install the Go compiler:
              <p>[https://go.dev/dl/](https://go.dev/dl/)</p>

            - |
              Download the source code for esbuild:
              <pre>
              git clone --depth 1 --branch vCURRENT_ESBUILD_VERSION https://github.com/evanw/esbuild.git
              cd esbuild
              </pre>

            - |
              Build the `esbuild` executable (it will be `esbuild.exe` on Windows):
              <pre>go build ./cmd/esbuild</pre>

          - p: >
              If you want to build for other platforms, you can just prefix the build
              command with the platform information. For example, you can build the
              32-bit Linux version using this command:

          - pre: |
              GOOS=linux GOARCH=386 go build ./cmd/esbuild

          - p: >
              **Why this is not recommended:**
              The native version can only be used via the command-line interface, which
              can be unergonomic for complex use cases and which does not support [plugins](/plugins/).
              You will need to write JavaScript or Go code and use [esbuild's API](/api/)
              to use plugins.
