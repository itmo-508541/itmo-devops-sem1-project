package reader

import "errors"

type MultiArchive struct {
	ZipArchive
	TarArchive
}

func NewMultiArchive() *MultiArchive {
	return &MultiArchive{}
}

func (a *MultiArchive) Contents(b []byte) ([]byte, error) {
	var err error
	var zipContents, tarContents []byte

	zipContents, err = a.ZipArchive.Contents(b)
	if err == nil {
		return zipContents, nil
	}

	tarContents, err = a.TarArchive.Contents(b)
	if err == nil {
		return tarContents, nil
	}

	return nil, errors.New("MultyAchive: not archive")
}
