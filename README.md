## Multithreaded file search utility

Golang utility to search the file in system by passing the directory path and filename,
it will show all the files with same name present under the given directory.

It's a very fast utility utilizing the multithread feature of go-lang. 

## Usage

- Download the utility
  - change the permission to make it executable for linux-based machine
  ```
    chmod +x mfs
  ```
- Run the utility below is the example
  ```
  mfs --directory /User/ --filename README.md
  ```
  - here **/User/** is the example for the directory name and **README.md** as filename.
