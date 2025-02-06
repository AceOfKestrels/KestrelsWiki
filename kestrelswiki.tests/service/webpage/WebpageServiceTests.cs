using kestrelswiki.logging.logger;
using kestrelswiki.service.file;
using kestrelswiki.service.webpage;
using Microsoft.AspNetCore.StaticFiles;
using Moq;

namespace kestrelswiki.tests.service.webpage;

public class WebpageServiceTests
{
    private Mock<IContentTypeProvider> _mockContentTypeProvider;
    private Mock<IFileReader> _mockFileReader;
    private Mock<ILogger> _mockLogger;
    private WebpageService _webpageService;

    [SetUp]
    public void Setup()
    {
        _mockLogger = new();
        _mockFileReader = new();
        _mockContentTypeProvider = new();
        _webpageService = new(_mockLogger.Object, _mockFileReader.Object, _mockContentTypeProvider.Object);
    }

    // TODO: Reimplement tests
}