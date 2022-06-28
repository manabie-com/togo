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
            unique: true,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        }
    }, {
        tableName: 'roles'
    })

    Role
        .create({
            code: 'Admin',
            status: true
        })

    return Role
}

module.exports = role