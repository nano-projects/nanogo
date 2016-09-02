package io

import (
	"bufio"
	"github.com/nano-projects/nanogo/log"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

type Properties struct {
	Path     string
	property map[string]string
}

func (p *Properties) Load() error {
	file, err := os.Open(p.Path)
	if err != nil {
		return err
	}

	defer file.Close()
	buf := bufio.NewReader(file)
	p.property = make(map[string]string)
	for {
		bytes, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		line := string(bytes)
		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
			return errors.New("Read file format error, cannot be start with ' ' or '\\t' ")
		}

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		eqIdx := strings.Index(line, "=")
		if eqIdx > 0 {
			key := line[:eqIdx]
			value := line[eqIdx+1:]
			log.Logger.Debugf("Read property: %v=%v", key, value)
			p.property[key] = value
		}
	}

	return nil
}
