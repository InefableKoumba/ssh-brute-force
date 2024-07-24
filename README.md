# SSH Brute Force Tool

This Go program attempts to connect to a given host using a list of passwords for brute-forcing SSH credentials. The program limits the number of concurrent connections to avoid overwhelming the server and getting rate-limited.

## Features

- Reads passwords from a file
- Attempts SSH connections using each password
- Limits concurrent connections to avoid rate limiting
- Logs successful username and password combinations to a file

## Prerequisites

- Go 1.15 or later
- SSH server to test against (with permission)
- Password list file (`p.txt`)

## Installation

1. Install [Go](https://golang.org/doc/install).
2. Download the required Go package:
   ```sh
   go get golang.org/x/crypto/ssh
   ```

## Usage

1. Create a file named `p.txt` containing the list of passwords to try, one per line.
2. Clone the repository or create a new Go file (`main.go`) and copy the provided code into it.
3. The program uses `root` as the default username, you can change it to match your need.
4. Run the program:
   ```sh
   go run main.go
   ```
5. Found passwords will be written in credentials_found.txt

## MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
