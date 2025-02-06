namespace kestrelswiki.logging;

public class Try<T>
{
    public Try(T result, Exception? exception = null)
    {
        Result = result;
        Exception = exception;
    }

    public Try(Exception exception)
    {
        Exception = exception;
    }

    public T? Result { get; }
    public Exception? Exception { get; }
    public bool Success => Result is not null; // nullable value types?

    public static Try<T> Fail(string errorMessage, Exception? innerException = null)
    {
        return new(new(errorMessage, innerException));
    }

    public Try<T> Then(Action<T> action)
    {
        if (Result is not null) action(Result);
        return this;
    }

    public Try<T> Catch(Action<Exception> action)
    {
        if (Exception is not null) action(Exception);
        return this;
    }
}