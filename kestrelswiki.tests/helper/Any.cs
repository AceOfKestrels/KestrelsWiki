using Moq;

namespace kestrelswiki.tests.helper;

public class Any
{
    public static string String => It.IsAny<string>();
}