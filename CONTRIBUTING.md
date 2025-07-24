# Contributing to Golang Laboratory

Thank you for your interest in contributing to this Go tutorial! This document provides guidelines for contributing to the project.

## How to Contribute

### Reporting Issues

1. **Check existing issues** first to avoid duplicates
2. **Use clear, descriptive titles**
3. **Provide detailed descriptions** including:
   - What you expected to happen
   - What actually happened
   - Steps to reproduce
   - Your environment (OS, Go version)
   - Error messages (if any)

### Suggesting Enhancements

1. **Open an issue** with the "enhancement" label
2. **Describe the enhancement** in detail
3. **Explain why** this would be useful
4. **Consider the learning progression** - how does it fit?

### Contributing Code

1. **Fork the repository**
2. **Create a feature branch** from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Make your changes**
4. **Test thoroughly**
5. **Commit with clear messages**
6. **Push to your fork**
7. **Create a Pull Request**

## Code Guidelines

### Go Code Standards

1. **Follow Go conventions**:
   - Use `gofmt` to format code
   - Follow Go naming conventions
   - Add comments for exported functions and types
   - Keep functions focused and small

2. **Educational focus**:
   - Code should be clear and easy to understand
   - Include comments explaining concepts
   - Prefer simple, explicit code over clever tricks
   - Show best practices

3. **Testing**:
   - Ensure all code examples run without errors
   - Test on multiple platforms if possible
   - Verify output matches documentation

### Documentation Standards

1. **README files**:
   - Clear learning objectives
   - Step-by-step explanations
   - Working code examples
   - "Try It Yourself" sections
   - Proper markdown formatting

2. **Code comments**:
   - Explain the "why", not just the "what"
   - Use clear, beginner-friendly language
   - Highlight Go-specific concepts

## Types of Contributions

### Highly Welcome

- **Bug fixes** in existing lessons
- **Improved explanations** of concepts
- **Additional exercises** and challenges
- **Better error handling examples**
- **Platform-specific setup instructions**
- **Performance optimization examples**
- **Real-world use cases**

### Consider Carefully

- **New lessons** - ensure they fit the progression
- **Advanced topics** - might belong in a separate advanced tutorial
- **Framework-specific content** - focus on standard library first

### Please Avoid

- **Breaking changes** to existing lesson structure
- **Overly complex examples** for beginners
- **Framework-heavy solutions** where standard library suffices
- **Platform-specific code** without alternatives

## Pull Request Process

1. **Update documentation** if needed
2. **Add tests** for new functionality
3. **Update the README** if adding new lessons
4. **Squash commits** if you have many small commits
5. **Write clear PR description**:
   - What changes you made
   - Why you made them
   - How to test them

### PR Review Criteria

- ✅ Code follows Go conventions
- ✅ Documentation is clear and helpful
- ✅ Examples are educational and correct
- ✅ Changes maintain learning progression
- ✅ All code runs without errors
- ✅ Follows project structure

## Lesson Structure Guidelines

### Required Files

- `main.go` - Working code examples
- `README.md` - Lesson documentation

### README Template

```markdown
# Lesson XX: Topic Name

## Learning Objectives
- Objective 1
- Objective 2

## Key Concepts

### Concept 1
Explanation with code examples

### Concept 2
Explanation with code examples

## Running the Code
```bash
cd lessonXX-topic
go run main.go
```

## Try It Yourself
1. Exercise 1
2. Exercise 2
```

### Code Structure

```go
// Lesson XX: Topic Name
// Brief description of what this lesson covers

package main

import (
    // imports
)

func main() {
    fmt.Println("=== Lesson XX: Topic Name ===")
    
    // Demonstrate concepts with clear sections
    demonstrateConcept1()
    demonstrateConcept2()
}

func demonstrateConcept1() {
    fmt.Println("\n--- Concept 1 ---")
    // Clear, commented examples
}
```

## Getting Help

- **Discord/Slack**: Join Go community channels
- **Issues**: Ask questions in GitHub issues
- **Documentation**: Read existing lessons for examples

## Recognition

Contributors will be:
- Added to the contributors list
- Mentioned in release notes for significant contributions
- Credited in the lesson if they created/significantly improved it

## Code of Conduct

- Be respectful and inclusive
- Focus on helping others learn
- Provide constructive feedback
- Assume good intentions
- Follow GitHub's community guidelines

Thanks for helping make Go more accessible to everyone! 🐹