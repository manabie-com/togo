import dotenv from "dotenv";
dotenv.config();

export const jwtConstants = {
  secret: process.env.JWT_SECRET,
};

export const saltRounds = parseInt(process.env.SALT_ROUND);
