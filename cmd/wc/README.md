# `wc` utility command

A simple command-line utility that reads and prints the number of lines, words, and characters from a text file.

## Description

This program reads a specified text file and counts the number of lines, words, and characters. It provides a quick way to analyze the content of a file without manually opening it, especially useful for large files.

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
   cd Coreutils-MohamedFadel/cmd/wc
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Run the program** by specifying the flags (`-l`, `-w`, `-c`) and the file path:

   ```bash
   ./wc -l -w -c path/to/your/textfile.txt
   ```

2. **Example**:
   ```bash
   ./wc -l -w sample.txt
   ```

This command will print the number of lines and words in sample.txt. If no flags are specified, it defaults to counting lines, words and characters.

```
Number of lines: 20
Number of words: 38
```

### Error Handling

- If no file path is specified, the program will return an error indicating that a valid path to a text file is required.

```bash
2024/09/03 17:48:43 invalid file path, you have to specify a valid path to text file
```

- If the file cannot be opened, an error message will be displayed indicating that the file could not be found or opened.

```bash
2024/09/03 17:50:04 cannot open this file for reading: No such file or directory
```

- If invalid flag value is specified like `-g`, an error message will be displayed indicating that it is invalid flag.

```bash
flag provided but not defined: -g
Usage of /tmp/go-build175744130/b001/exe/wc:
 -c    count characters flag
 -l    count lines flag
 -w    count words flag
exit status 2
```
