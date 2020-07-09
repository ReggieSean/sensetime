package main
import "testing"

func TestHello(t *testing.T) {

    matchTest := func (t *testing.T, got string, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }
    // t.Run("the test to match the name", func(t *testing.T){
    //     got := Hello("Chris")
    //     want := "Hello, Chris"
    //     matchTest(t,got, want)
    // })
    // t.Run("the test to check if it is a default name",func(t *testing.T) {
    //     got := Hello("")
    //     want := "Hello, Chris"
    //     matchTest(t,got, want)
    // })
    t.Run("Frech Greeting test", func(t *testing.T) {
        got := Hello("Joel", "french")
        want := "Bonjour, Joel"
        matchTest(t, got, want)
    })
    
    
}