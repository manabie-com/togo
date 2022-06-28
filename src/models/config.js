const config = (sequelize, DataTypes) => {
    const Config = sequelize.define('Config', {
        role: {
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        },
        limit: {
            type: DataTypes.INTEGER,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        }
    }, {
        tableName: 'configs'
    })

    Config
        .create({
            role: 'Admin',
            limit: 5
        })

    return Config
}

module.exports = config