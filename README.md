# quadchecker

A Go CLI tool that identifies which quad function generated a given ASCII rectangle pattern.

**Author:** Theodore Vairaktaris
**Institution:** Zone01 Athens
**License:** MIT

---

## Overview

**quadchecker** reads an ASCII rectangle from standard input and determines which quad generator function (`quadA`, `quadB`, `quadC`, `quadD`, or `quadE`) created it by:

1. Calculating the width and height of the input shape
2. Running each quad generator with those dimensions
3. Comparing outputs to find matches
4. Reporting all matching generators (or "Not a quad function" if none match)

---

## Usage

### Basic Usage

```bash
./quadA 5 3 | ./quadchecker
```

**Output:**
```
[quadA] [5] [3]
```

### Multiple Matches

Some dimensions produce identical outputs for different quad functions:

```bash
./quadC 1 1 | ./quadchecker
```

**Output:**
```
[quadA] [1] [1] || [quadB] [1] [1] || [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

### Invalid Input

```bash
echo "invalid shape" | ./quadchecker
```

**Output:**
```
Not a quad function
```

---

## Building

```bash
go build -o quadchecker
```

---

## Requirements

The following quad generator executables must be in your `PATH` or current directory:
- `quadA`
- `quadB`
- `quadC`
- `quadD`
- `quadE`

---

## How It Works

### Algorithm

1. **Read Input:** Scan all lines from stdin
2. **Validate Shape:** Check that all lines have equal width
3. **Extract Dimensions:** Calculate width (characters per line) and height (line count)
4. **Generate Comparisons:** Run each quad generator with the extracted dimensions
5. **Match Detection:** Compare generated output byte-for-byte with input
6. **Report Results:** Output all matching generators in sorted order

### Key Functions

- **`readStdin()`** - Reads and normalizes input from standard input
- **`findWidthAndHeight()`** - Validates rectangular shape and extracts dimensions
- **`runGenerator()`** - Executes a quad generator and captures output
- **`main()`** - Orchestrates the comparison and formats results

---

## Output Format

**Single match:**
```
[quadName] [width] [height]
```

**Multiple matches (separated by ` || `):**
```
[quadA] [5] [3] || [quadE] [5] [3]
```

**No match:**
```
Not a quad function
```

---

## Edge Cases Handled

- Empty input
- Irregular shapes (non-rectangular)
- Quad generators not found in PATH
- Large inputs (1MB buffer)
- Trailing newlines
- Multiple matching functions

---

## Example Session

```bash
# Generate a 3x3 quadB
./quadB 3 3 > test.txt

# Check what generated it
cat test.txt | ./quadchecker
# Output: [quadB] [3] [3]

# Try with quadC 2x2
./quadC 2 2 | ./quadchecker
# Output: [quadC] [2] [2]
```

---

## License

MIT License - Free to use for educational purposes.

---

## Acknowledgments

Zone01 Athens for the project specification and learning framework.
