package a2l

import (
	"fmt"
	"io"

	"github.com/flily/go-a2l/nodes"
)

const (
	ASAP2VersionMajor = 1
	ASAP2VersionMinor = 31
)

type Version struct {
	Major int
	Minor int
}

type Document struct {
	Information  string
	ASAP2Version Version
	Project      nodes.Node
}

func NewDocument() *Document {
	d := &Document{
		ASAP2Version: Version{
			Major: ASAP2VersionMajor,
			Minor: ASAP2VersionMinor,
		},
	}

	return d
}

func (d *Document) writeLine(out io.Writer, line string) (int, error) {
	n, err := out.Write([]byte(line + "\n"))
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (d *Document) writeBlock(out io.Writer, line string) (int, error) {
	return d.writeLine(out, line+"\n")
}

func (d *Document) WriteOn(out io.Writer) (int, error) {
	var err error
	result, n := 0, 0

	if len(d.Information) > 0 {
		n, err = d.writeBlock(out, nodes.StringCommentMultiLine(d.Information))
		if err != nil {
			return 0, err
		}

		result += n
	}

	version := fmt.Sprintf("ASAP2_VERSION  %d %d    /* Version %d.%d */",
		d.ASAP2Version.Major, d.ASAP2Version.Minor,
		d.ASAP2Version.Major, d.ASAP2Version.Minor)
	n, _ = d.writeBlock(out, version)
	result += n

	return result, nil
}
