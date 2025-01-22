using kestrelswiki.logging.logFormat;
using kestrelswiki.logging.loggerFactory;
using Microsoft.AspNetCore.Mvc;

namespace kestrelswiki.api;

public abstract class KestrelsController(ILoggerFactory loggerFactory, LogDomain logDomain) : ControllerBase
{
    protected ILogger logger = loggerFactory.Create(logDomain);
}