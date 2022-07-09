const role = (sequelize, DataTypes) => {
    const Role = sequelize.define('Role', {
        code: {
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        },
        status: {
            type: DataTypes.BOOLEAN,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        }
    }, {
        tableName: 'roles'
    })

    return Role
}

module.exports = role