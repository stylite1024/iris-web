package test

import (
	"iris-web/internal/router"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

// 测试文件以_test.go为后缀名
// 测试函数前缀为Test

// 测试
// go test

// 打印详细过程
// go test -v

// 只执行某一个测试方法
// go test -v -run TestAdd

// 根据testing.Short()标志，使用-short参数跳过测试
// go test -v -short

func TestIris(t *testing.T) {
	// parallel test 并行测试
	t.Parallel()
	if testing.Short() {
		t.Skip("skip this test case")
	}

	// case list
	type args struct {
		username string
	}
	type want struct {
		w int
	}
	testCases := []struct {
		name string // case name
		args args   // case args
		want want   // case want
	}{
		{"testCase1", args{"abc"}, want{200}},
		{"testCase2", args{"def"}, want{300}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := iris.New()
			router.InitRouter(app)
			e := httptest.New(t, app)
			e.GET("/api/v1/test").Expect().Status(tc.want.w)
		})
	}
}
