import { Schema, model } from 'mongoose';
import { IUserConfigurationEnum } from './user.enum';

const configurationSchema = new Schema({
  type: {
    type: String,
    enum: [...Object.values(IUserConfigurationEnum)],
    default: IUserConfigurationEnum.DAILY
  },
  limit: {
    type: Number,
    default: 0
  },
  count: {
    type: Number,
    default: 0
  }
});

const userSchema = new Schema({
  username: { type: String, required: true, unique: true },
  password: { type: String, required: true },
  configuration: { type: configurationSchema, required: true }
});

const userModel = model('user', userSchema);

export default userModel;
