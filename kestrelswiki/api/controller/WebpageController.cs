using kestrelswiki.environment;
using kestrelswiki.logging.logFormat;
using kestrelswiki.logging.loggerFactory;
using kestrelswiki.service.article;
using kestrelswiki.service.file;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.StaticFiles;

namespace kestrelswiki.api.controller;

[ApiController]
public class WebpageController(
    ILoggerFactory loggerFactory,
    IContentTypeProvider contentTypeProvider,
    IArticleService articleService,
    IFileReader fileReader)
    : KestrelsController(loggerFactory, LogDomain.WebpageController)
{
    protected WebpageInfo ArticlePage = new(Variables.Webpage.ArticleDirectory);
    protected WebpageInfo FrontPage = new(Variables.Webpage.FrontpageDirectory);
    protected WebpageInfo HomePage = new(Variables.Webpage.HomeDirectory);
    protected WebpageInfo NotFoundPage = new(Variables.Webpage.NotFoundDirectory);

    /*
     * Get webpage html at /, /wiki, /wiki/path
     * Get file at path
     *      File path starts after web path
     *      e.g.
     *          /js/scroll.js -> home/js/scroll.js
     *          /wiki/css/style.css -> frontpage/css/style.css
     *          /wiki/*article/css/style.css -> article/css/style.css
     * Get not found as fallback
     *
     */

    [HttpGet("")]
    public ActionResult GetHomepage()
    {
        return File(HomePage.HtmlPath, MimeType.TextHtml);
    }

    [HttpGet("*path")]
    public ActionResult GetHomeFile(string path)
    {
        if (fileReader.Exists(Path.Combine(HomePage.DirPath, path)).Success)
            return GetFile(Path.Combine(HomePage.DirPath, path));

        return GetNotFoundPage();
    }

    [HttpGet("wiki")]
    public ActionResult GetWikiFrontpage()
    {
        return File(FrontPage.HtmlPath, MimeType.TextHtml);
    }

    [HttpGet("wiki/*path")]
    public ActionResult GetWikiArticlePage(string path)
    {
        if (fileReader.Exists(Path.Combine(FrontPage.DirPath, path)).Success)
            return GetFile(Path.Combine(FrontPage.DirPath, path));

        if (fileReader.Exists(Path.Combine(ArticlePage.DirPath, path)).Success)
            return GetFile(Path.Combine(ArticlePage.DirPath, path));

        if (articleService.Exists(path))
            return File(ArticlePage.HtmlPath, MimeType.TextHtml);

        return GetNotFoundPage();
    }

    protected ActionResult GetFile(string path)
    {
        contentTypeProvider.TryGetContentType(path, out string? contentType);
        return File(path, contentType ?? MimeType.TextPlain);
    }

    protected ActionResult GetNotFoundPage()
    {
        return File(NotFoundPage.HtmlPath, MimeType.TextHtml);
    }

    protected class WebpageInfo(string dirPath)
    {
        public string DirPath => Path.Combine(Variables.WebRootPath, dirPath);
        public string HtmlPath => Path.Combine(DirPath, "index.html");
    }
}