import { Types } from 'mongoose';
import { StatusCode } from './enum';

interface IErrorByStatusCode {
  [key: string]: [
    {
      code: string;
      message: string;
    }
  ];
}

export const displayErrorsDescription = (errorList: {
  [key: string]: {
    statusCode: StatusCode;
    message: string;
  };
}) => {
  return `| HTTP Code | Error Code | Message |\n|---|---|---|\n${Object.entries(
    // group error by http statusCode
    Object.entries(errorList).reduce<IErrorByStatusCode>(
      (prevVal, [code, error]) => {
        if (!prevVal[error.statusCode]) {
          prevVal[error.statusCode] = [
            {
              code,
              message: error.message
            }
          ];
        } else {
          prevVal[error.statusCode].push({
            code,
            message: error.message
          });
        }
        return prevVal;
      },
      {}
    )
  )
    .map(([statusCode, errors]) => {
      // display each error message
      const errorsString = errors
        .map((error) => `${error.code} | ${error.message} |\n`)
        .join('| | ');
      return `| **${statusCode}** | ${errorsString}`;
    })
    .join('| | | |\n')}`;
};

export const isMongoObjectId = (id: string): boolean =>
  Types.ObjectId.isValid(id);
