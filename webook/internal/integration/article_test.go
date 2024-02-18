package integration

import (
	startup "GoBase/webook/internal/integration/startup"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ArticleTestSuite 测试套件
type ArticleTestSuite struct {
	suite.Suite
	server *gin.Engine
}

func (s *ArticleTestSuite) SetupSuite() {
	// 方式一：在所有测试执行之前，初始化一些内容
	s.server = startup.InitWebServer()

	// 方式二：方便对 server 进行定制
	//s.server = gin.Default()
	//artHdl := web.NewArticleHandler()
	//artHdl.RegisterRoutes(s.server)
}

func (s *ArticleTestSuite) TestPublish() {}

func (s *ArticleTestSuite) TestEdit() {
	t := s.T()
	testCases := []struct {
		name string
		// 集成测试准备数据
		before func(t *testing.T)
		// 集成测试验证数据
		after func(t *testing.T)

		// 预期的输入
		art Article
		// Http 响应码
		wantCode int64
		// 希望 http 响应，带上帖子的 id
		wantRes Result[int64]
	}{
		{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			reqBody, err := json.Marshal(tc.art)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/articles/edit", bytes.NewBuffer(reqBody))
			require.NoError(t, err)
			// 注意 json 数据格式
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			// 这是 http 请求进去，gin 框架的入口
			s.server.ServeHTTP(resp, req)

			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var webRes Result[int64]
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			require.NoError(t, err)
			assert.Equal(t, tc.wantRes, webRes)
			tc.after(t)
		})
	}
}

func (s *ArticleTestSuite) TestDemo() {
	s.T().Log("hello，这是测试套件")
}

func TestArticle(t *testing.T) {
	suite.Run(t, &ArticleTestSuite{})
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}
