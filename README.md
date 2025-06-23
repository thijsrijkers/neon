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


## Installation

> Work in progress


## Build a simple HTTP server

Neon lets you serve HTTP directly from JavaScript. You can create the server by making a server.js for example and put in the following code:
```js
neon.serve((req) => {
  return {
    status: 200,
    body: "Hello from Neon!",
  };
});
```

To start the server you can execute the file by running the following command:
```js
neon run server.js
```

Visit http://localhost:8080 in your browser.

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

## License

This project is licensed under the [MIT License](./LICENSE).

