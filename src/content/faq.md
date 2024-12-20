title: FAQ
body:
  - h1: FAQ

  - p: >
      This is a collection of common questions about esbuild. You can also
      ask questions on the [GitHub issue tracker](https://github.com/evanw/esbuild/issues).

  - toc: true

  - h2: Why is esbuild fast?

  - p: >
      Several reasons:

  - ul:
    - >
      <p>
      It's written in [Go](https://go.dev/) and compiles to native code.
      </p>
      <p>
      Most other bundlers are written in JavaScript, but a command-line
      application is a worst-case performance situation for a JIT-compiled
      language. Every time you run your bundler, the JavaScript VM is seeing
      your bundler's code for the first time without any optimization hints.
      While esbuild is busy parsing your JavaScript, node is busy parsing
      your bundler's JavaScript. By the time node has finished parsing your
      bundler's code, esbuild might have already exited and your bundler
      hasn't even started bundling yet.
      </p>
      <p>
      In addition, Go is designed from the core for parallelism while JavaScript
      is not. Go has shared memory between threads while JavaScript has to
      serialize data between threads. Both Go and JavaScript have parallel
      garbage collectors, but Go's heap is shared between all threads while
      JavaScript has a separate heap per JavaScript thread. This seems to cut
      the amount of parallelism that's possible with JavaScript worker threads
      in half [according to my testing](https://github.com/evanw/esbuild/issues/111#issuecomment-719910381),
      presumably since half of your CPU cores are busy collecting garbage for
      the other half.
      </p>

    - >
      <p>
      Parallelism is used heavily.
      </p>
      <p>
      The algorithms inside esbuild are carefully designed to fully saturate
      all available CPU cores when possible. There are roughly three phases:
      parsing, linking, and code generation. Parsing and code generation are
      most of the work and are fully parallelizable (linking is an inherently
      serial task for the most part). Since all threads share memory, work
      can easily be shared when bundling different entry points that import
      the same JavaScript libraries. Most modern computers have many cores
      so parallelism is a big win.
      </p>

    - >
      <p>
      Everything in esbuild is written from scratch.
      </p>
      <p>
      There are a lot of performance benefits with writing everything yourself
      instead of using 3rd-party libraries. You can have performance in mind
      from the beginning, you can make sure everything uses consistent data
      structures to avoid expensive conversions, and you can make wide
      architectural changes whenever necessary. The drawback is of course that
      it's a lot of work.
      </p>
      <p>
      For example, many bundlers use the official TypeScript compiler as a
      parser. But it was built to serve the goals of the TypeScript compiler team
      and they do not have performance as a top priority. Their code makes pretty
      heavy use of [megamorphic object shapes](https://mrale.ph/blog/2015/01/11/whats-up-with-monomorphism.html)
      and unnecessary [dynamic property accesses](https://github.com/microsoft/TypeScript/issues/39247)
      (both well-known JavaScript speed bumps). And the TypeScript parser
      appears to still run the type checker even when type checking is disabled.
      None of these are an issue with esbuild's custom TypeScript parser.
      </p>

    - >
      <p>
      Memory is used efficiently.
      </p>
      <p>
      Compilers are ideally mostly O(n) complexity in the length of the input.
      So if you are processing a lot of data, memory access speed is likely going
      to heavily affect performance. The fewer passes you have to make over your
      data (and also the fewer different representations you need to transform
      your data into), the faster your compiler will go.
      </p>
      <p>
      For example, esbuild only touches the whole JavaScript AST three times:
      </p>
      <ol>
      <li>A pass for lexing, parsing, scope setup, and declaring symbols</li>
      <li>A pass for binding symbols, minifying syntax, JSX/TS to JS, and ESNext-to-ES2015</li>
      <li>A pass for minifying identifiers, minifying whitespace, generating code, and generating source maps</li>
      </ol>
      <p>
      This maximizes reuse of AST data while it's still hot in the CPU cache.
      Other bundlers do these steps in separate passes instead of interleaving
      them. They may also convert between data representations to glue multiple
      libraries together (e.g. string→TS→JS→string, then string→JS→older JS→string,
      then string→JS→minified JS→string) which uses more memory and slows
      things down.
      </p>
      <p>
      Another benefit of Go is that it can store things compactly in memory,
      which enables it to use less memory and fit more in the CPU cache. All
      object fields have types and fields are packed tightly together so e.g.
      several boolean flags only take one byte each. Go also has value semantics
      and can embed one object directly in another so it comes "for free"
      without another allocation. JavaScript doesn't have these features and
      also has other drawbacks such as JIT overhead (e.g. hidden class slots)
      and inefficient representations (e.g. non-integer numbers are
      heap-allocated with pointers).
      </p>

  - p: >
      Each one of these factors is only a somewhat significant speedup, but
      together they can result in a bundler that is multiple orders of
      magnitude faster than other bundlers commonly in use today.

## Benchmark details

* _Example:_ time to do a production bundle of 10 copies of the [three.js](https://github.com/mrdoob/three.js)
![](statics/index.1.png)

TODO:
  - p: >
      This benchmark approximates a large JavaScript codebase by duplicating
      the [three.js](https://github.com/mrdoob/three.js) library 10 times
      and building a single bundle from scratch, without any caches. The
      benchmark can be run with `make bench-three` in the
      [esbuild repo](https://github.com/evanw/esbuild).

  - table: |
      | Bundler           |    Time | Relative slowdown | Absolute speed | Output size |
      | :---------------- | ------: | ----------------: | -------------: | ----------: |
      | esbuild           |   0.39s |                1x |  1403.7 kloc/s |      5.80mb |
      | parcel 2          |  14.91s |               38x |    36.7 kloc/s |      5.78mb |
      | rollup 4 + terser |  34.10s |               87x |    16.1 kloc/s |      5.82mb |
      | webpack 5         |  41.21s |              106x |    13.3 kloc/s |      5.84mb |

  - p: >
      Each time reported is the best of three runs. I'm running esbuild with
      <code>--bundle <wbr>--minify <wbr>--sourcemap</code>. I used the
      <a href="https://github.com/rollup/plugins/tree/master/packages/terser"><code>@rollup/<wbr>plugin-<wbr>terser</code></a>
      plugin because Rollup itself doesn't support minification. Webpack 5 uses
      <code>--mode=<wbr>production <wbr>--devtool=<wbr>sourcemap</code>.
      Parcel 2 uses the default options. Absolute speed is based on the total
      line count including comments and blank lines, which is currently 547,441.
      The tests were done on a 6-core 2019 MacBook Pro with 16gb of RAM and with
      [macOS Spotlight](https://en.wikipedia.org/wiki/Spotlight_(software)) disabled.

  - figcaption: TypeScript benchmark
  - benchmark:
      '[esbuild](https://esbuild.github.io/)': 0.10
      '[parcel 2](https://parceljs.org/)': 6.91
      '[webpack 5](https://webpack.js.org/)': 16.69

  - p: >
      This benchmark uses the old [Rome](https://github.com/rome/tools) code base
      (prior to their Rust rewrite) to approximate a large TypeScript codebase. All
      code must be combined into a single minified bundle with source maps and the
      resulting bundle must work correctly. The benchmark can be run with
      `make bench-rome` in the [esbuild repo](https://github.com/evanw/esbuild).

  - table: |
      | Bundler   |    Time | Relative slowdown | Absolute speed | Output size |
      | :-------- | ------: | ----------------: | -------------: | ----------: |
      | esbuild   |   0.10s |                1x |  1318.4 kloc/s |      0.97mb |
      | parcel 2  |   6.91ѕ |               69x |    16.1 kloc/s |      0.96mb |
      | webpack 5 |  16.69ѕ |              167x |     8.3 kloc/s |      1.27mb |

  - p: >
      Each time reported is the best of three runs. I'm running esbuild with
      <code>--bundle <wbr>--minify <wbr>--sourcemap <wbr>--platform=<wbr>node</code>.
      Webpack 5 uses [`ts-loader`](https://github.com/TypeStrong/ts-loader) with <code>transpileOnly: <wbr>true</code> and
      <code>--mode=<wbr>production <wbr>--devtool=<wbr>sourcemap</code>. Parcel 2 uses
      <code>"engines": <wbr>"node"</code> in <code>package.json</code>. Absolute
      speed is based on the total line count including comments and blank lines,
      which is currently 131,836. The tests were done on a 6-core 2019 MacBook Pro
      with 16gb of RAM and with [macOS Spotlight](https://en.wikipedia.org/wiki/Spotlight_(software)) disabled.

  - p: >
      The results don't include Rollup because I couldn't get it to work for
      reasons relating to TypeScript compilation. I tried
      <a href="https://github.com/rollup/plugins/tree/master/packages/typescript"><code>@rollup/<wbr>plugin-<wbr>typescript</code></a>
      but you can't disable type checking, and I tried
      <a href="https://github.com/rollup/plugins/tree/master/packages/sucrase"><code>@rollup/<wbr>plugin-<wbr>sucrase</code></a>
      but there's no way to provide a `tsconfig.json` file (which is required
      for correct path resolution).

  - h2: Upcoming roadmap

  - p: >
      These features are already in progress and are first priority:

  - ul:
    - 'Code splitting ([#16](https://github.com/evanw/esbuild/issues/16), [docs](/api/#splitting))'

  - p: >
      These are potential future features but may not happen or may happen
      to a more limited extent:

  - ul:
    - 'HTML content type ([#31](https://github.com/evanw/esbuild/issues/31))'

  - p: >
      After that point, I will consider esbuild to be relatively complete.
      I'm planning for esbuild to reach a mostly stable state and then stop
      accumulating more features. This will involve saying "no" to requests
      for adding major features to esbuild itself. I don't think esbuild
      should become an all-in-one solution for all frontend needs. In
      particular, I want to avoid the pain and problems of the "webpack
      config" model where the underlying tool is too flexible and usability
      suffers.

  - p: >
      For example, I am _not_ planning to include these features in esbuild's
      core itself:

  - ul:
    - 'Support for other frontend languages (e.g. [Elm](https://elm-lang.org/),
      [Svelte](https://svelte.dev/), [Vue](https://vuejs.org/),
      [Angular](https://angular.io/))'
    - TypeScript type checking (just run `tsc` separately)
    - An API for custom AST manipulation
    - Hot-module reloading
    - Module federation

  - p: >
      I hope that the extensibility points I'm adding to esbuild
      ([plugins](/plugins/) and the [API](/api/)) will make esbuild useful to
      include as part of more customized build workflows, but I'm not
      intending or expecting these extensibility points to cover all use
      cases. If you have very custom requirements then you should be using
      other tools. I also hope esbuild inspires other build tools to
      dramatically improve performance by overhauling their implementations
      so that everyone can benefit, not just those that use esbuild.

  - p: >
      I am planning to continue to maintain everything in esbuild's existing
      scope even after esbuild reaches stability. This means implementing
      support for newly-released JavaScript and TypeScript syntax features,
      for example.

  - h2: Production readiness

  - p: >
      This project has not yet hit version 1.0.0 and is still in active
      development. That said, it is far beyond the alpha stage and is pretty
      stable. I think of it as a late-stage beta. For some early-adopters
      that means it's good enough to use for real things. Some other people
      think this means esbuild isn't ready yet. This section doesn't try to
      convince you either way. It just tries to give you enough information
      so you can decide for yourself whether you want to use esbuild as your
      bundler.

  - p: >
      Some data points:

  - ul:
    - >
      **Used by other projects**
      <p>
      The API is already being used as a library within many other
      developer tools. For example, [Vite](https://vitejs.dev/)
      and [Snowpack](https://www.snowpack.dev/) are using
      esbuild to transform TypeScript into JavaScript and
      [Amazon CDK](https://aws.amazon.com/cdk/) (Cloud Development Kit)
      and [Phoenix](https://www.phoenixframework.org/) are using
      esbuild to bundle code.
      </p>

    - >
      **API stability**
      <p>
      Even though esbuild's version is not yet 1.0.0, effort is still made to
      keep the API stable. Patch versions are intended for backwards-compatible
      changes and minor versions are intended for backwards-incompatible changes.
      If you plan to use esbuild for something real, you should either pin the
      exact version (maximum safety) or pin the major and minor versions (only
      accept backwards-compatible upgrades).
      </p>

    - >
      **Only one main developer**
      <p>
      This tool is primarily built by [me](https://github.com/evanw). For
      some people this is fine, but for others this means esbuild is not a
      suitable tool for their organization. That's ok with me. I'm building
      esbuild because I find it fun to build and because it's the tool I'd
      want to use. I'm sharing it with the world because there are others
      that want to use it too, because the feedback makes the tool itself
      better, and because I think it will inspire the ecosystem to make
      better tools.
      </p>

    - >
      **Not always open to scope expansion**
      <p>
      I'm not planning on including major features that I'm not interested
      in building and/or maintaining. I also want to limit the project's
      scope so it doesn't get too complex and unwieldy, both from an
      architectural perspective, a testing and correctness perspective, and
      from a usability perspective. Think of esbuild as a "linker" for the
      web. It knows how to transform and bundle JavaScript and CSS. But the
      details of how your source code ends up as plain JavaScript or CSS
      may need to be 3rd-party code.
      </p>
      <p>
      I'm hoping that [plugins](/plugins/) will allow the community to add
      major features (e.g. WebAssembly import) without needing to contribute
      to esbuild itself. However, not everything is exposed in the plugin
      API and it may be the case that it's not possible to add a particular
      feature to esbuild that you may want to add. This is intentional;
      esbuild is not meant to be an all-in-one solution for all frontend
      needs.
      </p>

  - h2: Anti-virus software

  - p: >
      Since esbuild is written in native code, anti-virus software can sometimes
      incorrectly flag it as a virus. _This does not mean esbuild is a virus._ I
      do not publish malicious code and I take supply chain security very seriously.

  - p: >
      Virtually all of esbuild's code is first-party code except for [one dependency](https://github.com/evanw/esbuild/blob/main/go.mod)
      on Google's set of supplemental Go packages. My development work is done
      on different machine that is isolated from the one I use to publish builds.
      I have done additional work to ensure that esbuild's published builds are
      completely reproducible and after every release, published builds are
      [automatically compared](https://github.com/evanw/esbuild/blob/main/.github/workflows/validate.yml)
      to ones locally-built in an unrelated environment to ensure that they are
      bitwise identical (i.e. that the Go compiler itself has not been compromised).
      You can also build esbuild from source yourself and compare your build artifacts
      to the published ones to independently verify this.

  - p: >
      Having to deal with false-positives is an unfortunate reality of using
      anti-virus software. Here are some possible workarounds if your anti-virus
      won't let you use esbuild:

  - ul:
    - >
      Ignore your anti-virus software and remove esbuild from quarantine

    - >
      Report the specific esbuild native executable as a false-positive to your
      anti-virus software vendor

    - >
      Use [`esbuild-wasm`](/getting-started/#wasm) instead of `esbuild` to
      bypass your anti-virus software (which likely won't flag WebAssembly
      files the same way it flags native executables)

    - >
      Use another build tool instead of esbuild

  - h2#old-go-version: Outdated version of Go

  - p: >
      If you use an automated dependency vulnerability scanner, you may get a
      report that the version of the Go compiler that esbuild uses and/or the
      version of `golang.org/x/sys` (esbuild's only dependency) is outdated.
      These reports are benign and should be ignored.

  - p: >
      This happens because esbuild's code is deliberately intended to be
      compilable with Go 1.13. Later versions of Go have dropped support for
      certain older platforms that I want esbuild to be able to run on (e.g.
      older versions of macOS). While esbuild's published binaries are compiled
      with a much newer version of the Go compiler (and therefore don't work
      on older versions of macOS), you are currently still able to compile the
      latest version of esbuild for yourself with Go 1.13 and use it on older
      versions of macOS because esbuild's code can still be compiled with Go
      as far back as 1.13.

  - p: >
      People and/or automated tools sometimes see the `go 1.13` line in [`go.mod`](https://github.com/evanw/esbuild/blob/main/go.mod)
      and complain that esbuild's published binaries are built with Go 1.13, which
      is a really old version of Go. However, that's not true. That line in `go.mod`
      only specifies the minimum compiler version. It has nothing to do with the
      version of Go that esbuild's published binaries are built with, which is a
      much newer version of Go. [Please read the documentation.](https://go.dev/ref/mod#go-mod-file-go)

  - p: >
      People also sometimes want esbuild to update the `golang.org/x/sys` dependency
      because there is a known vulnerability in the version that esbuild uses
      (specifically [GO-2022-0493](https://pkg.go.dev/vuln/GO-2022-0493)
      about the `Faccessat` function). The problem that prevents esbuild from
      updating to a newer version of the `golang.org/x/sys` dependency is that
      newer versions have started using the `unsafe.Slice` function, which was
      first introduced in Go 1.17 (and therefore doesn't compile in older
      versions of Go). However, this vulnerability report is irrelevant because
      a) esbuild doesn't ever call that function in the first place and b)
      esbuild is a build tool, not a sandbox, and esbuild's file system access
      is not security-sensitive.

  - p: >
      I'm not going to drop compatibility with older platforms and prevent some
      people from being able to use esbuild just to work around irrelevant
      vulnerability reports. Please ignore any reports about the issues described
      above.

  - h2: Minified newlines

  - p: >
      People are sometimes surprised that esbuild's minifier typically changes
      the character escape sequence `\n` within JavaScript strings into a
      newline character in a template literal. But this is intentional.
      **This is not a bug with esbuild**. The job of a minifier is to generate as
      compact an output as possible that's equivalent to the input. The character
      escape sequence `\n` is two bytes long while a newline character is one
      byte long.

  - p: >
      For example, this code is 21 bytes long:

  - pre.js: |
      var text="a\nb\nc\n";

  - p: >
      While this code is 18 bytes long:

  - pre.js: |
      var text=`a
      b
      c
      `;

  - p: >
      So the second code is fully minified while the first one isn't. Minifying
      code does not mean putting it all on one line. Instead, minifying code
      means generating equivalent code that uses as few bytes as possible. In
      JavaScript, an untagged template literal is equivalent to a string literal,
      so esbuild is doing the correct thing here.

  - h2#top-level-name-collisions: Avoiding name collisions

  - p: >
      Top-level variables in an entry point module should never end up in the
      global scope when running esbuild's output in a browser. If that happens,
      it means you did not follow [esbuild's documentation about output formats](/api/#format)
      and are using esbuild incorrectly. **This is not a bug with esbuild.**

  - p: >
      Specifically, you must do either one of the following when running
      esbuild's output in a browser:

  - ol:
    - >
      <code>--format=<wbr>iife</code> with <code>&lt;script <wbr>src="..."&gt;</code>
      <p>
      If you are running your code in the global scope, then you should be using
      <code>--format=<wbr>iife</code>. This causes esbuild's output to wrap your
      code so that top-level variables are declared in a nested scope.
      </p>

    - >
      <code>--format=<wbr>esm</code> with <code>&lt;script <wbr>src="..." <wbr>type="module"&gt;</code>
      <p>
      If you are using <code>--format=<wbr>esm</code>, then you must run your
      code as a module. This causes the browser to wrap your code so that
      top-level variables are declared in a nested scope.
      </p>

  - p: >
      Using <code>--format=<wbr>esm</code> with <code>&lt;script <wbr>src="..."&gt;</code>
      will break your code in subtle and confusing ways (omitting
      <code>type="<wbr>module"</code> means that all top-level variables will end up
      in the global scope, which will then collide with top-level variables
      that have the same name in other JavaScript files).

  - h2: Top-level `var`

  - p: >
      People are sometimes surprised that esbuild sometimes rewrites top-level
      `let`, `const`, and `class` declarations as `var` declarations instead.
      This is done for a few reasons:

  - ul:
    - >
      **For correctness**
      <p>
      Bundling sometimes needs to lazily-initialize a module. For example, this
      happens when you call `require()` or `import()` using the path of a module
      within the bundle. Doing this involves separating the declaration and
      initialization of top-level symbols by moving the initialization into a
      closure. So for example `class` statements are rewritten as an assignment
      of a class expression to a variable. Keeping the declarations out of the
      lazy-initialization closure is important for performance, since it means
      other modules can reference them directly instead by name instead of
      indirectly via a slower property access.
      </p>
      <p>
      Another case where this is needed is when transforming top-level `using`
      declarations. This involves wrapping the entire module body in a `try`
      block, which also involves separating the declaration and initialization
      of top-level symbols. Top-level symbols may need to be exported, which
      means they cannot be declared within the `try` block.
      </p>
      <p>
      In both of these cases esbuild will fail with a build error if the source
      code contains a mutation of a `const` symbol, so it's not possible for
      esbuild's rewriting of top-level `const` into `var` to result in the
      mutation of a constant.
      </p>
      <p>
      Due to esbuild's current architecture, the part of esbuild that does this
      transformation (the parser) cannot know whether the current module will
      end up being lazily initialized or not. The information for this decision
      may only be discovered later on in the build, or may even change in future
      incremental builds that reuse the same AST (per-file ASTs are transformed
      once during parsing and then cached and reused across incremental builds).
      So this transformation is always done when bundling is active.
      </p>

    - >
      **For performance**
      <p>
      Multiple JavaScript VMs have had and continue to have performance issues
      with TDZ (i.e. "temporal dead zone") checks. These checks validate that a
      let, or const, or class symbol isn't used before it's initialized. Here
      are two issues with well-known VMs:
      </p>
      <ul>
        <li>
          V8: <a href="https://bugs.chromium.org/p/v8/issues/detail?id=13723">https://bugs.chromium.org/p/v8/issues/detail?id=13723</a> (10% slowdown)
        </li>
        <li>
          JavaScriptCore: <a href="https://bugs.webkit.org/show_bug.cgi?id=199866">https://bugs.webkit.org/show_bug.cgi?id=199866</a> (1,000% slowdown!)
        </li>
      </ul>
      <p>
      JavaScriptCore had a severe performance issue as their TDZ implementation
      had time complexity that was quadratic in the number of variables needing
      TDZ checks in the same scope (with the top-level scope typically being the
      worst offender). V8 has ongoing issues with TDZ checks being present
      throughout the code their JIT generates even when they have already been
      checked earlier in the same function or when the function in question has
      already been run (so the checks have already happened).
      </p>
      <p>
      In JavaScript, `let`, `const`, and `class` declarations all introduce TDZ
      checks while `var` declarations do not. Since bundling typically merges
      many modules into a single very large top-level scope, the performance
      impact of these TDZ checks can be pretty severe. Converting top-level
      `let`, `const`, and `class` declarations into `var` helps automatically
      make your code faster.
      </p>

  - p: >
      Note that esbuild doesn't preserve top-level TDZ side effects because
      modules may need to be lazily initialized (as described above), which
      means separating declaration from initialization. TDZ checks for
      top-level symbols could hypothetically still be supported by generating
      extra code that checks before each use of a top-level symbol and throws
      if it hasn't been initialized yet (effectively manually implementing
      what a real JavaScript VM would do). However, this seems like an
      excessive overhead for both code size and run time, and does not seem
      like something that a production-oriented bundler should do.
