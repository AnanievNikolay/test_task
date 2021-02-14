package file

import "path/filepath"

//NewPath ...
func NewPath(_dir, _path string) *Path {
	return &Path{
		path: _path,
		dir:  _dir,
	}
}

//Path ...
type Path struct {
	dir  string
	path string
}

//Abs ...
func (f *Path) Abs() string {
	dir := filepath.Dir(f.dir)
	absPath, _ := filepath.Abs(dir + f.path)
	return absPath
}
