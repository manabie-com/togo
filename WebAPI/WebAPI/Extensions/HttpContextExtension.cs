using Common;
using Microsoft.AspNetCore.Http;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace WebAPI.Extensions
{
    public static class HttpContextExtension
    {
        public static string GetClaimValue(HttpContext context, string claimName)
        {
            return context.User.Claims.FirstOrDefault(_ => _.Type == claimName)?.Value;
        }

        public static string GetUserId(this HttpContext context)
        {
            return GetClaimValue(context, ClaimType.USERID);
        }

        public static string GetUserName(this HttpContext context)
        {
            return GetClaimValue(context, ClaimType.USERNAME);
        }
    }
}
