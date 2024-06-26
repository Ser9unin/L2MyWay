package recursiveParser

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Ser9unin/L2MyWay/develop/dev09/pkg/linkParser"
)

// Site represents data needed to build sitemap
type Site struct {
	domain       string
	visitedLinks map[string]struct{}
	directory    string
}

// NewSite creates instance of Sitemap
func NewSite(link string, directory string) Site {
	v := map[string]struct{}{
		link: {},
	}

	site := Site{
		domain:       link,
		visitedLinks: v,
		directory:    directory,
	}

	return site
}

// DownloadSite recursively visits links in queue and downloads them.
// If depth is greater than zero it limit number of recursive calls.
func (s *Site) DownloadSite(queue []string, level int) error {
	if level == 0 {
		return nil
	}

	linksToParse := make([]string, level)

	for _, v := range queue {
		resp, err := http.Get(v)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		mediatype, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			fmt.Printf("can't parse link type: %s", err.Error())
		}
		ext, err := mime.ExtensionsByType(mediatype)
		if err != nil || len(ext) == 0 {
			fmt.Printf("can't parse link type: %s", err.Error())
			ext = append(ext, "")
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		r := bytes.NewReader(body)
		fileName := path.Join(s.directory, path.Base(resp.Request.URL.Path)+ext[0])
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		size, err := io.Copy(file, r)
		if err != nil {
			return err
		}

		fmt.Printf("\nDownloaded a file %s with size %d bytes\n", fileName, size)

		if _, err = r.Seek(0, 0); err != nil {
			return err
		}

		l, err := s.parseLinks(r)
		if err != nil {
			return err
		}
		linksToParse = append(linksToParse, l...)
	}

	if len(linksToParse) > 0 {
		level--
		return s.DownloadSite(linksToParse, level)
	}

	return nil
}

// parseLinks reads html data from io.Reader and creates slice of links.
// It only parses links with Sitemap.domain
func (s *Site) parseLinks(r io.Reader) ([]string, error) {
	res, err := linkParser.ParseHTML(r)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)

	for _, v := range res {
		href := v.Href.Host + v.Href.Path
		visited := true

		switch {
		case strings.HasPrefix(href, s.domain):
			visited = s.isVisited(href)
		case strings.HasPrefix(href, "/"):
			href = s.domain + href
			visited = s.isVisited(href)
		}

		if !visited {
			links = append(links, href)
		}
	}

	return links, nil
}

// isVisited checks if url visited
func (s *Site) isVisited(href string) (visited bool) {
	if _, visited = s.visitedLinks[href]; !visited {
		s.visitedLinks[href] = struct{}{}
	}

	return visited
}
