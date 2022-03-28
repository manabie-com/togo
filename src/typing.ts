type TTask = {
  id: number
  name: string
  isFinished: number
  assignedUserId: number
  assignedDate: number
}

type TUser = {
  id: number
  name: string
  limitTask: number
}

export {
  TTask,
  TUser,
}
