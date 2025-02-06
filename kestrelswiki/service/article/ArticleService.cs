using System.Collections.Generic;
using kestrelswiki.environment;
using kestrelswiki.service.file;

namespace kestrelswiki.service.article;

public class ArticleService(ILogger logger, IFileReader fileReader, IArticleStore store) : IArticleService
{
    public bool Exists(string path)
    {
        return store.Get(path) is not null;
    }

    public Try<bool> RebuildIndex()
    {
        Try<IEnumerable<FileInfo>> tri = fileReader.GetMarkdownFiles(Variables.ContentPath);
        List<Exception> exceptions = [];

        if (tri.Result is not null)
            foreach (FileInfo file in tri.Result)
                AddToIndex(file).Catch(exceptions.Add);

        if (tri.Exception is AggregateException aggEx)
            foreach (Exception ex in aggEx.InnerExceptions)
                logger.Write(ex.Message);
        else if (tri.Exception is not null) logger.Write(tri.Exception.Message);

        return new(tri.Result is not null, new AggregateException(exceptions));
    }

    protected Try<bool> AddToIndex(FileInfo file)
    {
        Try<string> tri = fileReader.TryReadAllText(file);
        // use fileReader to get article infos. we do not interact with physical file here
        // if (tri.Result is not null) store.Set(new(file.FullName, file.Name, tri.Result));
        return new(true);
    }
}