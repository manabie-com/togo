import { Base, TimeStamps } from "@typegoose/typegoose/lib/defaultClasses";
import { plugin } from "@typegoose/typegoose";
import { MessageConfig } from "../config";

interface MyBase extends Base {
}

class MyBase extends TimeStamps {
}

@plugin(require('mongoose-beautiful-unique-validation'), {defaultMessage: MessageConfig.UNIQUE_FIELD})
@plugin(require('mongoose-slug-updater'))
export class BaseEntity extends MyBase implements MyBase {
}
