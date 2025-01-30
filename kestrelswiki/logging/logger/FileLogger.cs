using kestrelswiki.environment;
using kestrelswiki.logging.logFormat;
using kestrelswiki.service.file;
using Microsoft.Extensions.Logging;

namespace kestrelswiki.logging.logger;

public class FileLogger(
    LogDomain logDomain,
    ILogFormatter logFormatter,
    string logFilePath,
    IFileWriter fileWriter) : ILogger
{
    public void Write(object message, LogLevel logLevel)
    {
        if (logLevel < Variables.LogLevel) return;
        if (Variables.DisabledLogDomains.Contains(logDomain.Name)) return;
        fileWriter.WriteLine(logFormatter.Format(logDomain, logLevel, message),
            Path.Combine(logFilePath, DateTime.Now.ToString("yyyy-MM-dd_HH-mm-ss") + ".log"));
    }
}