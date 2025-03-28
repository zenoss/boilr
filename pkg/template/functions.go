package template

import (
	"fmt"
	"os"
	"os/user"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	englishTitleCaser = cases.Title(language.English)

	// FuncMap contains the functions exposed to templating engine.
	FuncMap = template.FuncMap{
		// TODO confirmation prompt
		// TODO value prompt
		// TODO encoding utilities (e.g. toBinary)
		// TODO GET, POST utilities
		// TODO Hostname(Also accesible through $HOSTNAME), interface IP addr, etc.
		// TODO add validate for custom regex and expose validate package
		"env":      os.Getenv,
		"time":     CurrentTimeInFmt,
		"hostname": func() string { return os.Getenv("HOSTNAME") },
		"username": func() string {
			t, err := user.Current()
			if err != nil {
				return "Unknown"
			}

			return t.Name
		},
		"toBinary": func(s string) string {
			n, err := strconv.Atoi(s)
			if err != nil {
				return s
			}

			return fmt.Sprintf("%b", n)
		},

		"formatFilesize": func(value any) string {
			var size float64

			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				size = float64(v.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				size = float64(v.Uint())
			case reflect.Float32, reflect.Float64:
				size = v.Float()
			default:
				return ""
			}

			var KB float64 = 1 << 10
			var MB float64 = 1 << 20
			var GB float64 = 1 << 30
			var TB float64 = 1 << 40
			var PB float64 = 1 << 50

			filesizeFormat := func(filesize float64, suffix string) string {
				return strings.ReplaceAll(fmt.Sprintf("%.1f %s", filesize, suffix), ".0", "")
			}

			var result string
			if size < KB {
				result = filesizeFormat(size, "bytes")
			} else if size < MB {
				result = filesizeFormat(size/KB, "KB")
			} else if size < GB {
				result = filesizeFormat(size/MB, "MB")
			} else if size < TB {
				result = filesizeFormat(size/GB, "GB")
			} else if size < PB {
				result = filesizeFormat(size/TB, "TB")
			} else {
				result = filesizeFormat(size/PB, "PB")
			}

			return result
		},

		// String utilities
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
		"toTitle": strings.ToTitle,
		"title":   englishTitleCaser.String,

		"trimSpace":  strings.TrimSpace,
		"trimPrefix": strings.TrimPrefix,
		"trimSuffix": strings.TrimSuffix,

		"repeat":  strings.Repeat,
		"replace": strings.Replace,

		// camel splits a string by a separator and recombines while capitalizing
		// the first character in each part
		"camel": func(value, sep string) (result string) {
			parts := strings.Split(value, sep)
			result = parts[0]
			for _, part := range parts[1:] {
				result += englishTitleCaser.String(part)
			}
			return
		},
	}

	// Options contain the default options for the template execution.
	Options = []string{
		// TODO ignore a field if no value is found instead of writing <no value>
		"missingkey=invalid",
	}
)

// CurrentTimeInFmt returns the current time in the given format.
// See time.Time.Format for more details on the format string.
func CurrentTimeInFmt(fmt string) string {
	t := time.Now()

	return t.Format(fmt)
}
