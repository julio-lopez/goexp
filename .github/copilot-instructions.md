# Kopia Copilot Instructions

## When reviewing code, focus on:

### Security Critical Issues
- Check for hardcoded secrets, API keys, or credentials
- Verify proper input validation and sanitization
- Review authentication and authorization logic

### Code Quality Essentials
- Functions should be focused and appropriately sized
- Use clear, descriptive naming conventions
- Ensure proper error handling throughout

### Performance Issues
- Spot inefficient loops and algorithmic issues
- Check for memory leaks and resource cleanup
- Review caching opportunities for expensive operations

## Review Style
- Be specific and actionable in feedback
- Explain the rationale behind recommendations
- Acknowledge good patterns when you see them
- Ask clarifying questions when code intent is unclear

## Review Test Coverage
- Ensure there are tests that cover and exercise the new or changed functionality

Always prioritize security vulnerabilities and performance issues that could impact users.

Always suggest changes to improve readability.

## Repository Overview

**Key Technologies:**
- primary language: Go

## Build Commands

```
go build ./...
```

## Testing

### Unit Tests (Standard)
```bash
go test ./...
```

## Common Issues & Workarounds

### Build Issues

- **Go version mismatch:** Building requires the Go toolchain with the version specified in `go.mod`. The `go-version-file` is used in GitHub Actions.


### Configuration Files

- `.github/workflows/*.yml` - GitHub Actions workflows

### Code Style
- Uses golangci-lint with formatters: gci, gofumpt
- Imports organized: standard, default, localmodule
- Tests use the `stretchr/testify` packages

## Important Notes

- **Do not modify go.mod/go.sum manually** - Use `go get` to update dependencies. CI runs `git checkout go.mod go.sum` after ci-setup to revert local changes from tool downloads.

- Do not commit executables or binary artifacts to the repository. Do not modify `.gitignore`