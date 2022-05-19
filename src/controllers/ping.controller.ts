import format from "date-fns/format";
import { Request } from "express";

import { DATE_FORMATTER } from "@/utils/constant.util";

import { version } from "package.json";

export const pingResponse = async ({
   url,
}: Request): Promise<{ url: string; version: string; date: string }> => {
   return {
      date: format(new Date(), DATE_FORMATTER),
      url,
      version,
   };
};
