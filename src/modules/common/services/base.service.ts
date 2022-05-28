export abstract class BaseService {
  public baseRelations: string[] = [];

  public relations: string[] = [];

  constructor(relations: string[] = []) {
    this.baseRelations = relations;
    this.setRelations(relations, true);
  }

  setRelations(relations: string[], replace: boolean = false) {
    if (replace) {
      this.relations = [...relations];
    } else {
      this.relations = [...this.baseRelations];

      relations.forEach((rel) => {
        if (!this.relations.includes(rel)) {
          this.relations.push(rel);
        }
      });
    }
  }

  getRelations(): string[] {
    return this.relations;
  }
}
