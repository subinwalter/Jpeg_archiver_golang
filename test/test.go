/*
 * To test the util package function
 */
package main

import (
	"fmt"
	"util"
)

func main() {
	sbuf := util.JpegReader("source.jpg")
	simg := util.JpegDecode(sbuf)
	sgray := util.Grayscale(simg)
	dbuf := util.JpegEncode(simg, 85)
	dgray := util.JpegDecodeGray(dbuf)
	sgraybuf := util.JpegEncode(sgray, 100)
	dgraybuf := util.JpegEncode(dgray, 100)
	mpe := util.MeanPixelError(dgray, sgray)
	fmt.Println(mpe)
	fmt.Printf("%d %d %d %d", sbuf.Len()/1024, sgraybuf.Len()/1024, dbuf.Len()/1024, dgraybuf.Len()/1024)
	util.JpegWriter("dest.jpg", dbuf)
	util.JpegWriter("dest_gray.jpg", dgraybuf)
	util.JpegWriter("source_gray.jpg", sgraybuf)

}
