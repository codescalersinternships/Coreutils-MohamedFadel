# `tree` utility command

A simple command-line utility that displays the directory tree structure up to a specified depth.

## Description

This program prints the directory tree of a given path up to a specified depth level. It helps visualize the directory structure, making it easier to understand the hierarchy of files and folders without manually navigating each folder.

## Getting Started

### Dependencies

- Go programming language.

### Installing

1. **Clone the repository**:
   ```bash
   git clone https://github.com/codescalersinternships/Coreutils-MohamedFadel/tree/development
   ```
2. **Navigate to the project directory**:
   ```bash
   cd Coreutils-MohamedFadel/cmd/tree
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Run the program** by specifying the depth level `-L` and the directory path

   ```bash
   ./tree -L=2 path/to/your/directory
   ```

2. **Example**:

   ```bash
   ./tree -L=3 /home/user/documents
   ```

This command will print the directory structure of `/home/user/documents` up to 3 levels deep. If the `-L` flag is not specified, it defaults to 1 level.

### Error Handling

- If no path is specified, the program will return an error indicating that a valid path to a directory is required.

```bash
2024/09/04 22:13:22 invalid path, you have to specify a valid path to a directory
```

- If the file cannot be opened, an error message will be displayed indicating that the file could not be found or opened.

```bash
2024/09/04 22:15:46 /home/m/df: [error opening dir]
```

- If invalid flag value is specified like `-L=xyz` or `-L=-7`, an error message will be displayed indicating that it is invalid value.

```bash
invalid value "-3" for flag -L: parse error
Usage of /tmp/go-build1405114057/b001/exe/tree:
  -L uint
        Max display depth of the directory tree (default 1)
exit status 2
```
