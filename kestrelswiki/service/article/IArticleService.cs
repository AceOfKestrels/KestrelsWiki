using System.Threading.Tasks;

namespace kestrelswiki.service.article;

public interface IArticleService
{
    bool Exists(string path);

    Task<Try<bool>> RebuildIndex();
}