import { Injectable } from '@nestjs/common'
import { join } from 'path'
import { Low, JSONFile } from 'lowdb'
import { chain } from 'lodash'

const taskFile = join('src/database/task.json')
const userFile = join('src/database/user.json')
@Injectable()
export class AppService {
  private readonly taskDB
  private readonly userDB
  constructor() {
    const taskAdapter = new JSONFile(taskFile)
    const userAdapter = new JSONFile(userFile)
    this.taskDB = new Low(taskAdapter)
    this.userDB = new Low(userAdapter)
  }

  async takeTask(task = ''): Promise<any> {
    const userId = 2
    const tasksCount = await this.getUserTasksCount(userId)
    const user = await this.getUser(userId)
    return { tasksCount, user }
  }

  async getUserTasksCount(userId = 0) {
    await this.taskDB.read()
    this.taskDB.chain = chain(this.taskDB.data)
    const userTasks = this.taskDB.chain.find({ assignedUserId: userId }).value()
    return (userTasks || []).length
  }

  async getUser(userId = 0) {
    await this.userDB.read()
    this.userDB.chain = chain(this.userDB.data)
    const user = this.userDB.chain.find({ id: userId }).value()
    return user
  }
}
