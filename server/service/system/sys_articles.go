package system

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 文章相关api
type ArticleService struct{}

var articleService = new(ArticleService)

// 新增文章
func (articleService *ArticleService) AddArticle(article system.Article) (err error) {
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()

	if err := global.GVA_DB.Table("articles").Create(article).Error; err != nil {
		return err
	}
	return nil

}

// 编辑文章
func (articleService *ArticleService) EditArticle(article system.Article) (err error) {
	var articleStru system.Article
	if err := global.GVA_DB.Table("articles").Model(&articleStru).Where("article_id = ?", article.Id).Updates(article).Error; err != nil {
		return err
	}
	return nil
}

// 获取文章列表
func (articleService *ArticleService) GetArticlesList(page system.Page) (err error, params map[string]interface{}) {
	params = make(map[string]interface{})
	var articles []system.Article
	var total int64
	// articles := params["articles"].([]system.Article)
	// total := params["total"].(int64)
	offset := (page.Current - 1) * page.Size
	if err := global.GVA_DB.Table("articles").Model(&system.Article{}).Count(&total).Error; err != nil {
		return err, params
	}
	if err := global.GVA_DB.Table("articles").Order("article_id").Limit(page.Size).Offset(offset).Find(&articles).Error; err != nil {
		return err, params
	}
	params["articles"] = articles
	params["total"] = total
	// sqlWords := fmt.Sprintf("Select * From articles")
	// err = global.GVA_DB.Raw(sqlWords).Scan(&articles).Error
	return nil, params
}

// 删除文章
func (articleService *ArticleService) DelArticleById(id int) (err error) {
	var article system.Article
	if err := global.GVA_DB.Table("articles").Model(&article).Where("article_id = ?", id).Delete(&article).Error; err != nil {
		return err
	}
	return nil
}

// 解析markdown文件为article
func MD2Article(file *multipart.FileHeader) (err error, article system.Article) {
	fileContent, err := file.Open()
	if err != nil {
		return
	}
	defer fileContent.Close()
	content, err := io.ReadAll(fileContent)
	if err != nil {
		return
	}
	// 将 Markdown 内容解析为 HTML
	// htmlContent := blackfriday.Run(content)

	article.ArticleContent = string(content)
	article.ArticleTitle = file.Filename
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()

	return nil, article

	// htmlContent := blackfriday.Run(file)
}

// 上传markdown文件解析
func (articleService *ArticleService) UploadArticleMD(file *multipart.FileHeader) (err error) {
	var article system.Article
	if err, article = MD2Article(file); err != nil {
		return err
	}
	articleService.AddArticle(article)
	return nil
}

// 保存图片返回url
func (articleService *ArticleService) SaveArticleImg(file *multipart.FileHeader) (err error, url string) {
	uploadPath, err := filepath.Abs("./resource/uploads")
	if err != nil {
		return err, ""
	}

	// ext := filepath.Ext(file.Filename)
	// 创建上传目录（如果不存在）
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return err, ""
	}
	dst := filepath.Join(uploadPath, file.Filename)
	fileContent, err := file.Open()
	if err != nil {
		return
	}
	defer fileContent.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err, ""
	}
	defer out.Close()

	// 将上传的文件内容复制到目标文件
	if _, err = io.Copy(out, fileContent); err != nil {
		return err, ""
	}

	return nil, dst

}
