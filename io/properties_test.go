package io

import (
	"github.com/nano-projects/nanogo/test"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	data := `
# 组件服务上下文属性文件列表
context=

# 应用模式, DEV: 开发模式, PROD: 生产模式
context.mode=DEV

# 版本号
context.version=0.0.1-SNAPSHOT

# 服务根
context.root=/testgo

context.component-scan.base-package=org.nanoframework.nanogo.testgo.component
`

	test.DebugMode(t)
	if err := WriteFile("context.properties", data); err != nil {
		t.Error(err)
		return
	} else {
		t.Log("Initial context.properties file")
	}

	defer func() {
		if err := os.Remove("context.properties"); err != nil {
			t.Error(err)
			return
		} else {
			t.Log("Removed context.properties file")
		}
	}()

	p := &Properties{Path: "context.properties"}
	if err := p.Load(); err != nil {
		t.Error(err)
		return
	}

	t.Logf("Load context file Successful")
	t.Logf("values: %v", p.property)

	t.Logf("context: %v", p.property["context.component-scan.base-package"])
}
