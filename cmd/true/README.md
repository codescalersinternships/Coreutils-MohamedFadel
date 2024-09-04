# `true` utility command

A simple command-line utility that does nothing and always exits with a success status.

## Description

This program mimics the behavior of the Unix `true` command. It performs no operations and always exits successfully.

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
   cd Coreutils-MohamedFadel/cmd/true
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Example**:

   ```bash
   ./true; echo $?
   ```

This command will print 0 indicating success.

### Error Handling

- There is no error handling needed since the program performs no operations and always exits successfully.
