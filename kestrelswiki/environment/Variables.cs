using System.Collections.Generic;
using System.Linq;
using kestrelswiki.extensions;
using Microsoft.Extensions.Logging;
using ILogger = kestrelswiki.logging.logger.ILogger;

namespace kestrelswiki.environment;

public static class Variables
{
    public static readonly string LogPath = Environment.GetEnvironmentVariable("LOG_PATH") ?? ILogger.DefaultPath;

    public static readonly string ContentPath =
        Environment.GetEnvironmentVariable("CONTENT_PATH") ?? Path.Combine("..", "content");

    public static readonly string WebRootPath =
        Environment.GetEnvironmentVariable("WEBROOT_PATH") ?? Path.Combine("..", "wwwroot");

    public static readonly bool EnableFileLogging =
        Environment.GetEnvironmentVariable("FILE_LOGGING")?.ToLowerInvariant() == "true";

    public static readonly string
        WebPageRepo = Environment.GetEnvironmentVariable("WEBPAGE_REPOSITORY") ??
                      "https://github.com/AceOfKestrels/kestrelsnest.git";

    public static readonly string
        ContentRepository = Environment.GetEnvironmentVariable("CONTENT_REPOSITORY") ??
                            "https://github.com/AceOfKestrels/kestrelsnest.git";

    public static readonly string LogDateFormat =
        Environment.GetEnvironmentVariable("LOG_DATE_FORMAT") ?? ILogger.DefaultDateFormat;

    public static readonly LogLevel LogLevel =
        Environment.GetEnvironmentVariable("LOG_LEVEL")?.ToLogLevel() ?? LogLevel.Information;

    public static readonly HashSet<string> DisabledLogDomains =
        Environment.GetEnvironmentVariable("DISABLED_LOG_DOMAINS")?.Split(",").ToHashSet() ?? [];

    public static class Webpage
    {
        public static readonly string HomeDirectory = Environment.GetEnvironmentVariable("HOME_DIR") ?? "homepage";

        public static readonly string FrontpageDirectory =
            Environment.GetEnvironmentVariable("FRONTPAGE_DIR") ?? "wikifrontpage";

        public static readonly string ArticleDirectory =
            Environment.GetEnvironmentVariable("ARTICLE_DIR") ?? "wikiarticle";

        public static readonly string NotFoundDirectory =
            Environment.GetEnvironmentVariable("NOT_FOUND_DIR") ?? "error";

        public static readonly string GlobalFileDirectory =
            Environment.GetEnvironmentVariable("GLOBAL_DIR") ?? "global";
    }
}