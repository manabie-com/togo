module.exports = (sequelize, DataTypes) => {
  const Tasks = sequelize.define(
    'tasks',
    {
      id: {
        type: DataTypes.INTEGER,
        autoIncrement: true,
        primaryKey: true,
      },
      user_id: {
        type: DataTypes.INTEGER,
        allowNull: false,
      },
      content: {
        type: DataTypes.TEXT,
        allowNull: false,
      },
    },
    {
      underscored: true,
    },
  );

  return Tasks;
};
