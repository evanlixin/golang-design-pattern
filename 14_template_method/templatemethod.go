package templatemethod

import "fmt"

type Downloader interface {
	Download(uri string)
}

type implement interface {
	download()
	save()
}

type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}


type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}

	template := newTemplate(downloader)
	downloader.template = template

	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}

// http://neoyeelf.github.io/2019/04/07/golang%E4%B9%9F%E8%83%BD%E5%AE%9E%E7%8E%B0%E6%8A%BD%E8%B1%A1%E7%B1%BB%E4%BA%86%EF%BC%9F/
// 抽象类
// 抽象方法
type IGame interface {
	Name() string
}

// 父类
type Game struct {}
// 调用”子类”的方法来获取名字。从而间接地实现了在公共方法中调用不同”子类”的实现的抽象方法。
func (g *Game) play(game IGame) {
	fmt.Printf(fmt.Sprintf("%s is awesome!", game.Name()))
}

type Dota struct {
	Game
}
func (d *Dota) Name() string {
	return "Dota"
}

type LOL struct {
	Game
}
func (l *LOL) Name() string {
	return "LOL"
}

