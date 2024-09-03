
# `head` utility command

A simple command-line utility that reads and prints a specified number of lines from a text file.

## Description

This program reads a specified number of lines from a given text file and displays them on the console. It is a helpful tool when you want to quickly preview the contents of a file without opening it fully, especially when dealing with large files.

## Getting Started

### Dependencies

* Go programming language.

### Installing

1. **Clone the repository**:
    ```bash
    git clone https://github.com/codescalersinternships/Coreutils-MohamedFadel/tree/development
    ```
   
2. **Navigate to the project directory**:
    ```bash
    cd Coreutils-MohamedFadel/cmd/head
    ```
   
3. **Build the program**:
    ```bash
    go build
    ```

### Executing program

1. **Run the program** by specifying the number of lines (`-n`) and the file path:

    ```bash
    ./head -n=10 path/to/your/textfile.txt
    ```

2. **Example**:
    ```bash
    ./head -n=5 sample.txt
    ```

    This command will print the first 5 lines of `sample.txt`. If the `-n` flag is not specified, it defaults to 10 lines.

### Error Handling

 * If no file path is specified, the program will return an error indicating that a valid path to a text file is required.
```bash
2024/09/03 17:48:43 invalid file path, you have to specify a valid path to text file
```
 * If the file cannot be opened, an error message will be displayed indicating that the file could not be found or opened.
```bash
2024/09/03 17:50:04 cannot open this file for reading: No such file or directory
```
 * If unvalid flag value is specified like `-n=xyz` or `-n=-7`, an error message will be displayed indicating that it is unvalid value.
 ```bash
 invalid value "xyz" for flag -n: parse error
Usage of ./head:
  -n uint
        number of lines (default 10)
 ```
