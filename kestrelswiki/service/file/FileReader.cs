using System.Collections.Generic;
using System.Linq;

namespace kestrelswiki.service.file;

public class FileReader(ILogger logger) : IFileReader
{
    public Try<string> TryReadAllText(string path)
    {
        if (string.IsNullOrWhiteSpace(path))
            return new(new ArgumentException("Unable to read file: Path is empty."));

        return TryReadAllText(new FileInfo(Path.GetFullPath(path)));
    }

    public Try<string> TryReadAllText(FileInfo file)
    {
        if (!file.Exists) return Try<string>.Fail("File does not exist.");

        try
        {
            using FileStream stream = file.OpenRead();
            using StreamReader reader = new(stream);
            return new(reader.ReadToEnd());
        }
        catch (Exception e)
        {
            string errorMessage = $"Unable to read file at {file.FullName}: {e.Message}";
            logger.Write(errorMessage);
            return Try<string>.Fail(errorMessage, e);
        }
    }

    // TODO: Don't do this recursively
    public Try<IEnumerable<FileInfo>> GetMarkdownFiles(string path)
    {
        DirectoryInfo directory = new(path);
        if (!directory.Exists) return new([]);

        List<FileInfo> files = directory.GetFiles("*.md").ToList();
        List<Exception> exceptions = [];

        foreach (DirectoryInfo subDir in directory.GetDirectories().ToList().FindAll(d => d.Name != ".git"))
        {
            Try<IEnumerable<FileInfo>> tri = GetMarkdownFiles(subDir.FullName);
            if (tri.Result is not null) files.AddRange(tri.Result ?? []);

            switch (tri.Exception)
            {
                case null:
                    break;
                case AggregateException aggEx:
                    exceptions.AddRange(aggEx.InnerExceptions);
                    break;
                default:
                    exceptions.Add(tri.Exception);
                    break;
            }
        }

        return new(files, new AggregateException(exceptions));
    }

    public Try<bool> Exists(string path)
    {
        try
        {
            return new(File.Exists(path));
        }
        catch (Exception e)
        {
            return Try<bool>.Fail(e.Message);
        }
    }
}