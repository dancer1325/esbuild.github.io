title: Content Types
body:
  - h1: Content Types
  - p: >
      All of the built-in content types are listed below. Each content type
      has an associated "loader" which tells esbuild how to interpret the
      file contents. Some file extensions already have a loader configured
      for them by default, although the defaults can be overridden.

## JavaScript
  - p: 'Loader: `js`'

  - p: >
      This loader is enabled by default for `.js`, `.cjs`, and `.mjs` files.
      The `.cjs` extension is used by node for CommonJS modules and the `.mjs`
      extension is used by node for ECMAScript modules.

  - p: >
      Note that by default, esbuild's output will take advantage of all modern
      JS features. For example, <code>a !== <wbr>void 0 && <wbr>a !== <wbr>null ? <wbr>a : b</code>
      will become <code>a ?? b</code> when minifying is enabled which makes use
      of syntax from the [ES2020](https://262.ecma-international.org/11.0/#prod-CoalesceExpression)
      version of JavaScript. If this is undesired, you must specify esbuild's
      [target](/api/#target) setting to say in which browsers you need the
      output to work correctly. Then esbuild will avoid using JavaScript
      features that are too modern for those browsers.

  - p: >
      All modern JavaScript syntax is supported by esbuild. Newer syntax may
      not be supported by older browsers, however, so you may want to
      configure the [target](/api/#target) option to tell esbuild to convert
      newer syntax to older syntax as appropriate.

  - p: >
      These syntax features are always transformed for older browsers:

  - table: |
      | Syntax transform                                                                                                   | Language version | Example       |
      |--------------------------------------------------------------------------------------------------------------------|------------------|---------------|
      | [Trailing commas in function parameter lists and calls](https://github.com/tc39/proposal-trailing-function-commas) | `es2017`         | `foo(a, b, )` |
      | [Numeric separators](https://github.com/tc39/proposal-numeric-separator)                                           | `esnext`         | `1_000_000`   |

  - p: >
      These syntax features are conditionally transformed for older browsers depending on the configured language [target](/api/#target):

  - table: |
      | Syntax transform                                                                      | Transformed when `--target` is below | Example                            |
      |---------------------------------------------------------------------------------------|--------------------------------------|------------------------------------|
      | [Exponentiation operator](https://github.com/tc39/proposal-exponentiation-operator)   | `es2016`                             | `a ** b`                           |
      | [Async functions](https://github.com/tc39/ecmascript-asyncawait)                      | `es2017`                             | `async () => {}`                   |
      | [Asynchronous iteration](https://github.com/tc39/proposal-async-iteration)            | `es2018`                             | `for await (let x of y) {}`        |
      | [Async generators](https://github.com/tc39/proposal-async-iteration)                  | `es2018`                             | `async function* foo() {}`         |
      | [Spread properties](https://github.com/tc39/proposal-object-rest-spread)              | `es2018`                             | `let x = {...y}`                   |
      | [Rest properties](https://github.com/tc39/proposal-object-rest-spread)                | `es2018`                             | `let {...x} = y`                   |
      | [Optional catch binding](https://github.com/tc39/proposal-optional-catch-binding)     | `es2019`                             | `try {} catch {}`                  |
      | [Optional chaining](https://github.com/tc39/proposal-optional-chaining)               | `es2020`                             | `a?.b`                             |
      | [Nullish coalescing](https://github.com/tc39/proposal-nullish-coalescing)             | `es2020`                             | `a ?? b`                           |
      | [`import.meta`](https://github.com/tc39/proposal-import-meta)                         | `es2020`                             | `import.meta`                      |
      | [Logical assignment operators](https://github.com/tc39/proposal-logical-assignment)   | `es2021`                             | `a ??= b`                          |
      | [Class instance fields](https://github.com/tc39/proposal-class-fields)                | `es2022`                             | `class { x }`                      |
      | [Static class fields](https://github.com/tc39/proposal-static-class-features)         | `es2022`                             | `class { static x }`               |
      | [Private instance methods](https://github.com/tc39/proposal-private-methods)          | `es2022`                             | `class { #x() {} }`                |
      | [Private instance fields](https://github.com/tc39/proposal-class-fields)              | `es2022`                             | `class { #x }`                     |
      | [Private static methods](https://github.com/tc39/proposal-static-class-features)      | `es2022`                             | `class { static #x() {} }`         |
      | [Private static fields](https://github.com/tc39/proposal-static-class-features)       | `es2022`                             | `class { static #x }`              |
      | [Ergonomic brand checks](https://github.com/tc39/proposal-private-fields-in-in)       | `es2022`                             | `#x in y`                          |
      | [Class static blocks](https://github.com/tc39/proposal-class-static-block)            | `es2022`                             | `class { static {} }`              |
      | [Import assertions](https://github.com/tc39/proposal-import-assertions)               | `esnext`                             | `import "x" assert {}`<sup>1</sup> |
      | [Import attributes](https://github.com/tc39/proposal-import-attributes)               | `esnext`                             | `import "x" with {}`               |
      | [Auto-accessors](https://github.com/tc39/proposal-decorators#class-auto-accessors)    | `esnext`                             | `class { accessor x }`             |
      | [`using` declarations](https://github.com/tc39/proposal-explicit-resource-management) | `esnext`                             | `using x = y`                      |
      | [Decorators](https://github.com/tc39/proposal-decorators)                             | `esnext`                             | `@foo class Bar {}`                |

  - p: >
      <footer>
      <sup>1</sup> Import assertions never made it into the JavaScript
      specification. They are deprecated in favor of import attributes
      and are actively being removed from JavaScript runtimes.
      </footer>

  - p: >
      These syntax features are currently always passed through un-transformed:

  - table: |
      | Syntax transform                                                                                                       | Unsupported when `--target` is below | Example                     |
      |------------------------------------------------------------------------------------------------------------------------|--------------------------------------|-----------------------------|
      | [RegExp `dotAll` flag](https://github.com/tc39/proposal-regexp-dotall-flag)                                            | `es2018`                             | `/./s`<sup>1</sup>          |
      | [RegExp lookbehind assertions](https://github.com/tc39/proposal-regexp-lookbehind)                                     | `es2018`                             | `/(?<=x)y/`<sup>1</sup>     |
      | [RegExp named capture groups](https://github.com/tc39/proposal-regexp-named-groups)                                    | `es2018`                             | `/(?<foo>\d+)/`<sup>1</sup> |
      | [RegExp unicode property escapes](https://github.com/tc39/proposal-regexp-unicode-property-escapes)                    | `es2018`                             | `/\p{ASCII}/u`<sup>1</sup>  |
      | [BigInt](https://github.com/tc39/proposal-bigint)                                                                      | `es2020`                             | `123n`                      |
      | [Top-level await](https://github.com/tc39/proposal-top-level-await)                                                    | `es2022`                             | `await import(x)`           |
      | [Arbitrary module namespace identifiers](https://github.com/bmeck/proposal-arbitrary-module-namespace-identifiers)     | `es2022`                             | `export {foo as 'f o o'}`   |
      | [RegExp match indices](https://github.com/tc39/proposal-regexp-match-indices)                                          | `es2022`                             | `/x(.+)y/d`<sup>1</sup>     |
      | [RegExp set notation](https://github.com/tc39/proposal-regexp-v-flag)                                                  | `es2024`                             | `/[\w--\d]/v`<sup>1</sup>   |
      | [Hashbang grammar](https://github.com/tc39/proposal-hashbang)                                                          | `esnext`                             | `#!/usr/bin/env node`       |

  - p: >
      <footer>
      <sup>1</sup> Unsupported regular expression literals are transformed into a `new RegExp()`
      constructor call so you can bring your own polyfill library to get them to work anyway.
      </footer>

  - p: >
      See also [the list of finished ECMAScript proposals](https://github.com/tc39/proposals/blob/main/finished-proposals.md)
      and [the list of active ECMAScript proposals](https://github.com/tc39/proposals/blob/main/README.md).
      Note that while transforming code containing top-level await is
      supported, bundling code containing top-level await is only supported
      when the [output format](/api/#format) is set to [`esm`](/api/#format-esm).

  - h3: JavaScript caveats

  - p: >
      You should keep the following things in mind when using JavaScript with
      esbuild:

  - h4#es5: ES5 is not supported well
  - p: >
      Transforming ES6+ syntax to ES5 is not supported yet. However, if you're
      using esbuild to transform ES5 code, you should still set the
      [target](/api/#target) to `es5`. This prevents esbuild from introducing
      ES6 syntax into your ES5 code. For example, without this flag the object
      literal `{x: x}` will become `{x}` and the string `"a\nb"` will become
      a multi-line template literal when minifying. Both of these substitutions
      are done because the resulting code is shorter, but the substitutions
      will not be performed if the [target](/api/#target) is `es5`.

  - h4: Private member performance
  - p: >
      The private member transform (for the `#name` syntax) uses `WeakMap` and
      `WeakSet` to preserve the privacy properties of this feature. This is
      similar to the corresponding transforms in the Babel and TypeScript
      compilers. Most modern JavaScript engines (V8, JavaScriptCore, and
      SpiderMonkey but not ChakraCore) may not have good performance
      characteristics for large `WeakMap` and `WeakSet` objects.
  - p: >
      Creating many instances of classes with private fields or private
      methods with this syntax transform active may cause a lot of overhead
      for the garbage collector.
      This is because modern engines (other than ChakraCore) store weak
      values in an actual map object instead of as hidden properties on the
      keys themselves, and large map objects can cause performance issues with
      garbage collection. See [this reference](https://github.com/tc39/ecma262/issues/1657#issuecomment-518916579)
      for more information.

  - h4#real-esm-imports: Imports follow ECMAScript module behavior
  - p: >
      You might try to modify global state before importing a module which
      needs that global state and expect it to work. However, JavaScript
      (and therefore esbuild) effectively "hoists" all `import` statements
      to the top of the file, so doing this won't work:

  - pre.js: |
      window.foo = {}
      import './something-that-needs-foo'

  - p: >
      There are some broken implementations of ECMAScript modules out there
      (e.g. the TypeScript compiler) that don't follow the JavaScript
      specification in this regard. Code compiled with these tools may "work"
      since the `import` is replaced with an inline call to `require()`, which
      ignores the hoisting requirement. But such code will not work with real
      ECMAScript module implementations such as node, a browser, or esbuild,
      so writing code like this is non-portable and is not recommended.

  - p: >
      The way to do this correctly is to move the global state modification
      into its own import. That way it _will_ be run before the other import:

  - pre.js: |
      import './assign-to-foo-on-window'
      import './something-that-needs-foo'

  - h4#direct-eval: Avoid direct `eval` when bundling
  - p: >
      Although the expression `eval(x)` looks like a normal function call, it
      actually takes on special behavior in JavaScript. Using `eval` in this way
      means that the evaluated code stored in `x` can reference any variable in
      any containing scope by name. For example, the code <code>let y = <wbr>123; <wbr>return <wbr>eval('y')</code>
      will return `123`.
  - p: >
      This is called "direct eval" and is problematic when bundling your code
      for many reasons:
  - ul:
    - >
      <p>
      Modern bundlers contain an optimization called "scope hoisting" that
      merges all bundled files into a single file and renames variables to
      avoid name collisions. However, this means code evaluated by direct
      `eval` can read and write variables in any file in the bundle! This
      is a correctness issue because the evaluated code may try to access
      a global variable but may accidentally access a private variable
      with the same name from another file instead. It can potentially even
      be a security issue if a private variable in another file has sensitive
      data.
      </p>
    - >
      <p>
      The evaluated code may not work correctly when it references variables
      imported using an `import` statement. Imported variables are live
      bindings to variables in another file. They are not copies of those
      variables. So when esbuild bundles your code, your imports are replaced
      with a direct reference to the variable in the imported file. But that
      variable may have a different name, in which case the code evaluated by
      direct `eval` will be unable to reference it by the expected name.
      </p>
    - >
      <p>
      Using direct `eval` forces esbuild to deoptimize all of the code in
      all of the scopes containing calls to direct `eval`. For correctness,
      it must assume that the evaluated code might need to access any of the
      other code in the file reachable from that `eval` call. This means
      none of that code will be eliminated as dead code and none of that code
      will be minified.
      </p>
    - >
      <p>
      Because the code evaluated by the direct `eval` could need to reference
      any reachable variable by name, esbuild is prevented from renaming
      all of the variables reachable by the evaluated code. This means it
      can't rename variables to avoid name collisions with other variables
      in the bundle. So the direct `eval` causes esbuild to wrap the file
      in a CommonJS closure, which avoids name collisions by introducing
      a new scope instead. However, this makes the generated code bigger and
      slower because exported variables use run-time dynamic binding instead
      of compile-time static binding.
      </p>
  - p: >
      Luckily it is usually easy to avoid using direct `eval`. There are two
      commonly-used alternatives that avoid all of the drawbacks mentioned
      above:
  - ul:
    - >
      `(0, eval)('x')`
      <p>
      This is known as "indirect eval" because `eval` is not being called
      directly, and so does not trigger the grammatical special case for
      direct eval in the JavaScript VM. You can call indirect eval using any
      syntax at all except for an expression of the exact form `eval('x')`.
      For example, <code>var eval2 = <wbr>eval; <wbr>eval2('x')</code> and
      <code>[eval]\[0]('x')</code> and <code>window.<wbr>eval('x')</code> are
      all indirect eval calls. When you use indirect eval, the code is
      evaluated in the global scope instead of in the inline scope of the
      caller.
      </p>
    - >
      `new Function('x')`
      <p>
      This constructs a new function object at run-time. It is as if you wrote
      <code>function() <wbr>{ x }</code> in the global scope except that `x`
      can be an arbitrary string of code. This form is sometimes convenient
      because you can add arguments to the function, and use those arguments
      to expose variables to the evaluated code. For example,
      <code>(new Function('env', <wbr>'x'))(<wbr>someEnv)</code> is as if you
      wrote <code>(function(env) <wbr>{ x })(<wbr>someEnv)</code>. This is
      often a sufficient alternative for direct `eval` when the evaluated
      code needs to access local variables because you can pass the local
      variables in as arguments.
      </p>

  - h4#function-toString: >
      The value of `toString()` is not preserved on functions (and classes)
  - p: >
      It's somewhat common to call `toString()` on a JavaScript function
      object and then pass that string to some form of `eval` to get a new
      function object. This effectively "rips" the function out of the
      containing file and breaks links with all variables in that file.
      Doing this with esbuild is not supported and may not work. In
      particular, esbuild often uses helper methods to implement certain
      features and it assumes that JavaScript scope rules have not been
      tampered with. For example:
  - pre.js: |
      let pow = (a, b) => a ** b;
      let pow2 = (0, eval)(pow.toString());
      console.log(pow2(2, 3));
  - p: >
      When this code is compiled for ES6, where the `**` operator isn't
      available, the `**` operator is replaced with a call to the `__pow`
      helper function:
  - pre.js: |
      let __pow = Math.pow;
      let pow = (a, b) => __pow(a, b);
      let pow2 = (0, eval)(pow.toString());
      console.log(pow2(2, 3));
  - p: >
      If you try to run this code, you'll get an error such as
      <code>ReferenceError: <wbr>__pow <wbr>is <wbr>not <wbr>defined</code>
      because the function <code>(a, b) <wbr>=&gt; <wbr>__pow(a, b)</code>
      depends on the locally-scoped symbol `__pow` which is not available
      in the global scope. This is the case for many JavaScript language
      features including `async` functions, as well as some esbuild-specific
      features such as the [keep names](/api/#keep-names) setting.
  - p: >
      This problem most often comes up when people get the source code of a
      function with `.toString()` and then try to use it as the body of a
      [web worker](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API).
      If you are doing this and you want to use esbuild, you should instead build
      the source code for the web worker in a separate build step and then
      insert the web worker source code as a string into the code that
      creates the web worker. The [define](/api/#define) feature is one way
      to insert the string at build time.

  - h4#module-namespace-this: >
      The value of `this` is not preserved on functions called from a module namespace object
  - p: >
      In JavaScript, the value of `this` in a function is automatically filled
      in for you based on how the function is called. For example if a function
      is called using `obj.fn()`, the value of `this` during the function call
      will be `obj`. This behavior is respected by esbuild with one exception:
      if you call a function from a module namespace object, the value of `this`
      may not be correct. For example, consider this code that calls `foo` from
      the module namespace object `ns`:
  - pre.js: |
      import * as ns from './foo.js'
      ns.foo()
  - p: >
      If `foo.js` tries to reference the module namespace object using `this`,
      then it won't necessarily work after the code is bundled with esbuild:
  - pre.js: |
      // foo.js
      export function foo() {
        this.bar()
      }
      export function bar() {
        console.log('bar')
      }
  - p: >
      The reason for this is that esbuild automatically rewrites code most
      code that uses module namespace objects to code that imports things
      directly instead. That means the example code above will be converted
      to this instead, which removes the `this` context for the function call:
  - pre.js: |
      import { foo } from './foo.js'
      foo()
  - p: >
      This transformation dramatically improves [tree shaking](/api/#tree-shaking)
      (a.k.a. dead code elimination) because it makes it possible for esbuild
      to understand which exported symbols are unused. It has the drawback
      that this changes the behavior of code that uses `this` to access the
      module's exports, but this isn't an issue because no one should ever write
      bizarre code like this in the first place. If you need to access an exported
      function from the same file, just call it directly (i.e. `bar()` instead
      of `this.bar()` in the example above).

  - h4#default-interop: >
      The `default` export can be error-prone
  - p: >
      The ES module format (i.e. ESM) have a special export called `default`
      that sometimes behaves differently than all other export names. When
      code in the ESM format that has a `default` export is converted to the
      CommonJS format, and then that CommonJS code is imported into another
      module in ESM format, there are two different interpretations of what
      should happen that are both widely-used (the [Babel](https://babeljs.io/)
      way and the [Node](https://nodejs.org/) way). This is very unfortunate
      because it causes endless compatibility headaches, especially since
      JavaScript libraries are often authored in ESM and published as
      CommonJS.

  - p: >
      When esbuild [bundles](/api/#bundle) code that does this, it has to
      decide which interpretation to use, and there's no perfect answer. The
      heuristics that esbuild uses are the same heuristics that [Webpack](https://webpack.js.org/)
      uses (see below for details). Since Webpack is the most widely-used
      bundler, this means that esbuild is being the most compatible that it can
      be with the existing ecosystem regarding this compatibility problem. So
      the good news is that if you can get code with this problem to work with
      esbuild, it should also work with Webpack.

  - p: >
      Here's an example that demonstrates the problem:

  - pre.js: |
      // index.js
      import foo from './somelib.js'
      console.log(foo)

  - pre.js: |
      // somelib.js
      Object.defineProperty(exports, "__esModule", {
        value: true
      });
      exports["default"] = 'foo';

  - p: >
      And here are the two interpretations, both of which are widely-used:

  - ul:
    - |
      <p>**The Babel interpretation**</p>

      <p>
      If the Babel interpretation is used, this code will print `foo`. Their
      rationale is that `somelib.js` was converted from ESM into CommonJS (as
      you can tell by the `__esModule` marker) and the original code looked
      something like this:
      </p>

      <pre>
      // somelib.js
      export default 'foo'
      </pre>

      <p>
      If `somelib.js` hadn't been converted from ESM into CommonJS, then this
      code would print `foo`, so it should still print `foo` regardless of the
      module format. This is accomplished by detecting when a CommonJS module
      used to be an ES module via the `__esModule` marker (which all module
      conversion tools set including Babel, TypeScript, Webpack, and esbuild)
      and setting the default import to <code>exports.<wbr>default</code> if
      the `__esModule` marker is present. This behavior is important because
      it's necessary to run cross-compiled ESM correctly in a CommonJS environment,
      and for a long time that was the only way to run ESM code in Node before
      Node eventually added native ESM support.
      </p>

    - >
      <p>**The Node interpretation**</p>

      <p>
      If the Node interpretation is used, this code will print <code>{ default: <wbr>'foo' }</code>.
      Their rationale is that CommonJS code uses dynamic exports while ESM code
      uses static exports, so the fully general approach to importing CommonJS
      into ESM is to expose the CommonJS `exports` object itself somehow. For
      example, CommonJS code can do <code>exports[<wbr>Math.<wbr>random()<wbr>] =<wbr> 'foo'</code>
      which has no equivalent in ESM syntax. The `default` export is used for
      this because that's actually what it was originally designed for by the
      people who came up with the ES module specification. This interpretation
      is entirely reasonable for normal CommonJS modules. It only causes
      compatibility problems for CommonJS modules that used to be ES modules
      (i.e. when `__esModule` is present) in which case the behavior diverges
      from the Babel interpretation.
      </p>

  - p: >
      *If you are a library author:* When writing new code, you should strongly
      consider avoiding the `default` export entirely. It has unfortunately been
      tainted with compatibility problems and using it will likely cause problems
      for your users at some point.

  - p: >
      *If you are a library user:* By default, esbuild will use the Babel interpretation.
      If you want esbuild to use the Node interpretation instead, you need to
      either put your code in a file ending in `.mts` or `.mjs`, or you need to
      add <code>"type": <wbr>"module"</code> to your `package.json` file. The
      rationale is that Node's native ESM support can only run ESM code if the
      file extension is `.mjs` or <code>"type": <wbr>"module"</code> is present,
      so doing that is a good signal that the code is intended to be run in Node,
      and should therefore use the Node interpretation of `default` import. This
      is the same heuristic that Webpack uses.

## TypeScript
  - p: 'Loader: `ts` or `tsx`'

  - p: >
      This loader is enabled by default for `.ts`, `.tsx`, `.mts`, and `.cts`
      files, which means esbuild has built-in support for parsing TypeScript
      syntax and discarding the type annotations. However, esbuild _does not_
      do any type checking so you will still need to run `tsc -noEmit` in
      parallel with esbuild to check types. This is not something esbuild does
      itself.

  - p: >
      TypeScript type declarations like these are parsed and ignored (a
      non-exhaustive list):

  - table: |
      | Syntax feature              | Example                         |
      |-----------------------------|---------------------------------|
      | Interface declarations      | `interface Foo {}`              |
      | Type declarations           | `type Foo = number`             |
      | Function declarations       | `function foo(): void;`         |
      | Ambient declarations        | `declare module 'foo' {}`       |
      | Type-only imports           | `import type {Type} from 'foo'` |
      | Type-only exports           | `export type {Type} from 'foo'` |
      | Type-only import specifiers | `import {type Type} from 'foo'` |
      | Type-only export specifiers | `export {type Type} from 'foo'` |

  - p: >
      TypeScript-only syntax extensions are supported, and are always
      converted to JavaScript (a non-exhaustive list):

  - table: |
      | Syntax feature            | Example                    | Notes |
      |---------------------------|----------------------------|-------|
      | Namespaces                | `namespace Foo {}`         | |
      | Enums                     | `enum Foo { A, B }`        | |
      | Const enums               | `const enum Foo { A, B }`  | |
      | Generic type parameters   | `<T>(a: T): T => a`        | Must write `<T,>(`... with the `tsx` loader |
      | JSX with types            | `<Element<T>/>`            | |
      | Type casts                | `a as B` and `<B>a`        | |
      | Type imports              | `import {Type} from 'foo'` | Handled by removing all unused imports |
      | Type exports              | `export {Type} from 'foo'` | Handled by ignoring missing exports in TypeScript files |
      | Experimental decorators   | `@sealed class Foo {}`     | Requires [`experimentalDecorators`](https://www.typescriptlang.org/tsconfig#experimentalDecorators), <br>does not support [`emitDecoratorMetadata`](https://www.typescriptlang.org/tsconfig#emitDecoratorMetadata) |
      | Instantiation expressions | `Array<number>`            | TypeScript 4.7+ |
      | `extends` on `infer`      | `infer A extends B`        | TypeScript 4.7+ |
      | Variance annotations      | `type A<out B> = () => B`  | TypeScript 4.7+ |
      | The `satisfies` operator  | `a satisfies T`            | TypeScript 4.9+ |
      | `const` type parameters   | `class Foo<const T> {}`    | TypeScript 5.0+ |

  - h3: TypeScript caveats

  - p: >
      You should keep the following things in mind when using TypeScript with
      esbuild (in addition to the [JavaScript caveats](#javascript-caveats)):

  - h4#isolated-modules: Files are compiled independently
  - p: >
      Even when transpiling a single module, the TypeScript compiler actually
      still parses imported files so it can tell whether an imported name is
      a type or a value. However, tools like esbuild and Babel (and the
      TypeScript compiler's `transpileModule` API) compile each file in
      isolation so they can't tell if an imported name is a type or a value.
  - p: >
      Because of this, you should enable the [`isolatedModules`](https://www.typescriptlang.org/tsconfig#isolatedModules)
      TypeScript configuration option if you use TypeScript with esbuild.
      This option prevents you from using features which could cause mis-compilation
      in environments like esbuild where each file is compiled independently
      without tracing type references across files. For example, it prevents
      you from re-exporting types from another module using <code>export <wbr>{T} <wbr>from <wbr>'./types'</code>
      (you need to use <code>export <wbr>type <wbr>{T} <wbr>from <wbr>'./types'</code> instead).

  - h4#es-module-interop: Imports follow ECMAScript module behavior
  - p: >
      For historical reasons, the TypeScript compiler compiles ESM
      (ECMAScript module) syntax to CommonJS syntax by default. For example,
      <code>import *<wbr> as foo<wbr> from 'foo'</code> is compiled to
      <code>const foo =<wbr> require('foo')</code>. Presumably this happened
      because ECMAScript modules were still a proposal when TypeScript adopted
      the syntax. However, this is legacy behavior that doesn't match how this
      syntax behaves on real platforms such as node. For example, the `require`
      function can return any JavaScript value including a string but the
      `import * as` syntax always results in an object and cannot be a string.
  - p: >
      To avoid problems due to this legacy feature, you should enable the
      [`esModuleInterop`](https://www.typescriptlang.org/tsconfig#esModuleInterop)
      TypeScript configuration option if you use TypeScript with esbuild.
      Enabling it disables this legacy behavior and makes TypeScript's type
      system compatible with ESM. This option is not enabled by default because
      it would be a breaking change for existing TypeScript projects, but Microsoft
      [highly recommends applying it both to new and existing projects](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-2-7.html#support-for-import-d-from-cjs-from-commonjs-modules-with---esmoduleinterop)
      (and then updating your code) for better compatibility with the rest of the ecosystem.
  - p: >
      Specifically this means that importing a non-object value from a CommonJS
      module with ESM import syntax must be done using a default import instead
      of using `import * as`. So if a CommonJS module exports a function via
      <code>module.<wbr>exports =<wbr> fn</code>, you need to use
      <code>import <wbr>fn <wbr>from <wbr>'path'</code> instead of
      <code>import *<wbr> as<wbr> <wbr>fn <wbr>from <wbr>'path'</code>.

  - h4#no-type-system: Features that need a type system are not supported
  - p: >
      TypeScript types are treated as comments and are ignored by esbuild, so
      TypeScript is treated as "type-checked JavaScript." The interpretation of
      the type annotations is up to the TypeScript type checker, which you
      should be running in addition to esbuild if you're using TypeScript. This
      is the same compilation strategy that Babel's TypeScript implementation
      uses. However, it means that some TypeScript compilation features which
      require type interpretation to work do not work with esbuild.
  - p: >
      Specifically:
  - ul:
    - >
      <p>
      The [`emitDecoratorMetadata`](https://www.typescriptlang.org/tsconfig#emitDecoratorMetadata)
      TypeScript configuration option is not supported. This feature passes a
      JavaScript representation of the corresponding TypeScript type to the
      attached decorator function. Since esbuild does not replicate TypeScript's
      type system, it does not have enough information to implement this feature.
      </p>
    - >
      <p>
      The [`declaration`](https://www.typescriptlang.org/tsconfig#declaration)
      TypeScript configuration option (i.e. generation of `.d.ts` files) is
      not supported. If you are writing a library in TypeScript and you want to
      publish the compiled JavaScript code as a package for others to use, you
      will probably also want to publish type declarations. This is not something
      that esbuild can do for you because it doesn't retain any type information.
      You will likely either need to use the TypeScript compiler to generate them
      or manually write them yourself.
      </p>

  - h4#tsconfig-json: Only certain `tsconfig.json` fields are respected
  - p: >
      During bundling, the path resolution algorithm in esbuild will consider
      the contents of the `tsconfig.json` file in the closest parent directory
      containing one and will modify its behavior accordingly. It is also
      possible to explicitly set the `tsconfig.json` path with the build API
      using esbuild's [`tsconfig`](/api/#tsconfig) setting and to explicitly pass in
      the contents of a `tsconfig.json` file with the transform API using esbuild's
      [`tsconfigRaw`](/api/#tsconfig-raw) setting. However, esbuild currently
      only inspects the following fields in `tsconfig.json` files:

  - ul:
    - >
      [`experimentalDecorators`](https://www.typescriptlang.org/tsconfig#experimentalDecorators)
      <p>
      This option enables the transformation of decorator syntax in TypeScript
      files. The transformation follows the outdated decorator design that
      TypeScript itself follows when `experimentalDecorators` is enabled.
      </p>
      <p>
      Note that there is an updated design for decorators that is being added to
      JavaScript, as well as to TypeScript when `experimentalDecorators` is
      disabled. This is not something that esbuild implements yet, so esbuild
      will currently not transform decorators when `experimentalDecorators` is
      disabled.
      </p>

    - >
      [`target`](https://www.typescriptlang.org/tsconfig#target)
      <br>
      [`useDefineForClassFields`](https://www.typescriptlang.org/tsconfig#useDefineForClassFields)
      <p>
      These options control whether class fields in TypeScript files are compiled
      with "define" semantics or "assign" semantics:
      </p>
      <ul>
        <li><p>
        <b>Define semantics</b> (esbuild's default behavior): TypeScript class
        fields behave like normal JavaScript class fields. Field initializers
        do not trigger setters on the base class. You should write all new code
        this way going forward.
        </p></li>
        <li><p>
        <b>Assign semantics</b> (which you have to explicitly enable): esbuild
        emulates TypeScript's legacy class field behavior. Field initializers
        will trigger base class setters. This may be needed to get legacy code
        to run.
        </p></li>
      </ul>
      <p>
      The way to disable define semantics (and therefore enable assign semantics)
      with esbuild is the same way you disable it with TypeScript: by setting
      `useDefineForClassFields` to `false` in your `tsconfig.json` file.
      </p>
      <p>
      For compatibility with TypeScript, esbuild also copies TypeScript's behavior
      where when `useDefineForClassFields` is not specified, it defaults to
      `false` when `tsconfig.json` contains a `target` that is earlier than
      `ES2022`. But I recommend setting `useDefineForClassFields` explicitly
      if you need it instead of relying on this default value coming from the
      value of the `target` setting. Note that the `target` setting in
      `tsconfig.json` is only used by esbuild for determining the default value
      of `useDefineForClassFields`. It does *not* affect esbuild's own
      [`target`](/api/#target) setting, even though they have the same name.
      </p>

    - >
      [`baseUrl`](https://www.typescriptlang.org/tsconfig#baseUrl)
      <br>
      [`paths`](https://www.typescriptlang.org/tsconfig#paths)
      <p>
      These options affect esbuild's resolution of `import`/`require` paths to
      files on the file system. You can use it to define package aliases and to
      rewrite import paths in other ways. Note that using esbuild for import
      path transformation requires [`bundling`](/api/#bundle) to be enabled, as
      esbuild's path resolution only happens during bundling. Also note that
      esbuild also has a native [`alias`](/api/#alias) feature which you may
      want to use instead.
      </p>

    - >
      [`jsx`](https://www.typescriptlang.org/tsconfig#jsx)
      <br>
      [`jsxFactory`](https://www.typescriptlang.org/tsconfig#jsxFactory)
      <br>
      [`jsxFragmentFactory`](https://www.typescriptlang.org/tsconfig#jsxFragmentFactory)
      <br>
      [`jsxImportSource`](https://www.typescriptlang.org/tsconfig#jsxImportSource)
      <p>
      These options affect esbuild's transformation of JSX syntax into JavaScript.
      They are equivalent to esbuild's native options for these settings:
      [`jsx`](/api/#jsx),
      [`jsxFactory`](/api/#jsx-factory),
      [`jsxFragment`](/api/#jsx-fragment), and
      [`jsxImportSource`](/api/#jsx-import-source).
      </p>

    - >
      [`alwaysStrict`](https://www.typescriptlang.org/tsconfig#alwaysStrict)
      <br>
      [`strict`](https://www.typescriptlang.org/tsconfig#strict)
      <p>
      If either of these options are enabled, esbuild will consider all code in
      all TypeScript files to be in [strict mode](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Strict_mode)
      and will prefix generated code with `"use strict"` unless the output
      [`format`](/api/#format) is set to [`esm`](/api/#format-esm) (since all
      ESM files are automatically in strict mode).
      </p>

    - >
      [`verbatimModuleSyntax`](https://www.typescriptlang.org/tsconfig/#verbatimModuleSyntax)
      <br>
      [`importsNotUsedAsValues`](https://www.typescriptlang.org/tsconfig#importsNotUsedAsValues)
      <br>
      [`preserveValueImports`](https://www.typescriptlang.org/tsconfig/#preserveValueImports)
      <p>
      By default, the TypeScript compiler will delete unused imports when
      converting TypeScript to JavaScript. That way imports which turn out to
      be type-only imports accidentally don't cause an error at run-time. This
      behavior is also implemented by esbuild.
      </p>
      <p>
      These options allow you to disable this behavior and preserve unused
      imports, which can be useful if for example the imported file has useful
      side-effects. You should use `verbatimModuleSyntax` for this, as that
      replaces the older `importsNotUsedAsValues` and `preserveValueImports`
      settings (which TypeScript has now deprecated).
      </p>

    - >
      [`extends`](https://www.typescriptlang.org/tsconfig#extends)
      <p>
      This option allows you to split up your `tsconfig.json` file across
      multiple files. This value can be a string for single inheritance or an
      array for multiple inheritance (new in TypeScript 5.0+).
      </p>

  - p: >
      All other `tsconfig.json` fields (i.e. those that aren't in the above
      list) will be ignored.

  - h4#ts-vs-tsx: You cannot use the `tsx` loader for `*.ts` files
  - p: >
      The `tsx` loader is _not_ a superset of the `ts` loader. They are two
      different partially-incompatible syntaxes. For example, the character
      sequence `<a>1</a>/g` parses as `<a>(1 < (/a>/g))` with the `ts` loader
      and `(<a>1</a>) / g` with the `tsx` loader.
      </p>
      <p>The most common issue this causes is not being able to use generic
      type parameters on arrow function expressions such as `<T>() => {}`
      with the `tsx` loader. This is intentional, and matches the behavior of
      the official TypeScript compiler. That space in the `tsx` grammar is
      reserved for JSX elements.

## JSX
  - p: 'Loader: `jsx` or `tsx`'

  - p: >
      [JSX](https://facebook.github.io/jsx/) is an XML-like syntax extension
      for JavaScript that was created for [React](https://github.com/facebook/react).
      It's intended to be converted into normal JavaScript by your build tool.
      Each XML element becomes a normal JavaScript function call. For example,
      the following JSX code:

  - pre.jsx: |
      import Button from './button'
      let button = <Button>Click me</Button>
      render(button)

  - p: >
      Will be converted to the following JavaScript code:

  - pre.js: |
      import Button from "./button";
      let button = React.createElement(Button, null, "Click me");
      render(button);

  - p: >
      This loader is enabled by default for `.jsx` and `.tsx` files.
      Note that JSX syntax is not enabled in `.js` files by default. If you
      would like to enable that, you will need to configure it:

  - example:
      in:
        app.js: '<div/>'

      cli: |
        esbuild app.js --bundle --loader:.js=jsx

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.js': 'jsx' },
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
              ".js": api.LoaderJSX,
            },
            Write: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - h3: Auto-import for JSX

  - p: >
      Using JSX syntax usually requires you to manually import the JSX library
      you are using. For example, if you are using React, by default you will
      need to import React into each JSX file like this:

  - pre.jsx: |
      import * as React from 'react'
      render(<div/>)

  - p: >
      This is because the JSX transform turns JSX syntax into a call to
      <code>React.<wbr>createElement</code> but it does not itself import
      anything, so the `React` variable is not automatically present.

  - p: >
      If you would like to avoid having to manually `import` your JSX library
      into each file, you may be able to do this by setting esbuild's
      [JSX](/api/#jsx) transform to `automatic`, which generates import
      statements for you. Keep in mind that this also completely changes
      how the JSX transform works, so it may break your code if you are
      using a JSX library that's not React. Doing that looks like this:

  - example:
      in:
        app.jsx: '<a/>'

      cli: |
        esbuild app.jsx --jsx=automatic

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.jsx'],
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
            JSX:         api.JSXAutomatic,
            Outfile:     "out.js",
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - h3: Using JSX without React

  - p: >
      If you're using JSX with a library other than React (such as
      [Preact](https://preactjs.com/)), you'll likely need to configure the
      [JSX factory](/api/#jsx-factory) and [JSX fragment](/api/#jsx-fragment)
      settings since they default to <code>React<wbr>.createElement</code> and
      <code>React<wbr>.Fragment</code> respectively:

  - example:
      in:
        app.jsx: '<div/>'

      cli: |
        esbuild app.jsx --jsx-factory=h --jsx-fragment=Fragment

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.jsx'],
          jsxFactory: 'h',
          jsxFragment: 'Fragment',
          outfile: 'out.js',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.jsx"},
            JSXFactory:  "h",
            JSXFragment: "Fragment",
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      Alternatively, if you are using TypeScript, you can just configure JSX
      for TypeScript by adding this to your `tsconfig.json` file and esbuild
      should pick it up automatically without needing to be configured:

  - pre.json: |
      {
        "compilerOptions": {
          "jsxFactory": "h",
          "jsxFragmentFactory": "Fragment"
        }
      }

  - p: >
      You will also have to add <code>import <wbr>{h, <wbr>Fragment} <wbr>from <wbr>'preact'</code>
      in files containing JSX syntax unless you use auto-importing as described above.

  - h2: JSON
  - div: 'Loader: `json`'

  - p: >
      This loader is enabled by default for `.json` files. It parses the JSON
      file into a JavaScript object at build time and exports the object as
      the default export. Using it looks something like this:

  - pre.js: |
      import object from './example.json'
      console.log(object)

  - p: >
      In addition to the default export, there are also named exports for each
      top-level property in the JSON object. Importing a named export directly
      means esbuild can automatically remove unused parts of the JSON file from
      the bundle, leaving only the named exports that you actually used. For
      example, this code will only include the `version` field when bundled:

  - pre.js: |
      import { version } from './package.json'
      console.log(version)

## CSS
  - p: >
      Loader: `css` (also `global-css` and `local-css` for [CSS modules](#local-css))

  - p: >
      The `css` loader is enabled by default for `.css` files and the
      [`local-css`](#local-css) loader is enabled by default for `.module.css`
      files. These loaders load the file as CSS syntax. CSS is a first-class
      content type in esbuild, which means esbuild can [bundle](/api/#bundle)
      CSS files directly without needing to import your CSS from JavaScript
      code:

  - example:
      in:
        app.css: 'body {}'

      cli: |
        esbuild --bundle app.css --outfile=out.css

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.css'],
          bundle: true,
          outfile: 'out.css',
        })

      go: |
        package main

        import "github.com/evanw/esbuild/pkg/api"
        import "os"

        func main() {
          result := api.Build(api.BuildOptions{
            EntryPoints: []string{"app.css"},
            Bundle:      true,
            Outfile:     "out.css",
            Write:       true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      You can `@import` other CSS files and reference image and font files with
      `url()` and esbuild will bundle everything together. Note that you will
      have to configure a loader for image and font files, since esbuild
      doesn't have any pre-configured. Usually this is either the
      [data URL](#data-url) loader or the [external file](#external-file) loader.

  - p: >
      These syntax features are conditionally transformed for older browsers depending on the configured language [target](/api/#target):

  - table: |
      | Syntax transform                                                                     | Example                            |
      |--------------------------------------------------------------------------------------|------------------------------------|
      | [Nested declarations](https://www.w3.org/TR/css-nesting-1/)                          | `a { &:hover { color: red } }`     |
      | [Modern RGB/HSL syntax](https://www.w3.org/TR/css-color-4/#hex-notation)             | `#F008`                            |
      | [`inset` shorthand](https://developer.mozilla.org/en-US/docs/Web/CSS/Inset)          | `inset: 0`                         |
      | [`hwb()`](https://www.w3.org/TR/css-color-4/#the-hwb-notation)                       | `hwb(120 30% 50%)`                 |
      | [`lab()` and `lch()`](https://www.w3.org/TR/css-color-4/#specifying-lab-lch)         | `lab(60 -5 58)`                    |
      | [`oklab()` and `oklch()`](https://www.w3.org/TR/css-color-4/#specifying-oklab-oklch) | `oklab(0.5 -0.1 0.1)`              |
      | [`color()`](https://www.w3.org/TR/css-color-4/#color-function)                       | `color(display-p3 1 0 0)`          |
      | [Color stops with two positions](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Images/Using_CSS_gradients#creating_color_bands_stripes) | `linear-gradient(red 2% 4%, blue)` |
      | [Gradient transition hints](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Images/Using_CSS_gradients#gradient_hints) | `linear-gradient(red, 20%, blue)` <sup>1</sup> |
      | [Gradient color spaces](https://developer.mozilla.org/en-US/blog/css-color-module-level-4/#comparing_gradients_in_different_color_spaces) | `linear-gradient(in hsl, red, blue)` <sup>1</sup> |
      | [Gradient hue mode](https://developer.mozilla.org/en-US/blog/css-color-module-level-4/#using_hue_interpolation_modes_in_gradients) | `linear-gradient(in hsl longer hue, red, blue)` <sup>1</sup> |

  - p: >
      <footer>
      <sup>1</sup> This is demonstrated visually by esbuild's [gradient transformation tests](/gradient-tests/).
      </footer>

  - p: >
      Note that by default, esbuild's output will take advantage of modern
      CSS features. For example, <code>color: <wbr>rgba(255, <wbr>0, <wbr>0, <wbr>0.4)</code>
      will become <code>color: <wbr>#f006</code> when minifying is enabled
      which makes use of syntax from [CSS Color Module Level 4](https://www.w3.org/TR/css-color-4/#changes-from-3).
      If this is undesired, you must specify esbuild's [target](/api/#target)
      setting to say in which browsers you need the output to work correctly.
      Then esbuild will avoid using CSS features that are too modern for those
      browsers.

  - p: >
      When you provide a list of browser versions using the [target](/api/#target)
      setting, esbuild will also automatically insert vendor prefixes so that your
      CSS will work in those browsers at those versions or newer. Currently
      esbuild will do this for the following CSS properties:

  - ul:
    - '[`appearance`](https://caniuse.com/css-appearance)'
    - '[`backdrop-filter`](https://caniuse.com/css-backdrop-filter)'
    - '[`background-clip: text`](https://caniuse.com/background-clip-text)'
    - '[`box-decoration-break`](https://caniuse.com/css-boxdecorationbreak)'
    - '[`clip-path`](https://caniuse.com/css-clip-path)'
    - '[`font-kerning`](https://caniuse.com/font-kerning)'
    - '[`hyphens`](https://caniuse.com/css-hyphens)'
    - '[`initial-letter`](https://caniuse.com/css-initial-letter)'
    - '[`mask-composite`](https://caniuse.com/mdn-css_properties_mask-composite)'
    - '[`mask-image`](https://caniuse.com/mdn-css_properties_mask-image)'
    - '[`mask-origin`](https://caniuse.com/mdn-css_properties_mask-origin)'
    - '[`mask-position`](https://caniuse.com/mdn-css_properties_mask-position)'
    - '[`mask-repeat`](https://caniuse.com/mdn-css_properties_mask-repeat)'
    - '[`mask-size`](https://caniuse.com/mdn-css_properties_mask-size)'
    - '[`position: sticky`](https://caniuse.com/css-sticky)'
    - '[`print-color-adjust`](https://caniuse.com/css-color-adjust)'
    - '[`tab-size`](https://caniuse.com/css3-tabsize)'
    - '[`text-decoration-color`](https://caniuse.com/mdn-css_properties_text-decoration-color)'
    - '[`text-decoration-line`](https://caniuse.com/mdn-css_properties_text-decoration-line)'
    - '[`text-decoration-skip`](https://caniuse.com/mdn-css_properties_text-decoration-skip)'
    - '[`text-emphasis-color`](https://caniuse.com/mdn-css_properties_text-emphasis-color)'
    - '[`text-emphasis-position`](https://caniuse.com/mdn-css_properties_text-emphasis-position)'
    - '[`text-emphasis-style`](https://caniuse.com/mdn-css_properties_text-emphasis-style)'
    - '[`text-orientation`](https://caniuse.com/css-text-orientation)'
    - '[`text-size-adjust`](https://caniuse.com/text-size-adjust)'
    - '[`user-select`](https://caniuse.com/mdn-css_properties_user-select)'

  - h3#css-from-js: Import from JavaScript

  - p: >
      You can also import CSS from JavaScript. When you do this, esbuild will
      gather all CSS files referenced from a given entry point and bundle it
      into a sibling CSS output file next to the JavaScript output file for
      that JavaScript entry point. So if esbuild generates `app.js` it would
      also generate `app.css` containing all CSS files referenced by `app.js`.
      Here's an example of importing a CSS file from JavaScript:

  - pre.jsx: |
      import './button.css'

      export let Button = ({ text }) =>
        <div className="button">{text}</div>

  - p: >
      The bundled JavaScript generated by esbuild will not automatically import
      the generated CSS into your HTML page for you. Instead, you should import
      the generated CSS into your HTML page yourself along with the generated
      JavaScript. This means the browser can download the CSS and JavaScript
      files in parallel, which is the most efficient way to do it. That looks
      like this:

  - pre.html: |
      <html>
        <head>
          <link href="app.css" rel="stylesheet">
          <script src="app.js"></script>
        </head>
      </html>

  - p: >
      If the generated output names are not straightforward (for example if you
      have added `[hash]` to the [entry names](/api/#entry-names) setting and
      the output file names have content hashes) then you will likely want to
      look up the generated output names in the [metafile](/api/#metafile). To
      do this, first find the JS file by looking for the output with the
      matching `entryPoint` property. This file goes in the `<script>` tag. The
      associated CSS file can then be found using the `cssBundle` property.
      This file goes in the `<link>` tag.

  - h3#local-css: CSS modules

  - p: >
      [CSS modules](https://github.com/css-modules/css-modules) is a CSS
      preprocessor technique to avoid unintentional CSS name collisions. CSS
      class names are normally global, but CSS modules provides a way to make
      CSS class names local to the file they appear in instead. If two separate
      CSS files use the same local class name `.button`, esbuild will
      automatically rename one of them so that they don't collide. This is
      analogous to how esbuild automatically renames local variables with
      the same name in separate JS modules to avoid name collisions.

  - p: >
      There is support for bundling with CSS modules in esbuild. To use it,
      you need to enable [bundling](/api/#bundle), use the `local-css` loader
      for your CSS file (e.g. by using the `.module.css` file extension), and
      then import your CSS module code into a JS file. Each local CSS name in
      that file can be imported into JS to get the name that esbuild renamed
      it to. Here's an example:

  - pre.js: |
      // app.js
      import { outerShell } from './app.module.css'
      const div = document.createElement('div')
      div.className = outerShell
      document.body.appendChild(div)

  - pre.css: |
      /* app.module.css */
      .outerShell {
        position: absolute;
        inset: 0;
      }

  - p: >
      When you bundle this with <code>esbuild app.js <wbr>--bundle <wbr>--outdir=out</code>
      you'll get this (notice how the local CSS name `outerShell` has been renamed):

  - pre.js: |
      // out/app.js
      (() => {
        // app.module.css
        var outerShell = "app_outerShell";

        // app.js
        var div = document.createElement("div");
        div.className = outerShell;
        document.body.appendChild(div);
      })();

  - pre.css: |
      /* out/app.css */
      .app_outerShell {
        position: absolute;
        inset: 0;
      }

  - p: >
      This feature only makes sense to use when bundling is enabled both
      because your code needs to `import` the renamed local names so that
      it can use them, and because esbuild needs to be able to process all
      CSS files containing local names in a single bundling operation so
      that it can successfully rename conflicting local names to avoid collisions.

  - p: >
      The names that esbuild generates for local CSS names are an implementation
      detail and are not intended to be hard-coded anywhere. The only way you
      should be referencing the local CSS names in your JS or HTML is with an
      import statement in JS that is bundled with esbuild, as demonstrated above.
      For example, when [minification](/api/#minify) is enabled, esbuild will
      use a different name generation algorithm which generates names that are
      as short as possible (analogous to how esbuild minifies local identifiers
      in JS).

  - h4#global-css: Using global names

  - p: >
      The [`local-css`](#local-css) loader makes all CSS names in the file
      local by default. However, sometimes you want to mix local and global
      names in the same file. There are several ways to do this:

  - ul:
    - >
      You can wrap class names in `:global(...)` make them global and
      `:local(...)` to make them local.

    - >
      You can use `:global` to make names default to being global and
      `:local` to make names default to being local.

    - >
      You can use the `global-css` loader to still have local CSS features
      enabled but have names default to being global.

  - p: >
      Here are some examples:

  - pre.css: |
      /*
       * This is a local name with the "local-css" loader
       * and a global name with the "global-css" loader
       */
      .button {
      }

      /* This is a local name with both loaders */
      :local(.button) {
      }

      /* This is a global name with both loaders */
      :global(.button) {
      }

      /* "foo" is global and "bar" is local */
      :global .foo :local .bar {
      }

      /* "foo" is global and "bar" is local */
      :global {
        .foo {
          :local {
            .bar {}
          }
        }
      }

  - h4#composes: The `composes` directive

  - p: >
      The [CSS modules specification](https://github.com/css-modules/css-modules#composition)
      also describes a `composes` directive. It allows class selectors with
      local names to reference other class selectors. This can be used to split
      out common sets of properties to avoid duplicating them. And with the `from`
      keyword, it can also be used to reference class selectors with local names
      in other files. Here's an example:

  - pre.js: |
      // app.js
      import { submit } from './style.css'
      const div = document.createElement('div')
      div.className = submit
      document.body.appendChild(div)

  - pre.css: |
      /* style.css */
      .button {
        composes: pulse from "anim.css";
        display: inline-block;
      }
      .submit {
        composes: button;
        font-weight: bold;
      }

  - pre.css: |
      /* anim.css */
      @keyframes pulse {
        from, to { opacity: 1 }
        50% { opacity: 0.5 }
      }
      .pulse {
        animation: 2s ease-in-out infinite pulse;
      }

  - p: >
      Bundling this with <code>esbuild <wbr>app.js <wbr>--bundle <wbr>--outdir=<wbr>dist <wbr>--loader:<wbr>.css=<wbr>local-css</code>
      will give you something like this:

  - pre.js: |
      (() => {
        // style.css
        var submit = "anim_pulse style_button style_submit";

        // app.js
        var div = document.createElement("div");
        div.className = submit;
        document.body.appendChild(div);
      })();

  - pre.css: |
      /* anim.css */
      @keyframes anim_pulse {
        from, to {
          opacity: 1;
        }
        50% {
          opacity: 0.5;
        }
      }
      .anim_pulse {
        animation: 2s ease-in-out infinite anim_pulse;
      }

      /* style.css */
      .style_button {
        display: inline-block;
      }
      .style_submit {
        font-weight: bold;
      }

  - p: >
      Notice how using `composes` causes the string imported into JavaScript
      to become a space-separated list of all of the local names that were
      composed together. This is intended to be passed to the
      [`className`](https://developer.mozilla.org/en-US/docs/Web/API/Element/className)
      property on a DOM element. Also notice how using `composes` with `from`
      allows you to (indirectly) reference local names in other CSS files.

  - p: >
      Note that the order in which composed CSS classes from separate files
      appear in the bundled output file is deliberately _undefined_ by design
      (see [the specification](https://github.com/css-modules/css-modules#composing-from-other-files)
      for details). You are not supposed to declare the same CSS property in
      two separate class selectors and then compose them together. You are
      only supposed to compose CSS class selectors that declare non-overlapping
      CSS properties.

  - h3#css-caveats: CSS caveats

  - p: >
      You should keep the following things in mind when using CSS with esbuild:

  - h4#css-linting: Limited CSS verification
  - p: >
      CSS has a [general syntax specification](https://www.w3.org/TR/css-syntax-3/)
      that all CSS processors use and then [many specifications](https://www.w3.org/Style/CSS/current-work)
      that define what specific CSS rules mean. While esbuild understands
      general CSS syntax and can understand some CSS rules (enough to bundle
      CSS file together and to minify CSS reasonably well), esbuild does not
      contain complete knowledge of CSS. This means esbuild takes a "garbage
      in, garbage out" philosophy toward CSS. If you want to verify that your
      compiled CSS is free of typos, you should be using a CSS linter in
      addition to esbuild.

  - h4#css-import-order: >
      `@import` order matches the browser
  - p: >
      The `@import` rule in CSS behaves differently than the `import` keyword
      in JavaScript. In JavaScript, an `import` means roughly "make sure the
      imported file is evaluated before this file is evaluated" but in CSS,
      `@import` means roughly "re-evaluate the imported file again here"
      instead. For example, consider the following files:

  - ul:
    - '<code>entry.css</code><pre>@import "foreground.css";<br>@import "background.css";</pre>'
    - '<code>foreground.css</code><pre>@import "reset.css";<br>body {<br>  color: white;<br>}</pre>'
    - '<code>background.css</code><pre>@import "reset.css";<br>body {<br>  background: black;<br>}</pre>'
    - '<code>reset.css</code><pre>body {<br>  color: black;<br>  background: white;<br>}</pre>'

  - p: >
      Using your intuition from JavaScript, you might think that this code first
      resets the body to black text on a white background, and then overrides
      that to white text on a black background. _**This is not what happens.**_
      Instead, the body will be entirely black (both the foreground and the
      background). This is because `@import` is supposed to behave as if the
      import rule was replaced by the imported file (sort of like `#include`
      in C/C++), which leads to the browser seeing the following code:

  - pre.css: |
      /* reset.css */
      body {
        color: black;
        background: white;
      }

      /* foreground.css */
      body {
        color: white;
      }

      /* reset.css */
      body {
        color: black;
        background: white;
      }

      /* background.css */
      body {
        background: black;
      }

  - p: >
      which ultimately reduces down to this:

  - pre: |
      body {
        color: black;
        background: black;
      }

  - p: >
      This behavior is unfortunate, but esbuild behaves this way because that's
      how CSS is specified, and that's how CSS works in browsers. This is
      important to know about because some other commonly-used CSS processing
      tools such as [`postcss-import`](https://github.com/postcss/postcss-import/issues/462)
      incorrectly resolve CSS imports in JavaScript order instead of in CSS order.
      If you are porting CSS code written for those tools to esbuild (or even
      just switching over to running your CSS code natively in the browser),
      you may have appearance changes if your code depends on the incorrect
      import order.

  - h2: Text
  - p: 'Loader: `text`'

  - p: >
      This loader is enabled by default for `.txt` files. It loads the file
      as a string at build time and exports the string as the default export.
      Using it looks something like this:

  - pre.js: |
      import string from './example.txt'
      console.log(string)

  - h2: Binary
  - p: 'Loader: `binary`'

  - p: >
      This loader will load the file as a binary buffer at build time and
      embed it into the bundle using Base64 encoding. The original bytes of
      the file are decoded from Base64 at run time and exported as a
      `Uint8Array` using the default export. Using it looks like this:

  - pre.js: |
      import uint8array from './example.data'
      console.log(uint8array)

  - p: >
      If you need an `ArrayBuffer` instead, you can just access
      <code>uint8array<wbr>.buffer</code>. Note that this loader is not
      enabled by default. You will need to configure it for the appropriate
      file extension like this:

  - example:
      in:
        app.js: |
          import uint8array from './example.data'
          console.log(uint8array)
        example.data: |
          this is some data

      cli: |
        esbuild app.js --bundle --loader:.data=binary

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.data': 'binary' },
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
              ".data": api.LoaderBinary,
            },
            Write: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - h2: Base64
  - p: 'Loader: `base64`'

  - p: >
      This loader will load the file as a binary buffer at build time and
      embed it into the bundle as a string using Base64 encoding. This
      string is exported using the default export. Using it looks like this:

  - pre.js: |
      import base64string from './example.data'
      console.log(base64string)

  - p: >
      Note that this loader is not enabled by default. You will need to
      configure it for the appropriate file extension like this:

  - example:
      in:
        app.js: |
          import base64string from './example.data'
          console.log(base64string)
        example.data: |
          this is some data

      cli: |
        esbuild app.js --bundle --loader:.data=base64

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.data': 'base64' },
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
              ".data": api.LoaderBase64,
            },
            Write: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      If you intend to turn this into a `Uint8Array` or an `ArrayBuffer`, you
      should use the `binary` loader instead. It uses an optimized
      Base64-to-binary converter that is faster than the usual `atob`
      conversion process.

  - h2: Data URL
  - p: 'Loader: `dataurl`'

  - p: >
      This loader will load the file as a binary buffer at build time and
      embed it into the bundle as a Base64-encoded data URL. This string is
      exported using the default export. Using it looks like this:

  - pre.js: |
      import url from './example.png'
      let image = new Image
      image.src = url
      document.body.appendChild(image)

  - p: >
      The data URL includes a best guess at the MIME type based on the file
      extension and/or the file contents, and will look something like this
      for binary data:

  - pre: |
      data:image/png;base64,iVBORw0KGgo=

  - p: >
      ...or like this for textual data:

  - pre: |
      data:image/svg+xml,<svg></svg>%0A

  - p: >
      Note that this loader is not enabled by default. You will need to
      configure it for the appropriate file extension like this:

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
        esbuild app.js --bundle --loader:.png=dataurl

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.png': 'dataurl' },
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
            },
            Write: true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - h2: External file
  - p: >
      There are two different loaders that can be used for external files
      depending on the behavior you're looking for. Both loaders are
      described below:

  - h4#file: The `file` loader
  - p: 'Loader: `file`'

  - p: >
      This loader will copy the file to the output directory and embed the
      file name into the bundle as a string. This string is exported using
      the default export. Using it looks like this:

  - pre.js: |
      import url from './example.png'
      let image = new Image
      image.src = url
      document.body.appendChild(image)

  - p: >
      This behavior is intentionally similar to Webpack's
      [`file-loader`](https://v4.webpack.js.org/loaders/file-loader/)
      package. Note that this loader is not enabled by default. You will need
      to configure it for the appropriate file extension like this:

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
        esbuild app.js --bundle --loader:.png=file --outdir=out

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.png': 'file' },
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
            Outdir: "out",
            Write:  true,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      By default the exported string is just the file name. If you would like to
      prepend a base path to the exported string, this can be done with the
      [public path](/api/#public-path) API option.

  - h4#copy: The `copy` loader
  - p: 'Loader: `copy`'

  - p: >
      This loader will copy the file to the output directory and rewrite the
      import path to point to the copied file. This means the import will still
      exist in the final bundle and the final bundle will still reference the
      file instead of including the file inside the bundle. This might be useful
      if you are running additional bundling tools on esbuild's output, if you
      want to omit a rarely-used data file from the bundle for faster startup
      performance, or if you want to rely on specific behavior of your runtime
      that's triggered by an import. For example:

  - pre.js: |
      import json from './example.json' assert { type: 'json' }
      console.log(json)

  - p: >
      If you bundle the above code with the following command:

  - example:
      in:
        app.js: |
          import json from './example.json' assert { type: 'json' }
          console.log(json)
        example.json: |
          'this is some data'

      cli: |
        esbuild app.js --bundle --loader:.json=copy --outdir=out --format=esm

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.json': 'copy' },
          outdir: 'out',
          format: 'esm',
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
              ".json": api.LoaderCopy,
            },
            Outdir: "out",
            Write:  true,
            Format: api.FormatESModule,
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      the resulting `out/app.js` file might look something like this:

  - pre.js: |
      // app.js
      import json from "./example-PVCBWCM4.json" assert { type: "json" };
      console.log(json);

  - p: >
      Notice how the import path has been rewritten to point to the copied file
      `out/example-PVCBWCM4.json` (a content hash has been added due to the
      default value of the [asset names](/api/#asset-names) setting), and how
      the [import assertion](https://v8.dev/features/import-assertions) for JSON
      has been kept so the runtime will be able to load the JSON file.

  - h2: Empty file
  - p: 'Loader: `empty`'

  - p: >
      This loader tells esbuild to pretend that a file is empty. It can be
      a helpful way to remove content from your bundle in certain situations.
      For example, you can configure `.css` files to load with `empty` to prevent
      esbuild from bundling CSS files that are imported into JavaScript files:

  - example:
      in:
        app.js: |
          import './styles.css'
        styles.css: |
          .should-not-be-bundled {
            color: red;
          }

      cli: |
        esbuild app.js --bundle --loader:.css=empty

      js: |
        require('esbuild').buildSync({
          entryPoints: ['app.js'],
          bundle: true,
          loader: { '.css': 'empty' },
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
              ".css": api.LoaderEmpty,
            },
          })

          if len(result.Errors) > 0 {
            os.Exit(1)
          }
        }

  - p: >
      This loader also lets you remove imported assets from CSS files. For example,
      you can configure `.png` files to load with `empty` so that references to
      `.png` files in CSS code such as `url(image.png)` are replaced with `url()`.