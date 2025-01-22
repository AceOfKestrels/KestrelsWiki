namespace kestrelswiki.service.file;

public class FileReader(ILogger logger) : IFileReader
{
    public Try<string> TryReadAllText(string path)
    {
        if (string.IsNullOrWhiteSpace(path))
            return new(new ArgumentException("Unable to read file: Path is empty."));

        try
        {
            using FileStream stream = File.OpenRead(path);
            using StreamReader reader = new(stream);
            return new(reader.ReadToEnd());
        }
        catch (Exception e)
        {
            string errorMessage = $"Unable to read file at {path}: {e.Message}";
            logger.Write(errorMessage);
            return Try<string>.Fail(errorMessage, e);
        }
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