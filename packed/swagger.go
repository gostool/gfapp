package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
