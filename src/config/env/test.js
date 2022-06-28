class Config {
    constructor() {
        this.env = 'test'
        this.PORT = process.env.PORT || 3000
        this.API_BASE = '/api'
        this.DATABASE_HOST = 'localhost'
        this.DATABASE_PORT = process.env.DATABASE_PORT || 5433
        this.DATABASE = 'togo'
        this.DATABASE_USERNAME = 'postgres'
        this.DATABASE_PASSWORD = '2maupwd123'
        this.JWT_SECRET = 'HZADgA9ttB$S!dy!hu3Rauvg!L27'
        this.COOKIE_SECURE = false
        this.COOKIE_SAME_SITE = 'Strict'
        this.PUBLIC_ROUTES = ['/', '/api/auth/signup', '/api/auth/signin']
        this.CLIENT_URL = 'http://localhost:5500'
    }
}

module.exports = new Config()