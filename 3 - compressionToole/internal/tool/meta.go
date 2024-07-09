package tool

import (
	"io"
	"os"
)

type Meta struct {
	file        io.Reader
	newFileName string
}

func (m Meta) Do(command Command) error {
	newData, err := command.Execute(m.file)
	if err != nil {
		return err
	}

	resultFile, err := os.Create(m.newFileName)
	if err != nil {
		return err
	}

	result, err := io.ReadAll(newData)
	if err != nil {
		return err
	}

	if _, err = resultFile.Write(result); err != nil {
		return err
	}

	return nil
}
