using Microsoft.AspNetCore.Mvc.ModelBinding;
using Microsoft.OpenApi.Models;
using Swashbuckle.AspNetCore.SwaggerGen;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace MyTodo.BackendApi
{
    public class IgnorePropertyFilter : IOperationFilter
    {
        public void Apply(OpenApiOperation operation, OperationFilterContext context)
        {
            var excludedProperties = context.ApiDescription.ParameterDescriptions.Where(p =>
                p.Source.Equals(BindingSource.Form));

            if (excludedProperties.Any())
            {

                foreach (var excludedPropertie in excludedProperties)
                {
                    foreach (var customAttribute in excludedPropertie.CustomAttributes())
                    {
                        if (customAttribute.GetType() == typeof(JsonIgnoreAttribute))
                        {
                            for (int i = 0; i < operation.RequestBody.Content.Values.Count; i++)
                            {
                                for (int j = 0; j < operation.RequestBody.Content.Values.ElementAt(i).Encoding.Count; j++)
                                {
                                    if (operation.RequestBody.Content.Values.ElementAt(i).Encoding.ElementAt(j).Key ==
                                        excludedPropertie.Name)
                                    {
                                        operation.RequestBody.Content.Values.ElementAt(i).Encoding
                                            .Remove(operation.RequestBody.Content.Values.ElementAt(i).Encoding
                                                .ElementAt(j));
                                        operation.RequestBody.Content.Values.ElementAt(i).Schema.Properties.Remove(excludedPropertie.Name);


                                    }
                                }
                            }

                        }
                    }
                }

            }
        }
    }
}
