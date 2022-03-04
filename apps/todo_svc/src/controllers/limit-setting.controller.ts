import {
  Count,
  CountSchema,
  Filter,
  FilterExcludingWhere,
  repository,
  Where,
} from '@loopback/repository';
import {
  post,
  param,
  get,
  getModelSchemaRef,
  patch,
  put,
  del,
  requestBody,
  response,
} from '@loopback/rest';
import {LimitSetting} from '../models';
import {LimitSettingRepository} from '../repositories';

export class LimitSettingController {
  constructor(
    @repository(LimitSettingRepository)
    public limitSettingRepository: LimitSettingRepository,
  ) {}

  @post('/limit-settings')
  @response(200, {
    description: 'LimitSetting model instance',
    content: {'application/json': {schema: getModelSchemaRef(LimitSetting)}},
  })
  async create(
    @requestBody({
      content: {
        'application/json': {
          schema: getModelSchemaRef(LimitSetting, {
            title: 'NewLimitSetting',
            exclude: ['id'],
          }),
        },
      },
    })
    limitSetting: Omit<LimitSetting, 'id'>,
  ): Promise<LimitSetting> {
    return this.limitSettingRepository.create(limitSetting);
  }

  @get('/limit-settings/count')
  @response(200, {
    description: 'LimitSetting model count',
    content: {'application/json': {schema: CountSchema}},
  })
  async count(
    @param.where(LimitSetting) where?: Where<LimitSetting>,
  ): Promise<Count> {
    return this.limitSettingRepository.count(where);
  }

  @get('/limit-settings')
  @response(200, {
    description: 'Array of LimitSetting model instances',
    content: {
      'application/json': {
        schema: {
          type: 'array',
          items: getModelSchemaRef(LimitSetting, {includeRelations: true}),
        },
      },
    },
  })
  async find(
    @param.filter(LimitSetting) filter?: Filter<LimitSetting>,
  ): Promise<LimitSetting[]> {
    return this.limitSettingRepository.find(filter);
  }

  @patch('/limit-settings')
  @response(200, {
    description: 'LimitSetting PATCH success count',
    content: {'application/json': {schema: CountSchema}},
  })
  async updateAll(
    @requestBody({
      content: {
        'application/json': {
          schema: getModelSchemaRef(LimitSetting, {partial: true}),
        },
      },
    })
    limitSetting: LimitSetting,
    @param.where(LimitSetting) where?: Where<LimitSetting>,
  ): Promise<Count> {
    return this.limitSettingRepository.updateAll(limitSetting, where);
  }

  @get('/limit-settings/{id}')
  @response(200, {
    description: 'LimitSetting model instance',
    content: {
      'application/json': {
        schema: getModelSchemaRef(LimitSetting, {includeRelations: true}),
      },
    },
  })
  async findById(
    @param.path.string('id') id: string,
    @param.filter(LimitSetting, {exclude: 'where'})
    filter?: FilterExcludingWhere<LimitSetting>,
  ): Promise<LimitSetting> {
    return this.limitSettingRepository.findById(id, filter);
  }

  @patch('/limit-settings/{id}')
  @response(204, {
    description: 'LimitSetting PATCH success',
  })
  async updateById(
    @param.path.string('id') id: string,
    @requestBody({
      content: {
        'application/json': {
          schema: getModelSchemaRef(LimitSetting, {partial: true}),
        },
      },
    })
    limitSetting: LimitSetting,
  ): Promise<void> {
    await this.limitSettingRepository.updateById(id, limitSetting);
  }

  @put('/limit-settings/{id}')
  @response(204, {
    description: 'LimitSetting PUT success',
  })
  async replaceById(
    @param.path.string('id') id: string,
    @requestBody() limitSetting: LimitSetting,
  ): Promise<void> {
    await this.limitSettingRepository.replaceById(id, limitSetting);
  }

  @del('/limit-settings/{id}')
  @response(204, {
    description: 'LimitSetting DELETE success',
  })
  async deleteById(@param.path.string('id') id: string): Promise<void> {
    await this.limitSettingRepository.deleteById(id);
  }
}
