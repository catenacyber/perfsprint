package p

import (
	"errors"
	"fmt" // want "Fix imports"
	"io"
	"log"
	"net/url"
	"os"
	"time"
)

var errSentinel = errors.New("connection refused")

func positive() {
	var s string
	fmt.Sprintf("%s", "hello") // want "fmt.Sprintf can be replaced with just using the string"
	fmt.Sprintf("%v", "hello") // want "fmt.Sprintf can be replaced with just using the string"
	fmt.Sprintf("hello")       // want "fmt.Sprintf can be replaced with just using the string"
	fmt.Sprint("hello")        // want "fmt.Sprint can be replaced with just using the string"
	fmt.Sprintf("%s", s)       // want "fmt.Sprintf can be replaced with just using the string"
	fmt.Sprintf("%[1]s", s)    // want "fmt.Sprintf can be replaced with just using the string"
	fmt.Sprintf("%v", s)       // want "fmt.Sprintf can be replaced with just using the string"
	fmt.Sprint(s)              // want "fmt.Sprint can be replaced with just using the string"
	fmt.Errorf("hello")        // want "fmt.Errorf can be replaced with errors.New"

	fmt.Sprintf("Hello %s", s)         // want "fmt.Sprintf can be replaced with string concatenation"
	fmt.Sprintf("%s says Hello", s)    // want "fmt.Sprintf can be replaced with string concatenation"
	fmt.Sprintf("Hello says %[1]s", s) // want "fmt.Sprintf can be replaced with string concatenation"

	var err error
	fmt.Sprintf("%s", errSentinel) // want "fmt.Sprintf can be replaced with errSentinel.Error()"
	fmt.Sprintf("%v", errSentinel) // want "fmt.Sprintf can be replaced with errSentinel.Error()"
	fmt.Sprint(errSentinel)        // want "fmt.Sprint can be replaced with errSentinel.Error()"
	fmt.Sprintf("%s", io.EOF)      // want "fmt.Sprintf can be replaced with io.EOF.Error()"
	fmt.Sprintf("%v", io.EOF)      // want "fmt.Sprintf can be replaced with io.EOF.Error()"
	fmt.Sprint(io.EOF)             // want "fmt.Sprint can be replaced with io.EOF.Error()"
	fmt.Sprintf("%s", err)         // want "fmt.Sprintf can be replaced with err.Error()"
	fmt.Sprintf("%v", err)         // want "fmt.Sprintf can be replaced with err.Error()"
	fmt.Sprint(err)                // want "fmt.Sprint can be replaced with err.Error()"

	var b bool
	fmt.Sprintf("%t", true)  // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
	fmt.Sprintf("%v", true)  // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
	fmt.Sprint(true)         // want "fmt.Sprint can be replaced with faster strconv.FormatBool"
	fmt.Sprintf("%t", false) // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
	fmt.Sprintf("%v", false) // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
	fmt.Sprint(false)        // want "fmt.Sprint can be replaced with faster strconv.FormatBool"
	fmt.Sprintf("%t", b)     // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
	fmt.Sprintf("%v", b)     // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
	fmt.Sprint(b)            // want "fmt.Sprint can be replaced with faster strconv.FormatBool"

	var bs []byte
	var ba [3]byte
	fmt.Sprintf("%x", []byte{'a'})  // want "fmt.Sprintf can be replaced with faster hex.EncodeToString"
	fmt.Sprintf("%x", []uint8{'b'}) // want "fmt.Sprintf can be replaced with faster hex.EncodeToString"
	fmt.Sprintf("%x", bs)           // want "fmt.Sprintf can be replaced with faster hex.EncodeToString"
	fmt.Sprintf("%x", ba)           // want "fmt.Sprintf can be replaced with faster hex.EncodeToString"

	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	fmt.Sprintf("%d", i)         // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", i)         // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(i)                // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", 42)        // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", 42)        // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(42)               // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", i8)        // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", i8)        // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(i8)               // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", int8(42))  // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", int8(42))  // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(int8(42))         // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", i16)       // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", i16)       // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(i16)              // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", int16(42)) // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", int16(42)) // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(int16(42))        // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", i32)       // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", i32)       // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(i32)              // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", int32(42)) // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", int32(42)) // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(int32(42))        // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", i64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprintf("%v", i64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprint(i64)              // want "fmt.Sprint can be replaced with faster strconv.FormatInt"
	fmt.Sprintf("%d", int64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprintf("%v", int64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprint(int64(42))        // want "fmt.Sprint can be replaced with faster strconv.FormatInt"

	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	fmt.Sprintf("%d", ui)         // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", ui)         // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(ui)                // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", uint(42))   // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", uint(42))   // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(uint(42))          // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", ui8)        // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", ui8)        // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(ui8)               // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", uint8(42))  // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", uint8(42))  // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(uint8(42))         // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", ui16)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", ui16)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(ui16)              // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", uint16(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", uint16(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(uint16(42))        // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", ui32)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", ui32)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(ui32)              // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", uint32(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", uint32(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(uint32(42))        // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", ui64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", ui64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%x", ui64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%x", uint(42))   // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(ui64)              // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", uint64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", uint64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(uint64(42))        // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
}

func suggestedFixesTest() {
	_ = func() string {
		if false {
			return fmt.Sprint("replace me") // want "fmt.Sprint can be replaced with just using the string"
		}
		return fmt.Sprintf("%s", "replace me") // want "fmt.Sprintf can be replaced with just using the string"
	}

	fmt.Println(fmt.Sprint(errSentinel))        // want "fmt.Sprint can be replaced with errSentinel.Error()"
	fmt.Println(fmt.Sprintf("%s", errSentinel)) // want "fmt.Sprintf can be replaced with errSentinel.Error()"

	_ = func() string {
		switch 42 {
		case 1:
			return fmt.Sprint(false) // want "fmt.Sprint can be replaced with faster strconv.FormatBool"
		case 2:
			return fmt.Sprintf("%t", true) // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
		}
		return ""
	}

	var offset int
	params := url.Values{}
	params.Set("offset", fmt.Sprintf("%d", offset)) // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	params.Set("offset", fmt.Sprint(offset))        // want "fmt.Sprint can be replaced with faster strconv.Itoa"

	var pubKey []byte
	if verifyPubKey := true; verifyPubKey {
		log.Println("pubkey=" + fmt.Sprintf("%x", pubKey)) // want "fmt.Sprintf can be replaced with faster hex.EncodeToString"
	}

	var metaHash [16]byte
	fn := fmt.Sprintf("%s.%x", "coverage.MetaFilePref", metaHash)
	_ = "tmp." + fn + fmt.Sprintf("%d", time.Now().UnixNano()) // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	_ = "tmp." + fn + fmt.Sprint(time.Now().UnixNano())        // want "fmt.Sprint can be replaced with faster strconv.FormatInt"

	var change struct{ User struct{ ID uint } }
	var userStr string
	if id := change.User.ID; id != 0 {
		userStr = fmt.Sprintf("%d", id) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
		userStr = fmt.Sprint(id)        // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	}
	_ = userStr
}

func negative() {
	const val = "val%d"

	_ = int32(42)

	fmt.Scan(42)
	fmt.Scanf("%d", 42)
	fmt.Println("%d", 42)
	fmt.Printf("%d")
	fmt.Printf("%v")
	fmt.Printf("%d", 42)
	fmt.Printf("%s %d", "hello", 42)
	fmt.Errorf("this is %s", "complex")

	fmt.Fprint(os.Stdout, "%d", 42)
	fmt.Fprintf(os.Stdout, "test")
	fmt.Fprintf(os.Stdout, "%d")
	fmt.Fprintf(os.Stdout, "%v")
	fmt.Fprintf(os.Stdout, "%d", 42)
	fmt.Fprintf(os.Stdout, "%s %d", "hello", 42)

	fmt.Sprint("test", 42)
	fmt.Sprint(42, 42)
	fmt.Sprintf("%d", 42, 42)
	fmt.Sprintf("%#d", 42)
	fmt.Sprintf("value %d", 42)
	fmt.Sprintf(val, 42)
	fmt.Sprintf("%s %v", "hello", "world")
	fmt.Sprintf("%#v", 42)
	fmt.Sprintf("%T", struct{ string }{})
	fmt.Sprintf("%%v", 42)
	fmt.Sprintf("%3d", 42)
	fmt.Sprintf("% d", 42)
	fmt.Sprintf("%-10d", 42)
	fmt.Sprintf("%s %[1]s", "hello")
	fmt.Sprintf("%[1]s %[1]s", "hello")
	fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)

	// Integer.
	fmt.Sprintf("%#x", uint64(42))
	fmt.Sprintf("%#v", uint64(42))
	fmt.Sprintf("%#b", 42)
	fmt.Sprintf("%#o", 42)
	fmt.Sprintf("%#x", 42)
	fmt.Sprintf("%#X", 42)

	fmt.Sprintf("%b", 42)
	fmt.Sprintf("%c", 42)
	fmt.Sprintf("%o", 42)
	fmt.Sprintf("%O", 42)
	fmt.Sprintf("%q", 42)
	fmt.Sprintf("%x", 42)
	fmt.Sprintf("%X", 42)

	// Floating point.
	fmt.Sprintf("%9f", 42.42)
	fmt.Sprintf("%.2f", 42.42)
	fmt.Sprintf("%.2f", 42.42)
	fmt.Sprintf("%9.2f", 42.42)
	fmt.Sprintf("%9.f", 42.42)
	fmt.Sprintf("%.3g", 42.42)

	fmt.Sprintf("%b", float32(42.42))
	fmt.Sprintf("%e", float32(42.42))
	fmt.Sprintf("%E", float32(42.42))
	fmt.Sprintf("%f", float32(42.42))
	fmt.Sprintf("%F", float32(42.42))
	fmt.Sprintf("%g", float32(42.42))
	fmt.Sprintf("%G", float32(42.42))
	fmt.Sprintf("%x", float32(42.42))
	fmt.Sprintf("%X", float32(42.42))
	fmt.Sprintf("%v", float32(42.42))

	fmt.Sprintf("%b", 42.42)
	fmt.Sprintf("%e", 42.42)
	fmt.Sprintf("%E", 42.42)
	fmt.Sprintf("%f", 42.42)
	fmt.Sprintf("%F", 42.42)
	fmt.Sprintf("%g", 42.42)
	fmt.Sprintf("%G", 42.42)
	fmt.Sprintf("%x", 42.42)
	fmt.Sprintf("%X", 42.42)
	fmt.Sprintf("%v", 42.42)

	fmt.Sprintf("%b", 42i+42)
	fmt.Sprintf("%e", 42i+42)
	fmt.Sprintf("%E", 42i+42)
	fmt.Sprintf("%f", 42i+42)
	fmt.Sprintf("%F", 42i+42)
	fmt.Sprintf("%g", 42i+42)
	fmt.Sprintf("%G", 42i+42)
	fmt.Sprintf("%x", 42i+42)
	fmt.Sprintf("%X", 42i+42)
	fmt.Sprintf("%v", 42i+42)

	// String & slice of bytes.
	fmt.Sprintf("%q", "hello")
	fmt.Sprintf("%#q", `"hello"`)
	fmt.Sprintf("%+q", "hello")
	fmt.Sprintf("%X", "hello")

	// Slice.
	fmt.Sprintf("%x", []uint16{'d'})
	fmt.Sprintf("%x", []uint32{'d'})
	fmt.Sprintf("%x", []uint64{'d'})
	fmt.Sprintf("%x", []uint{'d'})
	fmt.Sprintf("%x", [1]byte{'c'})
	fmt.Sprintf("%x", [1]uint8{'d'})
	fmt.Sprintf("%x", [1]uint16{'d'})
	fmt.Sprintf("%x", [1]uint32{'d'})
	fmt.Sprintf("%x", [1]uint64{'d'})
	fmt.Sprintf("%x", [1]uint{'d'})
	fmt.Sprintf("%x", []int8{1})
	fmt.Sprintf("%x", []int16{1})
	fmt.Sprintf("%x", []int32{1})
	fmt.Sprintf("%x", []int64{1})
	fmt.Sprintf("%x", []int{1})
	fmt.Sprintf("%x", [...]int8{1})
	fmt.Sprintf("%x", [...]int16{1})
	fmt.Sprintf("%x", [...]int32{1})
	fmt.Sprintf("%x", [...]int64{1})
	fmt.Sprintf("%x", [...]int{1})
	fmt.Sprintf("%x", []string{"hello"})
	fmt.Sprintf("%x", []rune{'a'})

	fmt.Sprintf("% x", []byte{1, 2, 3})
	fmt.Sprintf("% X", []byte{1, 2, 3})
	fmt.Sprintf("%p", []byte{1, 2, 3})
	fmt.Sprintf("%#p", []byte{1, 2, 3})

	// Pointer.
	var ptr *int
	fmt.Sprintf("%v", ptr)
	fmt.Sprintf("%b", ptr)
	fmt.Sprintf("%d", ptr)
	fmt.Sprintf("%o", ptr)
	fmt.Sprintf("%x", ptr)
	fmt.Sprintf("%X", ptr)
}

func malformed() {
	fmt.Sprintf("%d", "example")
	fmt.Sprintf("%T", "example")
	fmt.Sprintf("%t", "example")
	fmt.Sprintf("%b", "example")
	fmt.Sprintf("%e", "example")
	fmt.Sprintf("%E", "example")
	fmt.Sprintf("%f", "example")
	fmt.Sprintf("%F", "example")
	fmt.Sprintf("%g", "example")
	fmt.Sprintf("%G", "example")
	fmt.Sprintf("%x", "example")
	fmt.Sprintf("%X", "example")

	fmt.Sprintf("%d", errSentinel)
	fmt.Sprintf("%T", errSentinel)
	fmt.Sprintf("%t", errSentinel)
	fmt.Sprintf("%b", errSentinel)
	fmt.Sprintf("%e", errSentinel)
	fmt.Sprintf("%E", errSentinel)
	fmt.Sprintf("%f", errSentinel)
	fmt.Sprintf("%F", errSentinel)
	fmt.Sprintf("%g", errSentinel)
	fmt.Sprintf("%G", errSentinel)
	fmt.Sprintf("%x", errSentinel)
	fmt.Sprintf("%X", errSentinel)

	fmt.Sprintf("%d", true)
	fmt.Sprintf("%T", true)
	fmt.Sprintf("%b", true)
	fmt.Sprintf("%e", true)
	fmt.Sprintf("%E", true)
	fmt.Sprintf("%f", true)
	fmt.Sprintf("%F", true)
	fmt.Sprintf("%g", true)
	fmt.Sprintf("%G", true)
	fmt.Sprintf("%x", true)
	fmt.Sprintf("%X", true)

	var bb []byte
	fmt.Sprintf("%d", bb)
	fmt.Sprintf("%T", bb)
	fmt.Sprintf("%t", bb)
	fmt.Sprintf("%b", bb)
	fmt.Sprintf("%e", bb)
	fmt.Sprintf("%E", bb)
	fmt.Sprintf("%f", bb)
	fmt.Sprintf("%F", bb)
	fmt.Sprintf("%g", bb)
	fmt.Sprintf("%G", bb)
	fmt.Sprintf("%X", bb)
	fmt.Sprintf("%v", bb)

	fmt.Sprintf("%T", 42)
	fmt.Sprintf("%t", 42)
	fmt.Sprintf("%b", 42)
	fmt.Sprintf("%e", 42)
	fmt.Sprintf("%E", 42)
	fmt.Sprintf("%f", 42)
	fmt.Sprintf("%F", 42)
	fmt.Sprintf("%g", 42)
	fmt.Sprintf("%G", 42)
	fmt.Sprintf("%x", 42)
	fmt.Sprintf("%X", 42)

	fmt.Sprintf("%T", uint(42))
	fmt.Sprintf("%t", uint(42))
	fmt.Sprintf("%b", uint(42))
	fmt.Sprintf("%e", uint(42))
	fmt.Sprintf("%E", uint(42))
	fmt.Sprintf("%f", uint(42))
	fmt.Sprintf("%F", uint(42))
	fmt.Sprintf("%g", uint(42))
	fmt.Sprintf("%G", uint(42))
	fmt.Sprintf("%X", uint(42))

	fmt.Sprintf("%d", 42.42)
	fmt.Sprintf("%d", map[string]string{})
	fmt.Sprint(make(chan int))
	fmt.Sprint([]int{1, 2, 3})
}
