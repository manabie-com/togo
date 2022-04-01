import { Type, applyDecorators } from '@nestjs/common';
import { ApiOkResponse, ApiProperty, getSchemaPath } from '@nestjs/swagger';
import { ResponseDto } from './ApiReponse.dto';

export class PaginatedDto<TData> {
  @ApiProperty()
  total: number;

  @ApiProperty()
  items: TData[];
}

export const ApiPaginatedResponse = <TModel extends Type<any>>(
  model: TModel,
) => {
  return applyDecorators(
    ApiOkResponse({
      schema: {
        allOf: [
          { $ref: getSchemaPath(ResponseDto) },
          {
            properties: {
              data: {
                properties: {
                  total: {
                    type: 'number',
                  },
                  data: {
                    items: {
                      type: 'array',
                      $ref: getSchemaPath(model),
                    },
                  },
                },
              },
            },
          },
        ],
      },
    }),
  );
};
