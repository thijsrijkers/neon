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

## Installation

> Work in progress

### Verify V8 Integration & GCC Setup

#### Test V8 Integration

Run the unit tests to verify V8 works:

```bash
go test ./unit/test_v8go.go
``` 

#### GCC Setup

- **Windows:**  
  Install MSYS2 with mingw64 toolchain, then add `C:\msys64\mingw64\bin` to your system PATH.

- **macOS:**  
  Install Xcode Command Line Tools using `xcode-select --install`.

## Project Structure

```graphql
neon/             
├── main.go
|           
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

## Testing:

### 1. Unit Testing
To run the unit tests, navigate to the source folder and run:

```bash
go test ./unit...
```

This will execute the unit tests and display the results in your terminal.

## License

This project is licensed under the [MIT License](./LICENSE).

