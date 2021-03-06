package kubebench

import (
	"encoding/json"
	"io"

	starboard "github.com/aquasecurity/starboard/pkg/apis/aquasecurity/v1alpha1"
)

type Converter interface {
	Convert(reader io.Reader) (report starboard.CISKubeBenchOutput, err error)
}

var DefaultConverter Converter = &converter{}

type converter struct {
}

func (c *converter) Convert(reader io.Reader) (report starboard.CISKubeBenchOutput, err error) {
	decoder := json.NewDecoder(reader)
	report = starboard.CISKubeBenchOutput{
		Scanner: starboard.Scanner{
			Name:    "kube-bench",
			Vendor:  "Aqua Security",
			Version: "latest",
		},
		Sections: []starboard.CISKubeBenchSection{},
	}

	for {
		var section starboard.CISKubeBenchSection
		de := decoder.Decode(&section)
		if de == io.EOF {
			break
		}
		if de != nil {
			err = de
			break
		}
		report.Sections = append(report.Sections, section)
	}
	return
}
