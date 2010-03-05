// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gobbler

import (
	. "gospec"
	"container/vector"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)


func TemporaryDirectorySpec(c Context) {
	
	c.Specify("When a temporary directory is created", func() {
		dir, _ := CreateTempDir()
		defer dir.Dispose()
		
		c.Specify("it exists", func() {
			c.Expect(dir, Satisfies, dir.Exists())
		})
		c.Specify("it is empty", func() {
			files := dir.AllFiles()
			c.Expect(files, Satisfies, len(files) == 0)
		})
		c.Specify("After disposing, the directory no longer exists", func() {
			dir.Dispose()
			c.Expect(dir, Not(Satisfies), dir.Exists())
		})
	})
	c.Specify("When multiple temporary directories are created in parallel", func() {
		dir1, _ := CreateTempDir()
		dir2, _ := CreateTempDir()
		defer dir1.Dispose()
		defer dir2.Dispose()
		
		c.Specify("they all are different directories", func() {
			c.Expect(dir1.Path(), Not(Equals), dir2.Path())
		})
	})
	
	c.Specify("When files are put into the directory", func() {
		dir, _ := CreateTempDir()
		defer dir.Dispose()
		
		dir.WriteFile("file.txt", "file contents")
		
		c.Specify("then the files are in the directory", func() {
			c.Expect(dir.AllFiles(), Contains, "file.txt")
		})
		c.Specify("and the files have the specified content", func() {
			c.Expect(dir.ReadFile("file.txt"), Equals, "file contents")
		})
		c.Specify("After disposing, the files no longer exist", func() {
			dir.Dispose()
			c.Expect(dir.ReadFile("file.txt"), Equals, "")
		})
	})
	c.Specify("When files are put into subdirectories", func() {
		dir, _ := CreateTempDir()
		defer dir.Dispose()
		
		dir.WriteFile("subdir/file.txt", "file contents")
		
		c.Specify("then the subdirectory is created", func() {
			c.Expect(dir.AllFiles(), Contains, "subdir")
		})
		c.Specify("and the files are in the subdirectory", func() {
			c.Expect(dir.AllFiles(), Contains, "subdir/file.txt")
		})
	})
}


// Implementations of the test utilities

var tempDirCounter = 0

const (
	defaultDirPerm  = 0777
	defaultFilePerm = 0666
)

func CreateTempDir() (dir *TempDir, err os.Error) {
	tempDirCounter++
	path := fmt.Sprintf("tempdir_%v_%v", time.Nanoseconds(), tempDirCounter)
	err = os.Mkdir(path, defaultDirPerm)
	if err != nil {
		return
	}
	dir = &TempDir{path}
	return
}

type TempDir struct {
	path string
}

func (this *TempDir) Exists() bool {
	if this == nil {
		return false
	}
	stat, err := os.Stat(this.path)
	if err != nil {
		return false
	}
	return stat.IsDirectory()
}

func (this *TempDir) Dispose() {
	if this == nil {
		return
	}
	os.RemoveAll(this.path)
}

func (this *TempDir) AllFiles() []string {
	collector := CollectFilesInDir(this.path)
	path.Walk(this.path, collector, nil)
	return collector.Files()
}

func (this *TempDir) WriteFile(name string, content string) {
	target := this.PathTo(name)
	createParentDirs(target)
	ioutil.WriteFile(target, []byte(content), defaultFilePerm)
}

func createParentDirs(file string) {
	parent, _ := path.Split(file)
	os.MkdirAll(parent, defaultDirPerm)
}

func (this *TempDir) ReadFile(name string) string {
	bytes, err := ioutil.ReadFile(this.PathTo(name))
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (this *TempDir) Path() string {
	return this.path
}

func (this *TempDir) PathTo(name string) string {
	return path.Join(this.path, name)
}



func CollectFilesInDir(rootDir string) *FileCollector {
	return &FileCollector{rootDir, new(vector.StringVector)}
}

type FileCollector struct {
	rootDir string
	paths *vector.StringVector
}

func (this *FileCollector) Files() []string {
	return *this.paths
}

func (this *FileCollector) VisitDir(path string, d *os.Dir) bool {
	this.addPath(path)
	return true
}

func (this *FileCollector) VisitFile(path string, d *os.Dir) {
	this.addPath(path)
}

func (this *FileCollector) addPath(path string) {
	path = this.relativeToRoot(path)
	if path != "" {
		this.paths.Push(path)
	}
}

func (this *FileCollector) relativeToRoot(path string) string {
	return removePrefix(this.rootDir, path)
}

func removePrefix(prefix string, filePath string) string {
	filePath = path.Clean(filePath)
	if filePath == prefix || filePath == "" {
		return ""
	}
	dir, file := path.Split(filePath)
	return path.Join(removePrefix(prefix, dir), file)
}

