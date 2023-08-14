package reader

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

func StdinToCsv(re *regexp.Regexp) (record [][]string, err error) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, errors.Wrap(err, "read fail")
	}

	// 첫번째 줄이 header 인지 판단.
	b = bytes.Trim(b, "\n")
	header := false
	if re.Match(b) {
		header = true
	}

	r := csv.NewReader(bytes.NewReader(b))
	if header {
		_, _ = r.Read()
	}

	records, err := r.ReadAll()
	return records, errors.Wrap(err, "csv read fail")
}
