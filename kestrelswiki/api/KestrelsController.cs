using System.Linq;
using System.Net;
using kestrelswiki.extensions;
using kestrelswiki.logging.logFormat;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using ILogger = kestrelswiki.logging.logger.ILogger;
using ILoggerFactory = kestrelswiki.logging.loggerFactory.ILoggerFactory;

namespace kestrelswiki.api;

public abstract class KestrelsController(ILoggerFactory loggerFactory, LogDomain logDomain) : ControllerBase
{
    protected readonly ILogger logger = loggerFactory.Create(logDomain);

    protected string ClientIp
    {
        get
        {
            if (string.IsNullOrEmpty(Request.Headers["X-Forwarded-For"]))
                return Request.HttpContext.Connection.RemoteIpAddress?.MapToIPv4().ToString() ??
                       IPAddress.Loopback.ToString();

            string forwardedFor = Request.Headers["X-Forwarded-For"].ToString();
            string? firstIp = forwardedFor.Split(',').FirstOrDefault()?.Trim();

            return firstIp ?? IPAddress.Loopback.ToString();
        }
    }

    protected void LogIncomingRequest(string message = "", LogLevel logLevel = LogLevel.Information)
    {
        logger.Write(
            $"{Request.Method} ({ClientIp}): {Request.Path}{Request.QueryString} {(message.IsNullOrWhiteSpace() ? "" : $" - {message}")}",
            logLevel);
    }
}