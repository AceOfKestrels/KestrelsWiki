using kestrelswiki.models;

namespace kestrelswiki.service.article;

public interface IArticleStore
{
    public Article? Get(string path);
    public bool Set(Article article);
    public bool Reset();
}