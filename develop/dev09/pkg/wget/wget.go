package mywget

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/Ser9unin/L2MyWay/develop/dev09/pkg/recursiveParser"
	"github.com/schollz/progressbar/v3"
)

type Mywget struct {
	url        *url.URL
	filePath   string
	level      int
	recursive  bool
	requisites bool
}

func (cfg *Mywget) ParseConfig(args []string) error {
	flags := flag.NewFlagSet("mywget", flag.ContinueOnError)
	flags.StringVar(&cfg.filePath, "O", "", "Set path to output file")
	flags.IntVar(&cfg.level, "l", 0, "Determine how many levels of links follow while download site.")
	flags.BoolVar(&cfg.recursive, "r", false, "Recursive work of wget")
	flags.BoolVar(&cfg.requisites, "p", false, "Download all resources necessary to properly display a given HTML page")

	err := flags.Parse(args)
	if err != nil {
		return err
	}

	cfg.url, err = url.Parse(flags.Arg(0))

	if err != nil {
		return err
	}

	if cfg.filePath == "" {
		cfg.filePath = path.Base(cfg.url.Path)
	}

	return nil
}

func (cfg *Mywget) Run() error {
	if cfg.recursive {
		queue := []string{cfg.url.String()}
		if err := os.Mkdir(cfg.url.Host, os.ModePerm); err != nil {
			return err
		}
		siteMap := recursiveParser.NewSite(cfg.url.String(), cfg.url.Host)
		err := siteMap.DownloadSite(queue, cfg.level)
		if err != nil {
			return err
		}
		return nil
	}

	err := downloadPage(cfg.url.String(), cfg.filePath)
	if err != nil {
		return err
	}
	return nil
}

func downloadPage(url, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	size, err := io.Copy(io.MultiWriter(file, bar), resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("\nDownloaded file %s with size %d bytes\n", filepath, size)

	return nil
}
