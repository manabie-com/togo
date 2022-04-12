import { Injectable } from '@nestjs/common';
import dotenv from 'dotenv';
import fs from 'fs';
import { EnvironmentConstants } from '../constants/environment.constant';

@Injectable()
export class EnvironmentService {
  private readonly settings: { [key: string]: string };

  constructor() {
    this.settings = dotenv.parse(
      fs.readFileSync(EnvironmentConstants.ENV_CONFIG_PATH),
    );
  }

  public getKey(key: string) {
    return this.settings[key];
  }
}
