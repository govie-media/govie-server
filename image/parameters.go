package image

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var PARAMETER_REGEX = regexp.MustCompile("([a-zA-Z]{1,2})(\\d+(?:\\.\\d+)?)(?:-+)?")

const (
	IMAGE_QUALITY          int = 76
	HIGH_DPI_IMAGE_QUALITY int = 56
)

type Parameters struct {
	maxWidth     int
	maxHeight    int
	x            int
	y            int
	width        int
	height       int
	rotate       int
	pixelDensity int
	quality      int
	computed     string
	original     bool
}

func (params *Parameters) Init(value string) {
	// Default the params
	params.pixelDensity = 1
	params.quality = IMAGE_QUALITY
	params.original = false

	// Parse the parameters from the query string
	if len(value) > 0 {
		match := PARAMETER_REGEX.FindAllStringSubmatch(value, -1)

		for i := 0; i < len(match); i++ {
			param := match[i]

			if len(param) == 3 {
				key := strings.ToLower(param[1])
				value, _ := strconv.Atoi(param[2])

				switch key {
				case "mw":
					params.maxWidth = value
				case "mh":
					params.maxHeight = value
				case "x":
					params.x = value
				case "y":
					params.y = value
				case "w":
					params.width = value
				case "h":
					params.height = value
				case "ro":
					params.rotate = value
				case "pd":
					params.pixelDensity = value
				case "q":
					params.quality = value
				case "o":
					params.original = true
				}
			}
		}

		if params.pixelDensity > 1 {
			params.quality = HIGH_DPI_IMAGE_QUALITY
		}

		params.Process()
	}
}

func (params *Parameters) Process() {
	var computed bytes.Buffer

	if params.x != 0 {
		computed.WriteString("-x")
		computed.WriteString(strconv.Itoa(params.x))
	}

	if params.y != 0 {
		computed.WriteString("-y")
		computed.WriteString(strconv.Itoa(params.y))
	}

	if params.width != 0 {
		computed.WriteString("-w")
		computed.WriteString(strconv.Itoa(params.width))
	}

	if params.height != 0 {
		computed.WriteString("-h")
		computed.WriteString(strconv.Itoa(params.height))
	}

	if params.maxWidth != 0 {
		computed.WriteString("-mw")
		computed.WriteString(strconv.Itoa(params.maxWidth))
	}

	if params.maxHeight != 0 {
		computed.WriteString("-mh")
		computed.WriteString(strconv.Itoa(params.maxHeight))
	}

	if params.pixelDensity > 1 {
		computed.WriteString("-pd")
		computed.WriteString(strconv.Itoa(params.pixelDensity))
	}

	if params.rotate > 0 && params.rotate < 360 {
		computed.WriteString("-ro")
		computed.WriteString(strconv.Itoa(params.rotate))
	}

	if params.quality > 0 && params.quality <= 100 && (params.quality != IMAGE_QUALITY && (params.pixelDensity > 1 && params.quality != HIGH_DPI_IMAGE_QUALITY)) {
		computed.WriteString("-q")
		computed.WriteString(strconv.Itoa(params.quality))
	}

	// Remove the first dash character
	if computed.Len() > 0 {
		params.computed = computed.String()[1:computed.Len()]
	} else {
		params.computed = ""
	}
}

func (params *Parameters) isRotating() bool {

	if params.rotate > 0 && params.rotate < 360 {
		return true
	}

	return false
}

func (params *Parameters) isCropping() bool {

	if params.width > 0 && params.height > 0 && params.x >= 0 && params.y >= 0 {
		return true
	}

	return false
}

func (params *Parameters) isResizing() bool {
	if params.maxWidth > 0 || params.maxHeight > 0 {
		return true
	}

	return false
}

func (params *Parameters) withinCropBounds(width, height int) bool {
	if width <= params.width+params.x && height <= params.height+params.y {
		return true
	}

	return true
}

func (params *Parameters) withinResizeBounds(width, height int) bool {
	if params.maxWidth > width || params.maxHeight > height {
		return false
	}

	return true
}
