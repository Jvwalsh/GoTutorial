# Structs
    # They use named fields
type Circle struct {
  x float64
  y float64
  r float64
}
==
type Circle struct {
  x, y, r float64
}
    //structs must begin with
        1) type
        2) name
        3) define 'struct' type
        4) series of names and their types

# Fields 
    See struct field info notes.go.
    They can ge accessved using '.' notation like propterties
    some short hand methods for defining as well
    example calling circle area with a circle Struct

# methods 
    functions on structs
    on the notes.go file see how area is assigned to c of type *Circle
    (func (c *Circle) area() float64{
        funciton logic
    })



# embedded types
leave iout stype within the struct to make it embedded. it's then callable directly fro the sub-Name or parent name. 
see robot/name example in notes.go
//People can talk, an android is a person, therefore an android can talk.

# interfaces
see shape in notes.go

    Defined same way as struct with the keyword interface instead of struct