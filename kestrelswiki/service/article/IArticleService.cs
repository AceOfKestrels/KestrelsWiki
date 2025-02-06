namespace kestrelswiki.service.article;

public interface IArticleService
{
    bool Exists(string path);

    Try<bool> RebuildIndex();
}