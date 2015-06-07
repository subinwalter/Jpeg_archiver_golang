/*
 * to test jpeg compress function with default parameters
 */
package main

import (
	"jpegcompress"
)

func main() {
	jpegcompress.Jpeg_compress("source.jpg", "dest.jpg", nil)
}
