using System.Collections.Generic;
using kestrelswiki.models;
using Microsoft.Extensions.Logging;
using ILogger = kestrelswiki.logging.logger.ILogger;

namespace kestrelswiki.service.article;

public class ArticleStore(ILogger logger) : IArticleStore
{
    private readonly Dictionary<string, Article> _articles = new();

    public Article? Get(string path)
    {
        if (_articles.TryGetValue(path, out Article? article))
            logger.Write($"Retrieved article: {article.Title} at {article.Path}", LogLevel.Debug);
        return article;
    }

    public bool Set(Article article)
    {
        _articles[article.Path] = article;
        logger.Write($"Added article: {article.Title} at {article.Path}");
        return true;
    }

    public bool Reset()
    {
        _articles.Clear();
        return true;
    }
}