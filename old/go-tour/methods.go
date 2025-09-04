package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ? This is good practice to use pointer receivers
// ? it will automatically do (&v) for you
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Shape interface {
	Area() float64
	Circumference() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r *Rectangle) Circumference() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Type: %T - Area: %0.2f, Circumference: %0.2f\n", s, s.Area(), s.Circumference())
}

func EmptyInterface() {
	var i interface{} //? Empty interface can hold any type

	i = 23
	fmt.Printf("i is of type %T with value %v\n", i, i)

	i = "Hello World"
	fmt.Printf("i is now of type %T with value %v\n", i, i)

	i = Vertex{9, 1}
	fmt.Printf("i is now of type %T with value %v\n", i, i)

	i = MyFloat(3.14)
	fmt.Printf("i is now of type %T with value %v\n", i, i)
}

func EmptyInterfaceTypeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) //? panic because there is no ok
	fmt.Println(f)
}

func TypeSwitch(i interface{}) {
	switch v := i.(type) { //? type keyword is used to switch on the type of the interface
	case int:
		fmt.Printf("Type is int with value %d\n", v)
	case string:
		fmt.Printf("Type is string with value %s\n", v)
	case float64:
		fmt.Printf("Type is float64 with value %f\n", v)
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

func (v Vertex) String() string {
	//? This is a Stringer method, which allows us to define how the Vertex type should be represented as a string.
	return fmt.Sprintf("Vertex(X: %0.2f, Y: %0.2f)", v.X, v.Y)
}

type HttpError struct {
	StatusCode int
	Message    string
}

// ? Every error type implements the following interface
// ? type error interface {
// ?     Error() string
// ? }

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

func Err() error {
	// ? since the method is defined on a pointer receiver, we will return the pointer
	// ? The garbage collector will put the HttpError in the heap
	return &HttpError{
		StatusCode: 404,
		Message:    "Not Found",
	}
}

func PrintError() {
	// ? This is good practice since after the if is executed, the variable will be garbage collected
	if err := Err(); err != nil {
		fmt.Println(err)
	}
}

func IoReader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) //? with cap of 8
	for {
		n, err := r.Read(b) //? n is the number of bytes read
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) //? it should be 8 and then 6
		if err == io.EOF {
			break
		}
	}
}

type IPAddress [4]byte

func (ip IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func PrintIPAddress() {
	hosts := map[string]IPAddress{
		"localhost": {127, 0, 0, 1},
		"google":    {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%s: %s\n", name, ip)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func SqrtWithError(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	//? You can use for range to iterate a fixed number of times
	//? can also do for i := range 10
	for range 10 {
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

// ? type Reader interface {
// ?	Read(b []byte) (n int, err error)
// ? }

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil //? return the number of bytes read and no error
}

func ReaderExample() {
	reader.Validate(MyReader{})
}

type Rot13Reader struct {
	r io.Reader
}

func (r *Rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i := range n {
		char := b[i]
		if char >= 'A' && char <= 'Z' {
			b[i] = 'A' + (char-'A'+13)%26
		}

		if char >= 'a' && char <= 'z' {
			b[i] = 'a' + (char-'a'+13)%26
		}
	}

	return n, nil
}

func Rot13ReaderExample() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// ? type Image interface {
// ?    ColorModel() color.Model
// ?    Bounds() Rectangle
// ?    At(x, y int) color.Color
// ? }

func ImageExample() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

type Image struct {
	Width, Height int
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img *Image) At(x, y int) color.Color {
	v := uint8((x + y) % 256) //? Just a simple color based on x and y
	return color.RGBA{v, v, 255, 255}
}

func AnotherImageExample() {
	m := &Image{Width: 100, Height: 100}
	pic.ShowImage(m)
}
