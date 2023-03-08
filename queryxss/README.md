# QueryXSS

This tools tests if a URL reflects a query key and value in the response body or in the response headers.

## Tests

The tool runs the following modifications to the URL before checking:

- [x] Request the URL as it is (no modifications)
- [x] Add a random token as the key and value
- [ ] Changes all the query's values to the random token
- [ ] Append the random token to all the query's keys
- [ ] Append the random token to all the query's values

After each test, the tool checks if the token is present in the response body or in the response headers.

## Usage

The tool receives a list of URLs from stdin and prints the results to stdout.

```bash
cat urls.txt | queryxss
```

## Install

### Using go install

```bash
go install github.com/vitorfhc/hacks/queryxss
```

### Clone the repository

```bash
git clone github.com/vitorfhc/hacks
cd hacks/queryxss
go install
```