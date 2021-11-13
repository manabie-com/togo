module.exports = (sequelize, DataTypes) => {
  const User = sequelize.define(
    'users',
    {
      id: {
        type: DataTypes.INTEGER,
        autoIncrement: true,
        primaryKey: true,
      },
      username: {
        type: DataTypes.STRING,
        allowNull: false,
      },
      password: {
        type: DataTypes.TEXT,
        allowNull: false,
      },
      max_todo: {
        type: DataTypes.INTEGER,
        allowNull: true,
        defaultValue: 5,
      },
    },
    {
      underscored: true,
      freezeTableName: true,
    },
  );

  return User;
};
