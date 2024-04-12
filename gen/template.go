package gen

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func changeBackendTemplate(name string, destDir string, backendConfig *backendConfig, srcDir string) error {
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}
	firstDir := true
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, path[len(srcDir):])

		if info.IsDir() {
			if firstDir {
				firstDir = false
				return nil
			}
			err := os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			err := copyBackendFile(path, name, destPath, backendConfig)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func copyBackendFile(path, name, destFile string, backendConfig *backendConfig) error {
	f, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(content), "bo-new-app", name)

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	t := template.Must(template.New("file").Parse(newContent))
	err = t.ExecuteTemplate(dest, "file", backendConfig)
	if err != nil {
		return err
	}
	return nil
}

func changeFrontendTemplate(name string, destDir string, frontendConfig *frontendConfig, srcDir string) error {
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}
	firstDir := true
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, path[len(srcDir):])

		if info.IsDir() {
			if firstDir {
				firstDir = false
				return nil
			}
			err := os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			err := copyFrontendFile(path, name, destPath, frontendConfig)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func copyFrontendFile(path, name, destFile string, frontendConfig *frontendConfig) error {
	f, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(content), "bo-new-app", name)

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	t := template.Must(template.New("file").Parse(newContent))
	err = t.ExecuteTemplate(dest, "file", frontendConfig)
	if err != nil {
		return err
	}
	return nil
}

func changeCITemplate(name string, destDir string, ciConfig *ciConfig, srcDir string) error {
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}
	firstDir := true
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, path[len(srcDir):])

		if info.IsDir() {
			if firstDir {
				firstDir = false
				return nil
			}
			err := os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			err := copyCIFile(path, name, destPath, ciConfig)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func copyCIFile(path, name, destFile string, ciConfig *ciConfig) error {
	f, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	t := template.Must(template.New("file").Parse(string(content)))
	err = t.ExecuteTemplate(dest, "file", ciConfig)
	if err != nil {
		return err
	}
	return nil
}
