
# Go-LZW

A minimalistic **LZW** (Lempel–Ziv–Welch) compression and decompression tool implemented in Go.  
Supports **dictionary growth** up to 15-bit codes (max code = 32767) and a “stop-growing dictionary” strategy once the dictionary is full.

## Features

- **Encoding (compression)**:
  - Reads from a file or `stdin`.
  - Produces an `.lzw` file or writes to `stdout` (depending on environment variables).
  - Uses a variable code size, starting at 9 bits, up to a maximum of 15 bits.
  - Stops adding new dictionary entries once `nextCode` hits 32767, but continues encoding with existing entries.

- **Decoding (decompression)**:
  - Can reads from any file, should be a `.lzw` file, or `stdin`.
  - Produces a decompressed output file (`output.out`) or writes to `stdout` based on `CLI` Env variable.
  - Mirrors the encoder’s logic (9-bit initial code size, grows to 15 bits).
  - Ignores any leftover bits that cannot form a valid code (to avoid out-of-range dictionary references).

## Requirements

- **Go 1.18+** (or a reasonably recent Go version).
- A POSIX-like shell (if using the provided Makefile or shell commands).

## Installation & Usage

### 1. Clone & Build

```bash
git clone https://github.com/yourusername/go-lzw.git
cd go-lzw
make
```

This will:
- Build the main Go binary (`go-lzw`).
- Create two symlinks: `encode` and `decode` pointing to that binary.

### 2. Compress (Encode)

To compress a file `input.txt`:

```bash
./encode input.txt
```

- By default, this writes a compressed file named `encoded.lzw`.
- Alternatively, set the environment variable `CLI=1` to write compressed data to `stdout`:

  ```bash
  export CLI=1
  ./encode input.txt > mycompressed.lzw
  ```

### 3. Decompress (Decode)

To decompress `output.lzw`:

```bash
./decode output.lzw
```

- By default, this writes the decompressed data to `output.out`.
- If `CLI=1` is set, it writes the uncompressed data to `stdout`:

  ```bash
  export CLI=1
  ./decode output.lzw > original_again.txt
  ```

### 4. Verify Round Trip

```bash
echo "Hello, World!" > sample.txt
./encode sample.txt
./decode output.lzw
diff sample.txt output.out
```

If they match, compression & decompression worked correctly.

## Makefile

- **`make`** or **`make all`**: Builds the `go-lzw` binary and creates the `encode` / `decode` symlinks.
- **`make clean`**: Removes the main binary, symlinks, and any leftover `.lzw` or `.out` files.

## Implementation Details

- **Dictionary Growth**  
  - Starts with 256 single-byte entries (`codes 0..255`).  
  - Grows code size from 9 bits up to 15 bits.  
  - Stops creating new dictionary entries once it reaches a max code of 32767.

- **File Output Strategy**  
  - By default, `encode` creates `output.lzw`; `decode` creates `output.out`.  
  - If the environment variable `CLI=1` is present, both commands write to `stdout` instead.

## Contributing

1. Fork the repository and create a new branch for your feature or bugfix.
2. Write clear commit messages and ensure the code passes any tests or lint checks.
3. Open a pull request, and describe the changes or issue you’re addressing.

## License

This project is distributed under the [MIT License](LICENSE). You’re free to use, modify, and distribute it with minimal restrictions.

## Acknowledgments

- **LZW Algorithm**: Terry Welch (1984).

---

Feel free to add more details (such as performance metrics, references, or advanced usage notes) as your project evolves.
