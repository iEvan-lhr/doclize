package doclize

import (
	"archive/zip"
)

type Docx struct {
	files   []*zip.File
	content string
	links   string
	headers map[string]string
	footers map[string]string
	images  map[string]string
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
		footers: r.footers,
		images:  r.images,
	}
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
