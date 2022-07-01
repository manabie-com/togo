using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Application.Common.Models
{
    public class Result
    {
        internal Result(bool succeeded, IEnumerable<string> errors)
        {
            Succeeded = succeeded;
            Errors = errors.ToArray();
        }

        public bool Succeeded { get; set; }

        public string[] Errors { get; set; }

        public int Code { get; set; }

        public static Result Success()
        {
            return new Result(true, Array.Empty<string>());
        }

        public static Result Failure(IEnumerable<string> errors)
        {
            return new Result(false, errors);
        }
    }

    public class Result<TData> : Result
    {
        internal Result(bool succeeded, IEnumerable<string> errors, TData data) : base(succeeded, errors)
        {
            Data = data;
        }

        public TData Data { get; set; }

        public static Result<TData> Success(TData data = default)
        {
            return new Result<TData>(true, Array.Empty<string>(), data );
        }

        public static Result<TData> Failure(IEnumerable<string> errors, TData data = default)
        {
            return new Result<TData>(false, errors, data);
        }
    }
}
