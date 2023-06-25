/**
channel utk ngirim data antar goroutine
channel hrs dideklarasiin tipe datanya dan cmn bs 1 aj
channel <- data ==> untuk mengirim data
data (var) <- channel ==> untuk menerima data
saat selesai pakai func close(channel)
*/

package my_goroutine

import "testing"

func TestCreateChannel(t *testing.T) {
  channel:= make(chan string)
  
  channel <- "eko"
  
  //var data string
  //data = <- channel
  
  // langsung kirim ke param suatu func
  //fmt.Println(<- channel)

  data := <- channel
  close(channel)
}