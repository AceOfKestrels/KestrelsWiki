namespace kestrelswiki.logging;

public class Try<T>
{
    public Try(T? result, Exception? exception = null)
    {
        Result = result;
        Exception = exception;
        Success = result is not null;
    }

    public Try(Exception exception)
    {
        Exception = exception;
        Success = false;
    }

    public T? Result { get; }
    public Exception? Exception { get; }
    public bool Success { get; }

    public static Try<T> Fail(string errorMessage, Exception? innerException = null)
    {
        return new(default, new(errorMessage, innerException));
    }

    public Try<T> Then(Action<T> action)
    {
        if (Success && Result is not null) action(Result);
        return this;
    }

    public Try<T> Catch(Action<Exception> action)
    {
        if (Exception is not null)
            action(
                Exception ??
                new NullReferenceException($"{GetType().Name}<{typeof(T).Name}> failed. No exception was thrown.")
            );
        return this;
    }
}