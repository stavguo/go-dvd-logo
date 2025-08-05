# Bouncing DVD Logo

A recreation of the classic bouncing DVD logo screensaver built in Go using the Ebiten game engine. This project serves as a template for 2D games and simulations using Entity Component System architecture and Finite State Machine scene management.

## Libraries Used

- **[Ebiten v2](https://github.com/hajimehoshi/ebiten)** - 2D game engine for Go
- **[Ark ECS](https://github.com/mlange-42/ark)** - Entity Component System framework
- **Go 1.24.2** - Programming language

## Project Structure

```
├── main.go              # Entry point and game loop
├── scenes/              # Scene management (title, game)
├── systems/             # ECS systems (rendering, physics, bounds checking)
├── components/          # ECS components (position, velocity, texture, etc.)
├── core/                # Core ECS setup and resource management
└── assets/              # Fonts and images
```

## Getting Started

### Prerequisites

- Go 1.24.2 or later

## Commands

```bash
# Run the program
go run .

# Download and organize dependencies
go mod tidy

# Format Go source code
gofmt -w .

# From your main branch, push builds directory to gh-pages
git subtree push --prefix builds origin gh-pages
```

## Building for WASM

### Compiling

On a Unix/Linux shell:

```bash
env GOOS=js GOARCH=wasm go build -o builds/main.wasm github.com/stavguo/go-dvd-logo
```

### Copying wasm_exec.js to execute the WebAssembly binary

On a Unix/Linux shell:

```bash
# Go 1.24 and newer
cp $(go env GOROOT)/lib/wasm/wasm_exec.js builds/.
```

If wasm_exec.js cannot be found, download the wasm_exec.js for the same version of Go you used to compile the project from the official repository or the GitHub repository. For example, if you used Go 1.24.5, the path to wasm_exec.js would be:

`https://raw.githubusercontent.com/golang/go/refs/tags/go1.24.5/lib/wasm/wasm_exec.js`

### Creating an HTML file

Create an HTML file to load and run your WASM file in the browser:

```html
<!DOCTYPE html>
<script src="wasm_exec.js"></script>
<script>
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(result => {
    go.run(result.instance);
});
</script>
```

## Architecture

This project uses an Entity Component System (ECS) architecture with the Ark library:

- **Components**: Data containers (Position, Velocity, Texture, Colors, Screen)
- **Systems**: Logic processors (ApplyVelocity, CheckBounds, RenderLogo)
- **Entities**: Combinations of components (the bouncing logo)

This demo features a scene management system that handles transitions between the title screen and the main game loop.