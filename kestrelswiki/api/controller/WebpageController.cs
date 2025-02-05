using kestrelswiki.environment;
using kestrelswiki.logging.logFormat;
using kestrelswiki.service.article;
using kestrelswiki.service.webpage;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using ILoggerFactory = kestrelswiki.logging.loggerFactory.ILoggerFactory;

namespace kestrelswiki.api.controller;

[ApiController]
public class WebpageController(
    ILoggerFactory loggerFactory,
    IArticleService articleService,
    IWebpageService webpageService)
    : KestrelsController(loggerFactory, LogDomain.WebpageController)
{
    protected readonly WebpageInfo articlePage = new(Variables.Webpage.ArticleDirectory);
    protected readonly WebpageInfo frontPage = new(Variables.Webpage.FrontpageDirectory);
    protected readonly WebpageInfo homePage = new(Variables.Webpage.HomeDirectory);
    protected readonly WebpageInfo notFoundPage = new(Variables.Webpage.NotFoundDirectory);

    /// <summary>
    ///     The homepage at /
    /// </summary>
    [HttpGet("")]
    public ActionResult GetHomepage()
    {
        if (!Variables.EnableWebpageApi) return NotFound();

        LogIncomingRequest();
        // logger.Write("GET at /");
        return File(homePage.HtmlPath) ?? GetNotFoundPage();
    }

    /// <summary>
    ///     A file relative to /
    /// </summary>
    [HttpGet("{*path}")]
    public ActionResult GetHomepageFile(string path)
    {
        if (!Variables.EnableWebpageApi) return NotFound();

        LogIncomingRequest();
        return File(Path.Combine(homePage.DirPath, path)) ?? GetNotFoundPage();
    }

    /// <summary>
    ///     The wiki frontpage at /wiki/
    /// </summary>
    [HttpGet("wiki")]
    public ActionResult GetWikiFrontpage()
    {
        if (!Variables.EnableWebpageApi) return NotFound();

        LogIncomingRequest();
        return File(frontPage.HtmlPath) ?? GetNotFoundPage();
    }

    /// <summary>
    ///     A file relative to /wiki/
    ///     Or a file relative to /wiki/*article/ if that does not exist
    ///     Or the article page at /wiki/article/ if that does not exist
    ///     Or the not found page if no article exists at this path
    /// </summary>
    [HttpGet("wiki/{*path}")]
    public ActionResult GetWikiArticlePage(string path)
    {
        if (!Variables.EnableWebpageApi) return NotFound();

        LogIncomingRequest();
        return File(Path.Combine(frontPage.DirPath, path))
               ?? File(Path.Combine(articlePage.DirPath, path))
               ?? (articleService.Exists(path) ? File(articlePage.HtmlPath) : GetNotFoundPage())
               ?? GetNotFoundPage();
    }

    /// <summary>
    ///     A global file relative to /global
    /// </summary>
    [HttpGet("global/{*path}")]
    public ActionResult GetGlobalFile(string path)
    {
        if (!Variables.EnableWebpageApi) return NotFound();

        LogIncomingRequest();
        return File(Path.Combine(
            Directory.GetCurrentDirectory(),
            Variables.WebRootPath,
            Variables.Webpage.GlobalFileDirectory,
            path)
        ) ?? GetNotFoundPage();
    }

    protected ActionResult? File(string physicalPath)
    {
        ActionResult? result = webpageService.TryGetFile(physicalPath).Result;
        if (result != null) logger.Write($"Returning file at {physicalPath}", LogLevel.Debug);
        return result;
    }

    protected ActionResult GetNotFoundPage()
    {
        ActionResult? result = File(notFoundPage.HtmlPath);
        if (result == null) logger.Write("Not found page not found!");
        return result ?? NotFound();
    }

    protected class WebpageInfo(string dirPath)
    {
        public string DirPath { get; } = Path.Combine(Directory.GetCurrentDirectory(), Variables.WebRootPath, dirPath);
        public string HtmlPath => Path.Combine(DirPath, "index.html");
    }
}