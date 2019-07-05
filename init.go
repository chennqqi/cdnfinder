package cdnfinder

//go:generate go-bindata -pkg resource -o resource/assets.go -nomemcopy assets/...
import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/turbobytes/cdnfinder/resource"
)

var (
	cdnmatches       [][]string
	resourcefinderjs = os.TempDir() + "/cdnfinder-resourcefinder.js"
	phantomjsbin     = ""
	initialized      = false
)

//Load cname chain
func populatecnamechain() {
	raw, err := resource.Asset("assets/cnamechain.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(raw, &cdnmatches)
	if err != nil {
		log.Fatal(err)
	}
}

func assettofile(key, dst string) {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		raw, err := resource.Asset(key)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(dst, raw, 0644)
		if err != nil {
			log.Fatal("WriteFile", err)
		}
	} else {
		log.Println(dst, " exists already...")
	}
}

//Make sure resourcefinder.js is present somewhere and get a handle to it
func ensureresourcefinder() {
	assettofile("assets/resourcefinder.js", resourcefinderjs)
}

func notsupported() {
	log.Fatal("OS/ARCH not supported:", runtime.GOOS, runtime.GOARCH)
}

func installphantomjs(url, dst, fname string) {
	log.Println("downloading", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if strings.HasSuffix(url, ".tar.bz2") {
		extracttar(resp.Body, fname, dst)
	}
	if strings.HasSuffix(url, ".zip") {
		extractzip(resp.Body, fname, dst)
	}
}

func extractzip(r io.Reader, fname, dst string) {

	tmpfile, err := ioutil.TempFile("", "phantomzip")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	_, err = io.Copy(tmpfile, r)
	if err != nil {
		log.Fatal(err)
	}
	tmpfile.Close()

	zrd, err := zip.OpenReader(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range zrd.File {
		info := f.FileInfo()
		if !info.IsDir() {
			if strings.HasSuffix(f.Name, fname) {
				log.Println(f.Name)
				file, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				rc, err := f.Open()
				if err != nil {
					log.Fatal(err)
				}
				_, err = io.Copy(file, rc)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func extracttar(r io.Reader, fname, dst string) {
	trd := tar.NewReader(bzip2.NewReader(r))
	for {
		header, err := trd.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		info := header.FileInfo()
		if !info.IsDir() {
			if strings.HasSuffix(header.Name, fname) {
				file, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				_, err = io.Copy(file, trd)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

//Install phantomjs
func loadphantomjs() {
	if phantomjsbin != "" {
		_, err := os.Stat(phantomjsbin)
		if err == nil {
			return
		}
	}

	fname := ""
	url := ""
	switch runtime.GOOS {
	case "windows":
		fname = "phantomjs.exe"
		url = "https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-windows.zip"
	case "linux":
		fname = "phantomjs"
		switch runtime.GOARCH {
		case "amd64":
			url = "https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2"
		case "386":
			url = "https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-i686.tar.bz2"
		default:
			notsupported()
		}
	case "darwin":
		fname = "phantomjs"
		url = "https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-macosx.zip"
	default:
		notsupported()
	}
	if phantomjsbin == "" {
		phantomjsbin = os.TempDir() + "/cdnfinder_2.1.1_" + fname
	}
	if _, err := os.Stat(phantomjsbin); os.IsNotExist(err) {
		installphantomjs(url, phantomjsbin, fname)
	} else {
		log.Println(phantomjsbin, "is already installed")
	}
}

// Init must be called once after the main package runs flag.Parse().
// The reason for doing this manually is to allow main package to set up its flags.
func Init(_phantomjsbin, _resourcefinderjs string) {
	if initialized {
		return
	}
	if _phantomjsbin != "" {
		phantomjsbin = _phantomjsbin
	}
	if _resourcefinderjs != "" {
		resourcefinderjs = _resourcefinderjs
	}

	initialized = true
	populatecnamechain()
	ensureresourcefinder()
	loadphantomjs()
}
