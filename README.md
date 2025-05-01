# 🧠 go-dsa-practice
Practice and master data structures and algorithms through Go programming. Each problem includes clear, idiomatic Go solutions with tests and benchmarks.

---

## 📂 Repository Structure
```
go-dsa-practice/
├── README.md
├── go.mod
├── problems/
│   ├── arrays/
│   │   ├── two_sum.go
│   │   ├── two_sum_test.go
│   │   ├── maximum_average_subarray.go
│   │   ├── maximum_average_subarray_test.go
│   ├── numbers/
│   │   ├── palindrome.go
│   │   ├── palindrome_test.go
│   ├── strings/
│   │   ├── is_palindrome.go
│   │   ├── is_palindrome_test.go
│   ├── linkedlist/
│   └── trees/
│
├── utils/
│   └── helpers.go  # Common utilities
│
├── benchmarks/
│   └── benchmark_suite_test.go
│
└── .gitignore
```
---

## 🚀 Getting Started

### 🛠 Requirements
- [Go 1.20+](https://golang.org/dl/)

### 🧪 Run All Tests
```bash
go test ./...
```

### 📏 Run Benchmarks
```bash
go test -bench=. ./...
```

### 🗂 Problem Categories
| Category    | Problems Included                             |
|-------------|-----------------------------------------------|
| Arrays      | Two Sum, Max Avg Subarray, Merge Sorted Arrays    |
<!--# | Strings     | Palindrome Check, Anagram, Longest Substring  |
# | Linked List | Reverse List, Detect Cycle, Merge Lists       |
# | Trees       | BFS, DFS, Max Depth, Invert Binary Tree       |
# | Sorting     | Bubble Sort, Merge Sort, Quick Sort           |
# | Graphs      | BFS, DFS, Cycle Detection, Dijkstra           | -->


# 🤝 Contribution
### Contributions are welcome! Please:

- Fork the repo
- Create a new branch: git checkout -b feature/problem-name
- Add your solution in the appropriate folder
- Add tests and optional benchmarks
- Open a Pull Request
