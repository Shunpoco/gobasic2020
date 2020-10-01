package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
)

const defaultEditor = "emacs"

func openFileInEditor(filename string) error {
	editor := defaultEditor

	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func FetchInputFromEditor() ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	defer os.Remove(filename)

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = openFileInEditor(filename); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
