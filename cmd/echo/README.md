# `echo` utility command

A simple command-line utility that prints the given strings to STDOUT, with an optional flag to omit the trailing newline.

## Description

This program prints the provided strings to the console. It includes an option to omit the trailing newline with the `-n` flag.

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
   cd Coreutils-MohamedFadel/cmd/echo
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Run the program** by specifying the `-n` flag (optional) and the strings you want to print:

   ```bash
   ./echo -n Hello World!
   ```

   If the `-n` flag is set, the program will print the strings without adding a trailing newline

2. **Example**:
   Without the `-n` flag:

   ```bash
   ./echo Hello World!
   ```

   Output:

   ```bash
   Hello World!
   ```

### Error Handling

- If no strings are provided, the program will simply print nothing.

- If invalid flag value is specified like `-r`, an error message will be displayed indicating that it is invalid flag.

  ```bash
  flag provided but not defined: -r
  Usage of /tmp/go-build229890464/b001/exe/echo:
  -n    omit trailing line after printing to STDOUT
  exit status 2
  ```
