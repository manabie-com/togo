const passportJWT = require('passport-jwt');
const UserRepository = require('../repositories/users');
const config = require('../../config/constants');

const { jwt: jwtConfig } = config;
const jwtOptions = {
  secretOrKey: jwtConfig.accessSecretKey,
  jwtFromRequest: passportJWT.ExtractJwt.fromAuthHeaderAsBearerToken(),
};

const jwtVerify = async (payload, cbDone) => {
  try {
    const user = await UserRepository.getUser({
      where: { id: payload.id },
      attributes: ['id', 'username', 'max_todo'],
    });

    cbDone(null, user && user.toJSON());
  } catch (error) {
    cbDone(error, false);
  }
};

const jwtStrategy = new passportJWT.Strategy(jwtOptions, jwtVerify);

module.exports = { jwtStrategy };
