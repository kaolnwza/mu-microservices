package helper

import (
	"fmt"
	"strings"
)

func PointToLatLong(geom string) (string, string) {
	spl := []string{"", ""}
	if geom != "" {
		geom = strings.Trim(geom[1:], "]")
		spl = strings.Split(geom, ",")
	}

	return string(spl[0]), string(spl[1])
}

func LatLongToPoint(lat string, long string) string {
	return fmt.Sprintf("POINT(%s %s)", lat, long)
}
