/*
 * Copyright (c) 2024 Huawei Technologies Co., Ltd.
 * openFuyao is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package v1beta1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"k8s.io/client-go/rest"

	"marketplace-service/pkg/server/runtime"
)

func TestHelloHandler(t *testing.T) {
	// 创建一个新的 go-restful Container 并注册路由
	container := restful.NewContainer()
	marketplaceServiceWebService := runtime.GetMarketplaceWebService()
	config := &rest.Config{
		Host: "https://fake-kubernetes-server",
	}
	_ = BindMarketPlaceRoute(marketplaceServiceWebService, config)
	container.Add(marketplaceServiceWebService)

	// 创建一个模拟请求
	req, err := http.NewRequest("GET", "/hello/world", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	container.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("返回错误的状态码: 得到 %v 期望 %v", status, http.StatusNotFound)
	}
}

// TestNewHandler 构造函数测试（无Mock，仅测对象创建）
func TestNewHandler(t *testing.T) {
	// 传入空接口（仅测构造逻辑，不依赖具体实现）
	handler := newHandler(nil)
	// 断言对象创建成功
	if handler == nil {
		t.Error("NewHandler创建对象失败，期望非nil")
	}
	if handler.HelmHandler != nil {
		t.Error("NewHandler初始化HelmHandler错误，期望nil")
	}
}
