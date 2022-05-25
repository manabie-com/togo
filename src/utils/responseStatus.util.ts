type AnyRecord = Record<string, unknown>;

function Code(status: number, message: string) {
   return (arg?: AnyRecord) => ({
      message,
      status,
      ...arg,
   });
}

export const CODE_200 = Code(200, "OK: Success");
export const CODE_201 = Code(201, "OK: New resource has been created");
export const CODE_400 = Code(400, "Bad Request: Please check parameter");
export const CODE_401 = Code(401, "Unauthorized: The request requires an user authentication");
export const CODE_404 = Code(404, "Not found: There is no resource behind the URI");
export const CODE_500 = Code(500, "Internal Server Error: API developers should avoid this error");
export const CODE_504 = Code(504, "Gateway Timeout");

export default {
   CODE_200,
   CODE_201,
   CODE_400,
   CODE_401,
   CODE_404,
   CODE_500,
   CODE_504,
};
