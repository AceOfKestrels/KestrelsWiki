using System.Net.Mime;
using kestrelswiki.extensions;
using kestrelswiki.service.file;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.StaticFiles;

namespace kestrelswiki.service.webpage;

public class WebpageService(ILogger logger, IFileReader fileReader, IContentTypeProvider contentTypeProvider)
    : IWebpageService
{
    public Try<PhysicalFileResult> TryGetFile(string path)
    {
        Try<PhysicalFileResult> tri;
        if (path.IsNullOrWhiteSpace())
        {
            tri = Try<PhysicalFileResult>.Fail(FormatErrorMessage(path, "Path is empty."));
        }
        else if (!fileReader.Exists(path).Result)
        {
            tri = Try<PhysicalFileResult>.Fail(FormatErrorMessage(path, "File not found."));
        }
        else
        {
            contentTypeProvider.TryGetContentType(path, out string? contentType);
            tri = new(new PhysicalFileResult(path, contentType ?? MediaTypeNames.Text.Plain));
        }

        tri.Catch(e => logger.Write(FormatErrorMessage(path, e.Message)));
        return tri.Result is not null
            ? new(tri.Result)
            : Try<PhysicalFileResult>.Fail(FormatErrorMessage(path, tri.Exception?.Message));
    }

    private string FormatErrorMessage(string? path, string? message)
    {
        return $"Error getting webpage{(path is not null ? $"at {path}" : "")}: {message}";
    }
}