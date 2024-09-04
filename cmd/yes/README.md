# `yes` utility command

A simple command-line utility that repeatedly prints command-line arguments or the letter "y" if no arguments are provided.

## Description

This program continuously prints the specified command-line arguments in an infinite loop. If no arguments are given, it prints the letter "y" repeatedly. This tool mimics the behavior of the Unix `yes` command and can be used to automate responses in scripts or tests.

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
   cd Coreutils-MohamedFadel/cmd/yes
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Example**:

   ```bash
   ./yes Go is awesome
   ```

This command will print `Go is awesome` continuously until interrupted.

2. **Example**: Run the program without arguments to print "y" indefinitely:

   ```bash
   ./yes
   ```

### Error Handling

- There is no error handling for command-line input as the program does not validate the provided arguments.

- Interrupt the program manually using `Ctrl+C` to stop it.
