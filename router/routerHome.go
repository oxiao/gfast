package router

import (
	"gfast/app/controller/home"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//前端路由处理
func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		cms := new(home.Index)
		cms.TemplatePath = "template/cms"
		group.GET("/cms", cms.Index)
		group.GET("/cms/list/:cateId/*page/*keyWords", cms.List)
		group.GET("/cms/show/:cateIds/:newsId", cms.Show)
		group.ALL("/cms/search/*page/*keyWords", cms.Search)

		site := new(home.Index)
		site.TemplatePath = "template/site"
		group.GET("/site", site.Index)
		group.GET("/site/list/:cateId/*page/*keyWords", site.List)
		group.GET("/site/show/:cateIds/:newsId", site.Show)
		group.ALL("/site/search/*page/*keyWords", site.Search)

	})
}
