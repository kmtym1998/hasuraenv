package services

import "os"

func ReplaceSymlink(dest, src string) error {
	os.Remove(src)
	return os.Symlink(dest, src)
}
