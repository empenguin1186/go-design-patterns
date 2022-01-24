package adapter

import "fmt"

type Banner struct {
	text string
}

func NewBanner(text string) *Banner {
	return &Banner{text: text}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)", b.text)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*", b.text)
}

type Print interface {
	PrintWeak()
	PrintStrong()
}

type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(banner *Banner) *PrintBanner {
	return &PrintBanner{banner: banner}
}

func (p *PrintBanner) PrintWeak() {
	p.banner.ShowWithParen()
}

func (p *PrintBanner) PrintStrong() {
	p.banner.ShowWithAster()
}
