using kestrelswiki.logging;
using kestrelswiki.models;
using kestrelswiki.service.article;
using kestrelswiki.service.file;
using kestrelswiki.tests.helper;
using Microsoft.Extensions.Logging;
using Moq;
using ILogger = kestrelswiki.logging.logger.ILogger;

namespace kestrelswiki.tests.service.article;

public class ArticleServiceTests
{
    private ExposedArticleService _articleService;
    private Mock<IArticleStore> _articleStoreMock;
    private Mock<IFileReader> _fileReaderMock;
    private Mock<ILogger> _loggerMock;

    [SetUp]
    public void Setup()
    {
        _loggerMock = new();
        _fileReaderMock = new();
        _articleStoreMock = new();
        _articleService = new(_loggerMock.Object, _fileReaderMock.Object, _articleStoreMock.Object);
    }

    [Test]
    public void Exists_CheckIfPathExistsInStore()
    {
        _articleService.Exists("path");

        _articleStoreMock.Verify(s => s.Get(Any.String), Times.AtLeastOnce);
    }

    [Test]
    public void Exists_ArticleExists_ReturnTrue()
    {
        _articleStoreMock.Setup(s => s.Get(Any.String)).Returns(new Article("", "", ""));

        bool exists = _articleService.Exists("path");

        Assert.That(exists, Is.True);
    }

    [Test]
    public void Exists_ArticleDoesNotExist_ReturnFalse()
    {
        _articleStoreMock.Setup(s => s.Get(Any.String)).Returns((Article?)null);

        bool exists = _articleService.Exists("path");

        Assert.That(exists, Is.False);
    }

    [Test]
    public void RebuildIndex_FetchesListOfArticles()
    {
        _fileReaderMock.Setup(s => s.GetMarkdownFiles(Any.String)).Returns(new Try<IEnumerable<FileInfo>>([]));
        _articleService.RebuildIndex();

        _fileReaderMock.Verify(s => s.GetMarkdownFiles(Any.String), Times.AtLeastOnce);
    }

    [Test]
    public void RebuildIndex_AggregateException_LogsErrors()
    {
        _fileReaderMock.Setup(s => s.GetMarkdownFiles(Any.String))
            .Returns(new Try<IEnumerable<FileInfo>>([], new AggregateException([new(), new()])));
        _articleService.RebuildIndex();

        _loggerMock.Verify(s => s.Write(Any.String, It.IsAny<LogLevel>()), Times.Exactly(2));
    }

    [Test]
    public void RebuildIndex_SingleException_LogsErrors()
    {
        _fileReaderMock.Setup(s => s.GetMarkdownFiles(Any.String))
            .Returns(new Try<IEnumerable<FileInfo>>([], new()));
        _articleService.RebuildIndex();

        _loggerMock.Verify(s => s.Write(Any.String, It.IsAny<LogLevel>()), Times.Once);
    }

    [Test]
    public void RebuildIndex_ReturnsErrorsDuringAdding()
    {
    }

    [Test]
    public void AddToIndex()
    {
    }

    protected class ExposedArticleService(ILogger logger, IFileReader fileReader, IArticleStore store)
        : ArticleService(logger, fileReader, store)
    {
        public new Try<bool> AddToIndex(FileInfo file)
        {
            return base.AddToIndex(file);
        }
    }
}