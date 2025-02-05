namespace kestrelswiki.logging.logFormat;

public record LogDomain(string Name)
{
    public static readonly LogDomain Startup = new("Startup");
    public static readonly LogDomain Logging = new("Logging");
    public static readonly LogDomain Files = new("Files");
    public static readonly LogDomain WebpageService = new("WebpageService");
    public static readonly LogDomain ArticleStore = new("ArticleStore");
    public static readonly LogDomain ArticleService = new("ArticleService");
    public static readonly LogDomain GitService = new("GitService");
    public static readonly LogDomain WebpageController = new("WebpageController");
    public static readonly LogDomain Testing = new("Testing");
}