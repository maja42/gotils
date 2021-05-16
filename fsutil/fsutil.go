package fsutil

import "os"

// PathExists checks if there is a file or directory at the given location.
func PathExists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

// DirExists checks if there is a directory at the given location.
// Returns false if the path does not exist, or if it points to a file.
func DirExists(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// FileExists checks if there is a file at the given location.
// Returns false if the path does not exist, or if it points to a directory.
func FileExists(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
