import { Checker } from "src/interfaces/checker.interface"
import { JSONResponse } from "src/interfaces/response.interface"
import { Constants } from "./constants"

export function handleChecker(checker: Checker): JSONResponse {
    if (checker.isSuccess) {
        return {
            code: Constants.SUCCESS_CODE,
            message: Constants.SUCCESS_MESSAGE,
            data: checker.data
        }
    }
    return {
        code: Constants.FAIL_CODE,
        message: Constants.FAIL_MESSAGE,
        data: null
    }
}
