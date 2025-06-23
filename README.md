# Neon

**Neon** is a fast, minimal JavaScript runtime built in **Go**, powered by **V8**.

It offers a modern alternative to Node.js and Deno by combining the raw performance of Google's V8 engine with the simplicity and efficiency of Go. Neon is designed to run JavaScript outside the browser with full npm support, a built-in HTTP server, and an architecture optimized for developer tooling and extensibility.


## Why Neon?

Neon is for developers who want:

- A minimal runtime for scripting, tooling, and lightweight services.
- Seamless npm integration without the weight of a full Node.js runtime.
- A Go-native implementation of JavaScript execution using V8.
- Flexibility to experiment with runtime behavior, permission models, and APIs.
- An ideal foundation for building developer tools, web frameworks, and edge runtimes.


## Core Principles

- **Performance-first**: Powered by V8, designed to stay lean and fast.
- **Simplicity over complexity**: Written in Go with modular, readable code.
- **Modern JavaScript**: ES6+ syntax, npm support, server APIs.
- **Extendability**: Built for embedding, customization, and future features like TypeScript, WASM, and sandboxing.

  
## Project Structure

```graphql
neon/             
├── main.go
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


## MVP Features

- **Powered by V8** – Uses the same JavaScript engine as Chrome and Node.js.
- **npm Package Support** – Install and run npm modules in your scripts.
- **Modular Runtime** – Go-based architecture with clean interfaces.
- **Built-in HTTP Server** – Create web servers directly from JS code.
- **Future-Oriented** – Plans for permissions, ESM, TypeScript, and WASM.


## Tech Stack

- **Go** – Core language used to build the runtime.
- **V8** – High-performance JavaScript engine (via CGo bindings).
- **npm** – Integrated via shell or native registry access.
- **ESBuild** – (Optional) Used for bundling, transforming code.
- **libuv** – (Planned) Event loop and async I/O support.
- **net/http** – Go’s built-in HTTP server for exposing runtime APIs.
  

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

