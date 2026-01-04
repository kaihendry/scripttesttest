# Script Test Harness

This directory contains script-based tests for the `kai` command using the `rsc.io/script` framework.

## Running Tests

To run the script tests, use the following command:

```bash
go test -v ./cmd/
```

## Writing Test Scripts

Test scripts are text files in this directory with a `.txt` extension. Each script contains commands and assertions.

### Basic Syntax

- `# Comment` - Comments start with `#`
- `kai <args>` - Execute the kai command with arguments
- `stdout <pattern>` - Assert that stdout matches the pattern
- `stderr <pattern>` - Assert that stderr matches the pattern
- `! <command>` - Assert that the command fails (non-zero exit code)
- `! stdout <pattern>` - Assert that stdout does NOT match the pattern
- `! stderr <pattern>` - Assert that stderr does NOT match the pattern

### Pattern Matching

Patterns can be:
- Plain text (matches if the text appears anywhere in the output)
- `.` (matches any non-empty output)
- Regular expressions

### Example Test Scripts

#### help.txt
```
# Test that kai --help works
kai --help
stdout 'A longer description'
! stderr .
! stderr 'Error'
```

#### version.txt
```
# Test that kai command runs without error
kai
! stderr 'Error'
```

#### invalid_flag.txt
```
# Test that invalid flags are rejected
! kai --invalid-flag-that-does-not-exist
stderr .
```

## Additional Commands

The script framework provides many built-in commands including:
- `env` - Set environment variables
- `cd` - Change directory
- `mkdir` - Create directories
- `cp`, `mv`, `rm` - File operations
- `grep`, `cat` - Text operations
- `exists` - Check if files exist
- And many more...

See the [rsc.io/script documentation](https://pkg.go.dev/rsc.io/script) for the complete list of commands.

## Benefits

- **Fast**: Tests run in parallel
- **Isolated**: Each test runs in its own temporary directory
- **Simple**: Easy to read and write test scenarios
- **Comprehensive**: Can test entire command-line workflows
- **Reproducible**: Tests are deterministic and don't depend on system state
