# gocrypt: lightweight encryption tool

[![Go](https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

For small to medium encryption jobs when you just need a quick and easy implementation of AES-256 with Galois Counter Mode.

## Motivation
Is this just a wrapper for Go's implementation of AES-256 GCM? Yes. But sometimes that's all you need. Why fuss with what algorithm you want to use
when you can just type a few arguments in the terminal and have an encrypted file that meets the gold standard for cryptography.

## Features

- Easy to use syntax
- Quickly encrypt and decrypt small to medium sized files (Successfully tested up to 2GB)
- Cool terminal effect that makes it look like encryption is happening

### Demo

Example output:
```bash
gocrypt main ❯ cat data/message.txt
This is a secret message.
Don't tell anyone!
The password is password.
Shhhhhh.....

gocrypt main ❯ ./gocrypt encrypt data/message.txt data/secretmessage
Enter secret:

Successfully encrypted file: data/secretmessage.gcrypt

gocrypt main ❯ cat data/secretmessage.gcrypt
'1��g LtA�W��
               ��.��xcp�ᐺODJ�*�a�$�K�n�-Dا)�T)�w�m^���&k�h�0�vܦ�����i�xP�3�RA���W��F;D��U�_t�A'wq����l��&D��
```

## Quick Start

### Prerequisites
go 1.25.7 or later

require (
        golang.org/x/crypto v0.49.0 // indirect
        golang.org/x/sys v0.42.0 // indirect
        golang.org/x/term v0.41.0 // indirect
)

### Build from source

```bash
# Clone the repo
git clone https://github.com/CodeZeroSugar/gocrypt.git
cd gocrypt

# Build the binary
go build .

# Or install globally (optional)
go install .
```

Now you can run ./gocrypt (or gocrypt if installed globally)

## Usage
A lightweight encryption tool that implements AES-256 GCM

```bash
gocrypt main ❯ ./gocrypt help

Usage: [command] [input filepath] [output filepath]
Commands:
        - encrypt : encrypt the file at the input filepath.
        - decrypt : decrypt the file at the input filepath.
        - help : display usage information.
Example: gocrypt encrypt file.txt file.txt
.gcrypt will be appended to the filename.
```

## Examples
**Encrypt a file:**
```bash
gocrypt encrypt message.txt secret
```
**Decrypt a file**
```bash
gocrypt decrypt secret.gcrypt message.txt
```

## Why This Project?
This was created to be a lightweight encrypt/decrypt tool and to learn about symmetric encryption algorithm implementation:
- Reading a plaintext or ciphertext file and applying AES-256 GCM, then writing to a new file.
- Securely storing the key with the ciphertext with Argon2.
- Using go routines to create terminal effects while the cryptographic function runs.

## Future Improvements (Roadmap)
- Implement chunking to enable processing of larger files.
- Allow user to declare multiple files in the arguments.
- Implement a wizard that walks the user through the options.
- Come up with a more unique name.

## Contributing
### Clone the repo

```bash
git clone https://github.com/CodeZeroSugar/gocrypt
cd gocrypt
```

### Build the compiled binary

```bash
go build .
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## License
This project is licensed under the MIT License — see the LICENSE file for details.

