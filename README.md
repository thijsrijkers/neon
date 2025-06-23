# Neon

**Neon** is a modern JavaScript runtime built with **Go** and powered by **V8**, designed for speed, simplicity, and developer productivity. It aims to provide a clean and extensible platform for running JavaScript outside the browser, with first-class support for **npm packages** and seamless integration of modern tooling.

## Features (MVP)

- Powered by V8 – High-performance JavaScript engine used by Chrome and Node.js.
- npm Support – Install and use npm packages directly in your Neon scripts.
- Modular Runtime – Built with Go for portability, speed, and simplicity.
- Built-in HTTP Server – Serve content or APIs using native JS syntax.
- Future-Proof – Sandboxing, permissions, and TypeScript support on the roadmap.

## Installation

> Work in progress

## Quick Start

Let's create your first Neon project.

### 1. Write your first script

Create a file called `main.js`:

```js
console.log("Welcome to Neon!");
``` 

Run it with:

```js
neon run main.js
```

### 2. Install an npm package

You can use packages from the npm registry:

```js
neon install lodash
```

This will download lodash and make it available in your project.

Update main.js:

```js
const _ = require("lodash");

console.log(_.shuffle(["neon", "is", "fast", "and", "modular"]));
```

Then run it again.

### 3. Build a simple HTTP server

Neon lets you serve HTTP directly from JavaScript.

Create server.js:

```js
neon.serve((req) => {
  return {
    status: 200,
    body: "Hello from Neon!",
  };
});
```

Start the server:

```js
neon run server.js
```

Visit http://localhost:8080 in your browser.

## Tech Stack

- **Go** – Core language used to build the runtime.
- **V8** – High-performance JavaScript engine (via CGo bindings).
- **npm** – Integrated via shell or native registry access.
- **ESBuild** – (Optional) Used for bundling, transforming code.
- **libuv** – (Planned) Event loop and async I/O support.
- **net/http** – Go’s built-in HTTP server for exposing runtime APIs.

## Project Structure

```graphql
neon/             
├──main.go
|           
├── v8/              # V8 bindings and engine bootstrap
├── runtime/         # Core runtime logic (script loading, eval)
├── npm/             # npm install logic, package resolution
├── server/          # HTTP server logic exposed to JavaScript
├── pkg/             # Shared libraries (if any)
├── scripts/         # Example or test JavaScript scripts
├── tests/           # Unit and integration tests
├── go.mod
├── README.md
└── Makefile         # Build and development commands
```

## Goals and Roadmap

- [ ] Execute JavaScript using V8
- [ ] CLI support (`run`, `install`, etc.)
- [ ] npm package install and resolution
- [ ] Simple HTTP server API
- [ ] ES module support
- [ ] Built-in permissions (like Deno)
- [ ] Built-in TypeScript support
- [ ] WASM support
- [ ] Plugin API
- [ ] Cross-platform static binaries
- [ ] Sandboxed execution with memory/CPU limits
