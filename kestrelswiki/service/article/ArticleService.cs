using System.Collections.Generic;
using System.Threading.Tasks;
using kestrelswiki.environment;
using kestrelswiki.service.file;

namespace kestrelswiki.service.article;

public class ArticleService(ILogger logger, IFileReader fileReader) : IArticleService
{
    public bool Exists(string path)
    {
        return path == "true";
    }

    public async Task<Try<bool>> RebuildIndex()
    {
        Try<IEnumerable<FileInfo>> tri = fileReader.GetMarkdownFiles(Variables.ContentPath);
        if (tri.Result is not null)
            foreach (FileInfo file in tri.Result)
                logger.Write(file.FullName);
        if (tri.Exception is AggregateException aggEx)
            foreach (Exception ex in aggEx.InnerExceptions)
                logger.Write(ex.Message);

        return new(tri.Success);
    }
}