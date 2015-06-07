/*
 * Jpeg compress function
 */

package jpegcompress

import (
	"bytes"
	"fmt"
	"util"
)

//comparison method only MPE mean pixel error method is implemented
type Comparison uint8

const (
	UNKNOWN Comparison = iota
	MPE
)

// to set quality of compressed image
type Quality uint8

const (
	UDEFINED Quality = iota
	VERYHIGH
	HIGH
	MEDIUM
	LOW
)

//parameters for compression
type Parameter struct {
	Preset    Quality
	Technique Comparison
	Jpegmax   int
	Jpegmin   int
	Attempts  int
}

var (
	method   Comparison
	preset   Quality
	attempts int
	jpegmin  int
	jpegmax  int
	target   float64
)

//to set parameters if provided,otherwise default parameters are used
func set_parameters(p *Parameter) {
	if method = p.Technique; method == 0 {
		method = MPE
	}

	if attempts = p.Attempts; attempts == 0 {
		attempts = 6
	}

	if preset = p.Preset; preset == 0 {
		preset = MEDIUM
	}

	if jpegmin = p.Jpegmin; jpegmin == 0 {
		jpegmin = 40
	}

	if jpegmax = p.Jpegmax; jpegmax == 0 {
		jpegmax = 95
	}
}

//to set target from quality parameter and  comparison technique
func set_target(m Comparison, p Quality) {

	switch m {
	case MPE:
		switch p {
		case LOW:
			target = 1.5
		case MEDIUM:
			target = 1
		case HIGH:
			target = 0.8
		case VERYHIGH:
			target = 0.6
		default:
			fmt.Print("problem in setting target\n")
		}
	case UNKNOWN:
		fmt.Print("unknown comparison technique\n")
	default:
		fmt.Print("default\n")
	}
}

//jpeg compress function which take source destn filename and parameters
func Jpeg_compress(source, dest string, param *Parameter) int {

	if param == nil {
		param = new(Parameter)
	}
	set_parameters(param)
	originalbuf := util.JpegReader(source)
	originalimg := util.JpegDecode(originalbuf)
	originalgray := util.Grayscale(originalimg)
	originalsize := originalbuf.Len()
	set_target(method, preset)
	min := jpegmin
	max := jpegmax
	var metric float64
	var comprsize int
	var comprbuf *bytes.Buffer
	for attempt := attempts - 1; attempt >= 0; attempt-- {
		quality := min + (max-min)/2
		comprbuf = util.JpegEncode(originalimg, quality)
		comprsize = comprbuf.Len()
		comprgray := util.JpegDecodeGray(comprbuf)

		if attempt == 0 {
			fmt.Println("final attemp")
		}

		switch method {
		case MPE:
			metric = util.MeanPixelError(originalgray, comprgray)
		default:
			fmt.Println("method not defined")
			return 0
		}
		fmt.Printf("quality %d metric %f\n", quality, metric)
		if metric < target {

			switch method {
			case MPE:
				max = quality - 1
			default:
				fmt.Println("method not defined")
				return 0
			}

		} else {
			switch method {
			case MPE:
				if comprsize > originalsize {
					fmt.Println("compressed size greater than original size abort")
					return 0
				}
				min = quality + 1
			default:
				fmt.Println("method not defined")
				return 0
			}

		}

	}

	util.JpegWriter(dest, comprbuf)
	return comprsize / 1024

}
