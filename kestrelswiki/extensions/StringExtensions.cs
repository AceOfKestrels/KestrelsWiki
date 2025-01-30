using Microsoft.Extensions.Logging;

namespace kestrelswiki.extensions;

public static class StringExtensions
{
    public static bool IsNullOrEmpty(this string? str)
    {
        return string.IsNullOrEmpty(str);
    }

    public static bool IsNullOrWhiteSpace(this string? str)
    {
        return string.IsNullOrWhiteSpace(str);
    }

    public static LogLevel ToLogLevel(this string str)
    {
        return Enum.TryParse(str, out LogLevel logLevel) ? logLevel : LogLevel.Information;
    }
}