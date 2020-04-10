// Code generated by go2go; DO NOT EDIT.


//line errors_test.go2:1
package errors

//line errors_test.go2:1
import "testing"

//line errors_test.go2:7
type e struct{}

func (a e) Error() string { return "error" }

func TestTryCatch(t *testing.T) {
//line errors_test.go2:11
 instantiate୦୦Try୦int(func() (int, error) {
//line errors_test.go2:13
  return 0, nil
	}).Catch(func(_ int, err error) int {
		t.Fatalf("catch error: %v", err)
		return 0
	}).Final(func(result int) {
		t.Log("errors: everything is good")
	})
//line errors_test.go2:19
 instantiate୦୦Try୦int(func() (int, error) {
//line errors_test.go2:22
  return 1, e{}
	}).Catch(func(result int, err error) int {
		t.Logf("captured result: %v", result)
		t.Logf("captured error: %v", err)
		return 1
	}).Final(func(result int) {
		if result != 1 {
			t.Fatalf("cannot capture error")
		}
	})
//line errors_test.go2:31
 instantiate୦୦Try୦int(func() (int, error) {
//line errors_test.go2:34
  return 1, nil
	}).Final(func(r int) {
		if r != 1 {
			t.Fatalf("result from try block is not as expected")
		}
	})
}
//line errors.go2:5
func instantiate୦୦Try୦int(e func() (

//line errors_test.go2:12
 int, error)) instantiate୦୦catch୦int {
//line errors.go2:6
 v, err := e()
			return instantiate୦୦catch୦int{err: err, result: v}
}

//line errors.go2:8
type instantiate୦୦catch୦int struct {
//line errors.go2:11
 err error
	result int
}

//line errors.go2:16
func (c instantiate୦୦catch୦int,) Catch(handler func(int,

//line errors.go2:16
 error) int,

//line errors.go2:16
) instantiate୦୦catch୦int {
	if c.err != nil {
		c.result = handler(c.result, c.err)
	}
	c.err = nil
	return c
}

//line errors.go2:25
func (c instantiate୦୦catch୦int,) Final(handler func(int,

//line errors.go2:25
)) {
	handler(c.result)
}

//line errors.go2:27
var _ = testing.AllocsPerRun
