import { NextFunction, Request, Response } from "express";

import handleAPI from "@/middlewares/handleAPI.middleware";
import { CODE_400 } from "@/utils/responseStatus.util";

let mockRequest: Request;
let mockResponse: Response;
const nextFunction: NextFunction = jest.fn();

beforeEach(() => {
   mockRequest = { query: {} } as Request;
   mockResponse = {
      json: jest.fn(),
      send: jest.fn().mockImplementation((body) => {
         mockResponse.statusMessage = body;
         return mockResponse;
      }),
      status: jest.fn().mockImplementation((status) => {
         mockResponse.statusCode = status;
         return mockResponse;
      }),
      statusCode: 200,
      statusMessage: "OK: Success",
   } as unknown as Response;
});

describe("handleAPI function", () => {
   it("Should return status 200", async () => {
      const expectedResponse = {
         body: { data: "testing", message: "OK: Success", status: 200 },
         status: 200,
      };

      // Call function
      const stubApi = async (): Promise<{ data: string }> => {
         return { data: "testing" };
      };

      await handleAPI(stubApi)(mockRequest, mockResponse, nextFunction);
      expect(mockResponse.statusMessage).toEqual(expectedResponse.body);
      expect(mockResponse.statusCode).toEqual(expectedResponse.status);
   });

   it("Should return status 400", async () => {
      const expectedResponse = {
         body: { message: "Bad Request: Please check parameter", status: 400 },
         status: 400,
      };

      // Call function
      const stubApi = async (): Promise<{ status: number; message: string }> => {
         return CODE_400();
      };

      await handleAPI(stubApi)(mockRequest, mockResponse, nextFunction);

      expect(mockResponse.statusMessage).toEqual(expectedResponse.body);
      expect(mockResponse.statusCode).toEqual(expectedResponse.status);
   });

   it("Should return status 500", async () => {
      const expectedResponse = {
         body: {
            error: "something bad happened",
            message: "Internal Server Error: API developers should avoid this error",
            status: 500,
         },
         status: 500,
      };

      // Call function
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const stubApi = async (_req: Request, _res: Response): Promise<Response> => {
         return Promise.reject(new Error("something bad happened"));
      };

      await handleAPI(stubApi)(mockRequest, mockResponse, nextFunction);
      expect(mockResponse.statusMessage).toEqual(expectedResponse.body);
      expect(mockResponse.statusCode).toEqual(expectedResponse.status);
   });
});
