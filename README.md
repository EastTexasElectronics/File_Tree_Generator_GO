
# File Tree Generator (ftg)

File Tree Generator (ftg) is a command-line tool for generating file trees. This tool allows you to exclude specific directories or files and offers an interactive mode to select items to exclude. It is compatible with multiple operating systems and can be installed using Homebrew, Go install, or by building from source.

## Compatibility

- **macOS**
- **Linux**
- **Windows**

## Installation

### Using Homebrew (macOS)

You can install `ftg` using Homebrew. Follow the steps below:

1. Tap the repository:

    ```sh
    brew tap EastTexasElectronics/ftg-go-tap
    ```

2. Install `ftg`:

    ```sh
    brew install ftg
    ```

### Using Go Install

You can also install `ftg` using Go install:

1. Make sure you have Go installed. If not, you can download and install it from [golang.org](https://golang.org/dl/).

2. Install `ftg`:

    ```sh
    go install github.com/EastTexasElectronics/File_Tree_Generator_GO@latest
    ```

### Building from Source

To build `ftg` from source, follow these steps:

1. Clone the repository:

    ```sh
    git clone https://github.com/EastTexasElectronics/File_Tree_Generator_GO.git
    cd File_Tree_Generator_GO
    ```

2. Build the binary:

    ```sh
    go build -o ftg
    ```

3. Move the binary to a directory in your PATH:

    ```sh
    mv ftg /usr/local/bin/
    ```

## Usage

Here are some examples of how to use `ftg`:

### Basic Usage

Generate a file tree in the current directory:

```sh
ftg
```

### Exclude Specific Patterns

Exclude specific directories or files (comma-separated):

```sh
ftg -e .git,node_modules,.vscode
```

### Specify Output Location

Specify an output location for the generated file tree:

```sh
ftg -o output_location.md
```

### Interactive Mode

Use interactive mode to select items to exclude:

```sh
ftg -i
```

### Clear Exclusion List

Clear the exclusion list:

```sh
ftg -c
```

### Show Help

Show the help message:

```sh
ftg -h
```

### Show Version

Show the version information:

```sh
ftg -v
```

## License

This project is licensed under the AGPL-3.0 license. See the [LICENSE] file for more details.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue on GitHub.

## Support

If you encounter any issues or have any questions, please open an issue on the GitHub repository.

## Donations

If you find this project useful, consider buying me a coffee:
[Buy me a coffee](https://www.buymeacoffee.com/rmhavelaar)
