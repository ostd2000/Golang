package main

import "fmt"

func main() {
  fmt.Printf("Hello %T %v", 5, 5)
  fmt.Println()

  var x string = fmt.Sprintf("Hello %T %v", 5, 5)
  println(x)

  // Boolean
  fmt.Printf("Hello %t", true)
  fmt.Println()

  // Binary
  fmt.Printf("Number: %b", 1025)
  fmt.Println()

  // Octal
  fmt.Printf("Number: %o", 1025)
  fmt.Println()

  // Decimal
  fmt.Println()
  fmt.Printf("Number: %d", 1025)

  // Hexadecimal
  fmt.Printf("Number: %x", 1025)
  fmt.Println()

  // Hexadecimal capital letters.
  fmt.Printf("Number: %X", 1025)
  fmt.Println()

  // Floating point number without scientific notation.
  fmt.Printf("Number: %f", 3.1415141514151415)
  fmt.Println()
  
  // Floating point number with scientific notation.
  fmt.Printf("Number: %e", 3.1415141514151415)
  fmt.Println()
 
  // Showing the whole floating point number.
  fmt.Printf("Number: %g", 3.1415141514151415)
  fmt.Println()

  // Rounding the floating point number to a certain precision.
  fmt.Printf("Number: %.2f", 3.145)    // 3.15
  fmt.Println()

  fmt.Printf("Number: %.f", 3.56789)    // 4
  fmt.Println()

  // String
  fmt.Printf("Hello %s", "Kaveh")
  fmt.Println()

  // String with double quotation.
  fmt.Printf("Hello %q", "Kaveh")
  fmt.Println()

  // String with padding to the left: Making the whole formatted string to the length we define.
  fmt.Printf("Hello %10s", "Kaveh")    // 10 - 5 = 5: We have 5 empty spaces.
  fmt.Println()

  // String with padding to the right: Making the whole formatted string to the length we define. 
  fmt.Printf("Hello %-10s are you good?", "Kaveh")    // 10 - 5 = 5: We have 5 empty spaces.
  fmt.Println()

  // Padding with "0" to the left: Making the whole formatted number to the length we define. 
  fmt.Printf("Number: %010d is cool", 5)
  fmt.Println()

  // "\n": new line
  fmt.Printf("Number: \n %d is cool", 5)
  fmt.Println()

  // "\t": tab
  fmt.Printf("Number: \t %d is cool", 5)
  fmt.Println()
}
