class ConfigService {
    constructor(Config) {
        this.config = Config
    }

    async getLimitByRole(roleName) {
        console.log('roleName', roleName)
        const config = await this.config.findOne({
            where: {
                role: roleName
            }
        })
        if (!config) {
            throw new Error('Config does not exist')
        }
        return config
    }
}

module.exports = ConfigService