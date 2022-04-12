export class UserAccount {
  id: string;
  username: string;

  constructor(user: any) {
    this.id = user?.id;
    this.username = user?.username;
  }
}
