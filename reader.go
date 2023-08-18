package doclize

import (
	"archive/zip"
	"bytes"
	tools "github.com/iEvan-lhr/exciting-tool"
	"io"
	"io/fs"
)

type Docx struct {
	files   []*zip.File
	content string
	links   string
	mapsByt *t
	headers map[string]string
	footers map[string]string
	images  map[string]string
}

// ReadDocxFromFS opens a docx file from the file system
func ReadDocxFromFS(file string, fs fs.FS) (*ReplaceDocx, error) {
	f, err := fs.Open(file)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer([]byte{})
	size, err := io.Copy(buff, f)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(buff.Bytes())
	return ReadDocxFromMemory(reader, size)
}

func ReadDocxFromMemory(data io.ReaderAt, size int64) (*ReplaceDocx, error) {
	reader, err := zip.NewReader(data, size)
	if err != nil {
		return nil, err
	}
	zipData := ZipInMemory{data: reader}
	return ReadDocx(zipData)
}

func ReadDocxFile(path string) (*ReplaceDocx, error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	zipData := ZipFile{data: reader}
	return ReadDocx(zipData)
}

func (r *ReplaceDocx) Editable() *Docx {

	return &Docx{
		files:   r.zipReader.files(),
		content: r.content,
		links:   r.links,
		headers: r.headers,
		mapsByt: getMapByt(r.content),
		footers: r.footers,
		images:  r.images,
	}
}

func getMapByt(content string) (m *t) {
	contents, other, steps := tools.Make(content).GetContentAll("<w:t>", "</w:t>")
	m = &t{
		content: contents,
		other:   other,
		steps: make(map[int]struct {
			model int
			index int
		}),
	}
	for i, s := range steps {
		m.steps[i] = struct {
			model int
			index int
		}{model: s.Model, index: s.Index}
	}
	return
}

func ReadDocx(reader ZipData) (*ReplaceDocx, error) {
	content, err := readText(reader.files())
	if err != nil {
		return nil, err
	}

	links, err := readLinks(reader.files())
	if err != nil {
		return nil, err
	}

	headers, footers, _ := readHeaderFooter(reader.files())
	images, _ := retrieveImageFilenames(reader.files())
	return &ReplaceDocx{zipReader: reader, content: content, links: links, headers: headers, footers: footers, images: images}, nil
}
