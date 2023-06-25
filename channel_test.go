/**
channel utk ngirim data antar goroutine
channel hrs dideklarasiin tipe datanya dan cmn bs 1 aj

*/

package my_goroutine

import "testing"

func TestCreateChannel(t *testing.T) {
  channel:= make(chan string)
  
  channel <- "eko"
}