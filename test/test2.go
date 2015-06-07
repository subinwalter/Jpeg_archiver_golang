/*
 * to test jpeg compress function by changing default parameters
 */

package main

import (
	"fmt"
	"jpegcompress"
)

func main() {
	param := jpegcompress.Parameter{Attempts: 10}
	size := jpegcompress.Jpeg_compress("source.jpg", "dest.jpg", &param)
	fmt.Println("size of compressed image %d\n", size)
}
