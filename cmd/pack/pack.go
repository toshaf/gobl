package pack

import (
    "archive/tar"
    "fmt"
    "os"
    "io/ioutil"
    "io"
)

const (
    GO_LIB_EXT = ".a"
)

func Pack(pkg string, w io.Writer) error {
    pkgdir := os.ExpandEnv("$GOPATH/pkg")
    files, err := ioutil.ReadDir(pkgdir)
    if err != nil {
        return err
    }

    writer := tar.NewWriter(w)
    defer func() {
        if err := writer.Close(); err != nil {
            panic(err)
        }
    }()

    for _,f := range files {
        if f.IsDir() {
            // these should be of the form $GOOS_$GOARCH
            name := fmt.Sprintf("%s/%s/%s%s", "pkg", f.Name(), pkg, GO_LIB_EXT)
            file, err := os.Open(os.ExpandEnv("$GOPATH/") + name)
            if err == nil {
                defer file.Close()

                body, err := ioutil.ReadAll(file)
                if err != nil {
                    panic(err)
                }

                header := tar.Header{
                    Name: name,
                    Mode: 0444,
                    Size: int64(len(body)),
                }

                writer.WriteHeader(&header)
                writer.Write(body)
            }
        }
    }

    return nil
}
