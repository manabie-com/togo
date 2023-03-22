module.exports = (sequelize, Sequelize) => {
  const Account = sequelize.define('account', {
    accountId: {
      allowNull: false,
      autoIncrement: true,
      primaryKey: true,
      type: Sequelize.INTEGER
    },
    userName: {
      type: Sequelize.STRING,
      allowNull: false,
      unique: true
    },
    passWord: {
      type: Sequelize.STRING
    },
    userId: {
      type: Sequelize.INTEGER
    },
    isActive: {
      type: Sequelize.BOOLEAN,
      allowNull: false,
      defaultValue: true
    }
  })

  Account.associate = models => {
    Account.belongsTo(models.user,
      {
        foreignKey: 'userId',
        as: 'loginUser'
      }
    )
  }

  return Account
}
