# go-leetcode

A CLI tool to fetch LeetCode problems and generate solution/test/README templates.

## Features

- Fetch LeetCode problem by slug
- Auto-generate:
    - `solution.go` (e.g., `two_sum.go`)
    - `solution_test.go` (e.g., `two_sum_test.go`)
    - `README.md` with metadata and content

## Usage

```bash
# Fetch a problem by slug
go run main.go fetch two-sum
```
