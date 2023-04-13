package katex

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"runtime"

	"github.com/lithdew/quickjs"
)

//go:embed katex.min.js
var code string

func Render(w io.Writer, src []byte, display bool) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	runtime := quickjs.NewRuntime()
	defer runtime.Free()

	context := runtime.NewContext()
	defer context.Free()

	globals := context.Globals()

	result, err := context.Eval(code)
	if err != nil {
		return err
	}
	defer result.Free()

	globals.Set("_EqSrc3120", context.String(string(src)))
	if display {
		result, err = context.Eval(`katex.renderToString(_EqSrc3120, { "displayMode": true, "output": "html"} )`)
	} else {
		result, err = context.Eval(`katex.renderToString(_EqSrc3120, { "output": "html"})`)
	}
	defer result.Free()

	_, err = io.WriteString(w, result.String())
	return err
}

func RenderWithOptions(w io.Writer, src []byte, options *Options, extra map[string]interface{}) error {
	m := make(map[string]interface{})
	if options != nil {
		data, _ := json.Marshal(options)
		err := json.Unmarshal(data, &m)
		if err != nil {
			return err
		}
	}
	if extra != nil {
		for k, v := range extra {
			m[k] = v
		}
	}
	if len(m) == 0 {
		return Render(w, src, false)
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	runtime := quickjs.NewRuntime()
	defer runtime.Free()

	context := runtime.NewContext()
	defer context.Free()

	globals := context.Globals()

	result, err := context.Eval(code)
	if err != nil {
		return err
	}
	defer result.Free()

	globals.Set("_EqSrc3120", context.String(string(src)))

	bs, err := json.Marshal(m)
	if err != nil {
		return err
	}
	result, err = context.Eval(fmt.Sprintf("katex.renderToString(_EqSrc3120, %s)", string(bs)))
	defer result.Free()

	_, err = io.WriteString(w, result.String())
	return err
}
