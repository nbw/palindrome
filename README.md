# Palindrome

Checks if a string is a palindrome.

If you complile a binary you then you can run it as a CLI tool. This assumes you have golang installed.

# Install

`go get -u github.com/nbw/palindrome`

## Code Summary

`palindorome/palindrome.go`

The steps are:
1. Remove any punctuation (operation: O(n))
2. Lower case the string (operation: O(n))
3. Run through the string and check if it's a palindrome (operation: O(n/2))

The complexity is O(n) + O(n) + O(n/2) = O(2.5n) => effectively O(n)

### Tests

Tests are in `palindrome_test.go`. You can run them with `Make test`

### Thoughts
I considered running through the string only once and performing punctuation removal and lowercasing on a per-letter basis, but you'd lose efficiency checking for punctation on a per-letter basis versus once with the entire string. Same thing with lowercase.

## CLI Usage

```
# Compile a binary
Make install

# Using the CLI tool
# - from the root of the project

palindrome "racecar"
palindrome" "A toyota's a Toyata"
```

## Compile binary

You can compile a binary with `Make install`

## Testing

Run `Make test`
