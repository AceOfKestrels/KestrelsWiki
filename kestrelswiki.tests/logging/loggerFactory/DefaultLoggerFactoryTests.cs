using kestrelswiki.logging.logFormat;
using kestrelswiki.logging.logger;
using kestrelswiki.logging.loggerFactory;
using kestrelswiki.service.file;
using Moq;

namespace kestrelswiki.tests.logging.loggerFactory;

public class DefaultLoggerFactoryTests
{
    private DefaultLoggerFactory _factory;
    private Mock<IFileWriter> _fileWriterMock;

    [SetUp]
    public void Setup()
    {
        _fileWriterMock = new();
        _factory = new(new LogFormatter(""), "", _fileWriterMock.Object);
    }

    [Test]
    public void Create_CreatesMultiLogger()
    {
        ILogger logger = _factory.Create(LogDomain.Testing);

        Assert.That(logger, Is.InstanceOf<MultiLogger>());
    }
}