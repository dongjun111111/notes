package pdf

import (
	"errors"
	"github.com/signintech/gopdf"
	"log"
	"zcm_crm/utils"
)

type ZCMPdfConfig struct {
	PAGEWIDE      float64
	PAGEHEIGHT    float64
	SIDE          float64
	TOP           float64
	BOTTOM        float64
	TABLELINESIZE float64
	LINESIZE      float64
}

type ZCMPdf struct {
	gopdf.GoPdf
	Config *ZCMPdfConfig
}

const FONTNAME = "ZCM"
const FONTPATH = "./pdf/ttf/SIMYOU.TTF"
const STAMPPATH = "./pdf/ttf/stamp.jpg"
const DEFAULTFONTSIZE = 10 //中文字，高和宽等于这个值，英文为一半
// const PAGEWIDE = 595.28    //595.28, 841.89 = A4
// const PAGEHEIGHT = 841.89
// const SIDE = 50
// const TOP = 90
// const BOTTOM = 90

var (
	defaultBr          float64
	defaultLineHight   float64
	defaultLineFontNum int
)

func NewZCMPdf(config *ZCMPdfConfig) *ZCMPdf {
	if config == nil {
		config = &ZCMPdfConfig{PAGEWIDE: 595.28, PAGEHEIGHT: 841.89, SIDE: 50, TOP: 90, BOTTOM: 90, LINESIZE: 1, TABLELINESIZE: 0.5}
	}
	defaultLineHight = float64(float64(DEFAULTFONTSIZE) + float64(DEFAULTFONTSIZE)/5)
	defaultBr = float64(float64(DEFAULTFONTSIZE) + float64(DEFAULTFONTSIZE)/2)
	num := (config.PAGEWIDE - config.SIDE*2) / DEFAULTFONTSIZE
	defaultLineFontNum = int(num)
	pdf := &ZCMPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: config.PAGEWIDE, H: config.PAGEHEIGHT}})
	pdf.SetTopMargin(config.TOP)
	pdf.SetLeftMargin(config.SIDE)
	pdf.SetLineWidth(config.SIDE)
	pdf.Config = config
	//设置默认字体
	var err error
	err = pdf.AddTTFFont(FONTNAME, FONTPATH)
	HandleError("load font error", err)
	err = pdf.SetFont(FONTNAME, "", DEFAULTFONTSIZE)
	HandleError("set font error", err)
	//设置线
	pdf.SetLineWidth(1)

	return pdf
}

//----zcmpdf 函数-------
func (pdf *ZCMPdf) DoBR(height float64) *ZCMPdf {
	if pdf.GetY() > pdf.Config.PAGEHEIGHT-pdf.Config.BOTTOM-DEFAULTFONTSIZE {
		pdf.AddPage()
	}
	pdf.Br(height)
	return pdf
}

func (pdf *ZCMPdf) DoDefaultBR() *ZCMPdf {
	if pdf.GetY() > pdf.Config.PAGEHEIGHT-pdf.Config.BOTTOM-DEFAULTFONTSIZE {
		pdf.AddPage()
	}
	pdf.Br(defaultBr)
	return pdf
}

func (pdf *ZCMPdf) WriteDefaultLine() *ZCMPdf {
	x, y := pdf.GetX(), pdf.GetY()
	// pdf.SetLineWidth(pdf.Config.TABLELINESIZE)
	pdf.Line(x, y, pdf.Config.PAGEWIDE-pdf.Config.SIDE, y)
	// pdf.SetY(y + 5)
	// pdf.SetLineWidth(pdf.Config.LINESIZE)
	return pdf
}
func (pdf *ZCMPdf) WriteVerticalLine(height float64) *ZCMPdf {
	x, y := pdf.GetX(), pdf.GetY()
	// pdf.SetLineWidth(pdf.Config.TABLELINESIZE)
	pdf.Line(x, y, x, y+height)
	// pdf.SetLineWidth(pdf.Config.LINESIZE)
	return pdf
}
func (pdf *ZCMPdf) WriteRightVerticalLine(height float64) *ZCMPdf {
	y := pdf.GetY()
	// pdf.SetLineWidth(pdf.Config.TABLELINESIZE)
	pdf.Line(pdf.Config.PAGEWIDE-pdf.Config.SIDE, y, pdf.Config.PAGEWIDE-pdf.Config.SIDE, y+height)
	// pdf.SetLineWidth(pdf.Config.LINESIZE)
	return pdf
}
func (pdf *ZCMPdf) ResetX() *ZCMPdf {
	pdf.SetX(pdf.Config.SIDE)
	return pdf
}

func (pdf *ZCMPdf) AddY(size float64) *ZCMPdf {
	y := pdf.GetY()
	pdf.SetY(y + size)
	return pdf
}
func (pdf *ZCMPdf) AddX(size float64) *ZCMPdf {
	x := pdf.GetX()
	pdf.SetX(x + size)
	return pdf
}
func (pdf *ZCMPdf) WriteInCenter(text string, fontSize int) *ZCMPdf {
	//1.设置字体
	pdf.SetFontSize(fontSize)
	wide, _ := pdf.MeasureTextWidth(text)
	//2.居中定位,写
	// offset := (PAGEWIDE - pdf.GetX() - wide) / 2
	offset := (pdf.Config.PAGEWIDE - wide) / 2
	pdf.SetX(offset)
	pdf.Cell(nil, text)
	pdf.DoBR(utils.GetBR(fontSize))
	if fontSize != DEFAULTFONTSIZE {
		//3.字体设置回默认的
		pdf.SetDefaultFontSize()
	}
	return pdf
}

//右对齐
func (pdf *ZCMPdf) WriteInRight(text string) *ZCMPdf {

	wide, _ := pdf.MeasureTextWidth(text)
	//2.居中定位,写
	// offset := (PAGEWIDE - pdf.GetX() - wide) / 2
	offset := pdf.Config.PAGEWIDE - pdf.Config.SIDE - wide
	pdf.SetX(offset)
	pdf.Cell(nil, text)
	pdf.DoBR(utils.GetBR(DEFAULTFONTSIZE))
	return pdf
}

func (pdf *ZCMPdf) WriteAnyPlace(percent float64, text string, fontSize int) *ZCMPdf {
	if fontSize == DEFAULTFONTSIZE {
		pdf.SetX(pdf.Config.PAGEWIDE * percent)
		pdf.Cell(nil, text)
	} else {
		pdf.SetFontSize(fontSize)

		pdf.SetX(pdf.Config.PAGEWIDE * percent)
		pdf.Cell(nil, text)

		pdf.SetDefaultFontSize()
	}
	return pdf
}

func (pdf *ZCMPdf) WriteWithLine(text string) *ZCMPdf {
	x, y := pdf.GetX(), pdf.GetY()
	// log.Println("x,y:", x, y)
	pdf.Cell(nil, text)
	wide, _ := pdf.MeasureTextWidth(text)
	pdf.Line(x, y+defaultLineHight, x+wide, y+defaultLineHight)
	// log.Println("wide:", wide)
	// log.Println(x, y+defaultLineHight, x+wide, y+defaultLineHight)
	//划线后x轴值不是线的末尾，设置回去
	pdf.SetX(x + wide)
	return pdf
}

func (pdf *ZCMPdf) WritePassage(text string) *ZCMPdf {
	length := utils.Strlen(text)
	wideFlag := pdf.Config.PAGEWIDE - pdf.Config.SIDE*2
	var i int
	var offset int
	var flag bool
	for i = 0; i <= length-defaultLineFontNum; i += defaultLineFontNum {
		offset = 0
		flag = true
		for flag {
			if i+defaultLineFontNum+offset > length {
				flag = false
			}
			strline := utils.Substr(text, i, defaultLineFontNum+offset)
			wide, _ := pdf.MeasureTextWidth(strline)
			if wide < wideFlag {
				offset++
			} else {
				flag = false
			}
		}
		pdf.Write(utils.Substr(text, i, defaultLineFontNum+offset))
		pdf.DoDefaultBR()
		i += offset
	}
	if i < length {
		pdf.Write(utils.Substr(text, i, length-i))
	}
	return pdf
}

func (pdf *ZCMPdf) Write(text string) *ZCMPdf {
	pdf.Cell(nil, text)
	return pdf
}

func (pdf *ZCMPdf) ZAddPage() *ZCMPdf {
	pdf.AddPage()
	return pdf
	// pdf.ZImage(BACKGROUNDPATH, 35, 0)
	// log.Println(pdf)
	// pdf.Image(BACKGROUNDPATH, 35, 0, nil)
}
func (pdf *ZCMPdf) ZImage(image string, x float64, y float64) *ZCMPdf {
	pdf.Image(image, x, y, nil)
	return pdf
}
func (pdf *ZCMPdf) Out(filename string) *ZCMPdf {
	pdf.WritePdf(filename)
	return pdf
}

func (pdf *ZCMPdf) WriteWithColor(text string, r uint8, g uint8, b uint8) *ZCMPdf {
	pdf.SetTextColor(r, g, b)
	pdf.Cell(nil, text)
	//颜色设置回去
	pdf.SetTextColor(0, 0, 0)
	pdf.SetGrayFill(0)
	return pdf
}

func (pdf *ZCMPdf) SetFontSize(fontSize int) *ZCMPdf {
	err := pdf.SetFont(FONTNAME, "", fontSize)
	HandleError("set font error", err)
	return pdf
}

func (pdf *ZCMPdf) SetDefaultFontSize() *ZCMPdf {
	err := pdf.SetFont(FONTNAME, "", DEFAULTFONTSIZE)
	HandleError("set font error", err)
	return pdf
}

func HandleError(prefix string, err error) bool {
	if err != nil {
		log.Fatalln(err.Error())
		panic(errors.New(prefix))
		return true
	} else {
		return false
	}
}

func (pdf *ZCMPdf) DrawCommonForm(x []float64, y []float64, value [][]string) *ZCMPdf {
	lenX := len(x)
	lenY := len(y)
	//画横线
	for i := 0; i < lenY; i++ {
		pdf.Line(x[0], y[i], x[lenX-1], y[i])
	}
	//画竖线
	for i := 0; i < lenX; i++ {
		pdf.Line(x[i], y[0], x[i], y[lenY-1])
	}
	//填值，暂不支持居中
	for i := 0; i < lenY-1; i++ {
		for j := 0; j < lenX-1; j++ {
			pdf.SetX(x[j] + 5)
			pdf.SetY(y[i] + 5)
			pdf.Cell(nil, value[i][j])
		}

	}
	return pdf
}
