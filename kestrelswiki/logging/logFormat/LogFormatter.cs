using kestrelswiki.extensions;
using Microsoft.Extensions.Logging;

namespace kestrelswiki.logging.logFormat;

public class LogFormatter(string dateFormat) : ILogFormatter
{
    public string Format(LogDomain logDomain, LogLevel logLevel, object message)
    {
        return $"{DateTime.Now.ToString(dateFormat)} [{logDomain.Name}{logLevel.DisplayName()}] {message}";
    }
}