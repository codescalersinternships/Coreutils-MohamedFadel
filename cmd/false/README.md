# `false` utility command

A simple command-line utility that does nothing and always exits with a failure status.

## Description

This program mimics the behavior of the Unix `false` command. It performs no operations and always exits with a failure status.

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
   cd Coreutils-MohamedFadel/cmd/false
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Example**:

   ```bash
   ./false; echo $?
   ```

This command will print 1 indicating failure status.

### Error Handling

- There is no error handling needed since the program performs no operations and always exits with a failure status.
