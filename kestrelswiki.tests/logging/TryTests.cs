using kestrelswiki.logging;

namespace kestrelswiki.tests.logging;

public class TryTests
{
    [Test]
    public void Constructor_Success_CreatesSuccessInstanceWithResult()
    {
        string value = "value";
        Try<string> tri = new(value);

        Assert.Multiple(() =>
        {
            Assert.That(tri.Success, Is.True);
            Assert.That(tri.Result, Is.EqualTo(value));
            Assert.That(tri.Exception, Is.Null);
        });
    }

    [Test]
    public void Constructor_Fail_CreatesFailedInstanceWithException()
    {
        Try<string> tri = new(new());

        Assert.Multiple(() =>
        {
            Assert.That(tri.Success, Is.False);
            Assert.That(tri.Result, Is.Null);
            Assert.That(tri.Exception, Is.Not.Null);
        });
    }

    [Test]
    public void Catch_ReturnsInstance()
    {
        Try<string> tri = new("value");

        Try<string> caught = tri.Catch(_ => { });

        Assert.That(caught, Is.EqualTo(tri));
    }

    [Test]
    public void Then_ReturnsInstance()
    {
        Try<string> tri = new("value");

        Try<string> caught = tri.Then(_ => { });

        Assert.That(caught, Is.EqualTo(tri));
    }

    [Test]
    public void Then_Success_RunsAction()
    {
        Try<string> tri = new("value");
        bool run = false;

        tri.Then(_ => run = true);

        Assert.That(run, Is.True);
    }

    [Test]
    public void Then_Fail_DoesNotRun()
    {
        Try<string> tri = new(new());
        bool run = false;

        tri.Then(_ => run = true);

        Assert.That(run, Is.False);
    }

    [Test]
    public void Catch_Fail_RunsAction()
    {
        Try<string> tri = new(new());
        bool run = false;

        tri.Catch(_ => run = true);

        Assert.That(run, Is.True);
    }

    [Test]
    public void Catch_Success_DoesNotRun()
    {
        Try<string> tri = new("value");
        bool run = false;

        tri.Catch(_ => run = true);

        Assert.That(run, Is.False);
    }
}