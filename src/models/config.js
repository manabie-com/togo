const config = (sequelize, DataTypes) => {
    const Config = sequelize.define('Config', {
        role: {
            type: DataTypes.STRING,
            unique: true,
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

    return Config
}

module.exports = config