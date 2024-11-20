# Pokedex CLI in Go

A Command Line Interface (CLI) application for exploring Pokémon data, built in Go. This application allows users to interact with the Pokémon API to view and explore Pokémon, their locations, and other details directly from the terminal.

## Features

- **Explore Pokémon Data**: View detailed information about Pokémon.
- **Inspect Pokémon Locations**: Explore where Pokémon can be found.
- **Catch and Cache Pokémon**: Simulate catching Pokémon and store them locally.
- **Command-Based Interaction**: Use various CLI commands to interact with the app, such as `help`, `map`, and `exit`.

## Project Structure

### Main Components

- **`main.go`**: The entry point for the application.
- **`commander.go`**: Handles CLI commands and their execution.
- **Commands**: Individual Go files implement specific commands (e.g., `command_map.go`, `command_help.go`, `command_exit.go`).
- **Internal API Clients**:
  - **`pokeapi`**: Interacts with the Pokémon API for fetching data.
  - **`pokecache`**: Implements caching mechanisms for storing and retrieving Pokémon data locally.

### Internal Packages

- **`internal/pokeapi`**:
  - **`client.go`**: Manages API interactions.
  - **`pokemon.go`**: Handles Pokémon-related API requests.
  - **`locations_list_area.go`**: Retrieves Pokémon locations.
- **`internal/pokecache`**:
  - **`cacheMethods.go`**: Implements caching logic.
  - **`cacheStruct.go`**: Defines cache structures.

### Go Modules

- **`go.mod`**: Defines module dependencies.

## Getting Started

### Prerequisites

- Go 1.20 or higher installed on your system.
- Access to the Pokémon API (`https://pokeapi.co`).

### Installation

1. Clone this repository:

   ```bash
   git clone <repository-url>
   cd pokedex-cli-go
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

### Running the Application

Run the CLI application with:

```bash
go run main.go
```

### Available Commands

- `help`: Display a list of available commands and their descriptions.
- `map`: Explore location areas in the Pokémon world.
- `inspect`: View detailed information about specific Pokémon.
- `catch`: Simulate catching Pokémon.
- `exit`: Exit the application.

### Testing

Run tests for the caching system:

```bash
go test ./internal/pokecache
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the project.

## License

This project is licensed under the MIT License. See the LICENSE file for details.