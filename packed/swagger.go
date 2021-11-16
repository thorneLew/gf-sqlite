package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDQickPZkACIgycDMXlienpqUX6UFovqzg/LzSElYGxxXVS4qpnQQEBAQFaev4+/v66QZuCAowuGQdcauxIe7b0qXakpsOVBtdJnz97CU32dmjamnAxasWlpiOTepI6CnQVGBj+/w/wZucw65yw0IGBgSGEgYEB5hQGDKewI5wCtx2kG1lNgDcjkwgzwivIJoO8AgP/G0EkXo8hjMLuFAgQYPjv2I0wCkk3KxtInomBiaGTgYHhJFg1IAAA//+bVSk5aAEAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
