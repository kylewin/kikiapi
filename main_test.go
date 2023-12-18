package main
import (
    "fmt"
    "testing"
)
func TestMain(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
t.Error* will report test failures but continue executing the test. t.Fatal* will report test failures and stop the test immediately.

        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}
