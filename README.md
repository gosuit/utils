# Utils

This repository contains some useful utils. Now it is coder and generator packages. 

## Installation

```zsh
go get github.com/gosuit/utils
```

## Coder Package

The coder package provides functionality for encryption, decryption, and hashing. It utilizes AES for encryption and bcrypt for hashing.

### Features

- **Encryption**: Securely encrypt plaintext strings using AES encryption.
- **Decryption**: Decrypt previously encrypted strings back to plaintext.
- **Hashing**: Generate bcrypt hashes of plaintext strings.
- **Hash Comparison**: Verify if a plaintext string matches a given bcrypt hash.

### Usage

```golang
import "github.com/gosuit/utils/coder"

func main() {
    cfg := coder.Config{
        Secret:   "your-32-byte-secret-key",
        HashCost: 10,
    }

    c, err := coder.New(cfg)
    if err != nil {
        // Handle error
    }

    encrypted, err = c.Encrypt("your plaintext")
    if err != nil {
        // Handle error
    }

    decrypted, err = c.Decrypt(encrypted)
    if err != nil {
        // Handle error
    }

    // decrypted == "your plaintext"

    hash, err = c.Hash("your plaintext")
    if err != nil {
        // Handle error
    }

    err = c.CompareHash(hash, "your plaintext")
    // Check comapre error
}
```
   
## Generator Package

The generator package provides functions for generating random numbers and secret keys.

### Features

- **Random Number Generation**: Create a string of random numbers of a specified length.
- **Secret Key Generation**: Generate a secure secret key of a specified length.

### Usage


```golang
import "github.com/gosuit/utils/generator"

func main() {
    randomNum := generator.GetRandomNum(10)

    secretKey, err := generator.GetSecret(32)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.