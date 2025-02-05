namespace kestrelswiki.models;

public class Article
{
    public Article(string path, string title, string content)
    {
        Path = path;
        Title = title;
        Content = content;
    }

    public string Path { get; private set; }

    public string Title { get; private set; }

    public string Content { get; private set; }
}