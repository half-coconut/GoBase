package integration

import (
	"GoBase/webook/internal/integration/startup"
	"GoBase/webook/internal/repository/dao"
	"GoBase/webook/internal/web"
	"GoBase/webook/ioc"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ArticleTestSuite 测试套件
type ArticleTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

func (s *ArticleTestSuite) SetupSuite() {
	// 方式一：在所有测试执行之前，初始化一些内容
	//s.server = ioc.InitWebServer()

	// 方式二：方便对 server 进行定制
	s.server = gin.Default()
	s.server.Use(func(ctx *gin.Context) {
		ctx.Set("claims", &web.UserClaims{
			Uid: 123,
		})
	})
	s.db = ioc.InitDB()
	artHdl := startup.InitArticleHandler()
	artHdl.RegisterRoutes(s.server)
}

// 每一个测试都会执行
func (s *ArticleTestSuite) TearDownTest() {
	// 清空所有数据，并且自增组件，恢复到 1
	s.db.Exec("TRUNCATE TABLE articles")
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
		wantCode int
		// 希望 http 响应，带上帖子的 id
		wantRes Result[int64]
	}{
		{
			name: "新建帖子，保存成功",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {
				// 验证数据库
				var art dao.Article
				err := s.db.Where("id=?", 1).First(&art).Error
				assert.NoError(t, err)
				assert.True(t, art.Ctime > 0)
				assert.True(t, art.Utime > 0)
				art.Ctime = 0
				art.Utime = 0
				assert.Equal(t, dao.Article{
					Id:       1,
					Title:    "我的标题",
					Content:  "我的内容",
					AuthorId: 123,
				}, art)
			},
			art: Article{
				Title:   "我的标题",
				Content: "我的内容",
			},
			wantCode: http.StatusOK,
			wantRes: Result[int64]{
				Data: 1,
				Msg:  "OK",
			},
		},
		{
			name: "修改已有的帖子，并保存",
			before: func(t *testing.T) {
				err := s.db.Create(dao.Article{
					Id:       2,
					Title:    "我的标题",
					Content:  "我的内容",
					AuthorId: 123,
					// time.Now(), 每次运行都不同，很难断言
					Ctime: 123,
					Utime: 234,
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				// 验证数据库
				var art dao.Article
				err := s.db.Where("id=?", 2).First(&art).Error
				assert.NoError(t, err)
				// 为了确保更新时间
				assert.True(t, art.Utime > 234)
				art.Utime = 0
				assert.Equal(t, dao.Article{
					Id:       2,
					Title:    "新的标题",
					Content:  "新的内容",
					AuthorId: 123,
					Ctime:    123,
				}, art)
			},
			art: Article{
				Id:      2,
				Title:   "新的标题",
				Content: "新的内容",
			},
			wantCode: http.StatusOK,
			wantRes: Result[int64]{
				Data: 2,
				Msg:  "OK",
			},
		},
		{
			name: "修改别人的帖子，不能更新成功",
			before: func(t *testing.T) {
				err := s.db.Create(dao.Article{
					Id:      3,
					Title:   "我的标题",
					Content: "我的内容",
					// 使用 789 用户创建老帖子
					AuthorId: 789,
					// time.Now(), 每次运行都不同，很难断言
					Ctime: 123,
					Utime: 234,
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				// 验证数据库
				var art dao.Article
				err := s.db.Where("id=?", 3).First(&art).Error
				assert.NoError(t, err)
				assert.Equal(t, dao.Article{
					Id:       3,
					Title:    "我的标题",
					Content:  "我的内容",
					AuthorId: 789,
					Ctime:    123,
					Utime:    234,
				}, art)
			},
			art: Article{
				Id:      3,
				Title:   "新的标题",
				Content: "新的内容",
			},
			wantCode: http.StatusOK,
			wantRes: Result[int64]{
				Code: 5,
				Data: 0,
				Msg:  "系统错误",
			},
		},
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
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Result[T any] struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}
