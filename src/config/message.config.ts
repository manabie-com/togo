export enum ErrorName {
    TokenExpiredError = 'TokenExpiredError',
    JsonWebTokenError = 'JsonWebTokenError',
}

export class MessageConfig {
    // General
    static readonly SUCCESS = 'SUCCESS';
    static readonly FAILURE = 'FAILURE';
    static readonly WELCOME = 'Bạn đang sử dụng hệ thống API {{SERVER_ENV}} của MyLife Company';
    static readonly URL_INVALID = 'URL_INVALID';
    static readonly VALIDATION_ERROR = 'ValidationError';
    static readonly CAST_ERROR = 'CastError';
    static readonly MONGO_ERROR = 'MongoError';
    static readonly MULTER_ERROR = 'MulterError';
    static readonly UNAUTHORIZED = 'Unauthorized';
    static readonly JSON_WEB_TOKEN_ERROR = 'JsonWebTokenError';
    static readonly TOKEN_EXPIRED_ERROR = 'TokenExpiredError';
    static readonly FORBIDDEN = 'FORBIDDEN';
    static readonly PERMISSION_DENIED = 'PERMISSION_DENIED';
    static readonly NOT_FOUND = 'NOT_FOUND';
    static readonly UNKNOWN_ERROR = 'UNKNOWN_ERROR';
    static readonly DELETION_REJECT = 'DELETION_REJECT';
    static readonly UNIQUE_FIELD = 'Value `{VALUE}` is not acceptable. Field `{PATH}` must be unique.';

    // User
    static readonly FULL_NAME_REQUIRED = 'FULL_NAME_REQUIRED';
    static readonly COUNTRY_CODE_REQUIRED = 'COUNTRY_CODE_REQUIRED';
    static readonly PHONE_REQUIRED = 'PHONE_REQUIRED';
    static readonly PHONE_INVALID = 'PHONE_INVALID';
    static readonly PHONE_EXISTED = 'PHONE_EXISTED';
    static readonly EMAIL_REQUIRED = 'EMAIL_REQUIRED';
    static readonly EMAIL_INVALID = 'EMAIL_INVALID';
    static readonly EMAIL_EXISTED = 'EMAIL_EXISTED';
    static readonly ROLE_REQUIRED = 'ROLE_REQUIRED';
    static readonly STATUS_REQUIRED = 'STATUS_REQUIRED';
    static readonly PASSWORD_REQUIRED = 'PASSWORD_REQUIRED';
    static readonly PASSWORD_INVALID = 'PASSWORD_INVALID';
    static readonly PASSWORD_WEAK = 'PASSWORD_WEAK';
    static readonly NEW_PASSWORD_REQUIRED = 'NEW_PASSWORD_REQUIRED';
    static readonly CONFIRMED_PASSWORD_REQUIRED = 'CONFIRMED_PASSWORD_REQUIRED';
    static readonly CONFIRMED_PASSWORD_NOT_MATCH = 'CONFIRMED_PASSWORD_NOT_MATCH';
    static readonly PASSWORD_NOT_SETUP = 'PASSWORD_NOT_SETUP';
    static readonly NEW_PASSWORD_MATCH_OLD_PASSWORD = 'NEW_PASSWORD_MATCH_OLD_PASSWORD';

    // Company
    static readonly COMPANY_CODE_REQUIRED = 'COMPANY_CODE_REQUIRED';
    static readonly COMPANY_CODE_EXIST = 'COMPANY_CODE_EXIST';
    static readonly COMPANY_TITLE_REQUIRED = 'COMPANY_TITLE_REQUIRED';
    static readonly COMPANY_SLUG_REQUIRED = 'COMPANY_SLUG_REQUIRED';
    static readonly COMPANY_TYPE_REQUIRED = 'COMPANY_TYPE_REQUIRED';
    static readonly COMPANY_TAX_TYPE_REQUIRED = 'COMPANY_TAX_TYPE_REQUIRED';
    static readonly COMPANY_STATUS_REQUIRED = 'COMPANY_STATUS_REQUIRED';
    static readonly COMPANY_ADDRESS_REQUIRED = 'COMPANY_ADDRESS_REQUIRED';
    static readonly COMPANY_NOT_FOUND = 'COMPANY_NOT_FOUND';
    static readonly SHOP_NOT_FOUND = 'SHOP_NOT_FOUND';
    static readonly COMPANY_REQUIRED = 'COMPANY_REQUIRED';

    // Area
    static readonly AREA_CODE_REQUIRED = 'AREA_CODE_REQUIRED';
    static readonly AREA_CODE_EXIST = 'AREA_CODE_EXIST';
    static readonly AREA_TITLE_REQUIRED = 'AREA_TITLE_REQUIRED';
    static readonly AREA_TYPE_REQUIRED = 'AREA_TYPE_REQUIRED';
    static readonly AREA_STATUS_REQUIRED = 'AREA_STATUS_REQUIRED';
    static readonly AREA_ADDRESS_REQUIRED = 'AREA_ADDRESS_REQUIRED';
    static readonly AREA_NOT_FOUND = 'AREA_NOT_FOUND';
    static readonly AREA_REQUIRED = 'AREA_REQUIRED';

    // Table
    static readonly TABLE_CODE_REQUIRED = 'TABLE_CODE_REQUIRED';
    static readonly TABLE_CODE_EXIST = 'TABLE_CODE_EXIST';
    static readonly TABLE_NAME_REQUIRED = 'TABLE_NAME_REQUIRED';
    static readonly TABLE_TYPE_REQUIRED = 'TABLE_TYPE_REQUIRED';
    static readonly TABLE_STATUS_REQUIRED = 'TABLE_STATUS_REQUIRED';
    static readonly TABLE_ADDRESS_REQUIRED = 'TABLE_ADDRESS_REQUIRED';
    static readonly TABLE_NOT_FOUND = 'TABLE_NOT_FOUND';
    static readonly TABLE_REQUIRED = 'TABLE_REQUIRED';

    // Category
    static readonly CATEGORY_CODE_REQUIRED = 'CATEGORY_CODE_REQUIRED';
    static readonly CATEGORY_NAME_REQUIRED = 'CATEGORY_NAME_REQUIRED';
    static readonly CATEGORY_NAME_JP_REQUIRED = 'CATEGORY_NAME_JP_REQUIRED';
    static readonly CATEGORY_NAME_KANJI_REQUIRED = 'CATEGORY_NAME_KANJI_REQUIRED';
    static readonly CATEGORY_TYPE_REQUIRED = 'CATEGORY_TYPE_REQUIRED';
    static readonly CATEGORY_STATUS_REQUIRED = 'CATEGORY_STATUS_REQUIRED';
    static readonly CATEGORY_CODE_EXIST = 'CATEGORY_CODE_EXIST';
    static readonly CATEGORY_NOT_FOUND = 'CATEGORY_NOT_FOUND';
    static readonly CATEGORY_REQUIRED = 'CATEGORY_REQUIRED';

    // Product
    static readonly PRODUCT_CODE_REQUIRED = 'PRODUCT_CODE_REQUIRED';
    static readonly PRODUCT_TEMPLATE_REQUIRED = 'PRODUCT_TEMPLATE_REQUIRED';
    static readonly PRODUCT_NAME_REQUIRED = 'PRODUCT_NAME_REQUIRED';
    static readonly PRODUCT_NAME_JP_REQUIRED = 'PRODUCT_NAME_JP_REQUIRED';
    static readonly PRODUCT_NAME_KANJI_REQUIRED = 'PRODUCT_NAME_KANJI_REQUIRED';
    static readonly PRODUCT_TYPE_REQUIRED = 'PRODUCT_TYPE_REQUIRED';
    static readonly PRODUCT_STATUS_REQUIRED = 'PRODUCT_STATUS_REQUIRED';
    static readonly PRODUCT_CODE_EXIST = 'PRODUCT_CODE_EXIST';
    static readonly PRODUCT_NOT_FOUND = 'PRODUCT_NOT_FOUND';
    static readonly PRODUCT_ORDER_MIN_0 = 'PRODUCT_ORDER_MIN_0';
    static readonly PRODUCT_OPTION_ORDER_MIN_0 = 'PRODUCT_OPTION_ORDER_MIN_0';
    static readonly PRODUCT_OPTION_REQUIRED = 'PRODUCT_OPTION_REQUIRED';
    static readonly PRODUCT_REQUIRED = 'PRODUCT_REQUIRED';

    // Price
    static readonly PRICE_CODE_REQUIRED = 'PRICE_CODE_REQUIRED';
    static readonly PRICE_TITLE_REQUIRED = 'PRICE_CODE_REQUIRED';
    static readonly PRICE_STATUS_REQUIRED = 'PRICE_STATUS_REQUIRED';

    // Item
    static readonly ITEM_PRICE_REQUIRED = 'ITEM_PRICE_REQUIRED';
    static readonly ITEM_PRICE_MIN_0 = 'ITEM_PRICE_MIN_0';
    static readonly ITEM_VAT_MIN_0 = 'ITEM_VAT_MIN_0';
    static readonly ITEM_VAT_MAX_20 = 'ITEM_VAT_MAX_20';
    static readonly ITEM_CURRENCY_REQUIRED = 'ITEM_CURRENCY_REQUIRED';
    static readonly ITEM_EXIST = 'ITEM_EXIST';
    static readonly ITEM_REQUIRED = 'ITEM_REQUIRED';

    // Item
    static readonly STOCK_CODE_REQUIRED = 'STOCK_CODE_REQUIRED';
    static readonly STOCK_MIN_0 = 'STOCK_MIN_0';
    static readonly STOCK_STATUS_REQUIRED = 'STOCK_STATUS_REQUIRED';
    static readonly STOCK_DUPLICATED = 'STOCK_DUPLICATED';

    // Location
    static readonly LOCATION_TYPE_REQUIRED = 'LOCATION_TYPE_REQUIRED';
    static readonly LOCATION_COORDINATES_REQUIRED = 'LOCATION_COORDINATES_REQUIRED';

    // Device
    static readonly DEVICE_CODE_REQUIRED = 'DEVICE_CODE_REQUIRED';
    static readonly DEVICE_UNIQUE_ID_REQUIRED = 'DEVICE_UNIQUE_ID_REQUIRED';
    static readonly DEVICE_NAME_REQUIRED = 'DEVICE_NAME_REQUIRED';
    static readonly DEVICE_TYPE_REQUIRED = 'DEVICE_TYPE_REQUIRED';
    static readonly DEVICE_OS_NAME_REQUIRED = 'DEVICE_OS_NAME_REQUIRED';
    static readonly DEVICE_OS_VERSION_REQUIRED = 'DEVICE_OS_VERSION_REQUIRED';
    static readonly DEVICE_STATUS_REQUIRED = 'DEVICE_STATUS_REQUIRED';
    static readonly DEVICE_CODE_EXIST = 'DEVICE_CODE_EXIST';
    static readonly DEVICE_ADDRESS_REQUIRED = 'DEVICE_ADDRESS_REQUIRED';
    static readonly DEVICE_LOCATION_REQUIRED = 'DEVICE_LOCATION_REQUIRED';
    static readonly DEVICE_NOT_FOUND = 'DEVICE_NOT_FOUND';
    static readonly DEVICE_BATTERY_MIN_0 = 'DEVICE_BATTERY_MIN_0';
    static readonly DEVICE_BATTERY_MAX_100 = 'DEVICE_BATTERY_MAX_100';
    static readonly DEVICE_BATTERY_STATE_REQUIRED = 'DEVICE_BATTERY_STATE_REQUIRED';
    static readonly DEVICE_APP_VERSION_REQUIRED = 'DEVICE_APP_VERSION_REQUIRED';
    static readonly DEVICE_BUILD_NUMBER_REQUIRED = 'DEVICE_BUILD_NUMBER_REQUIRED';
    static readonly DEVICE_IP_ADDRESS_INVALID = 'DEVICE_IP_ADDRESS_INVALID';
    static readonly DEVICE_MAC_ADDRESS_INVALID = 'DEVICE_MAC_ADDRESS_INVALID';
    static readonly DEVICE_REQUIRED = 'DEVICE_REQUIRED';

    // Page
    static readonly PAGE_CODE_REQUIRED = 'PAGE_CODE_REQUIRED';
    static readonly PAGE_ORDER_REQUIRED = 'PAGE_ORDER_REQUIRED';
    static readonly PAGE_STATUS_REQUIRED = 'PAGE_STATUS_REQUIRED';
    static readonly PAGE_REQUIRED = 'PAGE_REQUIRED';
    static readonly PAGE_NOT_FOUND = 'PAGE_NOT_FOUND';

    // Element
    static readonly ELEMENT_TYPE_REQUIRED = 'ELEMENT_TYPE_REQUIRED';
    static readonly ELEMENT_VALUE_REQUIRED = 'ELEMENT_VALUE_REQUIRED';
    static readonly ELEMENT_STATUS_REQUIRED = 'ELEMENT_STATUS_REQUIRED';
    static readonly ELEMENT_POSITION_REQUIRED = 'ELEMENT_POSITION_REQUIRED';
    static readonly ELEMENT_POSITION_RATE_REQUIRED = 'ELEMENT_POSITION_RATE_REQUIRED';

    // Position
    static readonly POSITION_LEFT_REQUIRED = 'POSITION_LEFT_REQUIRED';
    static readonly POSITION_TOP_REQUIRED = 'POSITION_TOP_REQUIRED';
    static readonly POSITION_WIDTH_REQUIRED = 'POSITION_WIDTH_REQUIRED';
    static readonly POSITION_HEIGHT_REQUIRED = 'POSITION_HEIGHT_REQUIRED';

    // Order
    static readonly LINE_STATUS_REQUIRED = 'LINE_STATUS_REQUIRED';
    static readonly ORDER_CODE_REQUIRED = 'ORDER_CODE_REQUIRED';
    static readonly ORDER_LINES_AT_LEAST_1 = 'ORDER_LINES_AT_LEAST_1';
    static readonly ORDER_CURRENCY_REQUIRED = 'ORDER_CURRENCY_REQUIRED';
    static readonly ORDER_EXISTED = 'ORDER_EXISTED';

    // Post
    static readonly POST_TITLE_REQUIRED = 'POST_TITLE_REQUIRED';
    static readonly POST_TITLE_EXIST = 'POST_TITLE_EXIST';
    static readonly SLUG_REQUIRED = 'SLUG_REQUIRED';
    static readonly SLUG_VALUE_REQUIRED = 'SLUG_VALUE_REQUIRED';
    static readonly POST_SUB_CONTENT_REQUIRED = 'POST_SUB_CONTENT_REQUIRED';
    static readonly POST_DESCRIPTION_REQUIRED = 'POST_DESCRIPTION_REQUIRED';
    static readonly POST_TYPE_REQUIRED = 'POST_TYPE_REQUIRED';
    static readonly POST_DISPLAY_REQUIRED = 'POST_DISPLAY_REQUIRED';
    static readonly POST_STATUS_REQUIRED = 'POST_STATUS_REQUIRED';
    static readonly POST_UNLIMITED_REQUIRED = 'POST_UNLIMITED_REQUIRED';
    static readonly POST_VALID_DATE_REQUIRED = 'POST_VALID_DATE_REQUIRED';
}
