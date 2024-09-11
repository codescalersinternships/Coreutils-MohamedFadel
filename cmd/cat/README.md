# `cat` utility command

A simple command-line utility that reads and displays the content of one or more text files, with an optional feature to display line numbers.

## Description

This program reads and displays the contents of specified text files on the console. It includes an option to print each line with its corresponding line number, making it easy to view and reference specific lines within the file. This tool is useful when you need to quickly preview the contents of files without opening them fully in an editor.

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
   cd Coreutils-MohamedFadel/cmd/cat
   ```
3. **Build the program**:
   ```bash
   go build
   ```

### Executing program

1. **Run the program** by specifying the file paths. Use the `-n` flag if you want to display line numbers:

   ```bash
   ./cat path/to/your/textfile1.txt path/to/your/textfile2.txt
   ```

2. To display line numbers, use the `-n` flag

   ```bash
   ./cat -n path/to/your/textfile.txt
   ```

3. **Example**:

   ```bash
   ./cat -n sample.txt
   ```

This command will print each line of `sample.txt` with its line number. If the `-n` flag is not used, the file content will be printed without line numbers.

### Error Handling

- If no file paths are specified, the program will return an error indicating that you must specify at least one valid file path:

  ```bash
  2024/09/03 17:48:43 invalid file(s) path(s), you have to specify at least one valid path to text file
  ```

- If the file cannot be opened, an error message will be displayed indicating that the file could not be found or opened.

  ```bash
  sample.txt: cannot open this file for reading: No such file or directory
  ```

- If invalid flag value is specified like `-g`, an error message will be displayed indicating that it is invalid flag.

  ```bash
  flag provided but not defined: -g
  Usage of /tmp/go-build130400896/b001/exe/cat:
  -n    print content with line numbers
  exit status 2
  ```
