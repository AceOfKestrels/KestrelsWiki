namespace kestrelswiki.service.file;

public interface IFileReader
{
    /// <summary>
    ///     Opens a file at path and reads its contents, returning a string containing it.
    /// </summary>
    /// <param name="path">The path to read at.</param>
    /// <param name="content"></param>
    /// <returns>The contents of the file, or null if an error occurs.</returns>
    Try<string> TryReadAllText(string path);

    Try<bool> Exists(string path);
}