export class CommonConstants {
  public static readonly AUTHORIZATION_HTTP_REQUEST_HEADER = 'authorization';
  public static readonly AUTHORIZATION_HEADER_BEARER = 'Bearer';
  public static readonly AUTHORIZATION_HEADER_REGEX = /(\S+)\s+(\S+)/;

  public static readonly JWT_SECRET_KEY = 'JWT_SECRET_KEY';
  public static readonly TOKEN_EXPIRE_TIME = '60m';
}
