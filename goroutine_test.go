/**
go func() ==> manggil goroutine
time default nya ns, time.Second adlh durasi = 10^9 * 1 ns
func goroutine akan ttp jln sambil system utk exec code selanjutnya (proses jl secara async),
dan jika program slesai dl maka goroutine akan dikill smua
func goroutine blh return value, tp return value nya engga
bisa ditangkap (mungkin keburu lewat func yg nangkep / slesai programnya)
*/

package my_goroutine

import (
  "fmt"
  "testing"
  "time"
)

func RunHelloWorld() {
  fmt.Println("hello world")
}

func TestCreateGoroutine(t *testing.T) {
  go RunHelloWorld()    // func goroutine
  fmt.Println("ups")

  time.Sleep(time.Second)
}

func DisplayNumber(number int) {
  fmt.Println("dislay", number)
}

// test bahwa goroutine ringan, 1 goroutine = 2KB
func TestManyGoroutine(t *testing.T) {
  for i:=0; i<100000; i++ {
    go DisplayNumber(i)
  }

  time.Sleep(5 * time.Second)
}