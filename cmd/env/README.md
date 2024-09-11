# `env` utility command

A simple command-line utility that prints all environment variables from the system.

## Description

This program displays all environment variables available in the current system. It can be useful for debugging, checking environment configurations, or understanding the variables that affect your program's execution.

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
   cd Coreutils-MohamedFadel/cmd/env
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Example**: Run the program to print all environment variables:

   ```bash
   ./env
   ```

This command will print all environment variables set in the current session.

### Error Handling

- The program does not include error handling for invalid inputs since it simply prints the environment variables.
