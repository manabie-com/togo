using System;

namespace TogoService.API.Infrastructure.Helper.MessageUtil
{
    public static class CommonMessages
    {
        public static string Ok { get { return "Ok"; } }
        public static string InvalidObject { get { return "Error occured: invalid object"; } }

        public static string GetCannotFindMsg(string entityName, Guid id)
        {
            return $"Cannot find entity {entityName} with id = {id.ToString()}";
        }

        public static string GetSuccessfulAddedItemsMsg(string entityName, int addedItems)
        {
            return $"Successfully added {addedItems} {entityName}.";
        }
    }
}
