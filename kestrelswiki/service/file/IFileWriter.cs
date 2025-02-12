namespace kestrelswiki.service.file;

public interface IFileWriter
{
    Try<bool> Write(string contents, string fileName);
    Try<bool> WriteLine(string contents, string fileName);

    Try<bool> CreateDirectory(string directoryName);
}