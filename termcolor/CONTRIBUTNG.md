# Contributing to TermColor

Thanks for your interest

## Development Setup

```bash
# Clone the repository
git clone https://github.com/phamio/termcolor.git
cd termcolor

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...
```

---

## Issue Guidelines
**One issue per problem**. This keeps discussions focused and makes tracking easier.

## Before Creating An Issue
- Check existing issues to avoid duplicates.
- Check if the problem has already being reported or fixed

---

## Issue Template
```markdown
**Go version:** `go version`
**OS:** (e.g., Linux, macOS, Windows 10/11)
**Termcolor version:** (e.g., v0.1.2)

**Problem description:**
(What happened? What did you expect?)

**Reproduction code:**
```go
// Minimal example that demonstrates the issue
```

---

## Pull Request Guidelines 
**Each pull request should address one issue**


---

## Areas for Improvement
- Windows testing across CMD, Powershell and Windows Terminal
- Performance optimizations
- Additional test coverage
- Documentation improvements

---

## Pull Request Guidelines
- Fork the repository
- Create a feature branch (git checkout -b feature/name)
- Ensure all tests pass (go test ./...)
- Follow Go conventions(go fmt, go lint)
- Submit a pull request to the `main` branch

---

## Reporting Issues
**Include**:
- Go version (go version)
- Operating system
- Minimal reproduction code

---

## Code of Conduct
Be respectful. Assume good intentions.

---

## License
By contributing, you agree your contributions will be licensed under **Apache License 2.0**
