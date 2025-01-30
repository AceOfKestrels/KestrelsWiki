using Microsoft.Extensions.Logging;

namespace kestrelswiki.extensions;

public static class LogLevelExtensions
{
    public static string DisplayName(this LogLevel level)
    {
        return level switch
        {
            LogLevel.Trace => "/TRACE",
            LogLevel.Debug => "/DEBUG",
            LogLevel.Information => "/INFO",
            LogLevel.Warning => "/WARN",
            LogLevel.Error => "/ERROR",
            LogLevel.Critical => "/CRITICAL",
            _ => ""
        };
    }
}