# Golang Laboratory - Hands-On Tutorial

A comprehensive, hands-on tutorial for learning the Go programming language from basics to building web applications and REST APIs.

## 🎯 What You'll Learn

This tutorial takes you through essential Go concepts with practical, runnable examples:

1. **Basic Syntax** - Hello World, packages, imports
2. **Data Types** - Variables, constants, and Go's type system
3. **Functions** - Function declaration, methods, closures
4. **OOP Concepts** - Structs, interfaces, composition
5. **Memory Management** - Pointers, memory allocation
6. **Control Flow** - Loops, conditionals, switch statements
7. **Error Handling** - Go's error model and best practices
8. **Concurrency** - Goroutines, channels, synchronization
9. **Web Development** - HTTP servers, routing, middleware
10. **APIs** - JSON handling, RESTful API development

## 📋 Prerequisites

- Basic programming knowledge (any language)
- Computer with Windows 11, macOS, or Linux
- Internet connection for downloading Go

## 🛠 Setup Instructions

### Windows 11 Setup

#### Option 1: Native Windows Installation

1. **Download Go**
   - Visit [https://golang.org/dl/](https://golang.org/dl/)
   - Download the Windows installer (`.msi` file)
   - Run the installer and follow the setup wizard

2. **Verify Installation**
   - Open Command Prompt or PowerShell
   - Run: `go version`
   - You should see output like: `go version go1.21.x windows/amd64`

3. **Set Up Environment (if needed)**
   - Go installer usually sets up PATH automatically
   - If not, add `C:\Go\bin` to your PATH environment variable

4. **Choose a Code Editor**
   - **VS Code** (Recommended): Download from [https://code.visualstudio.com/](https://code.visualstudio.com/)
     - Install the "Go" extension by Google
   - **GoLand**: JetBrains IDE specifically for Go
   - **Any text editor**: Notepad++, Sublime Text, etc.

#### Option 2: WSL (Windows Subsystem for Linux)

1. **Enable WSL**
   ```powershell
   # Run in PowerShell as Administrator
   wsl --install
   # Restart your computer when prompted
   ```

2. **Install Ubuntu**
   ```powershell
   wsl --install -d Ubuntu
   ```

3. **Update Ubuntu**
   ```bash
   sudo apt update && sudo apt upgrade -y
   ```

4. **Install Go in WSL**
   ```bash
   # Remove any old Go installation
   sudo rm -rf /usr/local/go
   
   # Download Go (check for latest version at https://golang.org/dl/)
   wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
   
   # Extract to /usr/local
   sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
   
   # Add Go to PATH
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

5. **Verify Installation**
   ```bash
   go version
   ```

### macOS Setup

1. **Using Homebrew (Recommended)**
   ```bash
   # Install Homebrew if you don't have it
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   
   # Install Go
   brew install go
   ```

2. **Manual Installation**
   - Download macOS installer from [https://golang.org/dl/](https://golang.org/dl/)
   - Run the installer package

3. **Verify Installation**
   ```bash
   go version
   ```

### Linux Setup

1. **Ubuntu/Debian**
   ```bash
   # Update package list
   sudo apt update
   
   # Install Go
   sudo apt install golang-go
   
   # Or install latest version manually:
   wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
   sudo rm -rf /usr/local/go
   sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

2. **Verify Installation**
   ```bash
   go version
   ```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
# Using HTTPS
git clone https://github.com/ly2xxx/golang_lab.git
cd golang_lab

# Or using SSH (if you have SSH keys set up)
git clone git@github.com:ly2xxx/golang_lab.git
cd golang_lab
```

### 2. Verify Your Setup

```bash
# Test with the first lesson
cd lesson01-hello-world
go run main.go
```

You should see:
```
Hello, World!
Hello from Go!
Hello Gopher!
Golang was first released in 2009
```

### 3. Initialize Go Module (Optional)

If you want to experiment with packages:
```bash
# In the root directory
go mod init golang-lab
go mod tidy
```

## 📚 Tutorial Structure

Each lesson is self-contained with:
- **main.go**: Working code examples
- **README.md**: Detailed explanations and concepts
- **Additional files**: Where needed (HTML, CSS, test scripts)

### Lesson Progression

| Lesson | Topic | Key Concepts | Estimated Time |
|--------|-------|--------------|----------------|
| [01](./lesson01-hello-world/) | Hello World | Package, imports, functions | 30 min |
| [02](./lesson02-variables-types/) | Variables & Types | Declaration, types, constants | 45 min |
| [03](./lesson03-functions-methods/) | Functions & Methods | Parameters, returns, methods | 60 min |
| [04](./lesson04-structs-interfaces/) | Structs & Interfaces | OOP, composition, polymorphism | 75 min |
| [05](./lesson05-pointers-memory/) | Pointers & Memory | Memory management, pointers | 60 min |
| [06](./lesson06-control-structures/) | Control Structures | Loops, conditionals, flow control | 45 min |
| [07](./lesson07-error-handling/) | Error Handling | Error types, panic/recover | 60 min |
| [08](./lesson08-concurrency/) | Concurrency | Goroutines, channels, sync | 90 min |
| [09](./lesson09-web-server/) | Web Servers | HTTP, routing, middleware | 75 min |
| [10](./lesson10-json-rest-api/) | JSON & REST APIs | JSON, REST, API design | 90 min |

**Total estimated time: 10-12 hours**

## 🏃‍♂️ How to Use This Tutorial

### For Beginners
1. **Follow lessons sequentially** - Each builds on previous concepts
2. **Type the code yourself** - Don't just copy-paste
3. **Experiment** - Modify examples and see what happens
4. **Read error messages** - Go has helpful error messages
5. **Practice** - Do the "Try It Yourself" sections

### For Experienced Programmers
1. **Jump to specific topics** - Use the lesson structure to find what you need
2. **Focus on Go-specific concepts** - Especially lessons 4, 5, 7, 8
3. **Build projects** - Combine concepts from multiple lessons

### Running Individual Lessons

```bash
# Navigate to any lesson
cd lesson01-hello-world

# Run the code
go run main.go

# Or build and run
go build -o hello main.go
./hello
```

### Running Web Server Lessons

```bash
# For lessons 9 and 10
cd lesson09-web-server
go run main.go
# Visit http://localhost:8080 in your browser

cd lesson10-json-rest-api
go run main.go
# Test API with curl or visit http://localhost:8080/api
```

## 🔧 Troubleshooting

### Common Issues

**1. "go: command not found"**
- Go is not installed or not in PATH
- Restart terminal after installation
- Check PATH: `echo $PATH` (Linux/Mac) or `echo $env:PATH` (Windows PowerShell)

**2. "package main is not in GOPATH"**
- You're using an old version of Go
- Update to Go 1.11+ which has module support
- Or run from correct directory

**3. "port already in use" (Lessons 9-10)**
- Another service is using port 8080
- Change port in main.go: `:8080` → `:8081`
- Or kill existing process

**4. Permission errors (Linux/Mac)**
```bash
sudo chown -R $USER:$USER /usr/local/go
```

### WSL-Specific Issues

**1. "cannot connect to server" when running web servers**
- Windows Firewall might be blocking
- Try accessing via `localhost` or `127.0.0.1`
- In WSL2, you might need the Windows IP

**2. File permission issues**
```bash
chmod +x test_api.sh  # Make script executable
```

## 📖 Additional Resources

### Official Documentation
- [Go Documentation](https://golang.org/doc/)
- [Go Tour](https://tour.golang.org/) - Interactive tutorial
- [Effective Go](https://golang.org/doc/effective_go.html) - Best practices
- [Go by Example](https://gobyexample.com/) - Code examples

### Books
- "The Go Programming Language" by Donovan & Kernighan
- "Go in Action" by Kennedy, Ketelsen & St. Martin
- "Learning Go" by Jon Bodner

### Online Courses
- [Go Fundamentals](https://www.pluralsight.com/courses/go-fundamentals)
- [Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)

### Community
- [Go Forum](https://forum.golangbridge.org/)
- [Reddit r/golang](https://www.reddit.com/r/golang/)
- [Gopher Slack](https://gophers.slack.com/)

## 🤝 Contributing

Found an issue or want to improve the tutorial?

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b improve-lesson-x`
3. **Make your changes**
4. **Test thoroughly**
5. **Submit a pull request**

### Areas for Contribution
- Additional exercises and challenges
- More real-world examples
- Performance optimization examples
- Advanced topics (testing, deployment, etc.)
- Translations to other languages

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🏆 Learning Goals Checklist

After completing this tutorial, you should be able to:

- [ ] Write basic Go programs with proper syntax
- [ ] Understand Go's type system and variable declarations
- [ ] Create and use functions and methods
- [ ] Design programs using structs and interfaces
- [ ] Manage memory with pointers effectively
- [ ] Use control structures for program flow
- [ ] Handle errors properly using Go's error model
- [ ] Write concurrent programs with goroutines and channels
- [ ] Build HTTP web servers and APIs
- [ ] Work with JSON data and create RESTful services
- [ ] Apply Go best practices and idioms
- [ ] Debug and troubleshoot Go programs

## 🎯 Next Steps

After completing this tutorial, consider:

1. **Build a real project**: Combine multiple concepts
2. **Learn testing**: Go has excellent testing support
3. **Explore frameworks**: Gin, Echo, Fiber for web development
4. **Study databases**: Working with SQL and NoSQL databases
5. **Learn deployment**: Docker, Kubernetes, cloud platforms
6. **Contribute to open source**: Practice with real Go projects

## ⭐ Feedback

If this tutorial helped you learn Go, please:
- ⭐ Star this repository
- 🍴 Fork and share with others
- 📝 Leave feedback in issues
- 📢 Share on social media

Happy coding with Go! 🐹

---

*Last updated: July 2025*
*Go version: 1.21+*