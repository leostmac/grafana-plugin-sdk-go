package numeric

import "github.com/leostmac/grafana-plugin-sdk-go/data"

func emptyFrameWithTypeMD(t data.FrameType) *data.Frame {
	return data.NewFrame("").SetMeta(&data.FrameMeta{Type: t})
}
