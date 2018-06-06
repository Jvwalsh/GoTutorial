Pointers

# this code won't change x
func zero(x int) {
  x = 0
}
func main() {
  x := 5
  zero(x)
  fmt.Println(x) // x is still 5
} 
# this code will change x // see the two "*"s.   Will reference location in memory not vlaue
func zero(xPtr *int) {
  *xPtr = 0
}
func main() {
  x := 5
  zero(&x)       //& knwos to fetch the *x instead of x main(){&x} and zero(){*Ptr} are the same!
  fmt.Println(x) // x is 0
}

# Second option for reference pointing using "new" // Functions exact same way as above with &x
func one(xPtr *int) {
  *xPtr = 1
}
func main() {
  xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr) // x is 1
}
