export default class Controller {

    #Repository = ''

    constructor() {
    }

    setRepository(Repo) {
        this.#Repository = Repo
    }

    getRepository() {
        return this.#Repository
    }

}
