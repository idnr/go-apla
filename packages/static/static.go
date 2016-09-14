// Copyright 2016 The go-daylight Authors
// This file is part of the go-daylight library.
//
// The go-daylight library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-daylight library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-daylight library. If not, see <http://www.gnu.org/licenses/>.

package static

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// static1block reads file data from disk. It returns an error on failure.
func static1block() (*asset, error) {
	path := "static/1block"
	name := "static/1block"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static1blockLocal reads file data from disk. It returns an error on failure.
func static1blockLocal() (*asset, error) {
	path := "static/1block-local"
	name := "static/1block-local"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticAlert_successHtml reads file data from disk. It returns an error on failure.
func staticAlert_successHtml() (*asset, error) {
	path := "static/alert_success.html"
	name := "static/alert_success.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticAnonym_money_transferHtml reads file data from disk. It returns an error on failure.
func staticAnonym_money_transferHtml() (*asset, error) {
	path := "static/anonym_money_transfer.html"
	name := "static/anonym_money_transfer.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticBackupHtml reads file data from disk. It returns an error on failure.
func staticBackupHtml() (*asset, error) {
	path := "static/backup.html"
	name := "static/backup.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticBlock_explorerHtml reads file data from disk. It returns an error on failure.
func staticBlock_explorerHtml() (*asset, error) {
	path := "static/block_explorer.html"
	name := "static/block_explorer.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticBlock_generationHtml reads file data from disk. It returns an error on failure.
func staticBlock_generationHtml() (*asset, error) {
	path := "static/block_generation.html"
	name := "static/block_generation.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticChange_node_keyHtml reads file data from disk. It returns an error on failure.
func staticChange_node_keyHtml() (*asset, error) {
	path := "static/change_node_key.html"
	name := "static/change_node_key.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticCssStyleCss reads file data from disk. It returns an error on failure.
func staticCssStyleCss() (*asset, error) {
	path := "static/css/style.css"
	name := "static/css/style.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticDashboard_anonymHtml reads file data from disk. It returns an error on failure.
func staticDashboard_anonymHtml() (*asset, error) {
	path := "static/dashboard_anonym.html"
	name := "static/dashboard_anonym.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsFiletypesRegularEot reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsFiletypesRegularEot() (*asset, error) {
	path := "static/fonts/glyphicons-filetypes-regular.eot"
	name := "static/fonts/glyphicons-filetypes-regular.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsFiletypesRegularSvg reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsFiletypesRegularSvg() (*asset, error) {
	path := "static/fonts/glyphicons-filetypes-regular.svg"
	name := "static/fonts/glyphicons-filetypes-regular.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsFiletypesRegularTtf reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsFiletypesRegularTtf() (*asset, error) {
	path := "static/fonts/glyphicons-filetypes-regular.ttf"
	name := "static/fonts/glyphicons-filetypes-regular.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsFiletypesRegularWoff reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsFiletypesRegularWoff() (*asset, error) {
	path := "static/fonts/glyphicons-filetypes-regular.woff"
	name := "static/fonts/glyphicons-filetypes-regular.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsFiletypesRegularWoff2 reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsFiletypesRegularWoff2() (*asset, error) {
	path := "static/fonts/glyphicons-filetypes-regular.woff2"
	name := "static/fonts/glyphicons-filetypes-regular.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsHalflingsRegularEot reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsHalflingsRegularEot() (*asset, error) {
	path := "static/fonts/glyphicons-halflings-regular.eot"
	name := "static/fonts/glyphicons-halflings-regular.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsHalflingsRegularSvg reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsHalflingsRegularSvg() (*asset, error) {
	path := "static/fonts/glyphicons-halflings-regular.svg"
	name := "static/fonts/glyphicons-halflings-regular.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsHalflingsRegularTtf reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsHalflingsRegularTtf() (*asset, error) {
	path := "static/fonts/glyphicons-halflings-regular.ttf"
	name := "static/fonts/glyphicons-halflings-regular.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsHalflingsRegularWoff reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsHalflingsRegularWoff() (*asset, error) {
	path := "static/fonts/glyphicons-halflings-regular.woff"
	name := "static/fonts/glyphicons-halflings-regular.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsHalflingsRegularWoff2 reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsHalflingsRegularWoff2() (*asset, error) {
	path := "static/fonts/glyphicons-halflings-regular.woff2"
	name := "static/fonts/glyphicons-halflings-regular.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsRegularEot reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsRegularEot() (*asset, error) {
	path := "static/fonts/glyphicons-regular.eot"
	name := "static/fonts/glyphicons-regular.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsRegularSvg reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsRegularSvg() (*asset, error) {
	path := "static/fonts/glyphicons-regular.svg"
	name := "static/fonts/glyphicons-regular.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsRegularTtf reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsRegularTtf() (*asset, error) {
	path := "static/fonts/glyphicons-regular.ttf"
	name := "static/fonts/glyphicons-regular.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsRegularWoff reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsRegularWoff() (*asset, error) {
	path := "static/fonts/glyphicons-regular.woff"
	name := "static/fonts/glyphicons-regular.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsRegularWoff2 reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsRegularWoff2() (*asset, error) {
	path := "static/fonts/glyphicons-regular.woff2"
	name := "static/fonts/glyphicons-regular.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsSocialRegularEot reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsSocialRegularEot() (*asset, error) {
	path := "static/fonts/glyphicons-social-regular.eot"
	name := "static/fonts/glyphicons-social-regular.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsSocialRegularSvg reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsSocialRegularSvg() (*asset, error) {
	path := "static/fonts/glyphicons-social-regular.svg"
	name := "static/fonts/glyphicons-social-regular.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsSocialRegularTtf reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsSocialRegularTtf() (*asset, error) {
	path := "static/fonts/glyphicons-social-regular.ttf"
	name := "static/fonts/glyphicons-social-regular.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsSocialRegularWoff reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsSocialRegularWoff() (*asset, error) {
	path := "static/fonts/glyphicons-social-regular.woff"
	name := "static/fonts/glyphicons-social-regular.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticFontsGlyphiconsSocialRegularWoff2 reads file data from disk. It returns an error on failure.
func staticFontsGlyphiconsSocialRegularWoff2() (*asset, error) {
	path := "static/fonts/glyphicons-social-regular.woff2"
	name := "static/fonts/glyphicons-social-regular.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticHistoryHtml reads file data from disk. It returns an error on failure.
func staticHistoryHtml() (*asset, error) {
	path := "static/history.html"
	name := "static/history.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgD_nullPng reads file data from disk. It returns an error on failure.
func staticImgD_nullPng() (*asset, error) {
	path := "static/img/D_null.png"
	name := "static/img/D_null.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgD_okPng reads file data from disk. It returns an error on failure.
func staticImgD_okPng() (*asset, error) {
	path := "static/img/D_ok.png"
	name := "static/img/D_ok.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUsSvg reads file data from disk. It returns an error on failure.
func staticImgUsSvg() (*asset, error) {
	path := "static/img/US.svg"
	name := "static/img/US.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg1Jpg reads file data from disk. It returns an error on failure.
func staticImgBg1Jpg() (*asset, error) {
	path := "static/img/bg1.jpg"
	name := "static/img/bg1.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg10Jpg reads file data from disk. It returns an error on failure.
func staticImgBg10Jpg() (*asset, error) {
	path := "static/img/bg10.jpg"
	name := "static/img/bg10.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg2Jpg reads file data from disk. It returns an error on failure.
func staticImgBg2Jpg() (*asset, error) {
	path := "static/img/bg2.jpg"
	name := "static/img/bg2.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg3Jpg reads file data from disk. It returns an error on failure.
func staticImgBg3Jpg() (*asset, error) {
	path := "static/img/bg3.jpg"
	name := "static/img/bg3.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg4Jpg reads file data from disk. It returns an error on failure.
func staticImgBg4Jpg() (*asset, error) {
	path := "static/img/bg4.jpg"
	name := "static/img/bg4.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg5Jpg reads file data from disk. It returns an error on failure.
func staticImgBg5Jpg() (*asset, error) {
	path := "static/img/bg5.jpg"
	name := "static/img/bg5.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg6Jpg reads file data from disk. It returns an error on failure.
func staticImgBg6Jpg() (*asset, error) {
	path := "static/img/bg6.jpg"
	name := "static/img/bg6.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg7Jpg reads file data from disk. It returns an error on failure.
func staticImgBg7Jpg() (*asset, error) {
	path := "static/img/bg7.jpg"
	name := "static/img/bg7.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg8Jpg reads file data from disk. It returns an error on failure.
func staticImgBg8Jpg() (*asset, error) {
	path := "static/img/bg8.jpg"
	name := "static/img/bg8.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgBg9Jpg reads file data from disk. It returns an error on failure.
func staticImgBg9Jpg() (*asset, error) {
	path := "static/img/bg9.jpg"
	name := "static/img/bg9.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgDummyPng reads file data from disk. It returns an error on failure.
func staticImgDummyPng() (*asset, error) {
	path := "static/img/dummy.png"
	name := "static/img/dummy.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgFaviconIco reads file data from disk. It returns an error on failure.
func staticImgFaviconIco() (*asset, error) {
	path := "static/img/favicon.ico"
	name := "static/img/favicon.ico"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgLockBgJpg reads file data from disk. It returns an error on failure.
func staticImgLockBgJpg() (*asset, error) {
	path := "static/img/lock-bg.jpg"
	name := "static/img/lock-bg.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgLogoSingleSvg reads file data from disk. It returns an error on failure.
func staticImgLogoSingleSvg() (*asset, error) {
	path := "static/img/logo-single.svg"
	name := "static/img/logo-single.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgLogoSvg reads file data from disk. It returns an error on failure.
func staticImgLogoSvg() (*asset, error) {
	path := "static/img/logo.svg"
	name := "static/img/logo.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgMbSampleJpg reads file data from disk. It returns an error on failure.
func staticImgMbSampleJpg() (*asset, error) {
	path := "static/img/mb-sample.jpg"
	name := "static/img/mb-sample.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgMockupPng reads file data from disk. It returns an error on failure.
func staticImgMockupPng() (*asset, error) {
	path := "static/img/mockup.png"
	name := "static/img/mockup.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgProfileBgJpg reads file data from disk. It returns an error on failure.
func staticImgProfileBgJpg() (*asset, error) {
	path := "static/img/profile-bg.jpg"
	name := "static/img/profile-bg.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgRadio_iconPng reads file data from disk. It returns an error on failure.
func staticImgRadio_iconPng() (*asset, error) {
	path := "static/img/radio_icon.png"
	name := "static/img/radio_icon.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgSquaresGif reads file data from disk. It returns an error on failure.
func staticImgSquaresGif() (*asset, error) {
	path := "static/img/squares.gif"
	name := "static/img/squares.gif"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser01Jpg reads file data from disk. It returns an error on failure.
func staticImgUser01Jpg() (*asset, error) {
	path := "static/img/user/01.jpg"
	name := "static/img/user/01.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser02Jpg reads file data from disk. It returns an error on failure.
func staticImgUser02Jpg() (*asset, error) {
	path := "static/img/user/02.jpg"
	name := "static/img/user/02.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser03Jpg reads file data from disk. It returns an error on failure.
func staticImgUser03Jpg() (*asset, error) {
	path := "static/img/user/03.jpg"
	name := "static/img/user/03.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser04Jpg reads file data from disk. It returns an error on failure.
func staticImgUser04Jpg() (*asset, error) {
	path := "static/img/user/04.jpg"
	name := "static/img/user/04.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser05Jpg reads file data from disk. It returns an error on failure.
func staticImgUser05Jpg() (*asset, error) {
	path := "static/img/user/05.jpg"
	name := "static/img/user/05.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser06Jpg reads file data from disk. It returns an error on failure.
func staticImgUser06Jpg() (*asset, error) {
	path := "static/img/user/06.jpg"
	name := "static/img/user/06.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser07Jpg reads file data from disk. It returns an error on failure.
func staticImgUser07Jpg() (*asset, error) {
	path := "static/img/user/07.jpg"
	name := "static/img/user/07.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser08Jpg reads file data from disk. It returns an error on failure.
func staticImgUser08Jpg() (*asset, error) {
	path := "static/img/user/08.jpg"
	name := "static/img/user/08.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser09Jpg reads file data from disk. It returns an error on failure.
func staticImgUser09Jpg() (*asset, error) {
	path := "static/img/user/09.jpg"
	name := "static/img/user/09.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser10Jpg reads file data from disk. It returns an error on failure.
func staticImgUser10Jpg() (*asset, error) {
	path := "static/img/user/10.jpg"
	name := "static/img/user/10.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser11Jpg reads file data from disk. It returns an error on failure.
func staticImgUser11Jpg() (*asset, error) {
	path := "static/img/user/11.jpg"
	name := "static/img/user/11.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser12Jpg reads file data from disk. It returns an error on failure.
func staticImgUser12Jpg() (*asset, error) {
	path := "static/img/user/12.jpg"
	name := "static/img/user/12.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgUser13Jpg reads file data from disk. It returns an error on failure.
func staticImgUser13Jpg() (*asset, error) {
	path := "static/img/user/13.jpg"
	name := "static/img/user/13.jpg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticIndexHtml reads file data from disk. It returns an error on failure.
func staticIndexHtml() (*asset, error) {
	path := "static/index.html"
	name := "static/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticInstall_step_0Html reads file data from disk. It returns an error on failure.
func staticInstall_step_0Html() (*asset, error) {
	path := "static/install_step_0.html"
	name := "static/install_step_0.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticInstall_step_1Html reads file data from disk. It returns an error on failure.
func staticInstall_step_1Html() (*asset, error) {
	path := "static/install_step_1.html"
	name := "static/install_step_1.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsAppJs reads file data from disk. It returns an error on failure.
func staticJsAppJs() (*asset, error) {
	path := "static/js/app.js"
	name := "static/js/app.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsClipboardJs reads file data from disk. It returns an error on failure.
func staticJsClipboardJs() (*asset, error) {
	path := "static/js/clipboard.js"
	name := "static/js/clipboard.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoAsn110MinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoAsn110MinJs() (*asset, error) {
	path := "static/js/crypto/asn1-1.0.min.js"
	name := "static/js/crypto/asn1-1.0.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoAsn1hex11MinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoAsn1hex11MinJs() (*asset, error) {
	path := "static/js/crypto/asn1hex-1.1.min.js"
	name := "static/js/crypto/asn1hex-1.1.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoCrypto11MinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoCrypto11MinJs() (*asset, error) {
	path := "static/js/crypto/crypto-1.1.min.js"
	name := "static/js/crypto/crypto-1.1.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoEcdsaModified10MinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoEcdsaModified10MinJs() (*asset, error) {
	path := "static/js/crypto/ecdsa-modified-1.0.min.js"
	name := "static/js/crypto/ecdsa-modified-1.0.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoEcparam10MinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoEcparam10MinJs() (*asset, error) {
	path := "static/js/crypto/ecparam-1.0.min.js"
	name := "static/js/crypto/ecparam-1.0.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtCjAesJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtCjAesJs() (*asset, error) {
	path := "static/js/crypto/ext/cj/aes.js"
	name := "static/js/crypto/ext/cj/aes.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtCjCryptojs312CoreFixMinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtCjCryptojs312CoreFixMinJs() (*asset, error) {
	path := "static/js/crypto/ext/cj/cryptojs-312-core-fix-min.js"
	name := "static/js/crypto/ext/cj/cryptojs-312-core-fix-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtCjEcPatchMinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtCjEcPatchMinJs() (*asset, error) {
	path := "static/js/crypto/ext/cj/ec-patch-min.js"
	name := "static/js/crypto/ext/cj/ec-patch-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtCjMd5_minJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtCjMd5_minJs() (*asset, error) {
	path := "static/js/crypto/ext/cj/md5_min.js"
	name := "static/js/crypto/ext/cj/md5_min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtCjSha1_minJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtCjSha1_minJs() (*asset, error) {
	path := "static/js/crypto/ext/cj/sha1_min.js"
	name := "static/js/crypto/ext/cj/sha1_min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtCjSha256_minJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtCjSha256_minJs() (*asset, error) {
	path := "static/js/crypto/ext/cj/sha256_min.js"
	name := "static/js/crypto/ext/cj/sha256_min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtCjX64Core_minJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtCjX64Core_minJs() (*asset, error) {
	path := "static/js/crypto/ext/cj/x64-core_min.js"
	name := "static/js/crypto/ext/cj/x64-core_min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtEcMinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtEcMinJs() (*asset, error) {
	path := "static/js/crypto/ext/ec-min.js"
	name := "static/js/crypto/ext/ec-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtEcPatchMinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtEcPatchMinJs() (*asset, error) {
	path := "static/js/crypto/ext/ec-patch-min.js"
	name := "static/js/crypto/ext/ec-patch-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtJsbnMinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtJsbnMinJs() (*asset, error) {
	path := "static/js/crypto/ext/jsbn-min.js"
	name := "static/js/crypto/ext/jsbn-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtJsbn2MinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtJsbn2MinJs() (*asset, error) {
	path := "static/js/crypto/ext/jsbn2-min.js"
	name := "static/js/crypto/ext/jsbn2-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtPrng4MinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtPrng4MinJs() (*asset, error) {
	path := "static/js/crypto/ext/prng4-min.js"
	name := "static/js/crypto/ext/prng4-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtRngMinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtRngMinJs() (*asset, error) {
	path := "static/js/crypto/ext/rng-min.js"
	name := "static/js/crypto/ext/rng-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCryptoExtYahooMinJs reads file data from disk. It returns an error on failure.
func staticJsCryptoExtYahooMinJs() (*asset, error) {
	path := "static/js/crypto/ext/yahoo-min.js"
	name := "static/js/crypto/ext/yahoo-min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDatetimeJs reads file data from disk. It returns an error on failure.
func staticJsDatetimeJs() (*asset, error) {
	path := "static/js/datetime.js"
	name := "static/js/datetime.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoDatatableJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoDatatableJs() (*asset, error) {
	path := "static/js/demo/demo-datatable.js"
	name := "static/js/demo/demo-datatable.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoFlotJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoFlotJs() (*asset, error) {
	path := "static/js/demo/demo-flot.js"
	name := "static/js/demo/demo-flot.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoFormsJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoFormsJs() (*asset, error) {
	path := "static/js/demo/demo-forms.js"
	name := "static/js/demo/demo-forms.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoJqcloudJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoJqcloudJs() (*asset, error) {
	path := "static/js/demo/demo-jqcloud.js"
	name := "static/js/demo/demo-jqcloud.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoJqgridJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoJqgridJs() (*asset, error) {
	path := "static/js/demo/demo-jqgrid.js"
	name := "static/js/demo/demo-jqgrid.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoNestableJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoNestableJs() (*asset, error) {
	path := "static/js/demo/demo-nestable.js"
	name := "static/js/demo/demo-nestable.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoPanelsJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoPanelsJs() (*asset, error) {
	path := "static/js/demo/demo-panels.js"
	name := "static/js/demo/demo-panels.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoRtlJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoRtlJs() (*asset, error) {
	path := "static/js/demo/demo-rtl.js"
	name := "static/js/demo/demo-rtl.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoSearchJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoSearchJs() (*asset, error) {
	path := "static/js/demo/demo-search.js"
	name := "static/js/demo/demo-search.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoSortableJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoSortableJs() (*asset, error) {
	path := "static/js/demo/demo-sortable.js"
	name := "static/js/demo/demo-sortable.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoUploadJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoUploadJs() (*asset, error) {
	path := "static/js/demo/demo-upload.js"
	name := "static/js/demo/demo-upload.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoVectorMapJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoVectorMapJs() (*asset, error) {
	path := "static/js/demo/demo-vector-map.js"
	name := "static/js/demo/demo-vector-map.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoWizardJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoWizardJs() (*asset, error) {
	path := "static/js/demo/demo-wizard.js"
	name := "static/js/demo/demo-wizard.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDemoDemoXeditableJs reads file data from disk. It returns an error on failure.
func staticJsDemoDemoXeditableJs() (*asset, error) {
	path := "static/js/demo/demo-xeditable.js"
	name := "static/js/demo/demo-xeditable.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsIndexJs reads file data from disk. It returns an error on failure.
func staticJsIndexJs() (*asset, error) {
	path := "static/js/index.js"
	name := "static/js/index.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsNprogressJs reads file data from disk. It returns an error on failure.
func staticJsNprogressJs() (*asset, error) {
	path := "static/js/nprogress.js"
	name := "static/js/nprogress.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticLang1Ini reads file data from disk. It returns an error on failure.
func staticLang1Ini() (*asset, error) {
	path := "static/lang/1.ini"
	name := "static/lang/1.ini"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticLang42Ini reads file data from disk. It returns an error on failure.
func staticLang42Ini() (*asset, error) {
	path := "static/lang/42.ini"
	name := "static/lang/42.ini"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticLangEnUsAllJson reads file data from disk. It returns an error on failure.
func staticLangEnUsAllJson() (*asset, error) {
	path := "static/lang/en-us.all.json"
	name := "static/lang/en-us.all.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticLangLocale_enUsIni reads file data from disk. It returns an error on failure.
func staticLangLocale_enUsIni() (*asset, error) {
	path := "static/lang/locale_en-US.ini"
	name := "static/lang/locale_en-US.ini"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticLangLocale_ruRuIni reads file data from disk. It returns an error on failure.
func staticLangLocale_ruRuIni() (*asset, error) {
	path := "static/lang/locale_ru-RU.ini"
	name := "static/lang/locale_ru-RU.ini"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticLoginHtml reads file data from disk. It returns an error on failure.
func staticLoginHtml() (*asset, error) {
	path := "static/login.html"
	name := "static/login.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticMenuHtml reads file data from disk. It returns an error on failure.
func staticMenuHtml() (*asset, error) {
	path := "static/menu.html"
	name := "static/menu.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticModal_anonymHtml reads file data from disk. It returns an error on failure.
func staticModal_anonymHtml() (*asset, error) {
	path := "static/modal_anonym.html"
	name := "static/modal_anonym.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPassHtml reads file data from disk. It returns an error on failure.
func staticPassHtml() (*asset, error) {
	path := "static/pass.html"
	name := "static/pass.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPswHtml reads file data from disk. It returns an error on failure.
func staticPswHtml() (*asset, error) {
	path := "static/psw.html"
	name := "static/psw.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticRequest_citizen_statusHtml reads file data from disk. It returns an error on failure.
func staticRequest_citizen_statusHtml() (*asset, error) {
	path := "static/request_citizen_status.html"
	name := "static/request_citizen_status.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticSignaturesHtml reads file data from disk. It returns an error on failure.
func staticSignaturesHtml() (*asset, error) {
	path := "static/signatures.html"
	name := "static/signatures.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticSignatures_newHtml reads file data from disk. It returns an error on failure.
func staticSignatures_newHtml() (*asset, error) {
	path := "static/signatures_new.html"
	name := "static/signatures_new.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTestHtml reads file data from disk. It returns an error on failure.
func staticTestHtml() (*asset, error) {
	path := "static/test.html"
	name := "static/test.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticUpdating_blockchainHtml reads file data from disk. It returns an error on failure.
func staticUpdating_blockchainHtml() (*asset, error) {
	path := "static/updating_blockchain.html"
	name := "static/updating_blockchain.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorAnimateCssAnimateMinCss reads file data from disk. It returns an error on failure.
func staticVendorAnimateCssAnimateMinCss() (*asset, error) {
	path := "static/vendor/animate.css/animate.min.css"
	name := "static/vendor/animate.css/animate.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorBootstrapDistCssBootstrapCss reads file data from disk. It returns an error on failure.
func staticVendorBootstrapDistCssBootstrapCss() (*asset, error) {
	path := "static/vendor/bootstrap/dist/css/bootstrap.css"
	name := "static/vendor/bootstrap/dist/css/bootstrap.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorBootstrapDistJsBootstrapJs reads file data from disk. It returns an error on failure.
func staticVendorBootstrapDistJsBootstrapJs() (*asset, error) {
	path := "static/vendor/bootstrap/dist/js/bootstrap.js"
	name := "static/vendor/bootstrap/dist/js/bootstrap.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorBootstrapFilestyleSrcBootstrapFilestyleJs reads file data from disk. It returns an error on failure.
func staticVendorBootstrapFilestyleSrcBootstrapFilestyleJs() (*asset, error) {
	path := "static/vendor/bootstrap-filestyle/src/bootstrap-filestyle.js"
	name := "static/vendor/bootstrap-filestyle/src/bootstrap-filestyle.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesFontawesomeIndexCss reads file data from disk. It returns an error on failure.
func staticVendorDatatablesFontawesomeIndexCss() (*asset, error) {
	path := "static/vendor/dataTables.fontAwesome/index.css"
	name := "static/vendor/dataTables.fontAwesome/index.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaCssDatatablesBootstrapCss reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaCssDatatablesBootstrapCss() (*asset, error) {
	path := "static/vendor/datatables/media/css/dataTables.bootstrap.css"
	name := "static/vendor/datatables/media/css/dataTables.bootstrap.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaImagesSortingIconsPsd reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaImagesSortingIconsPsd() (*asset, error) {
	path := "static/vendor/datatables/media/images/Sorting icons.psd"
	name := "static/vendor/datatables/media/images/Sorting icons.psd"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaImagesFaviconIco reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaImagesFaviconIco() (*asset, error) {
	path := "static/vendor/datatables/media/images/favicon.ico"
	name := "static/vendor/datatables/media/images/favicon.ico"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaImagesSort_ascPng reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaImagesSort_ascPng() (*asset, error) {
	path := "static/vendor/datatables/media/images/sort_asc.png"
	name := "static/vendor/datatables/media/images/sort_asc.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaImagesSort_asc_disabledPng reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaImagesSort_asc_disabledPng() (*asset, error) {
	path := "static/vendor/datatables/media/images/sort_asc_disabled.png"
	name := "static/vendor/datatables/media/images/sort_asc_disabled.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaImagesSort_bothPng reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaImagesSort_bothPng() (*asset, error) {
	path := "static/vendor/datatables/media/images/sort_both.png"
	name := "static/vendor/datatables/media/images/sort_both.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaImagesSort_descPng reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaImagesSort_descPng() (*asset, error) {
	path := "static/vendor/datatables/media/images/sort_desc.png"
	name := "static/vendor/datatables/media/images/sort_desc.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaImagesSort_desc_disabledPng reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaImagesSort_desc_disabledPng() (*asset, error) {
	path := "static/vendor/datatables/media/images/sort_desc_disabled.png"
	name := "static/vendor/datatables/media/images/sort_desc_disabled.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaJsDatatablesBootstrapJs reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaJsDatatablesBootstrapJs() (*asset, error) {
	path := "static/vendor/datatables/media/js/dataTables.bootstrap.js"
	name := "static/vendor/datatables/media/js/dataTables.bootstrap.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesMediaJsJqueryDatatablesMinJs reads file data from disk. It returns an error on failure.
func staticVendorDatatablesMediaJsJqueryDatatablesMinJs() (*asset, error) {
	path := "static/vendor/datatables/media/js/jquery.dataTables.min.js"
	name := "static/vendor/datatables/media/js/jquery.dataTables.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesColvisCssDatatablesColvisCss reads file data from disk. It returns an error on failure.
func staticVendorDatatablesColvisCssDatatablesColvisCss() (*asset, error) {
	path := "static/vendor/datatables-colvis/css/dataTables.colVis.css"
	name := "static/vendor/datatables-colvis/css/dataTables.colVis.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorDatatablesColvisJsDatatablesColvisJs reads file data from disk. It returns an error on failure.
func staticVendorDatatablesColvisJsDatatablesColvisJs() (*asset, error) {
	path := "static/vendor/datatables-colvis/js/dataTables.colVis.js"
	name := "static/vendor/datatables-colvis/js/dataTables.colVis.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorEonasdanBootstrapDatetimepickerBuildCssBootstrapDatetimepickerMinCss reads file data from disk. It returns an error on failure.
func staticVendorEonasdanBootstrapDatetimepickerBuildCssBootstrapDatetimepickerMinCss() (*asset, error) {
	path := "static/vendor/eonasdan-bootstrap-datetimepicker/build/css/bootstrap-datetimepicker.min.css"
	name := "static/vendor/eonasdan-bootstrap-datetimepicker/build/css/bootstrap-datetimepicker.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorEonasdanBootstrapDatetimepickerBuildJsBootstrapDatetimepickerMinJs reads file data from disk. It returns an error on failure.
func staticVendorEonasdanBootstrapDatetimepickerBuildJsBootstrapDatetimepickerMinJs() (*asset, error) {
	path := "static/vendor/eonasdan-bootstrap-datetimepicker/build/js/bootstrap-datetimepicker.min.js"
	name := "static/vendor/eonasdan-bootstrap-datetimepicker/build/js/bootstrap-datetimepicker.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorFontawesomeCssFontAwesomeMinCss reads file data from disk. It returns an error on failure.
func staticVendorFontawesomeCssFontAwesomeMinCss() (*asset, error) {
	path := "static/vendor/fontawesome/css/font-awesome.min.css"
	name := "static/vendor/fontawesome/css/font-awesome.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorFontawesomeFontsFontawesomeOtf reads file data from disk. It returns an error on failure.
func staticVendorFontawesomeFontsFontawesomeOtf() (*asset, error) {
	path := "static/vendor/fontawesome/fonts/FontAwesome.otf"
	name := "static/vendor/fontawesome/fonts/FontAwesome.otf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorFontawesomeFontsFontawesomeWebfontEot reads file data from disk. It returns an error on failure.
func staticVendorFontawesomeFontsFontawesomeWebfontEot() (*asset, error) {
	path := "static/vendor/fontawesome/fonts/fontawesome-webfont.eot"
	name := "static/vendor/fontawesome/fonts/fontawesome-webfont.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorFontawesomeFontsFontawesomeWebfontSvg reads file data from disk. It returns an error on failure.
func staticVendorFontawesomeFontsFontawesomeWebfontSvg() (*asset, error) {
	path := "static/vendor/fontawesome/fonts/fontawesome-webfont.svg"
	name := "static/vendor/fontawesome/fonts/fontawesome-webfont.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorFontawesomeFontsFontawesomeWebfontTtf reads file data from disk. It returns an error on failure.
func staticVendorFontawesomeFontsFontawesomeWebfontTtf() (*asset, error) {
	path := "static/vendor/fontawesome/fonts/fontawesome-webfont.ttf"
	name := "static/vendor/fontawesome/fonts/fontawesome-webfont.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorFontawesomeFontsFontawesomeWebfontWoff reads file data from disk. It returns an error on failure.
func staticVendorFontawesomeFontsFontawesomeWebfontWoff() (*asset, error) {
	path := "static/vendor/fontawesome/fonts/fontawesome-webfont.woff"
	name := "static/vendor/fontawesome/fonts/fontawesome-webfont.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorFontawesomeFontsFontawesomeWebfontWoff2 reads file data from disk. It returns an error on failure.
func staticVendorFontawesomeFontsFontawesomeWebfontWoff2() (*asset, error) {
	path := "static/vendor/fontawesome/fonts/fontawesome-webfont.woff2"
	name := "static/vendor/fontawesome/fonts/fontawesome-webfont.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorJqueryStorageApiJqueryStorageapiJs reads file data from disk. It returns an error on failure.
func staticVendorJqueryStorageApiJqueryStorageapiJs() (*asset, error) {
	path := "static/vendor/jQuery-Storage-API/jquery.storageapi.js"
	name := "static/vendor/jQuery-Storage-API/jquery.storageapi.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorJqueryDistJqueryJs reads file data from disk. It returns an error on failure.
func staticVendorJqueryDistJqueryJs() (*asset, error) {
	path := "static/vendor/jquery/dist/jquery.js"
	name := "static/vendor/jquery/dist/jquery.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorLoadersCssLoadersCss reads file data from disk. It returns an error on failure.
func staticVendorLoadersCssLoadersCss() (*asset, error) {
	path := "static/vendor/loaders.css/loaders.css"
	name := "static/vendor/loaders.css/loaders.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorModernizrModernizrCustomJs reads file data from disk. It returns an error on failure.
func staticVendorModernizrModernizrCustomJs() (*asset, error) {
	path := "static/vendor/modernizr/modernizr.custom.js"
	name := "static/vendor/modernizr/modernizr.custom.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorMomentMinMomentWithLocalesMinJs reads file data from disk. It returns an error on failure.
func staticVendorMomentMinMomentWithLocalesMinJs() (*asset, error) {
	path := "static/vendor/moment/min/moment-with-locales.min.js"
	name := "static/vendor/moment/min/moment-with-locales.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSelect2DistCssSelect2Css reads file data from disk. It returns an error on failure.
func staticVendorSelect2DistCssSelect2Css() (*asset, error) {
	path := "static/vendor/select2/dist/css/select2.css"
	name := "static/vendor/select2/dist/css/select2.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSelect2DistJsSelect2Js reads file data from disk. It returns an error on failure.
func staticVendorSelect2DistJsSelect2Js() (*asset, error) {
	path := "static/vendor/select2/dist/js/select2.js"
	name := "static/vendor/select2/dist/js/select2.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSelect2BootstrapThemeDistSelect2BootstrapCss reads file data from disk. It returns an error on failure.
func staticVendorSelect2BootstrapThemeDistSelect2BootstrapCss() (*asset, error) {
	path := "static/vendor/select2-bootstrap-theme/dist/select2-bootstrap.css"
	name := "static/vendor/select2-bootstrap-theme/dist/select2-bootstrap.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSimpleLineIconsCssSimpleLineIconsCss reads file data from disk. It returns an error on failure.
func staticVendorSimpleLineIconsCssSimpleLineIconsCss() (*asset, error) {
	path := "static/vendor/simple-line-icons/css/simple-line-icons.css"
	name := "static/vendor/simple-line-icons/css/simple-line-icons.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSimpleLineIconsFontsSimpleLineIconsEot reads file data from disk. It returns an error on failure.
func staticVendorSimpleLineIconsFontsSimpleLineIconsEot() (*asset, error) {
	path := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.eot"
	name := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSimpleLineIconsFontsSimpleLineIconsSvg reads file data from disk. It returns an error on failure.
func staticVendorSimpleLineIconsFontsSimpleLineIconsSvg() (*asset, error) {
	path := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.svg"
	name := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSimpleLineIconsFontsSimpleLineIconsTtf reads file data from disk. It returns an error on failure.
func staticVendorSimpleLineIconsFontsSimpleLineIconsTtf() (*asset, error) {
	path := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.ttf"
	name := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSimpleLineIconsFontsSimpleLineIconsWoff reads file data from disk. It returns an error on failure.
func staticVendorSimpleLineIconsFontsSimpleLineIconsWoff() (*asset, error) {
	path := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.woff"
	name := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSimpleLineIconsFontsSimpleLineIconsWoff2 reads file data from disk. It returns an error on failure.
func staticVendorSimpleLineIconsFontsSimpleLineIconsWoff2() (*asset, error) {
	path := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.woff2"
	name := "static/vendor/simple-line-icons/fonts/Simple-Line-Icons.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSpinkitCssSpinkitCss reads file data from disk. It returns an error on failure.
func staticVendorSpinkitCssSpinkitCss() (*asset, error) {
	path := "static/vendor/spinkit/css/spinkit.css"
	name := "static/vendor/spinkit/css/spinkit.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSweetalertDistSweetalertCss reads file data from disk. It returns an error on failure.
func staticVendorSweetalertDistSweetalertCss() (*asset, error) {
	path := "static/vendor/sweetalert/dist/sweetalert.css"
	name := "static/vendor/sweetalert/dist/sweetalert.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorSweetalertDistSweetalertMinJs reads file data from disk. It returns an error on failure.
func staticVendorSweetalertDistSweetalertMinJs() (*asset, error) {
	path := "static/vendor/sweetalert/dist/sweetalert.min.js"
	name := "static/vendor/sweetalert/dist/sweetalert.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorWhirlDistWhirlCss reads file data from disk. It returns an error on failure.
func staticVendorWhirlDistWhirlCss() (*asset, error) {
	path := "static/vendor/whirl/dist/whirl.css"
	name := "static/vendor/whirl/dist/whirl.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorXEditableDistBootstrap3EditableCssBootstrapEditableCss reads file data from disk. It returns an error on failure.
func staticVendorXEditableDistBootstrap3EditableCssBootstrapEditableCss() (*asset, error) {
	path := "static/vendor/x-editable/dist/bootstrap3-editable/css/bootstrap-editable.css"
	name := "static/vendor/x-editable/dist/bootstrap3-editable/css/bootstrap-editable.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorXEditableDistBootstrap3EditableImgClearPng reads file data from disk. It returns an error on failure.
func staticVendorXEditableDistBootstrap3EditableImgClearPng() (*asset, error) {
	path := "static/vendor/x-editable/dist/bootstrap3-editable/img/clear.png"
	name := "static/vendor/x-editable/dist/bootstrap3-editable/img/clear.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorXEditableDistBootstrap3EditableImgLoadingGif reads file data from disk. It returns an error on failure.
func staticVendorXEditableDistBootstrap3EditableImgLoadingGif() (*asset, error) {
	path := "static/vendor/x-editable/dist/bootstrap3-editable/img/loading.gif"
	name := "static/vendor/x-editable/dist/bootstrap3-editable/img/loading.gif"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticVendorXEditableDistBootstrap3EditableJsBootstrapEditableMinJs reads file data from disk. It returns an error on failure.
func staticVendorXEditableDistBootstrap3EditableJsBootstrapEditableMinJs() (*asset, error) {
	path := "static/vendor/x-editable/dist/bootstrap3-editable/js/bootstrap-editable.min.js"
	name := "static/vendor/x-editable/dist/bootstrap3-editable/js/bootstrap-editable.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/1block": static1block,
	"static/1block-local": static1blockLocal,
	"static/alert_success.html": staticAlert_successHtml,
	"static/anonym_money_transfer.html": staticAnonym_money_transferHtml,
	"static/backup.html": staticBackupHtml,
	"static/block_explorer.html": staticBlock_explorerHtml,
	"static/block_generation.html": staticBlock_generationHtml,
	"static/change_node_key.html": staticChange_node_keyHtml,
	"static/css/style.css": staticCssStyleCss,
	"static/dashboard_anonym.html": staticDashboard_anonymHtml,
	"static/fonts/glyphicons-filetypes-regular.eot": staticFontsGlyphiconsFiletypesRegularEot,
	"static/fonts/glyphicons-filetypes-regular.svg": staticFontsGlyphiconsFiletypesRegularSvg,
	"static/fonts/glyphicons-filetypes-regular.ttf": staticFontsGlyphiconsFiletypesRegularTtf,
	"static/fonts/glyphicons-filetypes-regular.woff": staticFontsGlyphiconsFiletypesRegularWoff,
	"static/fonts/glyphicons-filetypes-regular.woff2": staticFontsGlyphiconsFiletypesRegularWoff2,
	"static/fonts/glyphicons-halflings-regular.eot": staticFontsGlyphiconsHalflingsRegularEot,
	"static/fonts/glyphicons-halflings-regular.svg": staticFontsGlyphiconsHalflingsRegularSvg,
	"static/fonts/glyphicons-halflings-regular.ttf": staticFontsGlyphiconsHalflingsRegularTtf,
	"static/fonts/glyphicons-halflings-regular.woff": staticFontsGlyphiconsHalflingsRegularWoff,
	"static/fonts/glyphicons-halflings-regular.woff2": staticFontsGlyphiconsHalflingsRegularWoff2,
	"static/fonts/glyphicons-regular.eot": staticFontsGlyphiconsRegularEot,
	"static/fonts/glyphicons-regular.svg": staticFontsGlyphiconsRegularSvg,
	"static/fonts/glyphicons-regular.ttf": staticFontsGlyphiconsRegularTtf,
	"static/fonts/glyphicons-regular.woff": staticFontsGlyphiconsRegularWoff,
	"static/fonts/glyphicons-regular.woff2": staticFontsGlyphiconsRegularWoff2,
	"static/fonts/glyphicons-social-regular.eot": staticFontsGlyphiconsSocialRegularEot,
	"static/fonts/glyphicons-social-regular.svg": staticFontsGlyphiconsSocialRegularSvg,
	"static/fonts/glyphicons-social-regular.ttf": staticFontsGlyphiconsSocialRegularTtf,
	"static/fonts/glyphicons-social-regular.woff": staticFontsGlyphiconsSocialRegularWoff,
	"static/fonts/glyphicons-social-regular.woff2": staticFontsGlyphiconsSocialRegularWoff2,
	"static/history.html": staticHistoryHtml,
	"static/img/D_null.png": staticImgD_nullPng,
	"static/img/D_ok.png": staticImgD_okPng,
	"static/img/US.svg": staticImgUsSvg,
	"static/img/bg1.jpg": staticImgBg1Jpg,
	"static/img/bg10.jpg": staticImgBg10Jpg,
	"static/img/bg2.jpg": staticImgBg2Jpg,
	"static/img/bg3.jpg": staticImgBg3Jpg,
	"static/img/bg4.jpg": staticImgBg4Jpg,
	"static/img/bg5.jpg": staticImgBg5Jpg,
	"static/img/bg6.jpg": staticImgBg6Jpg,
	"static/img/bg7.jpg": staticImgBg7Jpg,
	"static/img/bg8.jpg": staticImgBg8Jpg,
	"static/img/bg9.jpg": staticImgBg9Jpg,
	"static/img/dummy.png": staticImgDummyPng,
	"static/img/favicon.ico": staticImgFaviconIco,
	"static/img/lock-bg.jpg": staticImgLockBgJpg,
	"static/img/logo-single.svg": staticImgLogoSingleSvg,
	"static/img/logo.svg": staticImgLogoSvg,
	"static/img/mb-sample.jpg": staticImgMbSampleJpg,
	"static/img/mockup.png": staticImgMockupPng,
	"static/img/profile-bg.jpg": staticImgProfileBgJpg,
	"static/img/radio_icon.png": staticImgRadio_iconPng,
	"static/img/squares.gif": staticImgSquaresGif,
	"static/img/user/01.jpg": staticImgUser01Jpg,
	"static/img/user/02.jpg": staticImgUser02Jpg,
	"static/img/user/03.jpg": staticImgUser03Jpg,
	"static/img/user/04.jpg": staticImgUser04Jpg,
	"static/img/user/05.jpg": staticImgUser05Jpg,
	"static/img/user/06.jpg": staticImgUser06Jpg,
	"static/img/user/07.jpg": staticImgUser07Jpg,
	"static/img/user/08.jpg": staticImgUser08Jpg,
	"static/img/user/09.jpg": staticImgUser09Jpg,
	"static/img/user/10.jpg": staticImgUser10Jpg,
	"static/img/user/11.jpg": staticImgUser11Jpg,
	"static/img/user/12.jpg": staticImgUser12Jpg,
	"static/img/user/13.jpg": staticImgUser13Jpg,
	"static/index.html": staticIndexHtml,
	"static/install_step_0.html": staticInstall_step_0Html,
	"static/install_step_1.html": staticInstall_step_1Html,
	"static/js/app.js": staticJsAppJs,
	"static/js/clipboard.js": staticJsClipboardJs,
	"static/js/crypto/asn1-1.0.min.js": staticJsCryptoAsn110MinJs,
	"static/js/crypto/asn1hex-1.1.min.js": staticJsCryptoAsn1hex11MinJs,
	"static/js/crypto/crypto-1.1.min.js": staticJsCryptoCrypto11MinJs,
	"static/js/crypto/ecdsa-modified-1.0.min.js": staticJsCryptoEcdsaModified10MinJs,
	"static/js/crypto/ecparam-1.0.min.js": staticJsCryptoEcparam10MinJs,
	"static/js/crypto/ext/cj/aes.js": staticJsCryptoExtCjAesJs,
	"static/js/crypto/ext/cj/cryptojs-312-core-fix-min.js": staticJsCryptoExtCjCryptojs312CoreFixMinJs,
	"static/js/crypto/ext/cj/ec-patch-min.js": staticJsCryptoExtCjEcPatchMinJs,
	"static/js/crypto/ext/cj/md5_min.js": staticJsCryptoExtCjMd5_minJs,
	"static/js/crypto/ext/cj/sha1_min.js": staticJsCryptoExtCjSha1_minJs,
	"static/js/crypto/ext/cj/sha256_min.js": staticJsCryptoExtCjSha256_minJs,
	"static/js/crypto/ext/cj/x64-core_min.js": staticJsCryptoExtCjX64Core_minJs,
	"static/js/crypto/ext/ec-min.js": staticJsCryptoExtEcMinJs,
	"static/js/crypto/ext/ec-patch-min.js": staticJsCryptoExtEcPatchMinJs,
	"static/js/crypto/ext/jsbn-min.js": staticJsCryptoExtJsbnMinJs,
	"static/js/crypto/ext/jsbn2-min.js": staticJsCryptoExtJsbn2MinJs,
	"static/js/crypto/ext/prng4-min.js": staticJsCryptoExtPrng4MinJs,
	"static/js/crypto/ext/rng-min.js": staticJsCryptoExtRngMinJs,
	"static/js/crypto/ext/yahoo-min.js": staticJsCryptoExtYahooMinJs,
	"static/js/datetime.js": staticJsDatetimeJs,
	"static/js/demo/demo-datatable.js": staticJsDemoDemoDatatableJs,
	"static/js/demo/demo-flot.js": staticJsDemoDemoFlotJs,
	"static/js/demo/demo-forms.js": staticJsDemoDemoFormsJs,
	"static/js/demo/demo-jqcloud.js": staticJsDemoDemoJqcloudJs,
	"static/js/demo/demo-jqgrid.js": staticJsDemoDemoJqgridJs,
	"static/js/demo/demo-nestable.js": staticJsDemoDemoNestableJs,
	"static/js/demo/demo-panels.js": staticJsDemoDemoPanelsJs,
	"static/js/demo/demo-rtl.js": staticJsDemoDemoRtlJs,
	"static/js/demo/demo-search.js": staticJsDemoDemoSearchJs,
	"static/js/demo/demo-sortable.js": staticJsDemoDemoSortableJs,
	"static/js/demo/demo-upload.js": staticJsDemoDemoUploadJs,
	"static/js/demo/demo-vector-map.js": staticJsDemoDemoVectorMapJs,
	"static/js/demo/demo-wizard.js": staticJsDemoDemoWizardJs,
	"static/js/demo/demo-xeditable.js": staticJsDemoDemoXeditableJs,
	"static/js/index.js": staticJsIndexJs,
	"static/js/nprogress.js": staticJsNprogressJs,
	"static/lang/1.ini": staticLang1Ini,
	"static/lang/42.ini": staticLang42Ini,
	"static/lang/en-us.all.json": staticLangEnUsAllJson,
	"static/lang/locale_en-US.ini": staticLangLocale_enUsIni,
	"static/lang/locale_ru-RU.ini": staticLangLocale_ruRuIni,
	"static/login.html": staticLoginHtml,
	"static/menu.html": staticMenuHtml,
	"static/modal_anonym.html": staticModal_anonymHtml,
	"static/pass.html": staticPassHtml,
	"static/psw.html": staticPswHtml,
	"static/request_citizen_status.html": staticRequest_citizen_statusHtml,
	"static/signatures.html": staticSignaturesHtml,
	"static/signatures_new.html": staticSignatures_newHtml,
	"static/test.html": staticTestHtml,
	"static/updating_blockchain.html": staticUpdating_blockchainHtml,
	"static/vendor/animate.css/animate.min.css": staticVendorAnimateCssAnimateMinCss,
	"static/vendor/bootstrap/dist/css/bootstrap.css": staticVendorBootstrapDistCssBootstrapCss,
	"static/vendor/bootstrap/dist/js/bootstrap.js": staticVendorBootstrapDistJsBootstrapJs,
	"static/vendor/bootstrap-filestyle/src/bootstrap-filestyle.js": staticVendorBootstrapFilestyleSrcBootstrapFilestyleJs,
	"static/vendor/dataTables.fontAwesome/index.css": staticVendorDatatablesFontawesomeIndexCss,
	"static/vendor/datatables/media/css/dataTables.bootstrap.css": staticVendorDatatablesMediaCssDatatablesBootstrapCss,
	"static/vendor/datatables/media/images/Sorting icons.psd": staticVendorDatatablesMediaImagesSortingIconsPsd,
	"static/vendor/datatables/media/images/favicon.ico": staticVendorDatatablesMediaImagesFaviconIco,
	"static/vendor/datatables/media/images/sort_asc.png": staticVendorDatatablesMediaImagesSort_ascPng,
	"static/vendor/datatables/media/images/sort_asc_disabled.png": staticVendorDatatablesMediaImagesSort_asc_disabledPng,
	"static/vendor/datatables/media/images/sort_both.png": staticVendorDatatablesMediaImagesSort_bothPng,
	"static/vendor/datatables/media/images/sort_desc.png": staticVendorDatatablesMediaImagesSort_descPng,
	"static/vendor/datatables/media/images/sort_desc_disabled.png": staticVendorDatatablesMediaImagesSort_desc_disabledPng,
	"static/vendor/datatables/media/js/dataTables.bootstrap.js": staticVendorDatatablesMediaJsDatatablesBootstrapJs,
	"static/vendor/datatables/media/js/jquery.dataTables.min.js": staticVendorDatatablesMediaJsJqueryDatatablesMinJs,
	"static/vendor/datatables-colvis/css/dataTables.colVis.css": staticVendorDatatablesColvisCssDatatablesColvisCss,
	"static/vendor/datatables-colvis/js/dataTables.colVis.js": staticVendorDatatablesColvisJsDatatablesColvisJs,
	"static/vendor/eonasdan-bootstrap-datetimepicker/build/css/bootstrap-datetimepicker.min.css": staticVendorEonasdanBootstrapDatetimepickerBuildCssBootstrapDatetimepickerMinCss,
	"static/vendor/eonasdan-bootstrap-datetimepicker/build/js/bootstrap-datetimepicker.min.js": staticVendorEonasdanBootstrapDatetimepickerBuildJsBootstrapDatetimepickerMinJs,
	"static/vendor/fontawesome/css/font-awesome.min.css": staticVendorFontawesomeCssFontAwesomeMinCss,
	"static/vendor/fontawesome/fonts/FontAwesome.otf": staticVendorFontawesomeFontsFontawesomeOtf,
	"static/vendor/fontawesome/fonts/fontawesome-webfont.eot": staticVendorFontawesomeFontsFontawesomeWebfontEot,
	"static/vendor/fontawesome/fonts/fontawesome-webfont.svg": staticVendorFontawesomeFontsFontawesomeWebfontSvg,
	"static/vendor/fontawesome/fonts/fontawesome-webfont.ttf": staticVendorFontawesomeFontsFontawesomeWebfontTtf,
	"static/vendor/fontawesome/fonts/fontawesome-webfont.woff": staticVendorFontawesomeFontsFontawesomeWebfontWoff,
	"static/vendor/fontawesome/fonts/fontawesome-webfont.woff2": staticVendorFontawesomeFontsFontawesomeWebfontWoff2,
	"static/vendor/jQuery-Storage-API/jquery.storageapi.js": staticVendorJqueryStorageApiJqueryStorageapiJs,
	"static/vendor/jquery/dist/jquery.js": staticVendorJqueryDistJqueryJs,
	"static/vendor/loaders.css/loaders.css": staticVendorLoadersCssLoadersCss,
	"static/vendor/modernizr/modernizr.custom.js": staticVendorModernizrModernizrCustomJs,
	"static/vendor/moment/min/moment-with-locales.min.js": staticVendorMomentMinMomentWithLocalesMinJs,
	"static/vendor/select2/dist/css/select2.css": staticVendorSelect2DistCssSelect2Css,
	"static/vendor/select2/dist/js/select2.js": staticVendorSelect2DistJsSelect2Js,
	"static/vendor/select2-bootstrap-theme/dist/select2-bootstrap.css": staticVendorSelect2BootstrapThemeDistSelect2BootstrapCss,
	"static/vendor/simple-line-icons/css/simple-line-icons.css": staticVendorSimpleLineIconsCssSimpleLineIconsCss,
	"static/vendor/simple-line-icons/fonts/Simple-Line-Icons.eot": staticVendorSimpleLineIconsFontsSimpleLineIconsEot,
	"static/vendor/simple-line-icons/fonts/Simple-Line-Icons.svg": staticVendorSimpleLineIconsFontsSimpleLineIconsSvg,
	"static/vendor/simple-line-icons/fonts/Simple-Line-Icons.ttf": staticVendorSimpleLineIconsFontsSimpleLineIconsTtf,
	"static/vendor/simple-line-icons/fonts/Simple-Line-Icons.woff": staticVendorSimpleLineIconsFontsSimpleLineIconsWoff,
	"static/vendor/simple-line-icons/fonts/Simple-Line-Icons.woff2": staticVendorSimpleLineIconsFontsSimpleLineIconsWoff2,
	"static/vendor/spinkit/css/spinkit.css": staticVendorSpinkitCssSpinkitCss,
	"static/vendor/sweetalert/dist/sweetalert.css": staticVendorSweetalertDistSweetalertCss,
	"static/vendor/sweetalert/dist/sweetalert.min.js": staticVendorSweetalertDistSweetalertMinJs,
	"static/vendor/whirl/dist/whirl.css": staticVendorWhirlDistWhirlCss,
	"static/vendor/x-editable/dist/bootstrap3-editable/css/bootstrap-editable.css": staticVendorXEditableDistBootstrap3EditableCssBootstrapEditableCss,
	"static/vendor/x-editable/dist/bootstrap3-editable/img/clear.png": staticVendorXEditableDistBootstrap3EditableImgClearPng,
	"static/vendor/x-editable/dist/bootstrap3-editable/img/loading.gif": staticVendorXEditableDistBootstrap3EditableImgLoadingGif,
	"static/vendor/x-editable/dist/bootstrap3-editable/js/bootstrap-editable.min.js": staticVendorXEditableDistBootstrap3EditableJsBootstrapEditableMinJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"static": &bintree{nil, map[string]*bintree{
		"1block": &bintree{static1block, map[string]*bintree{}},
		"1block-local": &bintree{static1blockLocal, map[string]*bintree{}},
		"alert_success.html": &bintree{staticAlert_successHtml, map[string]*bintree{}},
		"anonym_money_transfer.html": &bintree{staticAnonym_money_transferHtml, map[string]*bintree{}},
		"backup.html": &bintree{staticBackupHtml, map[string]*bintree{}},
		"block_explorer.html": &bintree{staticBlock_explorerHtml, map[string]*bintree{}},
		"block_generation.html": &bintree{staticBlock_generationHtml, map[string]*bintree{}},
		"change_node_key.html": &bintree{staticChange_node_keyHtml, map[string]*bintree{}},
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{staticCssStyleCss, map[string]*bintree{}},
		}},
		"dashboard_anonym.html": &bintree{staticDashboard_anonymHtml, map[string]*bintree{}},
		"fonts": &bintree{nil, map[string]*bintree{
			"glyphicons-filetypes-regular.eot": &bintree{staticFontsGlyphiconsFiletypesRegularEot, map[string]*bintree{}},
			"glyphicons-filetypes-regular.svg": &bintree{staticFontsGlyphiconsFiletypesRegularSvg, map[string]*bintree{}},
			"glyphicons-filetypes-regular.ttf": &bintree{staticFontsGlyphiconsFiletypesRegularTtf, map[string]*bintree{}},
			"glyphicons-filetypes-regular.woff": &bintree{staticFontsGlyphiconsFiletypesRegularWoff, map[string]*bintree{}},
			"glyphicons-filetypes-regular.woff2": &bintree{staticFontsGlyphiconsFiletypesRegularWoff2, map[string]*bintree{}},
			"glyphicons-halflings-regular.eot": &bintree{staticFontsGlyphiconsHalflingsRegularEot, map[string]*bintree{}},
			"glyphicons-halflings-regular.svg": &bintree{staticFontsGlyphiconsHalflingsRegularSvg, map[string]*bintree{}},
			"glyphicons-halflings-regular.ttf": &bintree{staticFontsGlyphiconsHalflingsRegularTtf, map[string]*bintree{}},
			"glyphicons-halflings-regular.woff": &bintree{staticFontsGlyphiconsHalflingsRegularWoff, map[string]*bintree{}},
			"glyphicons-halflings-regular.woff2": &bintree{staticFontsGlyphiconsHalflingsRegularWoff2, map[string]*bintree{}},
			"glyphicons-regular.eot": &bintree{staticFontsGlyphiconsRegularEot, map[string]*bintree{}},
			"glyphicons-regular.svg": &bintree{staticFontsGlyphiconsRegularSvg, map[string]*bintree{}},
			"glyphicons-regular.ttf": &bintree{staticFontsGlyphiconsRegularTtf, map[string]*bintree{}},
			"glyphicons-regular.woff": &bintree{staticFontsGlyphiconsRegularWoff, map[string]*bintree{}},
			"glyphicons-regular.woff2": &bintree{staticFontsGlyphiconsRegularWoff2, map[string]*bintree{}},
			"glyphicons-social-regular.eot": &bintree{staticFontsGlyphiconsSocialRegularEot, map[string]*bintree{}},
			"glyphicons-social-regular.svg": &bintree{staticFontsGlyphiconsSocialRegularSvg, map[string]*bintree{}},
			"glyphicons-social-regular.ttf": &bintree{staticFontsGlyphiconsSocialRegularTtf, map[string]*bintree{}},
			"glyphicons-social-regular.woff": &bintree{staticFontsGlyphiconsSocialRegularWoff, map[string]*bintree{}},
			"glyphicons-social-regular.woff2": &bintree{staticFontsGlyphiconsSocialRegularWoff2, map[string]*bintree{}},
		}},
		"history.html": &bintree{staticHistoryHtml, map[string]*bintree{}},
		"img": &bintree{nil, map[string]*bintree{
			"D_null.png": &bintree{staticImgD_nullPng, map[string]*bintree{}},
			"D_ok.png": &bintree{staticImgD_okPng, map[string]*bintree{}},
			"US.svg": &bintree{staticImgUsSvg, map[string]*bintree{}},
			"bg1.jpg": &bintree{staticImgBg1Jpg, map[string]*bintree{}},
			"bg10.jpg": &bintree{staticImgBg10Jpg, map[string]*bintree{}},
			"bg2.jpg": &bintree{staticImgBg2Jpg, map[string]*bintree{}},
			"bg3.jpg": &bintree{staticImgBg3Jpg, map[string]*bintree{}},
			"bg4.jpg": &bintree{staticImgBg4Jpg, map[string]*bintree{}},
			"bg5.jpg": &bintree{staticImgBg5Jpg, map[string]*bintree{}},
			"bg6.jpg": &bintree{staticImgBg6Jpg, map[string]*bintree{}},
			"bg7.jpg": &bintree{staticImgBg7Jpg, map[string]*bintree{}},
			"bg8.jpg": &bintree{staticImgBg8Jpg, map[string]*bintree{}},
			"bg9.jpg": &bintree{staticImgBg9Jpg, map[string]*bintree{}},
			"dummy.png": &bintree{staticImgDummyPng, map[string]*bintree{}},
			"favicon.ico": &bintree{staticImgFaviconIco, map[string]*bintree{}},
			"lock-bg.jpg": &bintree{staticImgLockBgJpg, map[string]*bintree{}},
			"logo-single.svg": &bintree{staticImgLogoSingleSvg, map[string]*bintree{}},
			"logo.svg": &bintree{staticImgLogoSvg, map[string]*bintree{}},
			"mb-sample.jpg": &bintree{staticImgMbSampleJpg, map[string]*bintree{}},
			"mockup.png": &bintree{staticImgMockupPng, map[string]*bintree{}},
			"profile-bg.jpg": &bintree{staticImgProfileBgJpg, map[string]*bintree{}},
			"radio_icon.png": &bintree{staticImgRadio_iconPng, map[string]*bintree{}},
			"squares.gif": &bintree{staticImgSquaresGif, map[string]*bintree{}},
			"user": &bintree{nil, map[string]*bintree{
				"01.jpg": &bintree{staticImgUser01Jpg, map[string]*bintree{}},
				"02.jpg": &bintree{staticImgUser02Jpg, map[string]*bintree{}},
				"03.jpg": &bintree{staticImgUser03Jpg, map[string]*bintree{}},
				"04.jpg": &bintree{staticImgUser04Jpg, map[string]*bintree{}},
				"05.jpg": &bintree{staticImgUser05Jpg, map[string]*bintree{}},
				"06.jpg": &bintree{staticImgUser06Jpg, map[string]*bintree{}},
				"07.jpg": &bintree{staticImgUser07Jpg, map[string]*bintree{}},
				"08.jpg": &bintree{staticImgUser08Jpg, map[string]*bintree{}},
				"09.jpg": &bintree{staticImgUser09Jpg, map[string]*bintree{}},
				"10.jpg": &bintree{staticImgUser10Jpg, map[string]*bintree{}},
				"11.jpg": &bintree{staticImgUser11Jpg, map[string]*bintree{}},
				"12.jpg": &bintree{staticImgUser12Jpg, map[string]*bintree{}},
				"13.jpg": &bintree{staticImgUser13Jpg, map[string]*bintree{}},
			}},
		}},
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
		"install_step_0.html": &bintree{staticInstall_step_0Html, map[string]*bintree{}},
		"install_step_1.html": &bintree{staticInstall_step_1Html, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"app.js": &bintree{staticJsAppJs, map[string]*bintree{}},
			"clipboard.js": &bintree{staticJsClipboardJs, map[string]*bintree{}},
			"crypto": &bintree{nil, map[string]*bintree{
				"asn1-1.0.min.js": &bintree{staticJsCryptoAsn110MinJs, map[string]*bintree{}},
				"asn1hex-1.1.min.js": &bintree{staticJsCryptoAsn1hex11MinJs, map[string]*bintree{}},
				"crypto-1.1.min.js": &bintree{staticJsCryptoCrypto11MinJs, map[string]*bintree{}},
				"ecdsa-modified-1.0.min.js": &bintree{staticJsCryptoEcdsaModified10MinJs, map[string]*bintree{}},
				"ecparam-1.0.min.js": &bintree{staticJsCryptoEcparam10MinJs, map[string]*bintree{}},
				"ext": &bintree{nil, map[string]*bintree{
					"cj": &bintree{nil, map[string]*bintree{
						"aes.js": &bintree{staticJsCryptoExtCjAesJs, map[string]*bintree{}},
						"cryptojs-312-core-fix-min.js": &bintree{staticJsCryptoExtCjCryptojs312CoreFixMinJs, map[string]*bintree{}},
						"ec-patch-min.js": &bintree{staticJsCryptoExtCjEcPatchMinJs, map[string]*bintree{}},
						"md5_min.js": &bintree{staticJsCryptoExtCjMd5_minJs, map[string]*bintree{}},
						"sha1_min.js": &bintree{staticJsCryptoExtCjSha1_minJs, map[string]*bintree{}},
						"sha256_min.js": &bintree{staticJsCryptoExtCjSha256_minJs, map[string]*bintree{}},
						"x64-core_min.js": &bintree{staticJsCryptoExtCjX64Core_minJs, map[string]*bintree{}},
					}},
					"ec-min.js": &bintree{staticJsCryptoExtEcMinJs, map[string]*bintree{}},
					"ec-patch-min.js": &bintree{staticJsCryptoExtEcPatchMinJs, map[string]*bintree{}},
					"jsbn-min.js": &bintree{staticJsCryptoExtJsbnMinJs, map[string]*bintree{}},
					"jsbn2-min.js": &bintree{staticJsCryptoExtJsbn2MinJs, map[string]*bintree{}},
					"prng4-min.js": &bintree{staticJsCryptoExtPrng4MinJs, map[string]*bintree{}},
					"rng-min.js": &bintree{staticJsCryptoExtRngMinJs, map[string]*bintree{}},
					"yahoo-min.js": &bintree{staticJsCryptoExtYahooMinJs, map[string]*bintree{}},
				}},
			}},
			"datetime.js": &bintree{staticJsDatetimeJs, map[string]*bintree{}},
			"demo": &bintree{nil, map[string]*bintree{
				"demo-datatable.js": &bintree{staticJsDemoDemoDatatableJs, map[string]*bintree{}},
				"demo-flot.js": &bintree{staticJsDemoDemoFlotJs, map[string]*bintree{}},
				"demo-forms.js": &bintree{staticJsDemoDemoFormsJs, map[string]*bintree{}},
				"demo-jqcloud.js": &bintree{staticJsDemoDemoJqcloudJs, map[string]*bintree{}},
				"demo-jqgrid.js": &bintree{staticJsDemoDemoJqgridJs, map[string]*bintree{}},
				"demo-nestable.js": &bintree{staticJsDemoDemoNestableJs, map[string]*bintree{}},
				"demo-panels.js": &bintree{staticJsDemoDemoPanelsJs, map[string]*bintree{}},
				"demo-rtl.js": &bintree{staticJsDemoDemoRtlJs, map[string]*bintree{}},
				"demo-search.js": &bintree{staticJsDemoDemoSearchJs, map[string]*bintree{}},
				"demo-sortable.js": &bintree{staticJsDemoDemoSortableJs, map[string]*bintree{}},
				"demo-upload.js": &bintree{staticJsDemoDemoUploadJs, map[string]*bintree{}},
				"demo-vector-map.js": &bintree{staticJsDemoDemoVectorMapJs, map[string]*bintree{}},
				"demo-wizard.js": &bintree{staticJsDemoDemoWizardJs, map[string]*bintree{}},
				"demo-xeditable.js": &bintree{staticJsDemoDemoXeditableJs, map[string]*bintree{}},
			}},
			"index.js": &bintree{staticJsIndexJs, map[string]*bintree{}},
			"nprogress.js": &bintree{staticJsNprogressJs, map[string]*bintree{}},
		}},
		"lang": &bintree{nil, map[string]*bintree{
			"1.ini": &bintree{staticLang1Ini, map[string]*bintree{}},
			"42.ini": &bintree{staticLang42Ini, map[string]*bintree{}},
			"en-us.all.json": &bintree{staticLangEnUsAllJson, map[string]*bintree{}},
			"locale_en-US.ini": &bintree{staticLangLocale_enUsIni, map[string]*bintree{}},
			"locale_ru-RU.ini": &bintree{staticLangLocale_ruRuIni, map[string]*bintree{}},
		}},
		"login.html": &bintree{staticLoginHtml, map[string]*bintree{}},
		"menu.html": &bintree{staticMenuHtml, map[string]*bintree{}},
		"modal_anonym.html": &bintree{staticModal_anonymHtml, map[string]*bintree{}},
		"pass.html": &bintree{staticPassHtml, map[string]*bintree{}},
		"psw.html": &bintree{staticPswHtml, map[string]*bintree{}},
		"request_citizen_status.html": &bintree{staticRequest_citizen_statusHtml, map[string]*bintree{}},
		"signatures.html": &bintree{staticSignaturesHtml, map[string]*bintree{}},
		"signatures_new.html": &bintree{staticSignatures_newHtml, map[string]*bintree{}},
		"test.html": &bintree{staticTestHtml, map[string]*bintree{}},
		"updating_blockchain.html": &bintree{staticUpdating_blockchainHtml, map[string]*bintree{}},
		"vendor": &bintree{nil, map[string]*bintree{
			"animate.css": &bintree{nil, map[string]*bintree{
				"animate.min.css": &bintree{staticVendorAnimateCssAnimateMinCss, map[string]*bintree{}},
			}},
			"bootstrap": &bintree{nil, map[string]*bintree{
				"dist": &bintree{nil, map[string]*bintree{
					"css": &bintree{nil, map[string]*bintree{
						"bootstrap.css": &bintree{staticVendorBootstrapDistCssBootstrapCss, map[string]*bintree{}},
					}},
					"js": &bintree{nil, map[string]*bintree{
						"bootstrap.js": &bintree{staticVendorBootstrapDistJsBootstrapJs, map[string]*bintree{}},
					}},
				}},
			}},
			"bootstrap-filestyle": &bintree{nil, map[string]*bintree{
				"src": &bintree{nil, map[string]*bintree{
					"bootstrap-filestyle.js": &bintree{staticVendorBootstrapFilestyleSrcBootstrapFilestyleJs, map[string]*bintree{}},
				}},
			}},
			"dataTables.fontAwesome": &bintree{nil, map[string]*bintree{
				"index.css": &bintree{staticVendorDatatablesFontawesomeIndexCss, map[string]*bintree{}},
			}},
			"datatables": &bintree{nil, map[string]*bintree{
				"media": &bintree{nil, map[string]*bintree{
					"css": &bintree{nil, map[string]*bintree{
						"dataTables.bootstrap.css": &bintree{staticVendorDatatablesMediaCssDatatablesBootstrapCss, map[string]*bintree{}},
					}},
					"images": &bintree{nil, map[string]*bintree{
						"Sorting icons.psd": &bintree{staticVendorDatatablesMediaImagesSortingIconsPsd, map[string]*bintree{}},
						"favicon.ico": &bintree{staticVendorDatatablesMediaImagesFaviconIco, map[string]*bintree{}},
						"sort_asc.png": &bintree{staticVendorDatatablesMediaImagesSort_ascPng, map[string]*bintree{}},
						"sort_asc_disabled.png": &bintree{staticVendorDatatablesMediaImagesSort_asc_disabledPng, map[string]*bintree{}},
						"sort_both.png": &bintree{staticVendorDatatablesMediaImagesSort_bothPng, map[string]*bintree{}},
						"sort_desc.png": &bintree{staticVendorDatatablesMediaImagesSort_descPng, map[string]*bintree{}},
						"sort_desc_disabled.png": &bintree{staticVendorDatatablesMediaImagesSort_desc_disabledPng, map[string]*bintree{}},
					}},
					"js": &bintree{nil, map[string]*bintree{
						"dataTables.bootstrap.js": &bintree{staticVendorDatatablesMediaJsDatatablesBootstrapJs, map[string]*bintree{}},
						"jquery.dataTables.min.js": &bintree{staticVendorDatatablesMediaJsJqueryDatatablesMinJs, map[string]*bintree{}},
					}},
				}},
			}},
			"datatables-colvis": &bintree{nil, map[string]*bintree{
				"css": &bintree{nil, map[string]*bintree{
					"dataTables.colVis.css": &bintree{staticVendorDatatablesColvisCssDatatablesColvisCss, map[string]*bintree{}},
				}},
				"js": &bintree{nil, map[string]*bintree{
					"dataTables.colVis.js": &bintree{staticVendorDatatablesColvisJsDatatablesColvisJs, map[string]*bintree{}},
				}},
			}},
			"eonasdan-bootstrap-datetimepicker": &bintree{nil, map[string]*bintree{
				"build": &bintree{nil, map[string]*bintree{
					"css": &bintree{nil, map[string]*bintree{
						"bootstrap-datetimepicker.min.css": &bintree{staticVendorEonasdanBootstrapDatetimepickerBuildCssBootstrapDatetimepickerMinCss, map[string]*bintree{}},
					}},
					"js": &bintree{nil, map[string]*bintree{
						"bootstrap-datetimepicker.min.js": &bintree{staticVendorEonasdanBootstrapDatetimepickerBuildJsBootstrapDatetimepickerMinJs, map[string]*bintree{}},
					}},
				}},
			}},
			"fontawesome": &bintree{nil, map[string]*bintree{
				"css": &bintree{nil, map[string]*bintree{
					"font-awesome.min.css": &bintree{staticVendorFontawesomeCssFontAwesomeMinCss, map[string]*bintree{}},
				}},
				"fonts": &bintree{nil, map[string]*bintree{
					"FontAwesome.otf": &bintree{staticVendorFontawesomeFontsFontawesomeOtf, map[string]*bintree{}},
					"fontawesome-webfont.eot": &bintree{staticVendorFontawesomeFontsFontawesomeWebfontEot, map[string]*bintree{}},
					"fontawesome-webfont.svg": &bintree{staticVendorFontawesomeFontsFontawesomeWebfontSvg, map[string]*bintree{}},
					"fontawesome-webfont.ttf": &bintree{staticVendorFontawesomeFontsFontawesomeWebfontTtf, map[string]*bintree{}},
					"fontawesome-webfont.woff": &bintree{staticVendorFontawesomeFontsFontawesomeWebfontWoff, map[string]*bintree{}},
					"fontawesome-webfont.woff2": &bintree{staticVendorFontawesomeFontsFontawesomeWebfontWoff2, map[string]*bintree{}},
				}},
			}},
			"jQuery-Storage-API": &bintree{nil, map[string]*bintree{
				"jquery.storageapi.js": &bintree{staticVendorJqueryStorageApiJqueryStorageapiJs, map[string]*bintree{}},
			}},
			"jquery": &bintree{nil, map[string]*bintree{
				"dist": &bintree{nil, map[string]*bintree{
					"jquery.js": &bintree{staticVendorJqueryDistJqueryJs, map[string]*bintree{}},
				}},
			}},
			"loaders.css": &bintree{nil, map[string]*bintree{
				"loaders.css": &bintree{staticVendorLoadersCssLoadersCss, map[string]*bintree{}},
			}},
			"modernizr": &bintree{nil, map[string]*bintree{
				"modernizr.custom.js": &bintree{staticVendorModernizrModernizrCustomJs, map[string]*bintree{}},
			}},
			"moment": &bintree{nil, map[string]*bintree{
				"min": &bintree{nil, map[string]*bintree{
					"moment-with-locales.min.js": &bintree{staticVendorMomentMinMomentWithLocalesMinJs, map[string]*bintree{}},
				}},
			}},
			"select2": &bintree{nil, map[string]*bintree{
				"dist": &bintree{nil, map[string]*bintree{
					"css": &bintree{nil, map[string]*bintree{
						"select2.css": &bintree{staticVendorSelect2DistCssSelect2Css, map[string]*bintree{}},
					}},
					"js": &bintree{nil, map[string]*bintree{
						"select2.js": &bintree{staticVendorSelect2DistJsSelect2Js, map[string]*bintree{}},
					}},
				}},
			}},
			"select2-bootstrap-theme": &bintree{nil, map[string]*bintree{
				"dist": &bintree{nil, map[string]*bintree{
					"select2-bootstrap.css": &bintree{staticVendorSelect2BootstrapThemeDistSelect2BootstrapCss, map[string]*bintree{}},
				}},
			}},
			"simple-line-icons": &bintree{nil, map[string]*bintree{
				"css": &bintree{nil, map[string]*bintree{
					"simple-line-icons.css": &bintree{staticVendorSimpleLineIconsCssSimpleLineIconsCss, map[string]*bintree{}},
				}},
				"fonts": &bintree{nil, map[string]*bintree{
					"Simple-Line-Icons.eot": &bintree{staticVendorSimpleLineIconsFontsSimpleLineIconsEot, map[string]*bintree{}},
					"Simple-Line-Icons.svg": &bintree{staticVendorSimpleLineIconsFontsSimpleLineIconsSvg, map[string]*bintree{}},
					"Simple-Line-Icons.ttf": &bintree{staticVendorSimpleLineIconsFontsSimpleLineIconsTtf, map[string]*bintree{}},
					"Simple-Line-Icons.woff": &bintree{staticVendorSimpleLineIconsFontsSimpleLineIconsWoff, map[string]*bintree{}},
					"Simple-Line-Icons.woff2": &bintree{staticVendorSimpleLineIconsFontsSimpleLineIconsWoff2, map[string]*bintree{}},
				}},
			}},
			"spinkit": &bintree{nil, map[string]*bintree{
				"css": &bintree{nil, map[string]*bintree{
					"spinkit.css": &bintree{staticVendorSpinkitCssSpinkitCss, map[string]*bintree{}},
				}},
			}},
			"sweetalert": &bintree{nil, map[string]*bintree{
				"dist": &bintree{nil, map[string]*bintree{
					"sweetalert.css": &bintree{staticVendorSweetalertDistSweetalertCss, map[string]*bintree{}},
					"sweetalert.min.js": &bintree{staticVendorSweetalertDistSweetalertMinJs, map[string]*bintree{}},
				}},
			}},
			"whirl": &bintree{nil, map[string]*bintree{
				"dist": &bintree{nil, map[string]*bintree{
					"whirl.css": &bintree{staticVendorWhirlDistWhirlCss, map[string]*bintree{}},
				}},
			}},
			"x-editable": &bintree{nil, map[string]*bintree{
				"dist": &bintree{nil, map[string]*bintree{
					"bootstrap3-editable": &bintree{nil, map[string]*bintree{
						"css": &bintree{nil, map[string]*bintree{
							"bootstrap-editable.css": &bintree{staticVendorXEditableDistBootstrap3EditableCssBootstrapEditableCss, map[string]*bintree{}},
						}},
						"img": &bintree{nil, map[string]*bintree{
							"clear.png": &bintree{staticVendorXEditableDistBootstrap3EditableImgClearPng, map[string]*bintree{}},
							"loading.gif": &bintree{staticVendorXEditableDistBootstrap3EditableImgLoadingGif, map[string]*bintree{}},
						}},
						"js": &bintree{nil, map[string]*bintree{
							"bootstrap-editable.min.js": &bintree{staticVendorXEditableDistBootstrap3EditableJsBootstrapEditableMinJs, map[string]*bintree{}},
						}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

