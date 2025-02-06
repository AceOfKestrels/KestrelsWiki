using kestrelswiki.logging.logger;
using kestrelswiki.models;
using kestrelswiki.service.article;
using Moq;

namespace kestrelswiki.tests.service.article;

public class ArticleStoreTests
{
    private ArticleStore _articleStore;
    private Mock<ILogger> _mockLogger;

    [SetUp]
    public void Setup()
    {
        _mockLogger = new();
        _articleStore = new(_mockLogger.Object);
    }

    [Test]
    public void Set_SetsValueMappedToPath_GetWorks()
    {
        Article value = new("path", "title", "content");

        _articleStore.Set(value);
        Article? get = _articleStore.Get(value.Path);

        Assert.That(get, Is.Not.Null);
        Assert.That(get, Is.EqualTo(value));
    }

    [Test]
    public void Reset_ClearsStore()
    {
        Article value = new("path", "title", "content");
        _articleStore.Set(value);

        _articleStore.Reset();
        Article? get = _articleStore.Get(value.Path);

        Assert.That(get, Is.Null);
    }
}