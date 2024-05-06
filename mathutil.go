// mathutil.go
package mathutil

// Factorial returns the factorial of n.
func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}

// Fibonacci returns the n-th Fibonacci number.
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

