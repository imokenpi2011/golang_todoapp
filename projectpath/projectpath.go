package projectpath

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// プロジェクトのrootフォルダを設定する
	Root = filepath.Join(filepath.Dir(b), "..")
)
