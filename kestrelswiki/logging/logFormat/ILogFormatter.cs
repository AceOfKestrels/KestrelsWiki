using Microsoft.Extensions.Logging;

namespace kestrelswiki.logging.logFormat;

public interface ILogFormatter
{
    string Format(LogDomain logDomain, LogLevel logLevel, object message);
}