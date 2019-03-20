// PACKAGE = collection of common source code
// types of packages executable - must have func called 'main'
//					 reusable - defines a package that can be used as a dependency
// main = secret name it means that file will be runnable
package main

// import standard libs (fmt) and reusable packages
import "fmt"

func main() {
	fmt.Println("Hi there!")
}
