// To execute Go code, please declare a func main() in a package "main"

/*
Scylla Pair Coding Task A
-------------------------

The task is to implement a command that is going to be poor-man-clone of the `wget` command. It should take 1 argument - the url of a file to download, and save it to local disk.

Each 1 second it should log current progress so far in bytes.

Example usage:

```
$ gowget http://releases.ubuntu.com/18.04.3/ubuntu-18.04.3-desktop-amd64.iso
Downloaded 123 bytes ...
Downloaded 456 bytes ...
Downloaded 789 bytes ...
...
```
*/

package main

import (
  "log"
  "os"
  "path"
  "net/http"
  "io"
  "time"
)

/*
type Reader interface {
  Read(p []byte) (n int, err error)
}
*/


type ProgressReader struct {
  r io.Reader
  totalBytes int
  exit bool
}

func (pr *ProgressReader) periodicLog()  {
  fmt.Printf("Downloaded %d ...\n", totalBytes)
}


func (pr *ProgressReader) Read(p []byte) (n int, err error) {
  n, err = pr.r.Read(p)
  pr.totalBytes += n
  p.exit = err == IOError
  return 
}


func InitProgressReader(r Reader) Reader {
  pr := &ProgressReader{r: r}
  ticker := time.NewTicker(1 * time.Second)
  go func() {
    for {
      <- ticker.C
      pr.periodicLog()
      if pr.exit {
        return
      }
    }
  } ()

  return pr
}


func wget(url string) {
  basename := path.Base(url)
  // open stream for reading
  response, err := http.Get(url)
  if err != nil {
    log.Printf("Failed to read %s %v", url, err)
    return
  }
  
  // open sream frow writeing
  outFile, err := os.Create(basename)
  if err != nil {
    log.Printf("Failed to open %s %v", basename, err)
    return
  }
  
  // read from stream block by block 
  // write the block to the output stream (file on the disk)
  progressReader := InitPorgressReader(response.Body)
  // I want progerss logs like 
  count, err := io.Copy(outFile, progressReader)
  
  if err != nil {
    log.Printf("Failed to move data between %s and %s %v", url, basename, err)
    return
  }
  log.Printf("I got %d byet from %s ", count, url)
}

// test
func main() {
  // https://gist.githubusercontent.com/rjeczalik/b3f0580912a8fea86032940079af8672/raw/f63bdb59c892f31911b9ff34bcf93e5391263930/scylla-pair-coding-task.md
  wget("https://repo.almalinux.org/almalinux/8/BaseOS/aarch64/os/Packages/time-1.9-3.el8.aarch64.rpm")
}
