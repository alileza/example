package version

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	AppName   string
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion = runtime.Version()
)

var versionInfoTmpl = `
{{.program}}, version {{.version}} (branch: {{.branch}}, revision: {{.revision}})
  build user:       {{.buildUser}}
  build date:       {{.buildDate}}
  go version:       {{.goVersion}}
`

func Print() string {
	m := map[string]string{
		"program":   AppName,
		"version":   Version,
		"revision":  Revision,
		"branch":    Branch,
		"buildUser": BuildUser,
		"buildDate": BuildDate,
		"goVersion": GoVersion,
	}
	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

// NewCollector returns a collector which exports metrics about current version information.
func NewCollector(program string) *prometheus.GaugeVec {
	buildInfo := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: program,
			Name:      "build_info",
			Help: fmt.Sprintf(
				"A metric with a constant '1' value labeled by version, revision, branch, and goversion from which %s was built.",
				program,
			),
		},
		[]string{"version", "revision", "branch", "goversion"},
	)
	buildInfo.WithLabelValues(Version, Revision, Branch, GoVersion).Set(1)
	return buildInfo
}