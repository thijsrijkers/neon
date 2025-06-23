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


