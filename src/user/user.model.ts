import { Schema, model } from 'mongoose';
import { UserConfigurationEnum } from './user.enum';

const configurationSchema = new Schema(
  {
    type: {
      type: String,
      enum: [...Object.values(UserConfigurationEnum)],
      default: UserConfigurationEnum.DAILY
    },
    limit: {
      type: Number,
      default: 0
    }
  },
  { _id: false }
);

const userSchema = new Schema({
  username: { type: String, required: true, unique: true },
  password: { type: String, required: true },
  configuration: { type: configurationSchema, required: true }
});

const userModel = model('user', userSchema);

export default userModel;
