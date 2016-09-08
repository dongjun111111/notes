package main

/*
原理:将数据的二进制形式写入图像红色通道数据二进制的低位
只支持png格式的输出
写入数据
go run t3.go -in="c.jpg" -data="需要隐藏的数据" -out="out.png"
读取数据
go run t3.go -in="out.png"
*/
import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

var FLAG = [4]byte{0x13, 0x14, 0x52, 0x00} //shadow flag.

//byte to 8 bits
func Byte2bits(b byte) (a [8]byte) {
	var c uint8 = 7
	var i uint8
	for i = 0; i < 8; i++ {
		a[i] = b >> (c - i) & 1
	}
	return
}

//8 bits to byte.
func Bits2Byte(a [8]byte) (b byte) {
	for i := 0; i < 8; i++ {
		b += a[i] * uint8(math.Pow(2, float64(7-i)))
	}
	return
}

//uint32 to 4 bytes.
func Uint32ToBytes(i uint32) (b [4]byte) {
	b[0] = uint8(i >> 24)
	b[1] = uint8(i >> 16 & 0xffff)
	b[2] = uint8(i >> 8 & 0xff)
	b[3] = uint8(i & 0xff)
	return
}

//4 bytes to uint32.
func Bytes2Uint32(b [4]byte) (i uint32) {
	var j uint32
	for ; j < 4; j++ {
		i += uint32(b[j]) << (24 - j*8)
	}
	return

}

func BuildShadowHeader(length uint32) (b [8]byte) {
	var i int
	for ; i < 4; i++ {
		b[i] = FLAG[i]
	}
	a := Uint32ToBytes(length)
	for ; i < 8; i++ {
		b[i] = a[i-4]
	}
	return

}
func WriteShadow(b []byte, im image.Image) (out image.Image, err error) {
	max := im.Bounds().Max.X*im.Bounds().Max.Y/8 - 64
	b_len := len(b)
	if len(b) > max {
		return nil, errors.New("image does not have enough space for shadow.")
	}
	head := BuildShadowHeader(uint32(b_len))
	var bb byte
	var bs [8]byte
	var i int
	out, err = SetImage(im, func(index, x, y int, in, out image.Image) {
		rgba := readRGBAColor(im.At(x, y))
		if index < b_len*8+64 {
			if index < 64 {
				bb = head[index/8]
			} else {
				bb = b[index/8-8]
			}
			bs = Byte2bits(bb)
			i = index % 8
			if bs[i] != rgba.R&1 {
				if bs[i] == 0 {
					rgba.R -= 1
				} else {
					rgba.R += 1
				}
			}
		}
		if v := out.(*image.RGBA); v != nil {
			v.SetRGBA(x, y, rgba)
		}

	})
	if err != nil {
		return nil, err
	}
	return
}
func ReadShadowData(im image.Image) (b []byte, err error) {
	head, err := ReadShadowHeader(im)
	if err != nil {
		return nil, err
	}
	length := int(ReadShadowLength(head))
	var bk []byte = make([]byte, length*8)
	b = make([]byte, length)
	_, err = SetImage(im, func(index, x, y int, in, out image.Image) {
		if index >= 64 && index < length*8+64 {
			R := readRGBAColor(im.At(x, y)).R
			bk[index-64] = uint8(R & 1)
		}
	})
	var bb [8]byte
	var bs []byte
	for i := 0; i < length; i++ {
		bs = bk[8*i : 8*(i+1)]
		for j := 0; j < 8; j++ {
			bb[j] = bs[j]
		}
		b[i] = Bits2Byte(bb)
	}
	return
}
func ReadShadowHeader(im image.Image) (b [8]byte, err error) {
	var bm [64]byte
	_, err = SetImage(im, func(index, x, y int, in, out image.Image) {
		rgba := readRGBAColor(im.At(x, y))
		if index < 64 {
			bm[index] = uint8(rgba.R & 1)
		}
	})
	if err != nil {
		return
	}
	var bb [8]byte
	var bs []byte
	for i := 0; i < 8; i++ {
		bs = bm[8*i : 8*(i+1)]
		for j := 0; j < 8; j++ {
			bb[j] = bs[j]
		}
		b[i] = Bits2Byte(bb)
	}
	return

}
func ReadShadowFlag(b [8]byte) (a [4]byte) {
	for i := 0; i < 4; i++ {
		a[i] = b[i]
	}
	return
}
func ReadShadowLength(b [8]byte) uint32 {
	var bb [4]byte
	for i := 4; i < 8; i++ {
		bb[i-4] = b[i]
	}
	return Bytes2Uint32(bb)
}

func OpenImage(path string) (image.Image, error) {
	im_read, err := os.Open(path)
	defer im_read.Close()
	if err != nil {
		return nil, err
	}
	im, _, err := image.Decode(im_read)
	if err != nil {
		return nil, err
	}
	return im, nil
}

//modify image
func SetImage(im image.Image, f func(index, x, y int, in, out image.Image)) (out image.Image, err error) {
	if f == nil {
		return im, nil
	}
	index := 0
	bounds := im.Bounds()
	out = image.NewRGBA(bounds)
	var m *image.RGBA = out.(*image.RGBA)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, im.At(x, y))
			f(index, x, y, im, out)
			index += 1
		}
	}
	return out, nil

}

//conert any color to RABGA color.
func readRGBAColor(from_color color.Color) color.RGBA {
	return color.RGBAModel.Convert(from_color).(color.RGBA)
}

//only write to jpeg formats.
func WriteImage(path string, im image.Image) error {
	out, err := os.OpenFile(path, os.O_CREATE, os.ModePerm)
	defer out.Close()
	if err != nil {
		return err
	}
	err = png.Encode(out, im)
	if err != nil {
		return err
	}
	return nil
}

var read_in string
var write_out string
var data string

func init() {
	flag.StringVar(&read_in, "in", "", "image path read in.")
	flag.StringVar(&write_out, "out", "out.jpg", "image path write out.")
	flag.StringVar(&data, "data", "", "data to shadow.")
}
func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	flag.Parse()
	if read_in == "" {
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}
	im, err := OpenImage(read_in)
	errHandle(err)
	if data != "" {
		out, err := WriteShadow([]byte(data), im)
		errHandle(err)
		err = WriteImage(write_out, out)
		errHandle(err)
	} else {
		head, err := ReadShadowHeader(im)
		errHandle(err)
		_flag := ReadShadowFlag(head)
		if _flag != FLAG {
			fmt.Println("image doesn't have shadow data.")
			return
		}
		data, err := ReadShadowData(im)
		errHandle(err)
		fmt.Println("shadow:", string(data))
	}
}
